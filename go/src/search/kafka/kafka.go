package kafka

import (
	"fmt"
	"context"
	"github.com/segmentio/kafka-go"

	com "common"
	sql "search/db"
)

var (
  RecvMovie chan com.Movie
)

func init() {
  RecvMovie = make( chan com.Movie, 100 )
}

func GoConsumer( tOffset int64 ) {

	r := kafka.NewReader(kafka.ReaderConfig{
    Brokers:   []string{"localhost:9092"},
    Topic:     com.TOPIC_MOVIE,
    //Partition: 0,
    //MinBytes:  10e3, // 10KB
    //MaxBytes:  10e6, // 10MB
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


