package controllers

import (
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

func ProxyArcApi(w http.ResponseWriter, r *http.Request) {

	var username string = os.Getenv("ARC_API_USERNAME")
	var passwd string = os.Getenv("ARC_API_PASSWORD")
	client := &http.Client{}
	req, err := http.NewRequest("GET", os.Getenv("ARC_API_URL")+r.RequestURI, nil)
	req.SetBasicAuth(username, passwd)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(resp.StatusCode)
	w.Write(StreamToByte(resp.Body))

}
