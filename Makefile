# exports file at internal/types
proto:
    protoc -I=internal/pb --go_out=internal/ internal/pb/*.proto
