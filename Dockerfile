# Compile stage
FROM golang:1.16 AS build-env
LABEL stage=builder
RUN mkdir -p /files_to_import/pve-exporter/config
ADD configuration.yml secrets.yml /files_to_import/pve-exporter/config/
ADD . /dockerdev
WORKDIR /dockerdev
RUN go build -o /files_to_import/pve-exporter/server

# Final stage
FROM debian:buster
EXPOSE 2122
COPY --from=build-env /files_to_import /
WORKDIR /
CMD ["/pve-exporter/server"]