package database

import (
	"encoding/json"
	"fmt"
	"sync"

	datajson "github.com/nayan9229/ad_prox_dsp/database/data_json"
	"github.com/nayan9229/ad_prox_dsp/models"
	"github.com/rs/zerolog/log"
)

var (
	db        []*models.Customer
	campaigns models.Campaigns
	mu        sync.Mutex
)
func init() {
	Connect()
}
// Connect with database
func Connect() {
	db = make([]*models.Customer, 0)
	cams := []models.Campaign{}
	err := json.Unmarshal(datajson.CampaignsData(), &cams)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to unmarshal campaigns data")
	}
	campaigns.Campaigns = cams
	fmt.Println("Connected with Database")
}

func Insert(user *models.Customer) {
	mu.Lock()
	db = append(db, user)
	mu.Unlock()
}

func Get() []*models.Customer {
	return db
}

func GetCampaigns() models.Campaigns {
	return campaigns
}
