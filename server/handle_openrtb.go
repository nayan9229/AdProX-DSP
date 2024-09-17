package server

import (
	"net/http"

	"github.com/nayan9229/ad_prox_dsp/database"
	"github.com/nayan9229/ad_prox_dsp/utils"
)

func (s *Server) openrtb(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	campaigns := database.GetCampaigns()
	adId := utils.AdID()
	bResp := utils.GenerateBidResponse(adId, campaigns.GetRandomCampaign())
	return bResp, nil
}
