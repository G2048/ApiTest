package logs

import (
    "log/slog"
    "time"

    "github.com/go-chi/httplog/v2"
)

var levelMap = map[string]slog.Level{
    "debug": slog.LevelDebug,
    "info":  slog.LevelInfo,
    "warn":  slog.LevelWarn,
    "error": slog.LevelError,
}

func NewHttpLogger(serviceName string, level string) *httplog.Logger {
    levelLog, ok := levelMap[level]
    if !ok {
        levelLog = slog.LevelDebug
    }
    return httplog.NewLogger(serviceName,
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
                "version": "v1.0-81aa4244d9fc8076a",
                "env":     "dev",
            },
            QuietDownRoutes: []string{
                "/",
                "/ping",
            },
            QuietDownPeriod: 10 * time.Second,
        })
}
