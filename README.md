# Description

The application allows you to read the syslog using the go-syslog library and print it to standard output using the Logrus library. Each event must have severity and facility records printed out using the WithFields mechanism.

The application consists of a server and a client. The server and client work with each other using the GRPC protocol. Initially, only the server is running. After the server receives a request, it processes request and outputs the result to the console. The client can connect to the application. After the client connects, the server sends syslog to the client, and the client is already processing these logs and output the result to the console. Also, the client can disconnect and only the server will work.

# How to use

## Docker

For the application to work, you first need to start the server, via docker-compose:

Run docker-compose: `docker-compose up`

## Sending syslog message via logger
Next, through bash, send a request to the server:

Command: `logger --udp -n 0.0.0.0 --port 1514 --rfc5424 "User 'username' Executed the 'Configure Term' Command."`

Congratulations, the app is working!!!

## Run integration tests

To run unit tests type: `go test ./...`
