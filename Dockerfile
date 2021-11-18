FROM golang:1.16-alpine as build
WORKDIR /app
COPY go.mod go.sum . ./
RUN  go mod download && \
 go build -o app && \
 chmod +x /app
EXPOSE 8080
ENTRYPOINT [ "/app/app" ]
