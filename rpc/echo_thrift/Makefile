THRIFT=~/workspace/local/bin/thrift

codegen: gen-go

gen-go: echo.thrift
	${THRIFT} -r --gen go echo.thrift

clean:
	-rm -rf gen-go
