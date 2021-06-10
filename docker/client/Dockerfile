FROM alpine:latest
RUN mkdir -p /app
WORKDIR /app
COPY client_cmd /app
ENTRYPOINT ["./client_cmd"]
EXPOSE 8081