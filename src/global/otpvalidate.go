package global

import (
	"context"
	// "log"

	"github.com/amitdotkr/sms-otp/src/pb"
)

func (*Server) OtpValidate(ctx context.Context, req *pb.OtpValidateRequest) (*pb.OtpValidateResponse, error) {

	mobile_number := req.GetMobileNumber()
	otp := req.GetOtp()

	otp_in_db := DB.Get(mobile_number).Val()
	res := otp_in_db == otp

	if res {
		DB.Del(mobile_number)
	}

	return &pb.OtpValidateResponse{Validated: res}, nil
}
