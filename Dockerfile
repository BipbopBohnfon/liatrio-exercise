# syntax=docker/dockerfile:1

FROM golang:1.25.1-alpine AS builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /liatrio-my-name-chris

# Remove everything but necessary components from final image
FROM scratch

COPY --from=builder /liatrio-my-name-chris .


# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 80

# Run
CMD ["/liatrio-my-name-chris"]
