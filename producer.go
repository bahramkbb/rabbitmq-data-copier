package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type RabbitmqProducer struct {
	conn         *amqp.Connection
	ch           *amqp.Channel
	exchangeName string
	routingKey 	 string
	amqpURI      string
}

func NewRabbitmqProducer(user string, password string, host string, vhost string, exchangeName string, routingKey string) *RabbitmqProducer {
	producer := new(RabbitmqProducer)

	producer.amqpURI = fmt.Sprintf("amqp://%s:%s@%s:5672/%s", user, password, host, vhost)
	producer.exchangeName = exchangeName
	producer.routingKey = routingKey

	var err error

	//Connect to rabbitmq server
	producer.conn, err = amqp.Dial(producer.amqpURI)
	if err != nil {
		log.Fatalf("producer error connecting to rabbitmq: %v", err)
	}

	//Open rabbitmq channel
	producer.ch, err = producer.conn.Channel()
	if err != nil {
		log.Fatalf("backend consumer error connecting to rabbitmq: %v", err)
	}

	log.Printf("Target URI: %s" , producer.amqpURI)

	return producer
}

func (rp *RabbitmqProducer) PublishMessage(msg amqp.Publishing) (bool, error) {

	log.Print("Publishing to : " + rp.exchangeName + " ," + rp.routingKey)

	//Publish message to queue
	if err := rp.ch.Publish(
		rp.exchangeName, // exchange
		rp.routingKey,      // routing key
		false,           // mandatory
		false,           // immediate
		msg); err != nil {
		log.Printf("Exchange Publish: %s", err)
		return false, err
	}
	return true, nil
}