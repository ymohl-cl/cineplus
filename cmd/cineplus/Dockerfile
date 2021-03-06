##
# Configure base image golang
##
FROM golang:1.15 AS base

RUN apt-get install -y ca-certificates
WORKDIR /go/cmd/app

##
# Get dependencies
##
FROM base AS deps

# Copy dependencies builder
COPY Makefile .
COPY go.mod .
COPY go.sum .

# Install dependencies
RUN make tools
RUN make install

##
# Running test and building
##
FROM deps AS build

COPY cmd/cineplus cmd/cineplus
COPY pkg pkg/
RUN make lint
RUN make test
RUN make build

FROM build AS conf

##
# Create a final distroless image: debian:strech
##
FROM gcr.io/distroless/base:latest AS distroless

COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/cmd/app/bin/app .

ENTRYPOINT ["./app"]