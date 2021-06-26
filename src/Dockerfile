FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o api main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/api /app/
ENTRYPOINT ./api