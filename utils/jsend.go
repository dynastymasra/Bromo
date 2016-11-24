package utils

// Jsend used to format JSEND
type Jsend struct {
	Status  string       `json:"status" binding:"required"`
	Message *string      `json:"message,omitempty"`
	Data    *interface{} `json:"data,omitempty"`
	Count   *int         `json:"count,omitempty"`
}

// FailResponse used to response if failure
func FailResponse(msg string) Jsend {
	return Jsend{Status: "failed", Message: &msg}
}

// SuccessResponse used to reponse if only return success information
func SuccessResponse() Jsend {
	return Jsend{Status: "success"}
}

// ObjectResponse used to response value if have object data to be reponses
func ObjectResponse(data interface{}) Jsend {
	return Jsend{Status: "success", Data: &data}
}
