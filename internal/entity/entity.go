package entity

type Request struct {
	Body Body
}

type Body struct {
	RequestID string `json:"requestId" validate:"required"`
}

type Response struct {
	ResponseID string `json:"responseId,omitemply"`
}
