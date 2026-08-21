FROM golang:alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /semver-sort ./cmd/semver-sort

FROM alpine
RUN apk add --no-cache jq curl
COPY --from=build /semver-sort /usr/local/bin/
COPY resource /opt/resource
