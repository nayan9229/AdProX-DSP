package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

func (s *Server) routes() chi.Router {
	r := chi.NewRouter()

	logs := func(r *http.Request, status, size int, duration time.Duration) {
		basicRequestLog(r, status, size, duration).Msg("")
	}
	r.Use(hlog.RequestIDHandler("req_id", "Request-Id"))
	r.Use(hlog.AccessHandler(logs))
	r.Use(hlog.RemoteAddrHandler("ip"))
	r.Use(hlog.UserAgentHandler("user_agent"))
	r.Use(hlog.RefererHandler("referer"))

	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/", Health)
	r.Get("/vast", XmlHandler(s.vast))
	r.Post("/openrtb", SimpleHandler(s.openrtb))
	r.Get("/vmap", XmlHandler(s.vast))
	r.Get("/vpaid", vpaid)
	r.Get("/aws_vast", XmlHandler(s.vast))
	r.Get("/tracking", SimpleHandler(s.tracking))
	return r
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// Basic HTTP request logging.
func basicRequestLog(r *http.Request, status, size int, duration time.Duration) *zerolog.Event {
	return hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Str("url", r.URL.String()).
		Int("status", status).
		Int("size", size).
		Dur("duration", duration)
}
