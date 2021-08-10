FROM golang:1.16 AS build
WORKDIR /compose/
COPY go.sum go.sum
COPY go.mod go.mod
COPY main.go main.go
RUN CGO_ENABLED=0 go build -o backend main.go

FROM alpine 
COPY --from=build /compose/ . 
CMD ["./backend"]
