package shiji_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetUsersMeUnits(t *testing.T) {
	client := client()

	req := client.NewGetUsersMeUnitsRequest()

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestGetUsersMeUnitsAll(t *testing.T) {
	client := client()

	req := client.NewGetUsersMeUnitsRequest()

	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
