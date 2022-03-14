package global

import (
	"encoding/json"
	"fmt"
	"log"

	"io"
	"net/http"
	"net/url"
	"strings"
)

type ResponseBody struct {
	Status     string `json:"status"`
	StatusCode string `json:"statusCode"`
	Reason     string `json:"reason"`
}

func SmsGatewayCenter(mobile, otp string) {

	defer wg.Done()

	gateway_url := "https://www.smsgateway.center/SMSApi/rest/send"
	// str1 := "Your SMSGatewayCenter OTP code is %s Please use the code within 2 minutes. - Demo Message."
	msg := fmt.Sprintf(SMS_OTP_MSG, otp)
	// msg1 := fmt.Sprintf(str1, otp)
	// fmt.Printf("msg1 string: %v", msg1)

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

	// payload := strings.NewReader("userId=amitdotkr&password=wdzo9poH&senderId=SMSGAT&sendMethod=simpleMsg&msgType=text&mobile=919467783277&msg=This%20is%20my%20first%20message%20with%20SMSGateway.Center&duplicateCheck=true&format=json")
	payload := strings.NewReader(data.Encode())

	req, _ := http.NewRequest("POST", gateway_url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	// res.Body()

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	// resp := []byte(body)

	var responseBody ResponseBody

	json.Unmarshal([]byte(body), &responseBody)

	if responseBody.Status != "success" {
		log.Printf("SmsGatewayCenter Api: Unable to send sms-otp, statusCode: %s & reason: %s", responseBody.StatusCode, responseBody.Reason)
	}

	// resbody := string(body)
	// var jsonMap map[string]interface{}
	// json.Unmarshal([]byte(resbody), &jsonMap)

	// status := jsonMap["status"]

	// fmt.Println(status)
	// fmt.Println(string(body))

}
