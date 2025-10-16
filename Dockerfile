FROM golang:1.25.1

WORKDIR /app

COPY XKCD-Client.go ./

RUN go build -o client XKCD-Client.go

CMD [ "./client" ]