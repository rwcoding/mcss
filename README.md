# 一个用Golang实现的Html组件编译工具

+ 生成 html 节点树
+ 组件命名请使用 **x-y** 的模式，如 **go-form**，文件命名：**go-form.html**，保持工程唯一
+ 支持组件嵌套
+ 支持组件变量传递、自定义全局变量（mcss）
+ 强大的、可定义的指令集
+ 自定义配置，参阅 **internal/config_sample.go**

# 它能做什么
> 如果使用传统前端技术，比如依赖 `jquery/bootstrap/layui` 等，同时又要在团队中实行 **前后端分离**，那么这将是一个有效的基础设施
#### 以前您的代码是这样的
```html 
<div class="mb-3">
  <label for="input1" class="form-label">Email address</label>
  <input type="text" class="form-control" id="input1" placeholder="name@example.com">
</div>
<div class="mb-3">
  <label for="textarea1" class="form-label">Example textarea</label>
  <textarea class="form-control" id="textarea1" rows="3"></textarea>
</div>
```
#### 您可以设置组件 `bs-form-item`
```html 
<div class="mb-3">
  <label for="{{ id }}" class="form-label">{{ title }}</label>
  {% if type == "textarea" %}
   <textarea class="form-control" id="{{ id }}" rows="3"></textarea>
  {% else %}
  <input type="text" class="form-control" id="{{ id }}" placeholder="{{ placeholder }}">
  {% endif %}
</div>
```
#### 现在您可以这样使用
```html 
<bs-form-item id="username" title="Your username">
<bs-form-item id="age" title="Your Age">
<bs-form-item id="email" title="Your email" placeholder="name@example.com">
<bs-form-item type="textarea" id="taid" title="Example textarea">
```

# 指令集
#### 指令
+ as 设置属性
+ ds 设置data属性
+ ts js模板自定义
+ ot 元素外部内容
+ in 元素内部内容
#### 配置文件中定义一些指令
```toml
[iset]
acl = [["ot", "{{ if verify('@v') }}", "{{ endif }}"]]
ext = [["ot", "🤣🤣🤣", "🌹🌹🌹"], ["in", "💖💖💖💖💖", "✨✨✨✨✨"], ["as", {class = "nc"}] ]
pop = ["as", {onclick = "alert('@v')"}]
```
#### 有一个元素
```html 
<h1 @ext @cds="user" @acl="title">Hello World</h1>
```
#### 结果
```html
{{ if  verify("title") }}
🤣🤣🤣<h1 class="user" data-name="xyz_user">💖💖💖💖💖Hello World✨✨✨✨✨</h1>🌹🌹🌹
{{ endif }}
```

## 测试
```shell
# html 测试请参阅 examples 目录
# 自定义配置可修改 mcss.yaml 测试
# http://127.0.0.1:8080

go run main.go
```

## Node
```
npm i -g @rwcoding/mcss
```