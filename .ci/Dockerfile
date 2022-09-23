FROM golang:1.18-alpine AS build_base

ARG ENV_FILE=.env.dev

ENV CGO_ENABLED=1
ENV GO111MODULE=on

RUN apk add --no-cache \
    git gcc g++ curl

# Set the Current Working Directory inside the container
WORKDIR /src

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY ./.ci/${ENV_FILE} /src/.env

# Build the Go app
RUN go build -o ./out/app ./main.go


# Start fresh from a smaller image
FROM alpine:3.15

RUN apk add ca-certificates curl tzdata

WORKDIR /app

COPY --from=build_base /src/out/app /app/restapi
COPY --from=build_base /src/.env /app/.env

RUN chmod +x restapi

# This container exposes port 3000 to the outside world
EXPOSE 3000

HEALTHCHECK --interval=5s --timeout=10s CMD curl -sf http://localhost:3000/api/v1/health-check || exit 1

# Run the binary program produced by `go install`
ENTRYPOINT ./restapi
