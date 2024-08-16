package main

import (
    "ApiTest/api"
    "ApiTest/api/routers/v1/users"
    "ApiTest/pkg/config"
    "github.com/danielgtaylor/huma/v2/humacli"
)

func main() {
    cli := humacli.New(func(hooks humacli.Hooks, options *config.ServerSettings) {
        server := api.NewServer(*options)
        server.Start()
        server.AddRouter(users.AddRouters)

        // Blocking operation
        server.Stop()
    })
    cli.Run()
}
