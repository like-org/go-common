package conf

import (
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type SrvConfig struct {
	Url  string
	Ip   string
	Port int
}
type GWConfig struct {
	Ip   string
	Port int
	Acc  SrvConfig
}
type Config struct {
	Version string
	Gw      GWConfig
}

var (
	Conf *Config
)

func New(filepath string) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("config file load failure, err: %v", err)
	}

	Conf = &Config{}
	err = toml.Unmarshal(data, Conf)
	if err != nil {
		log.Fatal("config file unmarshal failure")
	}

}
