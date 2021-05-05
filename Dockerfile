# Compile stage
FROM golang:1.16 AS build-env
LABEL stage=builder
ADD . /dockerdev
WORKDIR /dockerdev
RUN go build -o /server

# Final stage
FROM debian:buster
EXPOSE 2122
WORKDIR /
COPY --from=build-env /server dockerdev/configuration.yml dockerdev/secrets.yml /
CMD ["/server"]