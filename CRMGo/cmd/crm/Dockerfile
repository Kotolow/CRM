FROM golang:1.23-alpine

WORKDIR /crm

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o crm ./cmd/crm

CMD ["./crm"]