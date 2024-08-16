package main

import (
    "ApiTest/api"
    "ApiTest/api/routers/v1/users"
    "ApiTest/pkg/config"
)

const appName = "ApiTest"
const version = "1.0.0"

func main() {
    appSettings := config.AppSettings{appName, version}
    logSettings := config.LogSettings{appSettings, "debug"}
    settings := config.ServerSettings{logSettings, "3333"}

    server := api.NewServer(settings)
    server.Start()
    server.AddRouter(users.AddRouters)

    // Blocking operation
    server.Stop()
}
