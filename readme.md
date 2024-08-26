## ZxwyWebSite/SimpleServer
### 简介
+ 一个简易 HTTP 文件服务器
+ 我们在进行 Web 开发时经常会遇到需要预览文件的情况，而有些前端无法在文件模式下正常运行，就需要在本地起一个文件服务器

### 安装
+ 可直接下载 Release 版本或自行构建
+ `go install github.com/ZxwyWebSite/simpleserver@latest`

### 使用
+ 将可执行文件放入 $PATH，然后任意位置执行 simpleserver
+ 默认在运行目录启动服务，监听 0.0.0.0:8100
+ 启动参数
  ```
  PS C:\Users\root> simpleserver.exe -h
  flag needs an argument: -h
  Usage of C:\Programs\Path\simpleserver.exe:
  -f string
        Fdir (default ".")
  -h string
        Host (default "0.0.0.0")
  -p string
        Port (default "8100")
  -s string
        Sdir (default "/")
  ```
+ Fdir 文件目录，默认为当前运行位置
+ Host 监听地址，默认为任意 ipv4
+ Port 监听端口，默认为 8100
+ Sdir 路由地址，默认为 /

### 其它
+ 后续可能会美化文件列表，也可能直接停更