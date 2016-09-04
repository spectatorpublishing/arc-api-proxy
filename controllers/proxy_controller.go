package controllers

import (
	// "encoding/json"
	// "arc-api-proxy/helpers"
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
)

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

// Returns the apiUser object
func ProxyArcApi(w http.ResponseWriter, r *http.Request) {

	// return
	var username string = os.Getenv("ARC_API_USERNAME")
	var passwd string = os.Getenv("ARC_API_PASSWORD")
	client := &http.Client{}
	req, err := http.NewRequest("GET", os.Getenv("ARC_API_URL")+r.RequestURI, nil)
	req.SetBasicAuth(username, passwd)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(StreamToByte(resp.Body))
	// bodyText, err := ioutil.ReadAll(resp.Body)
	// s := string(bodyText)
	// return s

	// if err := helpers.RenderJson(w, http.StatusCreated, map[string]string{"test": "works"}); err != nil {
	// 	panic(err)
	// }

}
