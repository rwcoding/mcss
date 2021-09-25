# 一个用Golang实现的Html组件编译工具

+ 生成 html 节点树
+ 组件命名请使用 **x-y** 的模式，如 **go-form**，文件命名：**go-form.html**，保持工程唯一
+ 支持组件嵌套
+ 支持组件变量传递、自定义全局变量（mcss）
+ 自定义配置，参阅 **internal/config_sample.go**

## 测试
```shell
# 新建 src 目录
# 添加以下测试文档

go run main.go
```

## 示例-源文件
#### index.html
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ mcss.app }}</title>
</head>
<body>
<div style="text-align: center">
    <go-content arr="[arr-1,arr-2,arr-3]" obj="{a:obj-1, b:obj-2}"></go-content>
</div>

<script>
    let color = "red"
    setInterval(function () {
        color = color === "red" ? "black" : "red"
        document.getElementById('title').style.color = color
    }, 1000*2)
</script>
</body>
</html>
```
#### go-content.html
```html
<style>
    li {
        text-align: left;
    }
</style>
<go-header title="Go-Header - {{ mcss.app }}"></go-header>
<ul>
    <li>render: object and array</li>

    <li>{{obj.a}}</li>
    <li>{{obj.b}}</li>

    {% for v in arr%}
    <li>{{ v }}</li>
    {% endfor %}
</ul>
<go-footer></go-footer>
```
#### go-header.html
```html
<h1 id="title">{{ title }}</h1>
```

#### go-footer.html
```html
<hr>
<footer>
    Go-Footer: @2021 {{ mcss.app }}
</footer>
```

## 最终生成
```html 
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>mcss application</title>
</head>
<body>
<div style="text-align: center">
    <style>
    li {
        text-align: left;
    }
    </style>
    <h1 id="title">Go-Header - mcss application</h1>
    <ul>
        <li>render: object and array</li>
        <li>obj-1</li>
        <li>obj-2</li>
        <li>arr-1</li>   
        <li>arr-2</li>
        <li>arr-3</li>
    </ul>
    <hr>
    <footer>
        Go-Footer: @2021 mcss application
    </footer>
</div>
<script>
    let color = "red"
    setInterval(function () {
        color = color === "red" ? "black" : "red"
        document.getElementById('title').style.color = color
    }, 1000*2)
</script>
</body>
</html>
```