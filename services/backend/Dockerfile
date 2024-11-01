FROM golang:1.23.2-alpine3.20 AS build

# Copy libraries (check the service's go.mod)
WORKDIR /src/lib
COPY lib/db ./db

# Copy service
WORKDIR /src/services/backend
COPY services/backend/go.mod services/backend/go.sum ./
RUN go mod download
COPY services/backend .
RUN go build -o /app/app .

FROM alpine:3.20

WORKDIR /app
RUN apk --no-cache add ca-certificates tzdata curl
COPY --from=build /app/app .
ENTRYPOINT [ "/app/app" ]
