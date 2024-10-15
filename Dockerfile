# syntax=docker/dockerfile:1

FROM golang:1.23.2

# Install curl and unzip for Bun installation
RUN apt-get update && apt-get install -y curl unzip

# Install Bun
RUN curl -fsSL https://bun.sh/install | bash
RUN go install github.com/a-h/templ/cmd/templ@latest

# Add Bun to PATH
ENV PATH="/root/.bun/bin:${PATH}"
ENV OS="linux"
ENV PROD="true"

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
RUN cd /app
WORKDIR /app


# Build the Go app
RUN bun install
RUN go -C ./build/render_to_templ/ build -ldflags="-s -w" -o ./main.exe ./cmd/main.go
RUN go -C ./build/builder/ build -ldflags="-s -w" -o ./build.exe ./build.go
RUN ./build/builder/build.exe


EXPOSE 3000

ENTRYPOINT [ "./main.exe" ]
