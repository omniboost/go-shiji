package shiji_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-shiji"
)

func TestGetV0TravelAgentProfileByID(t *testing.T) {
	client := client()

	req := client.NewGetV0TravelAgentProfileByIDRequest()
	req.Headers().PropertyID = "f142d311-7fd8-46c8-afbe-6fd50c03244d"
	req.PathParams().ID = "d65615a6-1120-4f5e-876f-c7c5c2340c91"
	req.QueryParams().Extend = shiji.CommaSeparatedQueryParam{"LocalExternalAccountReceivable"}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
