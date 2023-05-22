package rabbitMq

import (
	"bookStoreUser/errors"

	"github.com/streadway/amqp"
)

func ConnectToRabbitMQ() (*amqp.Connection, *errors.RestError) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, errors.NewBadRequestError("Can not connect to rabbitMq")
	}
	return conn, nil
}

func DeclareQueue(conn *amqp.Connection, queueName string) (*amqp.Channel, *errors.RestError) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, errors.NewBadRequestError("rabbitMq channel error...")
	}
	_, err = channel.QueueDeclare(
		queueName, // Queue name
		false,     // Durable
		false,     // Delete when unused
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		return nil, errors.NewBadRequestError("rabbitMq queue declare error...")
	}
	return channel, nil
}

func PublishMessage(channel *amqp.Channel, queueName string, message string) *errors.RestError {
	err := channel.Publish(
		"",        // Exchange
		queueName, // Routing key
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	return errors.NewBadRequestError(err.Error())
}

func ConsumeMessages(channel *amqp.Channel, queueName string) (<-chan amqp.Delivery, *errors.RestError) {
	msgs, err := channel.Consume(
		queueName, // Queue name
		"",        // Consumer
		true,      // Auto-acknowledge
		false,     // Exclusive
		false,     // No-local
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}
	return msgs, nil
}
