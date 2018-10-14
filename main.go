package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
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
	igcs map[string]igcFile
}

func (db igcDB) add(igc igcFile, id string) {
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
			Idstr := "id"
			strValue := fmt.Sprintf("%d", idCount)
			newId := Idstr + strValue
			ids = append(ids, newId)
			idCount += 1
			//db.add(igc, newId)
			json.NewEncoder(w).Encode(newId)
		}
	case "GET":
		{
			//GET case
			http.Header.Add(w.Header(), "content-type", "application/json")
			parts := strings.Split(r.URL.Path, "/")
			fmt.Fprintf(w, "longueur : %d\n", len(parts))
			fmt.Fprintln(w, parts)
			if len(parts) > 5 || len(parts) < 3 {
				//deal with errors
				fmt.Fprintln(w, "wrong numbers of parameters")
				return
			}
			if len(parts) == 5 {
				fmt.Fprintln(w, "case 5")
				//deal with the id
				/*var igcWanted igcFile
				id := parts[4]
				igcWanted = db.Get(id)
				json.NewEncoder(w).Encode(igcWanted)*/

			}
			//fmt.Fprintln(w, parts)
			if len(parts) == 4 {
				//deal with the array
				json.NewEncoder(w).Encode(ids)
			}
		}
	default:
		http.Error(w, "not implemented yet", http.StatusNotImplemented)

	}
}

/*func idHandler(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")

}*/

var db igcDB
var ids []string
var idCount int

func main() {
	db = igcDB{}
	idCount = 0
	ids = nil
	port := os.Getenv("PORT")
	http.HandleFunc("/igcinfo/api", getApi)
	http.HandleFunc("/igcinfo/api/igc/", igcHandler)

	http.ListenAndServe(":"+port, nil)
}
