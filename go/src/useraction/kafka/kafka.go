package kafka

import (
	"fmt"
	"context"
	"github.com/segmentio/kafka-go"

	com "common"
)

var (
	SendMovie			chan com.Movie
	SendAvgRating	chan com.AvgRating
)

func init() {

	SendMovie = make( chan com.Movie, 100 )
	SendAvgRating	 = make( chan com.AvgRating, 100 )
}

func GoProducer() {

	kafka_movie_writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   com.TOPIC_MOVIE,
		Balancer: &kafka.LeastBytes{},
	})

	kafka_rating_writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   com.TOPIC_RATING,
		Balancer: &kafka.LeastBytes{},
	})

	for {
		select {

			case movie := <-SendMovie: {
				tRaw, tErr := movie.Marshal()
		    if tErr != nil {
				  panic( tErr )
				}

				tErr = kafka_movie_writer.WriteMessages( context.Background(),
					kafka.Message{
						Value: tRaw,
					},
				)
				if tErr != nil {
					panic(tErr)
				}

				fmt.Println("useraction movie Producer Done with err = ", tErr)
			}

			case rating := <-SendAvgRating: {
				tRaw, tErr := rating.Marshal()
		    if tErr != nil {
				  panic( tErr )
				}

				tErr = kafka_rating_writer.WriteMessages( context.Background(),
					kafka.Message{
						Value: tRaw,
					},
				)
				if tErr != nil {
					panic(tErr)
				}

				fmt.Println("useraction rating Producer Done with err = ", tErr)
			}
		}
	}

	kafka_movie_writer.Close()
	kafka_rating_writer.Close()
}

/*func GoProducer() {

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
}*/

