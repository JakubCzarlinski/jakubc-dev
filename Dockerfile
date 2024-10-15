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
RUN apk add tar
RUN apk add git

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Add Bun to PATH
# ENV PATH="/root/.bun/bin:${PATH}"
# Alias bun to npm

WORKDIR /app

COPY . .

RUN cd /app

RUN rm -rf ./build/render
RUN rm -rf ./build/render_to_templ

RUN git clone https://github.com/JakubCzarlinski/svelte-ssr ./build/render --quiet
RUN git clone https://github.com/JakubCzarlinski/svelte-ssr-to-templ ./build/render_to_templ --quiet

WORKDIR /app/build/render_to_templ
RUN cd /app/build/render_to_templ
RUN go mod download

WORKDIR /app/project
RUN cd /app/project
RUN go mod download

WORKDIR /app/build/builder
RUN cd /app/build/builder
RUN go mod download
RUN cd /app

WORKDIR /app
RUN bun install
RUN go -C ./build/render_to_templ/ build -ldflags="-s -w" -o ./main.exe ./cmd/main.go
RUN go -C ./build/builder/ build -ldflags="-s -w" -o ./build.exe ./build.go
RUN ./build/builder/build.exe



FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ .

EXPOSE 3000

ENTRYPOINT [ "./main.exe" ]
