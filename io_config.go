package godruid

type IOConfig interface{}

type ioConfigKafka struct {
	Type                      string              `json:"type"`
	ConsumerProperties        *ConsumerProperties `json:"consumerProperties"`
	Topic                     string              `json:"topic"`
	LatestMessageRejectPeriod string              `json:"lateMessageRejectionPeriod"`
}

func IOConfigKafka(topic, latestMessageRejectPeriod, bootstrapServers, saslMechanism, securityProtocal, sasLJAASConfig string) IOConfig {
	return &ioConfigKafka{
		Type:                      "kafka",
		Topic:                     topic,
		LatestMessageRejectPeriod: latestMessageRejectPeriod,
		ConsumerProperties: &ConsumerProperties{
			BootstrapServers: bootstrapServers,
			SASLMechanism:    saslMechanism,
			SecurityProtocal: securityProtocal,
			SASLJAASConfig:   sasLJAASConfig,
		},
	}
}

type ConsumerProperties struct {
	BootstrapServers string `json:"bootstrap.servers"`
	SASLMechanism    string `json:"sasl.mechanism"`
	SecurityProtocal string `json:"security.protocol"`
	SASLJAASConfig   string `json:"sasl.jaas.config"`
}
