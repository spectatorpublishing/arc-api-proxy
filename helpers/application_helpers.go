package helpers

import (
	"encoding/json"
	// "math/rand"
	"net/http"
	// "time"
)

// //Wrapper function for RenderJson
// func RenderJsonErr(w http.ResponseWriter, statusCode int, text string) error {
// 	return RenderJson(w, statusCode, JsonErr{StatusCode: statusCode, Error: text})
// }

func RenderJson(w http.ResponseWriter, statusCode int, object interface{}) error {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(object)
	return err

}

func RenderJsonString(w http.ResponseWriter, statusCode int, jsonStr string) error {

	WriteCommonHeaders(w, statusCode)
	w.Write([]byte(jsonStr))
	return nil

}

func WriteCommonHeaders(w http.ResponseWriter, statusCode int) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	return
}
