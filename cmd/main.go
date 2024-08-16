package main

import (
    "context"
    "time"

    "ApiTest/api"
    "ApiTest/pkg/logs"
)

const appName = "ApiTest"
const version = "1.0.0"

func main() {
    logger := logs.NewHttpLogger(appName, "debug")
    logger.Info("Start Application!")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    server := api.NewServer(appName, version, "3333", logger)
    defer server.Stop(ctx)
    server.Start()
}
