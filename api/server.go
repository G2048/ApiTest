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

    "ApiTest/api/routers/v1/users"
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
}

func NewServer(name, version, port string, logger *httplog.Logger) *Server {
    router := chi.NewRouter()
    server := &http.Server{
        Addr:    ":" + port,
        Handler: router,
    }

    return &Server{
        Name:    name,
        Version: version,
        Port:    port,
        logger:  logger,
        router:  router,
        server:  server,
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
    api := huma.NewAPI(config, adapter)

    users.AddRouters(api)
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
}

func (s *Server) AddMiddleware(fn func(http.Handler) http.Handler) {
    s.router.Use(fn)
}
