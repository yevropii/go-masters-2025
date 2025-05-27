FROM --platform=$BUILDPLATFORM golang:1.22 AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go-masters-2025 ./cmd/server


FROM gcr.io/distroless/static
USER nonroot:nonroot
COPY --from=build /go-masters-2025 /go-masters-2025
EXPOSE 8080
ENTRYPOINT ["/go-masters-2025"]
