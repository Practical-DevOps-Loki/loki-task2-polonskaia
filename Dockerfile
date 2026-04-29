FROM golang:1.21-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o webapp

FROM alpine:latest

WORKDIR /app

ENV PORT=3000
ENV LOG_PATH=/app/log/app.log

COPY --from=build-stage /app/webapp /app
COPY --from=build-stage /app/public /app/public

EXPOSE 3000

CMD ["/app/webapp"]