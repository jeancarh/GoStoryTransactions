FROM golang:1.21.3

# Set destination for COPY
WORKDIR /app
# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .
# Build
RUN go build -o storiapp
ENV \
    MongoUrl="mongodb+srv://jeancarlosh:EQ6ZdGJvAx7Y34dN@atlascluer.pccaqkn.mongodb.net" \
    MongoDb="stori" \
    MongoCollection="transactions" \
    emailRecipient="jeancarhz@gmail.com" \
    emailSender="jeancarhz@gmail.com" \
    emailPass="RfMA8cYEUNFd4tVh" \
    smtpServer="smtp-relay.brevo.com" \
    smtpPort=587

# Run
CMD ["./storiapp"]