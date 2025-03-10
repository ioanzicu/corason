# First stage - builder
FROM golang:1.24 AS builder

WORKDIR /usr/src/app

COPY . .

RUN CGO_ENABLE=0 GOOS=linux go build -a -installsuffix cgo -o corason ./cmd/main.go

# Second stage scratch 
FROM scratch 

COPY --from=builder /usr/src/app/corason ./corason

CMD ["./corason"]