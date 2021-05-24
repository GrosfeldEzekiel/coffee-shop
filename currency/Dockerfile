FROM dependencies AS builder
# Copy the application source code.
COPY . /go/src/app/currency
# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /go/bin/currency /go/src/app/currency/
ENTRYPOINT [ "/go/bin/currency" ]

FROM alpine:latest
COPY --from=builder /go/bin/currency /bin/currency
ENTRYPOINT [ "/bin/currency" ]