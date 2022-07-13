# syntax=docker/dockerfile:1.2

# Stage 1: Build the static files
FROM node:16.15.0-alpine3.15 as frontend-builder
WORKDIR /builder
COPY /frontend/package.json /frontend/package-lock.json ./
RUN npm ci
COPY /frontend .
RUN npm run build

# Stage 2: Build the binary
FROM --platform=$BUILDPLATFORM golang:1.18.3-alpine3.15 as binary-builder
ARG APP_NAME=http
RUN apk update && apk upgrade && \
  apk --update add git
WORKDIR /builder
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=frontend-builder /builder/build ./frontend/build/
ARG TARGETOS
ARG TARGETARCH
# for a normal app without frontend-builder it is worth to use RUN --mount=target=. \ below
# this way we don't need the COPY . . directive and it speeds up the build
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o engine cmd/${APP_NAME}/main.go

# Stage 3: Run the binary
FROM --platform=linux/arm64 gcr.io/distroless/static
ENV APP_PORT=5050
WORKDIR /app
COPY --from=binary-builder --chown=nonroot:nonroot /builder/engine .
EXPOSE $APP_PORT
ENTRYPOINT ["./engine"]