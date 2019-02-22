# set compiler, protoc use protoc-gen-go plugin, please make sure it!
PROTOC=~/go/bin/protoc/bin/protoc

codegen: echo.pb.go

echo.pb.go: echo.proto
	${PROTOC} --go_out=plugins=grpc:. echo.proto

clean:
	-rm -rf echo.pb.go
