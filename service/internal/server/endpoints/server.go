package endpoints

import (
	"hahathon/internal/config"
	"hahathon/internal/server"
	actionsEndpoints "hahathon/internal/server/endpoints/v1/actions"
	pingEndpoints "hahathon/internal/server/endpoints/v1/ping"
	"hahathon/internal/server/middleware"
	client "hahathon/internal/tabs-client"

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
	repoClient := client.NewClient(s.cfg.Token, s.log)
	router := chi.NewRouter()

	router.Use(limit.Limit)
	router.Route("/v1", func(r chi.Router) {
		router.Route("/ping", func(r chi.Router) {
			r.Get("/", pingEndpoints.Pong(s.log))
		})

		r.Route("/actions", func(r chi.Router) {
			r.Post("/", actionsEndpoints.Create(repoClient.ActionTypes, s.log))
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
