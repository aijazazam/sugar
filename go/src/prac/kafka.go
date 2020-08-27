package main

import (
	"fmt"
	"time"
	"context"
	"github.com/segmentio/kafka-go"
)

const (
	_TOPIC_MOVIE_ = "TOPIC-MOVIE"
)

func main() {

	go producer()
	go consumer()

	time.Sleep( 10 * time.Second )
}

func producer() {

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   _TOPIC_MOVIE_,
		Balancer: &kafka.LeastBytes{},
	})

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)

	fmt.Println("Producer Done with err = ", err)
	w.Close()
}

func consumer() {

	r := kafka.NewReader(kafka.ReaderConfig{
    Brokers:   []string{"localhost:9092"},
    Topic:     _TOPIC_MOVIE_,
    Partition: 0,
    MinBytes:  10e3, // 10KB
    MaxBytes:  10e6, // 10MB
	})

	r.SetOffset(0)

	for {
    m, err := r.ReadMessage(context.Background())
    if err != nil {
				fmt.Println( "consumer err = ", err )
        break
    }
    fmt.Printf("consumer message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	r.Close()
}

