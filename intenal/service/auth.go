package service

import (
	"github.com/alphaluqman/2FA/intenal/client"
	models "github.com/alphaluqman/2FA/intenal/pkg"
	"github.com/twilio/twilio-go"
	"log"
)

type AuthService struct {
	twilioClient *twilio.RestClient
}

func NewAuthService(twilioClient *twilio.RestClient) *AuthService {
	return &AuthService{
		twilioClient: twilioClient,
	}
}

func (s AuthService) SendOtp(request *models.SendOtpRequest) error {
	err := client.SendOtp(s.twilioClient, *request)
	if err != nil {
		log.Print("error sending otp ", err)
		return err
	}
	return nil
}

func (s AuthService) VerifyOtp(request *models.VerifyOtpRequest) error {
	err := client.CheckOtp(s.twilioClient, *request)
	if err != nil {
		log.Print("error verifying otp ", err)
		return err
	}
	return nil
}
