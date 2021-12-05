FROM golang:alpine

WORKDIR /loan-app

ADD . .

RUN go mod download

RUN go mod verify

RUN go build -o loan-app ./app/main.go

ENTRYPOINT ["./loan-app"]
