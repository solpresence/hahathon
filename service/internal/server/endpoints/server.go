package endpoints

import (
	"hahathon/internal/config"
	"hahathon/internal/server"
	"hahathon/internal/server/endpoints/v1/ping"
	"hahathon/internal/server/middleware"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

type Server struct {
	log *slog.Logger
	cfg *config.Config

	//Do not touch
	srv         *http.Server
	rateLimiter *middleware.IPRateLimiter
}

func (s *Server) CreateServer() {
	maxConn, _ := strconv.Atoi(s.cfg.Htppserver.MaxConn)
	idleTimeout, _ := strconv.Atoi(s.cfg.Htppserver.IddleTimeout)
	timeout, _ := strconv.Atoi(s.cfg.Htppserver.Timeout)
	limit := middleware.NewIPRateLimiter(maxConn, time.Minute, 10)

	router := chi.NewRouter()

	router.Use(limit.Limit)
	router.Route("/v1", func(r chi.Router) {
		router.Route("/ping", func(r chi.Router) {
			r.Get("/", ping.Pong(s.log))
		})
		r.Route("/employees", func(r chi.Router) {

		})
		r.Route("/positions", func(r chi.Router) {

		})
		r.Route("/actions", func(r chi.Router) {

		})
		r.Route("/locations", func(r chi.Router) {

		})
		r.Route("/action_types", func(r chi.Router) {

		})
	})

	srv := &http.Server{
		Addr:         s.cfg.Htppserver.Addr,
		Handler:      router,
		ReadTimeout:  time.Duration(timeout) * time.Second,
		WriteTimeout: time.Duration(timeout) * time.Second,
		IdleTimeout:  time.Duration(idleTimeout) * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			s.log.Error("failed to start server")
			return
		}
	}()

	s.srv = srv
	s.rateLimiter = limit
}

func (s *Server) Close() error {
	s.rateLimiter.Stop()
	return s.srv.Close()
}

func NewRestApi(cfg *config.Config, log *slog.Logger) server.Server {
	return &Server{
		cfg: cfg,
		log: log,
	}
}
