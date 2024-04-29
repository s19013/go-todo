package model

import (
	"encoding/json"
	"net/http"
)

type ErrorMessages struct {
	Messages []string `json:"messages"`
}

func NewErrorMessages() *ErrorMessages {
	return &ErrorMessages{}
}

func (model *ErrorMessages) AddErrorMessage(message string) {
	model.Messages = append(model.Messages, message)
}

func (model *ErrorMessages) CreateErrorResponse(writer http.ResponseWriter, statuscode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statuscode)

	// エラーメッセージをjson化
	err := json.NewEncoder(writer).Encode(model)

	// 上記のエラーは_で無視ができないらしいので処理が必要
	if err != nil {
		http.Error(writer, "server error", http.StatusInternalServerError)
		return
	}
}
