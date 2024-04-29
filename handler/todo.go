package handler

import (
	"encoding/json"
	"go-todo/model"
	"go-todo/service"
	"io"
	"net/http"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTODOHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{
		service: service,
	}
}

func (handler *TodoHandler) Create(writer http.ResponseWriter, request *http.Request) {
	errorResponse := model.NewErrorMessages()

	// body取り出し
	body, err := io.ReadAll(request.Body)
	if err != nil {
		errorResponse.AddErrorMessage("server error")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}

	// json変化
	var todoRequest model.CreateTodoRequest
	err2 := json.Unmarshal(body, &todoRequest)
	if err2 != nil {
		errorResponse.AddErrorMessage("Failed to parse JSON body")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}

	// subjectがあるか確認
	if todoRequest.Subject == "" {
		errorResponse.AddErrorMessage("subject not exist")
		errorResponse.CreateErrorResponse(writer, http.StatusBadRequest)
		return
	}

	// 登録してid,messageを受け取る
	response, err3 := handler.service.Create(request.Context(), todoRequest)
	if err3 != nil {
		errorResponse.AddErrorMessage("server error")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	err4 := json.NewEncoder(writer).Encode(response)
	if err4 != nil {
		errorResponse.AddErrorMessage("server error")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}
}
