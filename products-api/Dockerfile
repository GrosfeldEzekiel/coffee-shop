FROM dependencies AS builder
# Copy the application source code.
COPY . /go/src/app/products-api
# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /go/bin/products-api /go/src/app/products-api/
ENTRYPOINT [ "/go/bin/products-api" ]

FROM alpine:latest
COPY --from=builder /go/bin/products-api /bin/products-api
ENTRYPOINT [ "/bin/products-api" ]