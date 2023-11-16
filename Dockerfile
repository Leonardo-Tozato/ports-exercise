FROM golang:1.21.4-bookworm AS base
WORKDIR /usr/app

COPY go.mod go.sum ./
RUN go mod download -x

FROM base as test
COPY . ./
CMD ["go", "test", "./..."]

FROM base as compiler
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -ldflags "-w -s" -o /bin/app cmd/app/main.go


FROM alpine:3.18 AS release
WORKDIR /usr/app
RUN apk add --update --no-cache ca-certificates tzdata && \
    ln -fs /usr/share/zoneinfo/UTC /etc/localtime && \
    rm -rf /var/cache/apk/* /tmp/* /var/tmp/*

COPY --from=compiler /bin/app /bin/app
COPY ./ports.json ./

CMD ["/bin/app", "./ports.json"]