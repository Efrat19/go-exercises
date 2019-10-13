package main

import(
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func main()  {

	defaultUrl := "amqp://guest:guest@localhost:5672"
	url := os.Getenv("AMQP_URL")

	if url == "" {
		fmt.Println("missing env AMQP_URL. defaulting to "+ defaultUrl)
		url = defaultUrl
	}

	connection, err := amqp.Dial(url)

	if err != nil {
		panic("failed to connect to rabbitmq server:\n" + err.Error())
	}

	channel, err := connection.Channel()

	if err != nil {
		panic("channel couldnt be created")
	}

	err = channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)

    if err != nil {
        panic(err)
	}
	
	message := amqp.Publishing{
        Body: []byte("Hello World"),
    }

    // We publish the message to the exahange we created earlier
    err = channel.Publish("events", "random-key", false, false, message)

    if err != nil {
        panic("error publishing a message to the queue:" + err.Error())
    }

    // We create a queue named Test
    _, err = channel.QueueDeclare("test", true, false, false, false, nil)

    if err != nil {
        panic("error declaring the queue: " + err.Error())
    }

    // We bind the queue to the exchange to send and receive data from the queue
    err = channel.QueueBind("test", "#", "events", false, nil)

    if err != nil {
        panic("error binding to the queue: " + err.Error())
	}
	
	// We consume data in the queue named test using the channel we created in go.
    msgs, err := channel.Consume("test", "", false, false, false, false, nil)

    if err != nil {
        panic("error consuming the queue: " + err.Error())
    }

    // We loop through the messages in the queue and print them to the console.
    // The msgs will be a go channel, not an amqp channel
    for msg := range msgs {
    //print the message to the console
        fmt.Println("message received: " + string(msg.Body))
    // Acknowledge that we have received the message so it can be removed from the queue
        msg.Ack(false)
    }

    // We close the connection after the operation has completed.
    defer connection.Close()

}