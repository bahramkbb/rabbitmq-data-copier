package main

import "log"

var producer RabbitmqProducer
var consumer RabbitmqConsumer

func main(){
	targetUser := GetEnvVariable("TARGET_USER", "admin")
	targetPassword := GetEnvVariable("TARGET_PASSWORD", "admin")
	targetHost := GetEnvVariable("TARGET_HOST", "127.0.0.1")
	targetVHost := GetEnvVariable("TARGET_V_HOST", "/live")
	targetExchangeName := GetEnvVariable("TARGET_EXCHANGE_NAME", "app1")
	targetRoutingKey := GetEnvVariable("TARGET_ROUTING_KEY", "app1.*")

	sourceUser := GetEnvVariable("SOURCE_USER", "admin")
	sourcePassword := GetEnvVariable("SOURCE_PASSWORD", "admin")
	sourceHost := GetEnvVariable("SOURCE_HOST", "127.0.0.1")
	sourceVHost := GetEnvVariable("SOURCE_V_HOST", "app2")
	sourceQueueName := GetEnvVariable("SOURCE_QUEUE_NAME", "app2.*")



	log.Print("Initializing RabbitMQ Producer...")
	//user string, password string, host string, vhost string, exchangeName string, queueName string
	producer = *NewRabbitmqProducer(targetUser, targetPassword, targetHost, targetVHost, targetExchangeName, targetRoutingKey)

	log.Print("Initializing RabbitMQ Consumer...")
	//user string, password string, host string, vhost string, queueName string
	consumer = *NewRabbitmqConsumer(sourceUser, sourcePassword, sourceHost, sourceVHost, sourceQueueName)

	consumer.Start()
}
