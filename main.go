package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	igc "github.com/marni/goigc"
)

type igcTrack struct {
	H_date       string  //"H_date": <date from File Header, H-record>,
	Pilot        string  //"pilot": <pilot>,
	Glider       string  //"glider": <glider>,
	Glider_id    string  //"glider_id": <glider_id>,
	Track_length float64 //"track_length": <calculated total track length>
}

type API struct {
	Uptime  string //"uptime": <uptime>
	Info    string //"info": "Service for IGC tracks."
	Version string //"version": "v1"

}
type igcFile struct {
	Url string //a valid igc URL
}

type igcDB struct {
	igcs map[string]igcFile
}

func (db *igcDB) add(igc igcFile, id string) {
	for _, file := range db.igcs {
		if igc == file {
			return
		}
	}
	db.igcs[id] = igc
}

func (db igcDB) Count() int {
	return len(db.igcs)
}

func (db igcDB) Get(idWanted string) igcFile {
	for id, file := range db.igcs {
		if idWanted == id {
			return file
		}
	}
	return igcFile{}
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
			fmt.Fprintf(w, "URL : %s\n", igc.Url)
			Idstr := "id"
			strValue := fmt.Sprintf("%d", idCount)
			newId := Idstr + strValue
			ids = append(ids, newId)
			idCount += 1
			db.add(igc, newId)
			json.NewEncoder(w).Encode(newId)
		}
	case "GET":
		{
			//GET case
			http.Header.Add(w.Header(), "content-type", "application/json")
			parts := strings.Split(r.URL.Path, "/")
			fmt.Fprintf(w, "longueur : %d\n", len(parts))
			fmt.Fprintln(w, parts)
			if len(parts) != 5 {
				//deal with errors
				fmt.Fprintln(w, "wrong numbers of parameters")
				return
			}
			if parts[4] == "" {
				//deal with the array
				fmt.Fprintln(w, "case array")
				json.NewEncoder(w).Encode(ids)

			}
			if parts[4] != "" {
				fmt.Fprintln(w, "case id")
				//deal with the id
				var igcWanted igcFile
				id := parts[4]
				igcWanted = db.Get(id)
				//json.NewEncoder(w).Encode(igcWanted)

				//then encode the igcFile
				url := igcWanted.Url
				track, err := igc.ParseLocation(url)
				if err != nil {
					//fmt.Errorf("Problem reading the track", err)
				}
				igcT := igcTrack{}
				igcT.Glider = track.GliderType
				igcT.Glider_id = track.GliderID
				igcT.Pilot = track.Pilot
				igcT.Track_length = track.Task.Distance()
				igcT.H_date = track.Date.String()
				json.NewEncoder(w).Encode(igcT)
			}
		}
	default:

		http.Error(w, "not implemented yet", http.StatusNotImplemented)

	}
}

var db igcDB
var ids []string
var idCount int

func main() {
	db = igcDB{}
	db.igcs = map[string]igcFile{}
	idCount = 0
	ids = nil
	port := os.Getenv("PORT")
	http.HandleFunc("/igcinfo/api", getApi)
	http.HandleFunc("/igcinfo/api/igc/", igcHandler)
	http.ListenAndServe(":"+port, nil)
}
