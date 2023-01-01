package config

func NewDefaultLogConfig() LogConfig {
	return LogConfig{Level: 1, Path: "./dns_updater.log", MaxSize: 10, MaxBackups: 5, MaxAge: 7, Compress: true}
}
