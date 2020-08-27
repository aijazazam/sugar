package kafka

import (
	"fmt"
	"context"
	"github.com/segmentio/kafka-go"

	com "common"
)

var (
	SendCredential chan com.Credential
  RecvCredential chan com.Credential
)

func init() {

	SendCredential = make( chan com.Credential, 100 )
  RecvCredential = make( chan com.Credential, 100 )

	go GoProducer( SendCredential )
  go GoConsumer( 0, RecvCredential )
}

func GoProducer( cAuth <-chan com.Credential ) {

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   com.TOPIC_AUTHENTICATION,
		Balancer: &kafka.LeastBytes{},
	})

	// This automatically breaks when channel closes
	for tAuth := range cAuth {

		tRaw, tErr := tAuth.Marshal()
		if tErr != nil {
			panic( tErr )
		}

		tErr = w.WriteMessages( context.Background(),
			kafka.Message{
				Value: tRaw,
			},
		)
		if tErr != nil {
			panic(tErr)
		}

		fmt.Println("Producer Done with err = ", tErr)
	}

	w.Close()
}

func GoConsumer( tOffset int64, cAuth chan<- com.Credential ) {

	r := kafka.NewReader(kafka.ReaderConfig{
    Brokers:   []string{"localhost:9092"},
    Topic:     com.TOPIC_AUTHENTICATION,
    Partition: 0,
    MinBytes:  10e3, // 10KB
    MaxBytes:  10e6, // 10MB
	})

	r.SetOffset( tOffset )

	for {
    m, err := r.ReadMessage( context.Background() )
    if err != nil {
			panic( err )
    }

		tAuth := com.Credential{}
		tAuth.Unmarshal( m.Value )

		cAuth <- tAuth
    fmt.Printf("consumer message at offset %d: %s = %s\n", m.Offset, tAuth)
	}

	r.Close()
}


