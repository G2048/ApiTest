package logs

import (
    "log/slog"
    "time"

    "ApiTest/pkg/config"
    "github.com/go-chi/httplog/v2"
)

var levelMap = map[string]slog.Level{
    "debug": slog.LevelDebug,
    "info":  slog.LevelInfo,
    "warn":  slog.LevelWarn,
    "error": slog.LevelError,
}

func NewHttpLogger(settings config.LogSettings) *httplog.Logger {
    levelLog, ok := levelMap[settings.LogLevel]
    if !ok {
        levelLog = slog.LevelDebug
    }
    return httplog.NewLogger(settings.AppName,
        httplog.Options{
            LogLevel: levelLog,
            JSON:     true,
            // Concise:  true,
            // RequestHeaders:   true,
            // ResponseHeaders:  true,
            MessageFieldName: "message",
            LevelFieldName:   "logLevel",
            TimeFieldName:    "time",
            SourceFieldName:  "source",
            TimeFieldFormat:  time.RFC3339,
            Tags: map[string]string{
                "version": settings.AppVersion,
                "env":     "dev",
            },
            QuietDownRoutes: []string{
                "/",
                "/ping",
            },
            QuietDownPeriod: 10 * time.Second,
        })
}
