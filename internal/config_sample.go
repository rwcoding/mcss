package internal

// 自定义配置请在你的前端工程目录添加 mcss.yaml, 内容参阅 sample
// debug 是否显示完整的调试信息
// add 本地测试地址
// view 你的页面目录，默认 src
// tmp_path 构建临时目录，如 mcss.html.json 文件
// component 您的组件放置的目录，系统查找时会 逐一扫描
// void_tag 您的自定义的组件，如果不需要闭合，可以在这里配置
// mcss 系统全局变量，任何页面和组件都可以使用，如 {{ mcss.app }}
// script 运行外部命令，如 scss watch
//   _boot 随系统启动的外部命令

var sample = `
debug: true
addr: ":8080"
view: src
tmp_path: tmp
component:
  - src
void_tag:
  - go-button
mcss:
  app: mcss application
script:
  watch: npm start 
  _boot: watch

`
