FROM golang:alpine as builder

LABEL maintainer="Daniel Lwetabe <dannylwe11@gmail.com>"

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
# COPY go.mod .
# COPY go.sum .
COPY . .
RUN go mod download

# Copy the code into the container
# COPY . .

# Build the application
RUN go build -o main .

FROM scratch

# copy binary to app directory in second stage
COPY --from=builder /build/main /api/

EXPOSE 9001

ENTRYPOINT [ "/api/main" ]