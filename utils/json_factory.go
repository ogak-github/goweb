package utils

import (
	"encoding/json"
	"fmt"
	"goweb/model"
	"net/http"
	"reflect"
)

// This will do json encode
// also set header
func ResponseBody(w http.ResponseWriter, statusCode int, statusMessage string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode) // ini wajib kalau kode selain 200
	json.NewEncoder(w).Encode(model.ApiResponse{
		Code:          statusCode,
		StatusMessage: statusMessage,
		Data:          data,
	})
}

// This will do json decode
// make sure [requestBody] parameter receive value using a pointer
// for example &myRequestParam
func RequestBody(writer http.ResponseWriter, request *http.Request, requestBody any) bool {
	checkReq := reflect.ValueOf(requestBody)
	if checkReq.Kind() != reflect.Ptr {
		fmt.Println("Request Body must be pointer")
		ResponseBody(writer, http.StatusBadRequest, "Bad Request", "Invalid body format")
		return false
	}
	err := json.NewDecoder(request.Body).Decode(requestBody)

	if err != nil {
		ResponseBody(writer, http.StatusBadRequest, "Bad Request", "Invalid JSON")
		return false
	}

	return true
}
