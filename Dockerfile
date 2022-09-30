# docker build -t userut:v1 .
FROM golang

WORKDIR /src

COPY . .
RUN go env -w GOPROXY="https://goproxy.cn,direct" \
    && go mod tidy



RUN GOOS=linux GOARCH=amd64 go build -o gotest .


# Expose port 1323 to the outside world
EXPOSE 1323

# Command to run the executable
CMD ["./gotest"]