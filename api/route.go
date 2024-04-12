package api

import "github.com/gin-gonic/gin"

type Config struct {
	Router *gin.Engine
}

// modify below
func (app *Config) Routes() { //  *config type  as Reciever
	app.Router.POST("/otp", app.sendSMS())
	app.Router.POST("/verifyOTP", app.verifySMS())
}
