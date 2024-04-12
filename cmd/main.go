package main

import (
	"github.com/hemang/GO-OTP/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//initialize config (within route.go)
	app := api.Config{Router: router}
	//routes
	app.Routes()

	router.Run(":8000") // gin server on port 8000
}

// Curl commmands
//curl -H "Content-type:application/json" -X POST -d '{"phoneNumber":"+91##PhoneNumberHere##"}' http://localhost:8000/otp
//curl -H "Content-type:application/json" -X POST -d '{"user":{"phoneNumber":"+91##PhoneNumberHere##"},"code":"##OTPHERE##"}' http://localhost:8000/verifyOTP
