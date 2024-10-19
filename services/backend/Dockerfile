FROM golang:1.23.2-alpine3.20 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app .

FROM alpine:3.20

WORKDIR /app
RUN apk --no-cache add ca-certificates tzdata curl
COPY --from=build /app/app .
ENTRYPOINT [ "/app/app" ]
