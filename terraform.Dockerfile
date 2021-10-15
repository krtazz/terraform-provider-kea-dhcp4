FROM golang:alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN apk add git
ARG CGO_ENABLED=0
ARG GOARCH=amd64 
ARG GOOS=linux 
RUN go build -a -o terraform-provider-kea-dhcp4_v1.0-linux-amd64
ARG GOOS=darwin
RUN go build -a -o terraform-provider-kea-dhcp4_v1.0-darwin-amd64
ARG GOOS=windows
RUN go build -a -o terraform-provider-kea-dhcp4_v1.0-windows-amd64

FROM scratch AS exporter
COPY --from=builder /build/terraform-provider-kea* /

FROM hashicorp/terraform:0.12.20 AS runner
RUN mkdir -p /root/.terraform.d/plugins/linux_amd64/
COPY --from=builder /build/terraform-provider-kea-dhcp4_v1.0-linux-amd64 /root/.terraform.d/plugins/linux_amd64/
RUN mkdir /tffiles
WORKDIR /tffiles
COPY ./test-data/terraform/*.tf /tffiles/
COPY ./entrypoint.terraform.sh /docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]
