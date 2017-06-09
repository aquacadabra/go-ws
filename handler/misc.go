package handler

import (
	"encoding/json"
	"go-ws/msg"
	"log"
	"net/http"
)

func getJsonData(r *http.Request, v interface{}) bool {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func sendResposne(
	w http.ResponseWriter,
	code int,
	outdata interface{}) {

	data, _ := json.Marshal(outdata)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(data)
}

func sendResposneMessage(
	w http.ResponseWriter,
	code int,
	message string) {

	response := &msg.RespMessage{
		Message: message,
	}
	data, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(data)
}
