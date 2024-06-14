####################################################################
# Builder Stage                                                    #
####################################################################
FROM golang:1.22.3-alpine3.18 AS builder

# Set working directory.
WORKDIR /src

# Copy all the code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /src/bin/eth-parser .

####################################################################
# Final Stage                                                      #
##########################################################รื##########
FROM scratch as release
WORKDIR /app

# Copy the CA certificates to the production image from the builder stage.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary to the production image from the builder stage.
COPY --from=builder /src/bin/eth-parser .

# Run the application.
CMD ["/app/eth-parser"]
