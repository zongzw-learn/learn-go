# kqueue在golang语言下的使用实践

将kqueue的操作细节封装在NetPoller接口中，实现KqueuePoller的三个API：

* Start 启动基于kqueue的IO多路复用事件监听
* Close 停止kqueue
* SetHandler 设置可插入式的数据处理函数，目前只处理socket读取操作。亦可以根据需要添加写操作

使用方法参考 [main函数](./main.go)

## 原理

syscall + kqueue + kevent + socket API

## TODO

亦可以通过golang条件编译支持 linux环境下的epoll