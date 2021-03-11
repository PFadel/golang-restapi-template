FROM golang:1.15

WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./api ./main.go

EXPOSE 8080

ENTRYPOINT ./api
