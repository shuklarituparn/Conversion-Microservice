# Use a base image with the desired GLIBC version
FROM alpine:latest

# Install necessary dependencies
RUN apk update && apk add --no-cache libc6-compat

# Copy your program files into the container
COPY . /app

# Set the working directory
WORKDIR /app/cmd/video-convertor

# Run your program when the container starts
CMD ["./main"]
