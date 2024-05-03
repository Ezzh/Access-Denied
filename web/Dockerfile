FROM golang:1.21.5-alpine3.18

WORKDIR /usr/src/app

COPY go.mod go.sum* ./
RUN go mod download && go mod verify
EXPOSE 8080

COPY ./handlers ./handlers/
COPY ./models ./models/
COPY ./logic ./logic/
COPY ./database ./database/
COPY ./main.go .
COPY ./templates ./templates/
RUN ls
RUN go build -o /usr/local/bin/app

CMD ["app"]