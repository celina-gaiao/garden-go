package logger

import (
	"log"
	"time"
)

// NData number of entries of data to save per day
// sensores: one time per minute, 60min, 24h, 4 sensors
// actuators: turn on, turn off, x times per day, 3 actuators
// 60*24*4 + 2*4*3 = 5 784 < 6 000
const NData = 6000

var values []int
var data []Data
var allData []Data //TODO limit vector

// Data name, time and value
type Data struct {
	Name  string
	Time  time.Time
	Value int
}

// SaveData creats a Data struct with time stamp and add that to all data
func SaveData(name string, value int) {
	// create Data struct
	newData := Data{
		Name:  name,
		Time:  time.Now(),
		Value: value,
	}

	// add new data to all data
	allData = append(allData, newData)

	log.Println("- value added:", value)
}

// GetData para dar os dados a api logo em json
func GetData() []Data {
	return allData
}

// GetData para dar os dados a api logo em json
func DeletData() {
	allData = []Data{}
}

// mais tarde, ler e escrever
