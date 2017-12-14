# 目录说明
一个GO工程中主要包含以下三个目录：
* src：源代码文件
* pkg：包文件
* bin：相关bin文件

# 建立工程步骤
1. 建立工程文件夹 mygo
2. 在工程文件夹中建立 gopath/src, gopath/pkg, gopath/bin文件夹
3. 在GOPATH中添加 gopath 路径, mygo/gopath
4. 如工程中有自己的包 hellogo，那在src文件夹下建立以包名命名的文件夹 例 gopath/src/hellogo
5. 在bin文件架下编写主程序代码代码 hello.go
6. 在hellogo文件夹中编写 hellogo.go 和包测试文件 hellogo_test.go
7. 编译调试包
       go build hellogo  
       go test hellogo  
       go install hellogo  
       这时在pkg文件夹中可以发现会有一个相应的操作系统文件夹如linux_amd64, 在这个文件夹中会有 hellogo 文件夹，在该文件中有hellogo.a文件

8. 编译主程序
       go build gopath/bin/hello.go  
       成功后会生成 hello 文件

# 常用的package
Logging: see https://github.com/cihub/seelog/wiki/Quick-start
```bash
    go get github.com/cihub/seelog  
    go install github.com/cihub/seelog  
    go build gopath/bin/seelog.go  
```