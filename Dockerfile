FROM ubuntu:20.04
COPY red /app/red
WORKDIR /app
CMD ["/app/red"]