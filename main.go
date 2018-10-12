package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type API struct {
	Uptime  string //"uptime": <uptime>
	Info    string //"info": "Service for IGC tracks."
	Version string //"version": "v1"

}
type igcFile struct {
	Url string //a valid igc URL
}

type igcDB struct {
	igcs map[string]int
}

func (db igcDB) add(igc igcFile, idCount int) {
	//db.igcs = append(db.igcs, igc)
	for url, _ := range db.igcs {
		if url == igc.Url {
			return
		}
		db.igcs[igc.Url] = idCount
	}
}

func getApi(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	//io.WriteString(w, "Api information :\n")
	api := &API{}
	api.Uptime = "uptime"
	api.Info = "Service for IGC tracks."
	api.Version = "version : v1"
	//fmt.Fprintf(w, "%s\n%s\n%s", api.Uptime, api.Info, api.Version)
	json.NewEncoder(w).Encode(api)

}

func igcHandler(w http.ResponseWriter, r *http.Request) {
	db := &igcDB{}
	idCount := 0
	//var ids []int
	switch r.Method {
	case "POST":
		{
			if r.Body == nil {
				http.Error(w, "no JSON body", http.StatusBadRequest)
				return
			}
			var igc igcFile
			err := json.NewDecoder(r.Body).Decode(&igc)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			//TODO check correct igc URL
			db.add(igc, idCount)
			idCount += 1
			fmt.Fprintf(w, "map : %v", db.igcs)

			/*
				s := "http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc"
				track, err := igc.ParseLocation(s)
				if err != nil {
					fmt.Errorf("Problem reading the track", err)
				}

				fmt.Printf("Pilot: %s, gliderType: %s, date: %s",
					track.Pilot, track.GliderType, track.Date.String())*/
		}
	case "GET":
		{
			http.Header.Add(w.Header(), "content-type", "application/json")
			//GET case
		}
	default:
		http.Error(w, "not implemented yet", http.StatusNotImplemented)

	}
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/igcinfo/api", getApi)
	http.HandleFunc("/igcinfo/api/igc", igcHandler)
	http.ListenAndServe(":"+port, nil)
}
