package global

import (
	"encoding/json"
	"fmt"
	"log"

	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/amitdotkr/sms-otp/src/entities"
)

// type ResponseBody struct {
// 	Status     string `json:"status"`
// 	StatusCode string `json:"statusCode"`
// 	Reason     string `json:"reason"`
// }

func SmsGatewayCenter(mobile, otp string, response chan entities.ResponseBody) {

	// defer wg.Done()

	gateway_url := "https://www.smsgateway.center/SMSApi/rest/send"

	// Adding Random otp no. in DLT template sms string.
	msg := fmt.Sprintf(SMS_OTP_MSG, otp)

	data := url.Values{}
	data.Add("userId", SMSGAT_USER)
	data.Add("password", SMSGAT_PSWD)
	data.Add("senderId", SMSGAT_SENDERID)
	data.Add("sendMethod", "simpleMsg")
	data.Add("msgType", "text")
	data.Add("mobile", mobile)
	data.Add("msg", msg)
	data.Add("duplicateCheck", "true")
	data.Add("format", "json")

	payload := strings.NewReader(data.Encode())

	req, _ := http.NewRequest("POST", gateway_url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var responseBody entities.ResponseBody

	json.Unmarshal([]byte(body), &responseBody)

	if responseBody.Status != "success" {
		log.Printf("SmsGatewayCenter Api: Unable to send sms-otp, statusCode: %s & reason: %s", responseBody.StatusCode, responseBody.Reason)
	}

	response <- responseBody

}
