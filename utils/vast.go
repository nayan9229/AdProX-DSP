package utils

import (
	"fmt"
	"time"

	"github.com/haxqer/vast"
	"github.com/nayan9229/ad_prox_dsp/models"
)

func GenerateADM(campaign *models.Campaign) *vast.VAST {
	v := vast.VAST{
		Version: "3.0",
		Ads: []vast.Ad{
			{
				ID:     campaign.Ad.ID,
				AdType: "video",
				InLine: &vast.InLine{
					AdSystem: &vast.AdSystem{
						Version: "1.0",
						Name:    campaign.Ad.Title,
					},
					Description: &vast.CDATAString{CDATA: "VAST 2.0"},
					AdTitle:     vast.PlainString{CDATA: campaign.Ad.Title},
					Errors: []vast.CDATAString{{
						CDATA: fmt.Sprintf("%s/tracking?e=error", SERVER_BASE_URL),
					}},
					Impressions: []vast.Impression{{
						ID:  campaign.Ad.ID,
						URI: fmt.Sprintf("%s/tracking?e=imp", SERVER_BASE_URL),
					}},
					Pricing: &vast.Pricing{
						Value:    fmt.Sprintf("%f", campaign.Floor),
						Currency: "USD",
					},
					Creatives: []vast.Creative{{
						ID: campaign.Ad.ID,
						Linear: &vast.Linear{
							Duration: vast.Duration(campaign.Ad.Duration * int(time.Second)),
							VideoClicks: &vast.VideoClicks{
								ClickThroughs: []vast.VideoClick{
									{URI: ""},
								},
							},
							MediaFiles: &vast.MediaFiles{
								MediaFile: []vast.MediaFile{{
									Width:    campaign.Creative.W,
									Height:   campaign.Creative.H,
									Delivery: "progressive",
									Type:     "video/mp4",
									URI:      campaign.Creative.MediaFile,
								}},
							},
							TrackingEvents: &vast.TrackingEvents{
								Tracking: []vast.Tracking{
									{
										Event: "start",
										URI:   fmt.Sprintf("%s/tracking?e=start", SERVER_BASE_URL),
									},
									{
										Event: "firstQuartile",
										URI:   fmt.Sprintf("%s/tracking?e=firstQuartile", SERVER_BASE_URL),
									},
									{
										Event: "midpoint",
										URI:   fmt.Sprintf("%s/tracking?e=midpoint", SERVER_BASE_URL),
									},
									{
										Event: "thirdQuartile",
										URI:   fmt.Sprintf("%s/tracking?e=thirdQuartile", SERVER_BASE_URL),
									},
									{
										Event: "complete",
										URI:   fmt.Sprintf("%s/tracking?e=complete", SERVER_BASE_URL),
									},
								},
							},
						},
					}},
				},
			},
		},
		Errors: []vast.CDATAString{{
			CDATA: fmt.Sprintf("%s/tracking?e=error", SERVER_BASE_URL),
		}},
	}
	return &v
}
