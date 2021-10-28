package internal

// 自定义配置请在你的前端工程目录添加 mcss.yaml, 内容参阅 sample
// debug 是否显示完整的调试信息
// add 本地测试地址
// view 你的页面目录，可配置多个，默认 src
// tmp_path 构建临时目录，如 mcss.html.json 文件
// component 您的组件放置的目录，系统查找时会 逐一扫描
// void_tag 您的自定义的组件，如果不需要闭合，可以在这里配置
// mcss 系统全局变量，任何页面和组件都可以使用，如 {{ mcss.app }}
// iset 指令集，可多个，用 || 分隔，其中  @v 会分别替换为 值， 所有键名必须以 @ 符号开头
//   语法：指令|参数 || 指令|参数
//    - as: 设置元素属性，参数形如：参数名:值， 如： class:@v， 默认可以省略 :@v
//    - ds: 设置元素 data 属性
//    - ts: 声明JS模板块, 参数形如：模板开始块|结束块
//    - ot: 元素外部首尾声明语句，参数形如：开始内容|结束内容，其中内容可以只设置一个
//    - in: 元素内部首尾声明语句，和 ht 对应，参数形如：开始内容|结束内容，其中内容可以只设置一个
//   例：
//    - 配置：@eg: as|class || as|name:@v || ds|name || ds|title:hello
//    - From: <div class="c1" @eg="c2"></div>
//    - To:   <div class="c1 c2" name="c2" data-name="c2" data-title="hello"></div>
//
//   例：
//    - 配置：@eg: ts|{{ if @v }}|{{ endif }}
//    - From: <div @eg="user"></div>
//    - To:   {{ if user }} <div></div> {{ endif }}
//
//   例：
//    - 配置：@eg: ot|<label id="@v">|</label>
//    - From: <input @eg="doc">
//    - To:   <label id="doc"><input></label>
//
// script 运行外部命令，如 scss watch
//   _boot 随系统启动的外部命令

var sample = `
debug = false
addr = ":8080"
view = ["src"]
tmp_path = ["tmp"]
component = ["src"]
void_tag = ["go-button"]
[mcss]
app = "mcss application"
`
