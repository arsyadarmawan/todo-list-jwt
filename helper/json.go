package helper

import (
	"encoding/json"
	"net/http"
	"task/model/web"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicHandling(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicHandling(err)
	return
}

func WriteToResponseErrorBody(writer http.ResponseWriter, response interface{}, r *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicHandling(err)
}

func ValidationToken(token string, validation bool) bool {
	var writer http.ResponseWriter

	if token == "" || validation == false {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "Bad Request",
			Data:   "Invalidate Token",
		}
		WriteToResponseBody(writer, webResponse)
		return true
	}
	return false
}
