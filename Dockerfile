FROM golang:1.14-alpine AS build

WORKDIR /app
COPY ./src /app
RUN CGO_ENABLED=0 go build -o /bin/demo

EXPOSE 5000

FROM scratch
COPY --from=build /bin/demo /bin/demo
ENTRYPOINT ["/bin/demo"]