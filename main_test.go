package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Make sure nano node is up

func TestGeoLocation(t *testing.T) {

	ts := httptest.NewServer(initServer())
	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/getGeoLocations", ts.URL), nil)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

}

func TestGetPeers(t *testing.T) {

	ts := httptest.NewServer(initServer())
	defer ts.Close()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/getPeers", ts.URL), nil)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

}

func TestGetBalance(t *testing.T) {

	ts := httptest.NewServer(initServer())
	defer ts.Close()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/getBalance/xrb_1111111111111111111111111111111111111111111111111111hifc8npp", ts.URL), nil)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

}
