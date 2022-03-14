package main

import (
	"log"

	"github.com/amitdotkr/sms-otp/src/global"
)

func main() {
	log.Println("SMS Otp MicroService Started")
	global.GrpcServer()
}
