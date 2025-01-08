FROM golang:1.23.2-alpine as Builder

WORKDIR /app

COPY ../go.mod ../go.sum ./

RUN go mod download

COPY docs/ docs/

COPY backend/ backend/

RUN go build -o /app/main ./backend/cmd/main.go

FROM alpine:latest

COPY --from=Builder /app/main /bin/main

CMD ["/bin/main"]