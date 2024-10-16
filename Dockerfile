# syntax=docker/dockerfile:1

FROM golang:1.23.2-alpine AS golang

FROM oven/bun:alpine AS builder

COPY --from=golang /usr/local/go/ /usr/local/go/
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

ENV PATH="$PATH:/usr/local/go/bin"
ENV OS="linux"
ENV PROD="true"

RUN apk update
RUN apk add git

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

RUN git clone https://github.com/JakubCzarlinski/svelte-ssr /app/build/render --quiet
RUN git clone https://github.com/JakubCzarlinski/svelte-ssr-to-templ /app/build/render_to_templ --quiet
RUN go -C /app/build/render_to_templ/ mod download

WORKDIR /app
COPY package.json /app/package.json
RUN bun install

COPY ./project/go.mod /app/project/go.mod
COPY ./project/go.sum /app/project/go.sum
RUN go -C /app/project mod download

COPY ./build/builder/go.mod /app/build/builder/go.mod
COPY ./build/builder/go.sum /app/build/builder/go.sum
RUN go -C /app/build/builder mod download

COPY . /app

RUN go -C /app/build/render_to_templ/ build -ldflags="-s -w" -o ./main.exe ./cmd/main.go
RUN go -C /app/build/builder/ build -ldflags="-s -w" -o ./build.exe ./build.go
RUN /app/build/builder/build.exe


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/dist /app/dist
COPY --from=builder /app/main.exe /app/main.exe

EXPOSE 3000

ENTRYPOINT [ "./main.exe" ]
