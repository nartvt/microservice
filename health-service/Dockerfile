FROM golang:1.17-alpine

WORKDIR app/

COPY . .

RUN go mod tidy

RUN go run *.go

CMD ["go","version" ]

EXPOSE 3000