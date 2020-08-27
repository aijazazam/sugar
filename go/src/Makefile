
.PHONY: api search recommend useraction clean

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
