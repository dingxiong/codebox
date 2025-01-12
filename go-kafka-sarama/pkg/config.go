package pkg

type Config struct {
	BootstrapBrokers []string // host_name:port
}

var AppConfig *Config

func init() {
	AppConfig = &Config{}
}
