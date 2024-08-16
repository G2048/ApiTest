package config

type AppSettings struct {
    AppName    string
    AppVersion string
}
type LogSettings struct {
    AppSettings
    LogLevel string
}

type ServerSettings struct {
    LogSettings
    Port string
}
