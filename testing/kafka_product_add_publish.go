package main

import (
	"dev/testing/kafka"
	"fmt"
	"github.com/gogo/protobuf/types"
	productBase "vietnt.me/protobuf/internal-apis-go/product/base"
)

func main() {
	brokerHosts := "localhost:9092"
	productAddTopic := "es.product.added"

	kafkaPublishClients := kafka.InitKafkaPublisher(brokerHosts,
		productAddTopic,
	)

	pubClient := kafka.NewClientKafka(kafkaPublishClients)
	if pubClient == nil {
		return
	}

	productData := &productBase.ChangedProductV2{
		ProductId: 16688752,
		Fields:    &types.FieldMask{Paths: []string{}},
	}

	pubClient.PubES7AddDataV2(productData, productAddTopic)

	fmt.Println("DONE.!")
}
