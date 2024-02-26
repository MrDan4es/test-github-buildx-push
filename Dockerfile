FROM golang:1.22 as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o main

FROM gcr.io/distroless/static-debian12:latest

COPY --from=builder /app/main /main

ENTRYPOINT ["/main"]

