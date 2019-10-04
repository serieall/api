package models

type Config struct {
	Secret      string `env:"SERIEALL_SECRET" envDefault:"secret"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"Debug"`
	Port        int    `env:"PORT" envDefault:"8080"`
	ImageFolder string `env:"IMAGE_FOLDER" envDefault:"/home/bmayelle/perso/AveDeux/public/images"`
	ImagePath   string `env:"IMAGE_PATH" envDefault:"/images"`
	NatsHost    string `env:"NATS_HOST" envDefault:"localhost"`
}
