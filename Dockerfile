FROM golang:latest AS build

WORKDIR /gowork

ARG GOPROXY
ARG CGO_ENABLED=0

COPY . .

RUN echo $(git rev-parse HEAD) >internal/version/commit

RUN go build -v -o . ./...

FROM alpine:latest

WORKDIR /app

ENV PATH=$PATH:/app

COPY --from=build /gowork/fsctl .

WORKDIR /workspace