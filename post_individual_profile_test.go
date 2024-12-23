package shiji_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestPostIndividualProfile(t *testing.T) {
	client := client()

	req := client.NewPostIndividualProfileRequest()
	req.Headers().PropertyID = "806c52a0-753d-4a03-a4ad-0d5b577a11dc"
	body := req.NewRequestBody()
	body.Details.FirstName = "Emma"

	req.SetRequestBody(body)
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
