# Start with an Ubuntu base image
FROM --platform=linux/amd64 ubuntu:20.04

# Avoid prompts from apt
ENV DEBIAN_FRONTEND=noninteractive

# Install Python and other essential tools
RUN apt-get update && apt-get install -y \
    python3 \
    python3-pip \
    curl 

# Install Go
ENV GO_VERSION 1.20
RUN curl -LO https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz \
    && rm go${GO_VERSION}.linux-amd64.tar.gz
ENV PATH $PATH:/usr/local/go/bin

# Set the working directory in the container
WORKDIR /app

# Copy the Python script and requirements file
COPY . /app/

# Install Python dependencies
RUN python3 -m pip install --no-cache-dir -r requirements.txt

RUN go mod download
RUN go mod tidy

# Build the Go API
RUN go build -o go-api main.go

# Expose port 3000
EXPOSE 3000

# Command to run on container start (Adjust as per your Go API executable)
CMD [ "./go-api" ]