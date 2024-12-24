package shiji_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetProfileDashboard(t *testing.T) {
	client := client()

	req := client.NewGetProfileDashboardRequest()
	req.Headers().PropertyID = "806c52a0-753d-4a03-a4ad-0d5b577a11dc"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestGetProfileDashboardAll(t *testing.T) {
	client := client()

	req := client.NewGetProfileDashboardRequest()
	req.Headers().PropertyID = "806c52a0-753d-4a03-a4ad-0d5b577a11dc"

	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
