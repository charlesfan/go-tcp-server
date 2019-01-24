package config

import "os"

type Config struct {
	Pixabay *Pixabay
}

type Pixabay struct {
	Key string
	URL string
}

type Service struct {
	Config *Config
}

func (s *Service) Init() bool {
	s.Config.Pixabay = &Pixabay{
		Key: os.Getenv("PKey"),
		URL: "https://pixabay.com/api/",
	}

	return true
}

func NewConfig() *Service {
	return &Service{
		Config: &Config{},
	}
}
