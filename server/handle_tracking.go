package server

import "net/http"

func (s *Server) tracking(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	type res struct{
		Status string `json:"status"`
		Message string `json:"message"`
	}
	return res{
		Status: "success",
		Message: "tracking",
	}, nil
}
