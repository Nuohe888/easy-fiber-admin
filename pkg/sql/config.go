package sql

type Config struct {
	User         string `toml:"user"`
	Pass         string `toml:"pass"`
	Host         string `toml:"host"`
	DbName       string `toml:"dbName"`
	Port         int    `toml:"port"`
	MaxIdleConns int    `toml:"maxIdleConns"`
	MaxOpenConns int    `toml:"maxOpenConns"`
}
