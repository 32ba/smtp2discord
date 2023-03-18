FROM golang:1.19.0 AS builder
WORKDIR /app
ENV GO111MODULE=on
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o smtp2discord .
RUN strip smtp2discord

FROM gcr.io/distroless/static-debian11
WORKDIR /
COPY --from=builder /app/smtp2discord /smtp2discord
USER nonroot
CMD ["/smtp2discord"]
