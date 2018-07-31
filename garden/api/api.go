package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	logg "../logger"
)

// Data name, time and value
type Data struct {
	Name  string    `json:"name"`
	Time  time.Time `json:"time"`
	Value int       `json:"value"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// TODO pars request
	// water plants or get json, get last hour, get last day
	// eg 192.168.1.33:3000/solenoids/green/off
	// if on, create backup timeout
	// fmt.Fprintf(w, "Hi %s", r.URL.Path[1:])

	json := getData()
	fmt.Fprintf(w, "%s", json)
	fmt.Println("json sent")
}

func getData() []byte {

	dataJSON := []Data{}

	data := logg.GetData()

	//amo-te!!!

	for i := 0; i < len(data); i++ {
		newData := Data{
			Name:  data[i].Name,
			Time:  data[i].Time,
			Value: data[i].Value,
		}
		dataJSON = append(dataJSON, newData)
	}

	// sending data through json
	b, _ := json.Marshal(dataJSON)
	fmt.Println("JSON")
	fmt.Printf("%+v\n", string(b))

	return b
}

func Start() {
	fmt.Println("Start...")
	http.HandleFunc("/", handler)
	fmt.Println("... middle")
	http.ListenAndServe(":8080", nil)

	// req, err := http.NewRequest(http.MethodGet, "localhost:8080/celina", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// print(req)
}
