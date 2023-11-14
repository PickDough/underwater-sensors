# syntax=docker/dockerfile:1
FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY ./go.mod ./go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY src/sensor_channels/ ./src/sensor_channels/
COPY src/common/ ./src/common/
COPY wait-for-it.sh ./wait-for-it.sh

RUN chmod +x ./wait-for-it.sh

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /sensor-channels ./src/sensor_channels

# Run
CMD ["/sensor-channels"]