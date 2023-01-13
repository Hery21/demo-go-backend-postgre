package models

type JSON struct {
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
	Error        bool        `json:"error"`
	ErrorMessage string      `json:"errorMessage"`
}
