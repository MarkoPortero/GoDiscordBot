FROM golang:1.16.15

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go build -o ./out/GoDiscordBot .

CMD ["./out/GoDiscordBot"]