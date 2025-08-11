package sqlite

type Config struct {
	Path         string `toml:"path"`
	MaxIdleConns int    `toml:"maxIdleConns"`
	MaxOpenConns int    `toml:"maxOpenConns"`
	EnableWAL    bool   `toml:"enableWAL"`
	BusyTimeout  int    `toml:"busyTimeout"` // 毫秒
}