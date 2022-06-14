FROM golang:1.18-alpine

#RUN go install github.com/maxcnunes/gaper/cmd/gaper@latest

WORKDIR /go/src/wallet

COPY internal ./internal
COPY cmd/wallet/ ./
COPY go.mod go.sum ./

EXPOSE 8080

RUN go install wallet

CMD ["wallet"]

#ENTRYPOINT gaper --disable-default-ignore