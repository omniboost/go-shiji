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
	req.Headers().PropertyID = "c6a95ecb-be35-4619-875b-cae5d0df5fd2"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
