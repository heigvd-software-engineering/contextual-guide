FROM golang:alpine as build
ENV CGO_ENABLED=0

WORKDIR /app

COPY . /app
RUN go build -o server .

FROM scratch
COPY ./src/assets /src/assets
COPY ./src/views /src/views
COPY --from=build /app/server /app/server
EXPOSE 3000
CMD ["/app/server"]
