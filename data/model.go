package data

type OTPData struct {
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
	//The json struct tag in the declaration (json:"phoneNumber,omitempty") specifies how the field
	//should be serialized and deserialized when encoding or decoding the struct to and from JSON.
}

type VerifyData struct {
	User *OTPData `json:"user,omitempty" validate:"required"`
	Code string   `json:"code,omitempty" validate:"required"`
}
