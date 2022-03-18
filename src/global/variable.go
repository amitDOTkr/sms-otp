package global

import (
	// "log"
	"os"
	"strconv"
	// "github.com/joho/godotenv"
)

var REDIS_URL string
var MIN_OTP int
var MAX_OTP int

var SMSGAT_USER string
var SMSGAT_PSWD string
var SMSGAT_API string
var SMSGAT_SENDERID string
var SMS_OTP_MSG string

// var DB_NAME string

func settingVariable() {

	// Godotenv code needs to be commented for production
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Println("Error: Using Godotenv package, No .env file present.")
	// }

	// Redis Database
	if redis_url := os.Getenv("REDIS_URL"); redis_url == "" {
		REDIS_URL = "localhost:6379"
	} else {
		REDIS_URL = redis_url
	}

	// OTP Valuve
	if min_otp := os.Getenv("MIN_OTP"); min_otp == "" {
		MIN_OTP = 123456
	} else {
		MIN_OTP, _ = strconv.Atoi(min_otp)
	}
	if max_otp := os.Getenv("MAX_OTP"); max_otp == "" {
		MAX_OTP = 999999
	} else {
		MAX_OTP, _ = strconv.Atoi(max_otp)
	}

	// SMSGATEWAY.CENTER API Variables
	if smsgat_api := os.Getenv("SMSGAT_API"); smsgat_api == "" {
		SMSGAT_USER = os.Getenv("SMSGAT_USER")
		SMSGAT_PSWD = os.Getenv("SMSGAT_PSWD")
	} else {
		SMSGAT_API = smsgat_api
	}

	SMSGAT_SENDERID = os.Getenv("SMSGAT_SENDERID")
	SMS_OTP_MSG = os.Getenv("SMS_OTP_MSG")

	// Initializing Redis DB Connection.
	ConnectToDB()
}

func init() {
	settingVariable()
}
