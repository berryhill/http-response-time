FROM iron/go:dev

# Working Directory
WORKDIR /app

# Build App
ENV SRC_DIR=/go/src/github.com/berryhill/http-response-time
ADD . $SRC_DIR
RUN cd $SRC_DIR; go build -o http-response-time; cp http-response-time /app/

# Run App
ENTRYPOINT ./http-response-time