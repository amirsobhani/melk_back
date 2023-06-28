package clientOtp

type OtpValidator struct {
	Token  int    `json:"token" validate:"required"`
	Mobile string `json:"mobile" validate:"required,len=10"`
}
