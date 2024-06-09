FROM golang:1.22-alpine AS builder

WORKDIR /usr/local/src

COPY ["go.mod", "./"]
RUN go mod download

COPY . ./
RUN go build -o ./bin/healthcheck main.go

FROM alpine AS runner

RUN apk --no-cache add bash

COPY --from=builder /usr/local/src/bin/healthcheck /

EXPOSE 8000

ENTRYPOINT /healthcheck
