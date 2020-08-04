#### 日志收集项目

##### logAdmin模块：负责收集日志
##### logAgent模块：从etcd中获取日志收集信息，发送至nsq
##### logTransfer模块：nsq消费日志，发送至elasticsearch

##### 使用技术:
* nsq
* etcd
* elasticsearch

logagent
用etcd代替zookeeper做服务注册
用nsq代替kafka做消息队列
都使用docker部署运行调用