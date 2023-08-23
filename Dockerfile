FROM golang:alpine as build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o leaky .

FROM alpine:latest as runtime
COPY --from=build /app/leaky /usr/local/bin/leaky
ENTRYPOINT ["/usr/local/bin/leaky"]
