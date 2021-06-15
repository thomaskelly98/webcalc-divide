package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_divide(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/?x=12&y=6", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler(res, req)

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	exp := `{"error": 0, "x": 12, "y": 6, "answer": 2}`
	act := res.Body.String()
	require.JSONEq(t, exp, act)
}

func Test_divide_no_x(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/?y=6", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler(res, req)

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	exp := `{"error": "Url Param x is missing"}`
	act := res.Body.String()
	require.JSONEq(t, exp, act)
}

func Test_divide_x_no_val(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/?x=&y=6", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler(res, req)

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	exp := `{"error": "Url Param x is missing"}`
	act := res.Body.String()
	require.JSONEq(t, exp, act)
}

func Test_divide_no_y(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/?x=12", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler(res, req)

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	exp := `{"error": "Url Param y is missing"}`
	act := res.Body.String()
	require.JSONEq(t, exp, act)
}

func Test_divide_y_no_val(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/?x=12&y=", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler(res, req)

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	exp := `{"error": "Url Param y is missing"}`
	act := res.Body.String()
	require.JSONEq(t, exp, act)
}

func Test_divide_x_not_int(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/?x=the&y=6", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler(res, req)

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	exp := `{"error": "Url Param x is not an integer"}`
	act := res.Body.String()
	require.JSONEq(t, exp, act)
}

func Test_divide_y_not_int(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/?x=12&y=the", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler(res, req)

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	exp := `{"error": "Url Param y is not an integer"}`
	act := res.Body.String()
	require.JSONEq(t, exp, act)
}

func Test_divide_x_is_0(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/?x=0&y=6", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler(res, req)

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	exp := `{"error": "Url Param x cannot be 0"}`
	act := res.Body.String()
	require.JSONEq(t, exp, act)
}

func Test_divide_y_is_0(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/?x=12&y=0", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler(res, req)

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	exp := `{"error": "Url Param y cannot be 0"}`
	act := res.Body.String()
	require.JSONEq(t, exp, act)
}

func Test_divide_no_params(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler(res, req)

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	exp := `{"error": "Url Param x is missing"}`
	act := res.Body.String()
	require.JSONEq(t, exp, act)
}

func Test_divide_params_empty(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/?x=&y=", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler(res, req)

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	exp := `{"error": "Url Param x is missing"}`
	act := res.Body.String()
	require.JSONEq(t, exp, act)
}
