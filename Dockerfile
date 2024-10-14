# syntax=docker/dockerfile:1

FROM golang:1.23.2

# Install curl and unzip for Bun installation
RUN apt-get update && apt-get install -y curl unzip

# Install Bun
RUN curl -fsSL https://bun.sh/install | bash

# Add Bun to PATH
ENV PATH="/root/.bun/bin:${PATH}"
ENV OS="linux"

# Print the GOOS
RUN echo "OS: $OS"

WORKDIR /app

COPY . .

# Change directory to /app
RUN cd /app

# RUN git clone https://github.com/JakubCzarlinski/svelte-ssr ./build/render
# RUN git clone https://github.com/JakubCzarlinski/svelte-ssr-to-templ ./build/render_to_templ

# Build the Go app
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN bun install
RUN go -C ./build/render_to_templ/ build -ldflags="-s -w" -o ./main.exe ./cmd/main.go
RUN go -C ./build/builder/ build -ldflags="-s -w" -o ./build.exe ./build.go
RUN ./build/builder/build.exe


EXPOSE 3000

ENTRYPOINT [ "./main.exe" ]
