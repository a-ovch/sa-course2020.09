FROM golang:1.15.6 AS builder
WORKDIR /go/src/user
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /go/src/user/cmd/user
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/src/user/bin/user

FROM alpine:3.12.3
RUN adduser -D app-executor
USER app-executor
WORKDIR /app
COPY --from=builder /go/src/user/bin/user /app/user
ENTRYPOINT ["/app/user"]