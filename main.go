package main

import (
	"net/http"
	"os"
	"encoding/json"
	"strings"
)

type Response struct {
	Hostname string `json:"hostname"`
	Http_headers map[string]string `json:"http_headers"`
	Url string `json:"url"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	headers := make(map[string]string)
	for k, v := range r.Header {
		headers[k] = strings.Join(v[:], ",")
	}

	data := Response{
		os.Getenv("HOSTNAME"),
		headers,
		r.URL.String()}

	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	w.Write([]byte("\n"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
