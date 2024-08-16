package config

import _ "github.com/danielgtaylor/huma/v2/formats/cbor"
import _ "github.com/fxamacker/cbor/v2"

type AppSettings struct {
    AppName    string `help:"Application name" default:"ApiTest"`
    AppVersion string `help:"Application version" default:"1.0.0"`
}
type LogSettings struct {
    AppSettings
    LogLevel string `help:"Log level" default:"debug"`
}

type ServerSettings struct {
    LogSettings
    Port string `help:"Server port" default:"3333"`
}
