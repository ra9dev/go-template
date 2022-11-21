FROM golang:alpine as builder

ENV CGO_ENABLED=0

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o app cmd/*.go


FROM alpine

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app .
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

ENTRYPOINT ["./app"]
