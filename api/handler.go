package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hemang/GO-OTP/data"

	"github.com/gin-gonic/gin"
)

const appTimeout = time.Second * 10

func (app *Config) sendSMS() gin.HandlerFunc { // handler function (handles HTTP requests for Gin)
	return func(c *gin.Context) { // Request and Response context of HTTP request
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.OTPData // accessing data/model.go  struct for Data entered by user
		defer cancel()           // cancel timeout when exiting the fucntion

		app.validateBody(c, &payload) // ensure correct format of user input data

		newData := data.OTPData{ // object to send to twilio service
			PhoneNumber: payload.PhoneNumber,
		}

		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJSON(c, err)
			return
		}

		app.writeJSON(c, http.StatusAccepted, "OTP sent successfully") // success message to user
	}
}

func (app *Config) verifySMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.VerifyData
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.VerifyData{ // to verify the input OTP by user to TWilio service
			User: payload.User,
			Code: payload.Code,
		}

		err := app.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
		fmt.Println("err: ", err)
		if err != nil {
			app.errorJSON(c, err)
			return
		}

		app.writeJSON(c, http.StatusAccepted, "OTP verified successfully")
	}
}
