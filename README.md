# tros_example_server

#### 介绍
应用级开发框架tros的使用案例

#### 使用方法
- 安装 google protobuf 的编译工具 buf, 我这里使用的是 [buf v1.27.2](https://github.com/bufbuild/buf/releases/tag/v1.27.2)
- 选择对应的平台版本，然后设置环境变量，在IDE中执行，出现如下结果，即为安装成功
```shell
buf --version
```

```text
$ buf --version
1.27.2
```

- 在项目根目录执行以下命令，即可分别在 internal/proto 和根目录的 proto 目录下，生成对应的 接口代码和openapi 文档
```shell
buf lint
buf generate
```

- 目录结构：
    - |--cmd: 启动、参数和controller初始化管理
    - |--conf: 配置文件定义，以及自定义配置项的方法
    - |--docker: 本地化docker镜像打包和运行（如果要迁移到云对象存储，可自行增加上传过程，将img改成远程地址即可）
    - |--internal: 服务内部的应用逻辑
        - |--api: controller层方法
        - |--proto: .proto定义的接口的golang代码实现。此目录为 buf 命令自动生成**请勿修改！！**
    - |--openapi: .proto定义的swagger接口文档，可用网页版管理或postman/apifox。此目录为 buf 明林自动生成**请勿修改！！**
    - |--proto: 以 .proto 文件的形式定义 grpc或http 接口。（**tros支持同时使用这两种协议**）
    - |--main.go: 一般为程序启动入口（可自定义）
**<注意>**：根目录和proto目录下的 buf.* 文件，为 buf 命令编译和检查的相关参数，可根据实际需要自行定义。