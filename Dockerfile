# development
FROM golang:1.14-alpine AS dev

RUN apk add --no-cache make git
RUN go get -u github.com/cosmtrek/air
RUN mkdir -p /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8000

CMD ["make", "dev"]

# builder image
FROM golang:1.14-alpine AS builder

RUN mkdir -p /app

WORKDIR /app

COPY --from=dev /app ./

RUN make build

# production
FROM golang:1.14-alpine AS prod

RUN mkdir -p /app

WORKDIR /app

COPY --from=builder /app/build ./build
COPY --from=builder /app/Makefile ./Makefile

CMD ["make", "start"]