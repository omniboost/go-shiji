package shiji_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetTravelAgentProfileByID(t *testing.T) {
	client := client()

	req := client.NewGetTravelAgentProfileByIDRequest()
	req.Headers().PropertyID = "806c52a0-753d-4a03-a4ad-0d5b577a11dc"
	req.PathParams().ID = "ef578132-1cb9-457f-905d-b137e7d2e886"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
