package config

func NewDefaultLogConfig() LogConfig {
	return LogConfig{Level: 1, MaxSize: 10, MaxBackups: 5, MaxAge: 7, Compress: true}
}
