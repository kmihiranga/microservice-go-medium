FROM alpine:latest

RUN mkdir /app
COPY authApp /app
COPY ops /app/ops

CMD ["/app/authApp"]