Launch

`docker-compose up -d`

Tests

`docker exec -it -w /go/src/wallet -e CGO_ENABLED=0 server_container go test ./...`

Routes

* `POST /deposit {"receiver":1,"amount":100}`
* `POST /transfer {"sender":1,"receiver":3,"amount":100}`