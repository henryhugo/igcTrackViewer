package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type API struct {
	Uptime  string //"uptime": <uptime>
	Info    string //"info": "Service for IGC tracks."
	Version string //"version": "v1"

}

func getApi(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	io.WriteString(w, "Api information :\n")
	api := &API{}
	api.Uptime = "uptime"
	api.Info = "Service for IGC tracks."
	api.Version = "version : v1"
	//fmt.Fprintf(w, "%s\n%s\n%s", api.Uptime, api.Info, api.Version)
	json.NewEncoder(w).Encode(api)

}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/igcinfo/api", getApi)
	http.ListenAndServe(":"+port, nil)
}
