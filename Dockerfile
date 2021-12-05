# syntax=docker/dockerfile:1

FROM golang:1.14

ENV GO111MODULE=on

# Download necessary Go modules
COPY . ./IDT-messaging
WORKDIR IDT-messaging
RUN git config --global url."https://".insteadOf git://
RUN go mod download

# Build the app
RUN go build -o /IDT-messaging-core-app

# Expose port 8080
EXPOSE 8080

# Run the app
CMD [ "/IDT-messaging-core-app" ]