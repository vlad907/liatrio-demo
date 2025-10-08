FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

FROM gcr.io/distroless/base-debian12:nonroot

WORKDIR /app

COPY --from=builder /app/server .

ENV PORT=8080

EXPOSE 8080

USER nonroot

ENTRYPOINT ["./server"]
