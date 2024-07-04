package pkg

type SendOtpRequest struct {
	To string `json:"to"`
}

type VerifyOtpRequest struct {
	Code string `json:"code"`
	To   string `json:"to"`
}
