
# See: https://protobuf.dev/reference/go/go-generated/

PREFIX=github.com/RexGreenway/CollectionApp

{
    protoc \
        --go_out="./" \
        --go_opt="module=$PREFIX" \
        --go-grpc_out="./" \
        --go-grpc_opt="module=$PREFIX" \
        ./protos/collection.proto
}