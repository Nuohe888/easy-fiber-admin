package server

type Config struct {
	Port    int    `toml:"port"`
	Domain  string `toml:"domain"`
	Storage string `toml:"storage"`
	Env     int    `toml:"env"` //0开发 1生产 (开发环境可以绕过一些不必要的验证)
}
