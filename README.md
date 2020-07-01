# DIRService
我的运行环境为windows，本代码支持Linux
运行步骤：
DIRService、InfoServer两个服务的ip和端口号在Handle/PubConst 文件中设置

1、开启DIRService：

  	go build Server/ServerHttp.go
	
  	运行：ServerHttp.exe D:\Soft\GoLand
	
  	浏览器打开：http://127.0.0.1:8091/?path=doc
	
2、开启InfoServer：

  	go build InfoServer/InfoServerHttp.go
	
  	运行：InfoServerHttp.exe
	
  	浏览器打开：http://127.0.0.1:8092/?path=doc





需求如下：

1. 使用 Golang 编写一个 HTTP Web Service（DIRService）并完成调试：

服务描述
此HTTP 服务用于远程目录信息查询，以HTTP GET 形式提供服务；
形式	
HTTP GET

启动参数
服务启动时可指定一个启动参数root， 以此作为根目录， 用于查询本机root 路径下的文件信息；

服务命令行启动
参数	类型	描述
root	string	作为服务根目录
的本机路径

查询参数
客户端 (浏览器或Postman) 发送一个GET 请求到服务端， 并包含以下参数：

参数名	类型	描述
path	string	相对路径

服务端需完成的工作
•	解析path 参数， 并结合 root + path 得到完整本机路径；
•	获取完整路径下的目录信息；
•	返回以下形式的JSON 对象： (size 单位为字节)
{

“path”: “path/to/target”, “dirs”: [
{“name”:”file1”, “isDir”: FALSE, “size”: 100},
{“name”:”file2”, “isDir”: FALSE, “size”: 50},
{“name”:”dir1”, “isDir”: TRUE, “size”: 0},

]
}

2.	使用 Golang 编写一个 HTTP Web Service （DIRInfoService）， 调用以上（DIRService的 HTTP 接口）

服务描述
此HTTP 服务通过调用之前开发的DIRService 的HTTP 接口，用于统计远程目录信息，以HTTP GET
形式提供服务。 统计的信息包括：
1.	文件计数；
2.	目录计数（包括起始目录）；
3.	累计所用磁盘空间节数；

统计范围包括：
1.	起始目录下 (包括各层子目录下) 的所有文件和目录；

形式
HTTP GET
查询参数
客户端 (浏览器或Postman) 发送一个GET 请求到服务端， 并包含以下参数：

参数名	类型	描述
path	string	相对路径

服务端需完成的工作
•	解析path 参数， 并以此依次调用DIRService (包括各层子目录)；
•	需要多线程调用（DIRService）；
•	限制最多 5 个同时并发的调用协程；
•	统计文件计数及累计所用磁盘空间节数；
•	返回以下JSON 对象：
{
“path”: ”path/to/target”,
“dirCount”: 2,	//包括起始目录

“fileCount”: 3,
“totalSize”: 200
}
