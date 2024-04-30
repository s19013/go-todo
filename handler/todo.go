package handler

import (
	"encoding/json"
	"go-todo/model"
	"go-todo/service"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
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
		log.Printf("error body取り出し失敗:%v\n", err)
		errorResponse.AddErrorMessage("server error")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}

	// json変化
	var todoRequest model.CreateTodoRequest
	err2 := json.Unmarshal(body, &todoRequest)
	if err2 != nil {
		log.Printf("Error:Failed to parse JSON body:%v", err2)
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
		log.Printf("Error データベース処理失敗:%v\n", err3)
		errorResponse.AddErrorMessage("server error")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	err4 := json.NewEncoder(writer).Encode(response)
	if err4 != nil {
		log.Printf("Error jsonに変換失敗:%v\n", err4)
		errorResponse.AddErrorMessage("server error")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}
}

func (handler *TodoHandler) Get(writer http.ResponseWriter, request *http.Request) {
	errorResponse := model.NewErrorMessages()

	// パラメータを取得する関数がないので自分でどうにかするしかない
	// URLのPath部分を"/"で分割
	parts := strings.Split(request.URL.Path, "/")

	// 先にidパラメーターが空文字どうかを確認
	// [,"todo",id]の3つの要素が入った配列が帰ってくる｡
	// "/todo/"だとidは空文字になる
	// "/todo/1"だとidは1になる
	id := parts[len(parts)-1]

	if id == "" {
		log.Println("Error:id not found")
		errorResponse.AddErrorMessage("id not found")
		errorResponse.CreateErrorResponse(writer, http.StatusBadRequest)
		return
	}

	// 今は文字列のままなので数値型に変換する必要がある
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Error 数値変換失敗:%v\n", err)
		errorResponse.AddErrorMessage("server error")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}

	// note:ぶっちゃけわざわざリクエスト型にいれる必要があるかどうかは疑問
	// まあ､今後パラメータ増えるかも?だし｡
	todoRequest := model.GetTodoRequest{
		ID: intId,
	}

	// todoを受け取る
	response, err3 := handler.service.Get(request.Context(), todoRequest)
	if err3 != nil {
		log.Printf("Error データベース処理失敗:%v\n", err3)
		errorResponse.AddErrorMessage("server error")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	err4 := json.NewEncoder(writer).Encode(response)
	if err4 != nil {
		log.Printf("Error json変換に失敗:%v\n", err4)
		errorResponse.AddErrorMessage("server error")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}
}

func (handler *TodoHandler) Update(writer http.ResponseWriter, request *http.Request) {
	errorResponse := model.NewErrorMessages()

	// body取り出し
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("error body取り出し失敗:%v\n", err)
		errorResponse.AddErrorMessage("server error")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}

	// json変化
	var todoRequest model.UpdateTodoRequest
	err2 := json.Unmarshal(body, &todoRequest)
	if err2 != nil {
		log.Printf("Error:Failed to parse JSON body:%v", err2)
		errorResponse.AddErrorMessage("Failed to parse JSON body")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}

	// idがあるか確認
	if todoRequest.ID == 0 {
		log.Println("Error:Id not exist")
		errorResponse.AddErrorMessage("ID not exist")
		errorResponse.CreateErrorResponse(writer, http.StatusBadRequest)
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
	response, err3 := handler.service.Update(request.Context(), todoRequest)
	if err3 != nil {
		log.Printf("Error データベース処理失敗:%v\n", err3)
		errorResponse.AddErrorMessage("server error")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	err4 := json.NewEncoder(writer).Encode(response)
	if err4 != nil {
		log.Printf("Error jsonに変換失敗:%v\n", err4)
		errorResponse.AddErrorMessage("server error")
		errorResponse.CreateErrorResponse(writer, http.StatusInternalServerError)
		return
	}
}
