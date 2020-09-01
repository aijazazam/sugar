package kafka

import (
	"fmt"
	"context"
	"github.com/segmentio/kafka-go"

	com "common"
	sql "recommendation/db"
)

var (
  RecvMovie		chan com.Movie
	RecvRating	chan com.AvgRating
)

func init() {
  RecvMovie = make( chan com.Movie, 100 )
	RecvRating = make( chan com.AvgRating, 100 )
}

func GoMovieConsumer( tOffset int64 ) {

	r := kafka.NewReader(kafka.ReaderConfig{
    Brokers:   []string{"localhost:9092"},
    Topic:     com.TOPIC_MOVIE,
	})

	r.SetOffset( tOffset )

	db := sql.GetDB()

	for {
    m, err := r.ReadMessage( context.Background() )
    if err != nil {
			panic( err )
    }

		tMovie := com.Movie{}
		tMovie.Unmarshal( m.Value )

		fmt.Println(tMovie)

		if err := db.InsertMovie( tMovie ); err != nil {
      panic( err )
    }
	}

	r.Close()
}

func GoRatingConsumer( tOffset int64 ) {

  r := kafka.NewReader(kafka.ReaderConfig{
    Brokers:   []string{"localhost:9092"},
    Topic:     com.TOPIC_RATING,
  })

  r.SetOffset( tOffset )

  db := sql.GetDB()

  for {
    m, err := r.ReadMessage( context.Background() )
    if err != nil {
      panic( err )
    }

    tRating := com.AvgRating{}
    tRating.Unmarshal( m.Value )

    fmt.Println(tRating)

    if err := db.UpdateMovieRating( tRating ); err != nil {
      panic( err )
    }
  }

  r.Close()
}

