package main

import (
    "ApiTest/api"
    "ApiTest/api/routers/v1/users"
    "ApiTest/pkg/logs"
)

const appName = "ApiTest"
const version = "1.0.0"

func main() {
    logger := logs.NewHttpLogger(appName, "debug")
    logger.Info("Start Application!")

    server := api.NewServer(appName, version, "3333", logger)
    server.Start()
    server.AddRouter(users.AddRouters)

    // Blocking operation
    server.Stop()
}
