
## build

FROM golang AS build

WORKDIR /go/src/blinkr

COPY go.sum go.mod ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o /blinkr

## deploy

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /blinkr /blinkr

EXPOSE 8000

CMD ["/blinkr"]
