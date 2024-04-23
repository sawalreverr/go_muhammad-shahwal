FROM golang:1.22 AS build-stage

WORKDIR /go/src/simple-api

COPY . . 

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /simple-api

FROM gcr.io/distroless/base-debian11 AS build-release-stage

COPY --from=build-stage /go/src/simple-api/.env /.env
COPY --from=build-stage /simple-api /simple-api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./simple-api"]