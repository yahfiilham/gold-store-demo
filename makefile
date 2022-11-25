all: validate clean generate build

validate:
	swagger validate ./api/swagger.yml

spec:
	swagger generate spec -o ./api/swagger-gen.yml

clean:
	rm -rf gold-store-demo-server
	rm -rf gold-store-demo-migrator
	go clean -i .

generate:validate clean
	swagger -q generate server -f api/swagger.yml -A gold-store-demo -s internal/apis -m pkg/models
	swagger -q generate client -f api/swagger.yml -A gold-store-demo -c pkg/client -m pkg/models

doc:
	swagger validate ./api/swagger.yml
	swagger serve api/swagger.yml --no-open --host=0.0.0.0 --port=8081 --base-path=/
	
build: clean
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix cgo ./cmd/gold-store-demo-server

run: 
	./gold-store-demo-server --port=8080 --host=0.0.0.0

.PHONY: validate spec clean generate doc build run 
