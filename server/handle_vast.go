package server

import (
	"net/http"

	"github.com/nayan9229/ad_prox_dsp/utils"
	"github.com/nayan9229/ad_prox_dsp/database"
)

func (s *Server) vast(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	camp := database.GetCampaigns()
	admVast := utils.GenerateADM(camp.GetRandomCampaign())
	return admVast, nil
}
