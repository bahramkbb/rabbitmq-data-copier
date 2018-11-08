package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

type RabbitmqConsumer struct {
	amqpURI      string
	queueName    string
	consumerName string
}

// Consumer constructor which initializes variables required in the struct
func NewRabbitmqConsumer(user string, password string, host string, vhost string, queueName string) *RabbitmqConsumer {
	consumer := new(RabbitmqConsumer)

	consumer.consumerName = "ReplicatorConsumer"
	consumer.amqpURI = fmt.Sprintf("amqp://%s:%s@%s:5672/%s", user, password, host, vhost)
	consumer.queueName = queueName

	return consumer
}

// Start the consumer
func (bc *RabbitmqConsumer) Start() (bool, error) {
	//Connect to rabbitmq server
	conn, err := amqp.Dial(bc.amqpURI)
	log.Printf("Source URI: %s" , bc.amqpURI)
	defer conn.Close()
	if err != nil {
		log.Printf("Recipe consumer error connecting to rabbitmq: %v", err)
		return false, nil
	}

	//Open rabbitmq channel
	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		log.Printf("Recipe consumer error while opening channel on rabbitmq: %v", err)
		return false, nil
	}

	//Initialize consumer to read messages and put them into msgs channel
	msgs, err := ch.Consume(
		bc.queueName,
		bc.consumerName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Printf("error while registering consumer on rabbitmq: %v", err)
		return false, nil
	}

	//Making a forever loop so the consumer stays listening for ever
	forever := make(chan bool)

	//Consumer function
	go func() {
		for d := range msgs {

			data := amqp.Publishing{
				Headers:         d.Headers,
				ContentType:     d.ContentType,
				ContentEncoding: d.ContentEncoding,
				DeliveryMode:    d.DeliveryMode,
				Priority:        d.Priority,
				CorrelationId:   d.CorrelationId,
				ReplyTo:         d.ReplyTo,
				Expiration:      d.Expiration,
				MessageId:       d.MessageId,
				Timestamp:       d.Timestamp,
				Type:            d.Type,
				UserId:          d.UserId,
				AppId:           d.AppId,
				Body:            d.Body,
			}

			if _, err := producer.PublishMessage(data); err == nil {
				//d.Ack(false)
				log.Print("New message pushed to live")
			} else {
				log.Print(err)
				d.Reject( true)
				os.Exit(1)
			}

			time.Sleep(time.Millisecond * 50)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return true, nil
}