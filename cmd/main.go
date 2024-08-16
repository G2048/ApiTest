package main

import (
    "ApiTest/api"
    "ApiTest/pkg/logs"
)

const appName = "ApiTest"
const version = "1.0.0"

func main() {
    logger := logs.NewHttpLogger(appName, "debug")
    logger.Info("Start Application!")

    server := api.NewServer(appName, version, "3333", logger)
    server.Start()
    // Blocking operation
    server.Stop()
}
