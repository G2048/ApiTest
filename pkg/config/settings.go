package config

import _ "github.com/danielgtaylor/huma/v2/formats/cbor"
import _ "github.com/fxamacker/cbor/v2"

type AppSettings struct {
    AppName    string `short:"n" help:"Application name" default:"ApiTest"`
    AppVersion string `short:"v" help:"Application version" default:"1.0.0"`
}
type LogSettings struct {
    AppSettings
    LogLevel string `short:"l" help:"Log level" default:"debug"`
}

type ServerSettings struct {
    LogSettings
    Port string `short:"p" help:"Server port" default:"3333"`
}
