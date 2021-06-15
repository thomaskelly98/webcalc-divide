package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// IF I NEED TO CHANGE RETURN
// type ReturnJson struct {
// 	error  bool
// 	string string
// 	answer int
// }

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// return json. Allow CORS
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	x, ok_x := r.URL.Query()["x"]
	y, ok_y := r.URL.Query()["y"]

	// check x and y exist
	if !ok_x || len(x[0]) < 1 {
		json.NewEncoder(w).Encode(map[string]string{"error": "Url Param x is missing"})
		return
	}
	if !ok_y || len(y[0]) < 1 {
		json.NewEncoder(w).Encode(map[string]string{"error": "Url Param y is missing"})
		return
	}

	val_x := x[0]
	val_y := y[0]

	// try to convert x and y to integers
	int_x, err := strconv.Atoi(val_x)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Url Param x is not an integer"})
		return
	}
	int_y, err := strconv.Atoi(val_y)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Url Param y is not an integer"})
		return
	}

	// ensure x is not 0 and y is not 0
	if int_x == 0 {
		json.NewEncoder(w).Encode(map[string]string{"error": "Url Param x cannot be 0"})
		return
	}
	if int_y == 0 {
		json.NewEncoder(w).Encode(map[string]string{"error": "Url Param y cannot be 0"})
		return
	}

	// get division answer
	div_ans := divide(int_x, int_y)

	json.NewEncoder(w).Encode(map[string]int{"error": 0, "x": int_x, "y": int_y, "answer": div_ans})
}

func divide(x int, y int) int {
	return x / y
}
