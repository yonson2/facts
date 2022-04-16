FROM golang:bullseye
ENV GOPROXY http://proxy.golang.org
ENV GO_ENV=production
WORKDIR /var/www/app
COPY . ./
RUN go mod download
RUN go build -o /app ./main.go

EXPOSE 3000

CMD ["/app"]
