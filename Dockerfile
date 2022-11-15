FROM docker.io/golang:1.17-alpine AS builder

WORKDIR /go/src/app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11

COPY --from=builder /go/bin/app /
ENTRYPOINT ["/app"]