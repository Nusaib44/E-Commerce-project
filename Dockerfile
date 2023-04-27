
FROM golang:1.20.2-alpine

RUN mkdir /apk
WORKDIR /app
ADD . /app


COPY  go.mod .
COPY go.sum .
COPY  .env .

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

RUN go mod download
RUN go build -o main .

EXPOSE 8080
CMD [ "/app/main" ]


























# FROM golang:1.17-alpine AS build

# # Set the working directory to /app
# WORKDIR /app

# # Copy the go.mod and go.sum files to the working directory
# COPY go.mod go.sum ./

# # Install the dependencies
# RUN go mod download

# # Copy the rest of the application code to the working directory
# COPY . .

# # Build the application
# RUN go build -o app

# # Use a smaller base image for the final image
# FROM alpine:3.14

# # Set the working directory to /app
# WORKDIR /app

# # Copy the application binary from the build stage to the final image
# COPY --from=build /app/app .

# # Expose port 8080 for the application
# EXPOSE 8080

# # Start the application
# CMD ["./app"]