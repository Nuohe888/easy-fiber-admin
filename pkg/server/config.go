package server

type Config struct {
	Port    int    `toml:"port"`
	Domain  string `toml:"domain"`
	Storage string `toml:"storage"`
}
