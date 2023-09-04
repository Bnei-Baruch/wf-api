package cmd

import (
	"context"
	"github.com/Bnei-Baruch/wf-api/api"
	"github.com/Bnei-Baruch/wf-api/models"
	"github.com/Bnei-Baruch/wf-api/utils"
	"github.com/Bnei-Baruch/wf-api/version"
	"github.com/spf13/viper"
	"net/http"
	"os"

	"github.com/coreos/go-oidc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	if viper.GetString("mqtt.debug") == "true" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	file, err := os.OpenFile(viper.GetString("server.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	log.Infof("Starting WF API server version %s", version.Version)

	// cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = append(corsConfig.AllowMethods, http.MethodDelete)
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	corsConfig.AllowAllOrigins = true

	// Authentication
	var oidcIDTokenVerifier *oidc.IDTokenVerifier
	if viper.GetBool("authentication.enable") {
		log.Info("Initializing Auth System")
		issuer := viper.GetString("authentication.issuer")
		oidcProvider, err := oidc.NewProvider(context.TODO(), issuer)
		if err != nil {
			log.Infof("KC init error: %s", err)
			return
		}
		oidcIDTokenVerifier = oidcProvider.Verifier(&oidc.Config{
			SkipClientIDCheck: true,
		})
	}

	// Setup mqtt
	if err := api.InitMQTT(); err != nil {
		log.Errorf("MQTT Init error: %s", err)
	}

	// Setup http
	gin.SetMode(viper.GetString("server.mode"))
	router := gin.New()
	router.Use(
		cors.New(corsConfig),
		utils.MdbLoggerMiddleware(),
		utils.EnvMiddleware(models.DB, oidcIDTokenVerifier),
		utils.ErrorHandlingMiddleware(),
		utils.AuthenticationMiddleware(),
		utils.RecoveryMiddleware())

	api.SetupRoutes(router)

	srv := &http.Server{
		Addr:    viper.GetString("server.addr"),
		Handler: router,
	}

	// service connections
	log.Infoln("Running application")
	if err := srv.ListenAndServe(); err != nil {
		log.Infof("Server listen: %s", err)
	}
}
