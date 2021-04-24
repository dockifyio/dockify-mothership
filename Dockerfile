FROM golang:1.15.2 AS build
WORKDIR /$GOPATH/src/github.com/dockifyio/dockify-mothership
COPY . .
# Download all the dependencies
RUN go get -d -t -v ./...

# Install the package
RUN go install -v ./...

RUN FIREBASE_API=ENTER_FIREBASE_API_FOR_TESTING go test -v ./test/...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-s' cmd/dockify-mothership/main.go

FROM golang:1.15-alpine

# Add go user
RUN adduser --system gorole \
    && mkdir -p /app \
    && chown -R gorole /app

# Change the work directory
WORKDIR /app

# Set Env Vars
ENV SERVER_PORT 8080

COPY --from=build /$GOPATH/src/github.com/dockifyio/dockify-mothership/main .

EXPOSE $SERVER_PORT

# Run
CMD ./main
