package push

type Interface interface {
	Publish(topic string, data interface{}) error
}

type Config struct {
	Type  string
	Host  string
	Port  int
	Param map[string]interface{}
}

func New(cfg *Config) Interface {
	switch cfg.Type {
	case "rabbitmq":
		return NewRabbitMQ()
	default:
		return NewRabbitMQ()
	}
}
