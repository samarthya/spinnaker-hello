FROM golang:1.17.6-bullseye AS GOLANG
WORKDIR /src
ADD go.mod go.mod
ADD go.sum go.sum
RUN go mod download
ADD main.go main.go
RUN go build -o /server .

ENV HTTP_PORT=8181

FROM gcr.io/distroless/base-debian10
COPY --from=GOLANG /server /server
EXPOSE ${HTTP_PORT:-8181}/tcp
# USER ${USER:-root}
ENTRYPOINT [ "/server"]