FROM golang:alpine as build
ENV CGO_ENABLED=0

WORKDIR /app

COPY . /app
RUN cd /app/src/cmd/server &&  go build -o server .

FROM scratch
COPY --from=build /app/src/cmd/server/server /app/server
EXPOSE 3000
CMD ["/app/server"]
