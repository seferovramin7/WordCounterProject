FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

COPY endg-urls /app/endg-urls

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/endg-urls .

CMD ["./main"]


#docker build -t wordcounterproject .
#docker run -d -p 8080:8080 wordcounterproject

#docker tag wordcounterproject seferovramin7/wordcounterproject
#docker push seferovramin7/wordcounterproject
