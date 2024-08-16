package api

import (
    "context"
    "errors"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "ApiTest/pkg/config"
    "ApiTest/pkg/logs"
    "github.com/danielgtaylor/huma/v2"
    "github.com/danielgtaylor/huma/v2/adapters/humachi"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/httplog/v2"
)

type Server struct {
    Name    string
    Version string
    Port    string
    logger  *httplog.Logger
    router  *chi.Mux
    server  *http.Server
    api     huma.API
}

func NewServer(settings config.ServerSettings) *Server {
    logger := logs.NewHttpLogger(settings.LogSettings)
    logger.Info("Start Application!")

    router := chi.NewRouter()
    server := &http.Server{
        Addr:    ":" + settings.Port,
        Handler: router,
    }

    return &Server{
        Name:    settings.AppName,
        Version: settings.AppVersion,
        Port:    settings.Port,
        logger:  logger,
        router:  router,
        server:  server,
        api:     nil,
    }
}

func (s *Server) Start() {
    s.logger.Info("Start Server!")

    s.router.Use(middleware.RequestID)
    s.router.Use(middleware.RealIP)
    s.router.Use(middleware.Recoverer)
    s.router.Use(httplog.RequestLogger(s.logger))

    config := huma.DefaultConfig(s.Name, s.Version)
    adapter := humachi.NewAdapter(s.router)
    s.api = huma.NewAPI(config, adapter)

    go func() {
        err := s.server.ListenAndServe()
        if !errors.Is(err, http.ErrServerClosed) {
            s.logger.Error(fmt.Sprintf("Error %s by running server...", err))
        }
    }()
}

// Gracefully shutdown the server
func (s *Server) Stop() {
    sigc := make(chan os.Signal, 1)
    signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
    <-sigc

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err := s.server.Shutdown(ctx)
    if err != nil {
        s.logger.Error(fmt.Sprintf("Error %s by stopping server...", err))
    }
    s.logger.Info("Stop Server!")
    os.Exit(0)
}

func (s *Server) AddMiddleware(fn func(http.Handler) http.Handler) {
    s.router.Use(fn)
}
func (s *Server) AddRouter(f func(api *huma.API)) {
    f(&s.api)
}
