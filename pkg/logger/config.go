package logger

type Config struct {
	// Level: Debug Info Warn Error Fatal
	Level            string `toml:"level"`
	Directory        string `toml:"directory"`
	MaxSize          int    `toml:"maxSize"`
	MaxAge           int    `toml:"maxAge"`
	StacktraceLevel  string `toml:"stacktraceLevel"`
	EnableStacktrace bool   `toml:"enableStacktrace"`
	EnableFileOut    bool   `toml:"enableFileOut"`
	EnableMixedSave  bool   `toml:"enableMixedSave"`
	EnableConsoleOut bool   `toml:"enableConsoleOut"`
}
