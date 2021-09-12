FROM golang:1.16

WORKDIR /stone
COPY . .

EXPOSE 8080 5432

RUN go mod download
RUN go build -o "stone.out" src/*.go

CMD ["./stone.out"]