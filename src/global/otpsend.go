package global

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/amitdotkr/sms-otp/src/entities"
	"github.com/amitdotkr/sms-otp/src/pb"
)

var wg sync.WaitGroup

func (*Server) OtpSend(ctx context.Context, req *pb.OtpRequest) (*pb.OtpResponse, error) {
	mobile_number := req.GetMobileNumber()

	otp := GenerateOtp()

	wg.Add(1)
	go OtpSaveInRedis(mobile_number, otp)

	response := make(chan entities.ResponseBody)
	go SmsGatewayCenter(mobile_number, strconv.Itoa(otp), response)

	wg.Wait()
	responseBody := <-response
	res := &pb.OtpResponse{
		Status:     responseBody.Status,
		StatusCode: responseBody.StatusCode,
		Reason:     responseBody.Reason,
	}
	close(response)
	return res, nil
}

// Generate Random Number for OTP.
func GenerateOtp() int {
	rand.Seed(time.Now().UnixNano())
	min := MIN_OTP
	max := MAX_OTP
	return rand.Intn(max-min) + min
}

// Saving OTP in Redis Database for Authentication
func OtpSaveInRedis(mobile_number string, otp int) {
	defer wg.Done()
	err := DB.Set(mobile_number, otp, time.Second*120).Err()
	if err != nil {
		log.Printf("Redis DB Error: %v", err)
	}
}
