package client

import (
	"fmt"
	models "github.com/alphaluqman/2FA/intenal/pkg"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
	"os"
)

func InitTwilioClient() *twilio.RestClient {
	TWILIO_ACCOUNT_SID := os.Getenv("TWILIO_ACCOUNT_SID")
	TWILIO_AUTH_TOKEN := os.Getenv("TWILIO_AUTH_TOKEN")
	//var VERIFY_SERVICE_SID string = os.Getenv("VERIFY_SERVICE_SID")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TWILIO_ACCOUNT_SID,
		Password: TWILIO_AUTH_TOKEN,
	})
	return client
}
func SendOtp(client *twilio.RestClient, request models.SendOtpRequest) error {
	params := &openapi.CreateVerificationParams{}
	params.SetTo(request.To)
	params.SetChannel("sms")
	VERIFY_SERVICE_SID := os.Getenv("VERIFY_SERVICE_SID")
	resp, err := client.VerifyV2.CreateVerification(VERIFY_SERVICE_SID, params)
	if err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		fmt.Printf("Sent verification '%s'\n", *resp.Sid)
	}
	return nil
}
func CheckOtp(client *twilio.RestClient, request models.VerifyOtpRequest) error {
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(request.To)
	params.SetCode(request.Code)
	VERIFY_SERVICE_SID := os.Getenv("VERIFY_SERVICE_SID")
	resp, err := client.VerifyV2.CreateVerificationCheck(VERIFY_SERVICE_SID, params)

	if err != nil {
		fmt.Println(err.Error())
		return err
	} else if *resp.Status == "approved" {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect!")
	}
	return nil
}
