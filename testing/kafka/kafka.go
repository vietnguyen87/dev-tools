package kafka

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"vietnt.me/core/sen-kit/pubsub"
	kafka_go "vietnt.me/core/sen-kit/pubsub/kafka-go"
	productBase "vietnt.me/protobuf/internal-apis-go/product/base"
)

type Client struct {
	PublishClients map[string]pubsub.Publisher
}

type ClientKafka interface {
	PubES7AddDataV2(changedProduct *productBase.ChangedProductV2, topic string) error
}

func NewClientKafka(clients map[string]pubsub.Publisher) ClientKafka {
	return &Client{
		PublishClients: clients,
	}
}

//var kafkaPublishClients map[string]pubsub.Publisher

func (c *Client) kafkaPublish(topic string) pubsub.Publisher {
	pub, ok := c.PublishClients[topic]
	if !ok || pub == nil {
		return nil
	}

	return pub
}

func InitKafkaPublisher(brokerHosts string, topics ...string) (pubs map[string]pubsub.Publisher) {
	pubs = make(map[string]pubsub.Publisher)
	for _, topic := range topics {
		kafkaConfig := kafka_go.PublisherConfig{
			Topic:   topic,
			Logger:  nil,
			Brokers: strings.Split(brokerHosts, ","),
		}
		pub, err := kafka_go.NewPublisher(&kafkaConfig)
		if err != nil {
			fmt.Printf("connect to %s kafka topic fail, err: %s", topic, err)
			os.Exit(1)
		}

		pubs[topic] = pub
	}

	fmt.Printf("connect to kafka topics success: %s", topics)

	return pubs
}

func (c *Client) PubES7AddDataV2(changedProduct *productBase.ChangedProductV2, topic string) error {
	data, _ := json.Marshal(changedProduct)

	pubClient := c.kafkaPublish(topic)
	err := pubClient.PublishRaw(
		context.Background(),
		fmt.Sprintf("%d", changedProduct.ProductId),
		data,
	)
	if err != nil {
		fmt.Printf("publish fail to kafka with topic %s, %v", topic, map[string]interface{}{
			"err":  err,
			"data": string(data),
		})
	} else {
		//log success
		hashData := fmt.Sprintf("%x", sha1.Sum([]byte(data)))
		fmt.Printf("publish success to kafka with topic %s, %v", topic, map[string]interface{}{
			"hash": hashData,
			"data": string(data),
			"func": "PubES7AddDataV2",
		})
	}

	return err
}
