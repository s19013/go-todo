package handler

import (
	"encoding/json"
	"go-todo/model"
	"go-todo/service"
	"io"
	"log"
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
		log.Printf("Error:%v\n", err)
		errorResponse.AddErrorMessage("server error")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}

	// json変化
	var todoRequest model.CreateTodoRequest
	err2 := json.Unmarshal(body, &todoRequest)
	if err2 != nil {
		log.Println("Error:Failed to parse JSON body")
		errorResponse.AddErrorMessage("Failed to parse JSON body")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}

	// subjectがあるか確認
	if todoRequest.Subject == "" {
		log.Println("Error:subject not exist")
		errorResponse.AddErrorMessage("subject not exist")
		errorResponse.CreateErrorResponse(writer, http.StatusBadRequest)
		return
	}

	// 登録してid,messageを受け取る
	response, err3 := handler.service.Create(request.Context(), todoRequest)
	if err3 != nil {
		log.Printf("Error:%v\n", err)
		errorResponse.AddErrorMessage("server error")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	err4 := json.NewEncoder(writer).Encode(response)
	if err4 != nil {
		log.Printf("Error:%v\n", err)
		errorResponse.AddErrorMessage("server error")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}
}
