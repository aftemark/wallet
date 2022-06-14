FROM golang:1.18-alpine

WORKDIR /usr/src/consumers

COPY internal ./internal
COPY cmd/deposit_consumer/main.go ./
COPY go.mod go.sum ./

RUN go install wallet

CMD ["wallet"]