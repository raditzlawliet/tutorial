# syntax=docker/dockerfile:1
# Stage build
FROM golang:1.21 as build
WORKDIR /src
COPY <<EOF /src/main.go
package main

import "fmt"

func main() {
  fmt.Println("hello, world")
}
EOF
RUN go build -o /bin/hello ./main.go

# Last Stage
FROM scratch
COPY --from=build /bin/hello /bin/hello
CMD ["/bin/hello"]

# Example how to ...
### docker build -t go-app .
### docker run go-app