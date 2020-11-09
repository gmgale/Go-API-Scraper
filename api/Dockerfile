#builder image
FROM golang:latest AS builder
RUN mkdir -p /build
ADD *.go /build/
WORKDIR /build
RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq
RUN CGO_ENABLED=0 GOOS=linux go build -a -o myAPI

# generate clean, final image for end users
FROM alpine:latest
COPY --from=builder /build/myAPI .
EXPOSE 8080:9000

# executable
ENTRYPOINT [ "./myAPI" ]
# arguments that can be overridden
CMD [ "-port", "8080" ]
