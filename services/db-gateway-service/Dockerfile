FROM golang:1.23.7-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./src ./src
WORKDIR /app/src

RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]