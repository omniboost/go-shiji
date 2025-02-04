package shiji_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetIndividualProfiles(t *testing.T) {
	client := client()

	req := client.NewGetIndividualProfilesRequest()
	req.Headers().PropertyID = "c6a95ecb-be35-4619-875b-cae5d0df5fd2"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestGetIndividualProfilesAll(t *testing.T) {
	client := client()

	req := client.NewGetIndividualProfilesRequest()
	req.Headers().PropertyID = "806c52a0-753d-4a03-a4ad-0d5b577a11dc"

	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
