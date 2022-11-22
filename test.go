package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetEntries(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/Verify", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(VerifiLogin)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	//expected := `[{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb@g.com","phone_number":"0987654321"},{"id":2,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"},{"id":6,"first_name":"FirstNameSample","last_name":"LastNameSample","email_address":"lr@gmail.com","phone_number":"1111111111"}]`
	expected := `{"Success":true,"ErrorID":2,"Message":"Backend connected Successfully"}`

	//bodyrespon := rr.Body.String()
	//
	//fmt.Println(bodyrespon)
	//fmt.Println(expected)
	//fmt.Println(bodyrespon)
	//fmt.Println(expected)
	//
	//fmt.Printf("Size of i is: %d", len(bodyrespon))
	//fmt.Printf("Size of i is: %d", len(expected))

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		//if expected != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)

	}

}
