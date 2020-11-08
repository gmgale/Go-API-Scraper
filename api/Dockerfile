FROM golang:latest

RUN mkdir -p /app/src
WORKDIR /app/src
RUN go get github.com/gorilla/mux
COPY . 	.
EXPOSE 5432 8080
RUN go build
CMD [ "./wait-for", "db:5432" ]
CMD [ "./goapiscraper" ]


# FROM alpine:latest
# COPY --from=0 . /app/src
# WORKDIR /app/src
# EXPOSE 8080:9000
