package shiji_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-shiji"
)

func TestGetV0IndividualProfileByID(t *testing.T) {
	client := client()

	req := client.NewGetV0IndividualProfileByIDRequest()
	req.Headers().PropertyID = "f142d311-7fd8-46c8-afbe-6fd50c03244d"
	req.PathParams().ID = "066c2ecd-8ea5-4be2-a601-a8f894a66c21"
	req.QueryParams().Extend = shiji.CommaSeparatedQueryParam{"LocalExternalAccountReceivable"}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
