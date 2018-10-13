package main

import (
	"encoding/json"
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
	igcs map[string]igcFile
}

func (db *igcDB) add(igc igcFile, Id string) {
	db.igcs[Id] = igc
}

func (db *igcDB) Count() int {
	return len(db.igcs)
}

/*func (db igcDB) Get(i int) igcFile {
	if i < 0 || i >= len(db.igcs) {
		return igcFile{}
	}
	return db.igcs[i]
}*/

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
	/*idCount := 0
	var ids []string*/
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
			newId := "id"
			idCount += 1
			ids = append(ids, newId)
			db.add(igc, newId)
			json.NewEncoder(w).Encode(igc)
		}
	case "GET":
		{
			//GET case
			//http.Header.Add(w.Header(), "content-type", "application/json")
			//json.NewEncoder(w).Encode(ids)
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
