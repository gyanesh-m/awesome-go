# Build stage
FROM golang:1.22-alpine AS builder


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk add --no-cache make

RUN make dependencies setup build-migration

# Run stage
FROM alpine:20240315

WORKDIR /app

COPY --from=builder /app/bin/migration /app/

COPY --from=builder /app/configs/ /app/configs/

COPY --from=builder /app/internal/db/migrations/ /app/internal/db/migrations/

CMD ["./migration"]

