package kafka

import (
	"fmt"
	"context"
	"github.com/segmentio/kafka-go"

	com "common"
)

var (
	SendMovie chan com.Movie
)

func init() {

	SendMovie = make( chan com.Movie, 100 )

	go GoProducer()
}

func GoProducer() {

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   com.TOPIC_MOVIE,
		Balancer: &kafka.LeastBytes{},
	})

	// This automatically breaks when channel closes
	for tMovie := range SendMovie {

		tRaw, tErr := tMovie.Marshal()
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

		fmt.Println("useraction Producer Done with err = ", tErr)
	}

	w.Close()
}

