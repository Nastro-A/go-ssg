FROM --platform=$BUILDPLATFORM golang:1.24.5-alpine3.22 AS build

ARG TARGETOS

ARG TARGETARCH

WORKDIR /usr/src/go-ssg

COPY . .

RUN go mod download && GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o /usr/local/bin/go-ssg .

FROM alpine:3.22.1

ENV PORT=3000

WORKDIR /go-ssg

ENV MD_DIR="/go-ssg/markdown"

ENV HTML_DIR="/go-ssg/html"

COPY --from=build /usr/local/bin/go-ssg /go-ssg/go-ssg

CMD ["/go-ssg/go-ssg"]
