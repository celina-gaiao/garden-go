package api

import (
	"fmt"
	"testing"

	logg "../logger"
)

func TestFoo(t *testing.T) {
	//////////////////////////////////
	// test getData() with 1 sample
	//////////////////////////////////

	fmt.Println("//////////////////////////////////")
	fmt.Println("// test with 1 sample")

	fmt.Println("Setting up data...")
	logg.SaveData("Light", 14)

	res := getData()
	// TODO parse
	resName := "Light"
	resValue := 14

	expectedName := "Light"
	expectedValue := 14
	if resName != expectedName {
		t.Errorf("getData error in name: %s", string(res))
	}
	if resValue != expectedValue {
		t.Errorf("getData error in value: %s", string(res))
	}

	// clear data
	fmt.Println("Cleaning up data...")
	logg.DeletData()

	//////////////////////////////////
	// test getData() with multiple samples
	//////////////////////////////////

	fmt.Println("//////////////////////////////////")
	fmt.Println("test with multiple samples")

	fmt.Println("Setting up data...")
	logg.SaveData("Light", 15)
	logg.SaveData("Temperature", 25)

	res = getData()
	// TODO parse
	// TODO array
	resName = "Light"
	resValue = 15

	expectedName = "Light"
	expectedValue = 15
	if resName != expectedName {
		t.Errorf("getData error in name: %s", string(res))
	}
	if resValue != expectedValue {
		t.Errorf("getData error in value: %s", string(res))
	}

	// clear data
	fmt.Println("Cleaning up data...")
	logg.DeletData()
}

// func TestHandler(t *testing.T) {
// 	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
// 	// pass 'nil' as the third parameter.
// 	req, err := http.NewRequest("GET", "/", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
// 	rr := httptest.NewRecorder()
// 	handler1 := http.HandlerFunc(handler)

// 	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
// 	// directly and pass in our Request and ResponseRecorder.
// 	handler1.ServeHTTP(rr, req)

// 	// Check the status code is what we expect.
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	// Check the response body is what we expect.
// 	expected := `[]`
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }

// // layout := "2006-01-02 15:04:05.000000000 +0000 UTC"
// // str := "2017-09-25 20:20:20.0000000 -0700 PDT"
// // t, _ := time.Parse(layout, str)
