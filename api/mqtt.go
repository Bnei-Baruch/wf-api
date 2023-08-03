package api

import (
	"encoding/json"
	"fmt"
	"github.com/Bnei-Baruch/wf-api/models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var MQTT mqtt.Client

type MqttPayload struct {
	Action  string      `json:"action,omitempty"`
	ID      string      `json:"id,omitempty"`
	Name    string      `json:"name,omitempty"`
	Source  string      `json:"src,omitempty"`
	Error   error       `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Result  string      `json:"result,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type PahoLogAdapter struct {
	level log.Level
}

func NewPahoLogAdapter(level log.Level) *PahoLogAdapter {
	return &PahoLogAdapter{level: level}
}

func (a *PahoLogAdapter) Println(v ...interface{}) {
	log.Infof("MQTT: %s", fmt.Sprint(v...))
}

func (a *PahoLogAdapter) Printf(format string, v ...interface{}) {
	log.Infof("MQTT: %s", fmt.Sprintf(format, v...))
}

func InitMQTT() error {
	log.Info("MQTT: Init")
	mqtt.DEBUG = NewPahoLogAdapter(log.DebugLevel)
	mqtt.WARN = NewPahoLogAdapter(log.WarnLevel)
	mqtt.CRITICAL = NewPahoLogAdapter(log.PanicLevel)
	mqtt.ERROR = NewPahoLogAdapter(log.ErrorLevel)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(viper.GetString("mqtt.url"))
	opts.SetClientID(viper.GetString("mqtt.client_id"))
	opts.SetUsername(viper.GetString("mqtt.user"))
	opts.SetPassword(viper.GetString("mqtt.password"))
	opts.SetAutoReconnect(true)
	//opts.SetOnConnectHandler(SubMQTT)
	opts.SetConnectionLostHandler(LostMQTT)
	MQTT = mqtt.NewClient(opts)
	if token := MQTT.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func SubMQTT(c mqtt.Client) {
	if token := MQTT.Subscribe(viper.GetString("mqtt.topic"), byte(1), gotMessage); token.Wait() && token.Error() != nil {
		log.Infof("MQTT: Subscribed to: %s", viper.GetString("mqtt.topic"))
	} else {
		log.Errorf("MQTT: Subscribe error: %s", token.Error())
	}
}

func LostMQTT(c mqtt.Client, err error) {
	log.Errorf("MQTT: Lost connection: %s", err)
}

func gotMessage(c mqtt.Client, m mqtt.Message) {
	log.Debugf("MQTT: Received message: %s from topic: %s\n", m.Payload(), m.Topic())
}

func SendRespond(id string, m *MqttPayload) {
	var topic string

	if id == "false" {
		topic = viper.GetString("mqtt.srv_topic")
	} else {
		topic = viper.GetString("mqtt.srv_topic") + "/" + id
	}
	message, err := json.Marshal(m)
	if err != nil {
		log.Errorf("MQTT: Message parsing error: %s", err)
	}

	if token := MQTT.Publish(topic, byte(1), false, message); token.Wait() && token.Error() != nil {
		log.Errorf("MQTT: Publish error: %s, reason: %s", topic, token.Error())
	}

	log.Debugf("MQTT: Send message: %s to topic: %s\n", string(message), topic)
}

func SendMessage(id string) {
	var topic string
	var m interface{}
	date := time.Now().Format("2006-01-02")

	switch id {
	case "ingest":
		topic = viper.GetString("mqtt.monitor_ingest_topic")
		m, _ = models.V1FindByKV("date", date, []models.Ingest{})
	case "trimmer":
		topic = viper.GetString("mqtt.monitor_trimmer_topic")
		m, _ = models.V1FindByKV("date", date, []models.Trimmer{})
	case "archive":
		topic = viper.GetString("mqtt.monitor_archive_topic")
		m, _ = models.V1FindByKV("date", date, []models.Kmedia{})
	case "trim":
		topic = viper.GetString("mqtt.state_trimmer_topic")
		m, _ = models.FindTrimmed([]models.Trimmer{})
	case "drim":
		topic = viper.GetString("mqtt.state_dgima_topic")
		m, _ = models.FindTrimmed([]models.Dgima{})
	case "bdika":
		topic = viper.GetString("mqtt.state_aricha_topic")
		m, _ = models.FindAricha([]models.Aricha{})
	case "jobs":
		topic = viper.GetString("mqtt.state_jobs_topic")
		m, _ = models.FindJobs([]models.Job{})
	case "langcheck":
		topic = viper.GetString("mqtt.state_langcheck_topic")
		m, _ = models.GetState("langcheck")
	}

	message, err := json.Marshal(m)
	if err != nil {
		log.Errorf("MQTT: Message parsing error: %s", err)
	}

	if token := MQTT.Publish(topic, byte(1), true, message); token.Wait() && token.Error() != nil {
		log.Errorf("MQTT: Publish error: %s, reason: %s", topic, token.Error())
	}

	log.Debugf("MQTT: Send message from: %s to topic: %s\n", id, topic)
}
