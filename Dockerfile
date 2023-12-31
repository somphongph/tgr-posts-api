FROM golang:1.21-alpine3.18 AS build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go test --tags=unit -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o ./out/goapp ./

FROM alpine:3.18
COPY --from=build-base /app/out/goapp /app/goapp
COPY ./configs ./configs

RUN adduser -D -u 1000 appuser
USER appuser

CMD ["/app/goapp"]
