
.PHONY: api search recommend useraction clean

init:
	go get github.com/dchest/siphash
	go get github.com/go-sql-driver/mysql
	go get github.com/tylertreat/BoomFilters
	go get github.com/dgrijalva/jwt-go
	go get github.com/segmentio/kafka-go
	go get github.com/agnivade/levenshtein

all: api search recommend useraction

api:
	cd apigateway; make

search:
	cd search; make

recommend:
	cd recommendation; make

useraction:
	cd useraction; make

clean:
	rm $$GOBIN/apigateway $$GOBIN/recommendation $$GOBIN/search $$GOBIN/useraction 2> /dev/null

run:
	apigateway &
	recommendation &
	search &
	useraction &

kill:
	pkill -9 apigateway 2> /dev/null
	pkill -9 recommendation 2> /dev/null
	pkill -9 search 2> /dev/null
	pkill -9 useraction 2> /dev/null

