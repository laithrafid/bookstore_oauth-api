package app

import (
	"github.com/gin-gonic/gin"
	"github.com/laithrafid/bookstore_oauth-api/src/domain/access_token"
	"github.com/laithrafid/bookstore_oauth-api/src/http"
	"github.com/laithrafid/bookstore_oauth-api/src/repository/db"
	"github.com/laithrafid/bookstore_oauth-api/src/utils/config_utils"
	"github.com/laithrafid/bookstore_oauth-api/src/utils/logger_utils"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	router = gin.Default()
)

func StartApplication() {
	config, err := config_utils.LoadConfig(".")
	if err != nil {
		logger_utils.Error("cannot load config of application:", err)
	}

	atHandler := http.NewAccessTokenHandler(
		access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	logger_utils.Info("starting the application ....")
	router.Run(config.OauthApiAddress)
}
