package models

import (
	"errors"
	"math/rand"
	"time"
)

type Campaign struct {
	ID    int     `json:"id,omitempty"`
	Floor float64 `json:"floor"`
	Ad    struct {
		ID       string `json:"id,omitempty"`
		Title    string `json:"title,omitempty"`
		Duration int    `json:"duration,omitempty"`
		Crid     int    `json:"crid,omitempty"`
	} `json:"ad,omitempty"`
	Creative struct {
		ID        string `json:"id,omitempty"`
		MediaFile string `json:"mediaFile,omitempty"`
		W         int    `json:"w,omitempty"`
		H         int    `json:"h,omitempty"`
	} `json:"creative,omitempty"`
}

type Campaigns struct {
	Campaigns []Campaign
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (c *Campaigns) GetCampaignById(ID int) (*Campaign, error) {
	for _, cam := range c.Campaigns {
		if cam.ID == ID {
			return &cam, nil
		}
	}
	return &Campaign{}, errors.New("no campaign found")
}

func (c *Campaigns) GetRandomCampaign() *Campaign {
	// Generate a random index within the range of the slice length
	randomIndex := rand.Intn(len(c.Campaigns))
	// Retrieve the random element from the slice
	return &c.Campaigns[randomIndex]
}
