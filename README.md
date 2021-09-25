# 一个用Golang实现的Html组件编译工具

+ 生成 html 节点树
+ 组件命名请使用 **x-y** 的模式，如 **go-form**，文件命名：**go-form.html**，保持工程唯一
+ 支持组件嵌套
+ 支持组件变量传递、自定义全局变量（mcss）
+ 自定义配置，参阅 **internal/config_sample.go**

## 测试
```shell
# html 测试请参阅 examples 目录
# 自定义配置可修改 mcss.yaml 测试
# http://127.0.0.1:8080

go run main.go
```