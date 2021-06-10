FROM alpine:latest
RUN mkdir -p /app
WORKDIR /app
COPY server_cmd /app
ENTRYPOINT ["./server_cmd"]