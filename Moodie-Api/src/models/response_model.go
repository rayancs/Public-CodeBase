package models

type ResponseModel struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
	Error  []string    `json:"Error"`
}
