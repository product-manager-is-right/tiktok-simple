#  Tiktok (Simple version) 抖音基础版

**简介:**

本项目采用微服务架构，将项目服务端拆分为User模块和Video模块，并使用网关来实现网络互连。实现了抖音基础功能，包括但不限于：用户登录，视频上传，点赞关注，评论，聊天等功能

使用技术栈：hertz框架，Mysql数据库，jwt，Nacos，minIO文件系统，redis，rabbitMQ，ffmpeg，docker

## 目录

+ 项目演示
  + 配置要求
  + 演示界面
  + 演示视频
  + 项目部署
  
+ 项目设计

  + 文件目录

  + 项目架构设计

  + 数据库设计

  + 服务模块

    + 用户模块

    + 视频模块
    + 点赞模块
    + 关注模块
    + 评论模块
    + 聊天模块

  + 性能优化

    + 微服务架构

      + 网关

      + 服务注册
      + 服务发现

    + Redis架构设计

    + 消息中间件架构设计

    + 数据库安全以及索引优化

  + 技术总结

+ 未来展望

  + 完全分布式系统

+ 贡献者以及分工

+ 项目地址

+ 鸣谢

### 1 项目演示

#### 1.1 项目环境

+ go 1.19

+ MySQL数据库

+ 服务端搭建Nacos、Redis、RabbitMQ、ffmpeg环境

+ 配置文件服务器minio

#### 1.2 项目页面

+ 登录页面
      <img src="https://raw.githubusercontent.com/product-manager-is-right/tiktok-simple/master/resource/screenShot/%E7%99%BB%E5%BD%95%E9%A1%B5%E9%9D%A2.png" alt="Logo" width="200" height="400">



+ 视频页面

<img src="https://raw.githubusercontent.com/product-manager-is-right/tiktok-simple/master/resource/screenShot/%E8%A7%86%E9%A2%91%E6%96%87%E4%BB%B6.png" alt="Logo" width="200" height="400"/>

+ 视频上传

<img src="https://raw.githubusercontent.com/product-manager-is-right/tiktok-simple/master/resource/screenShot/%E5%8F%91%E5%B8%83%E8%A7%86%E9%A2%91.png" alt="Logo" width="200" height="400" />

+ 评论列表

<img src="https://raw.githubusercontent.com/product-manager-is-right/tiktok-simple/master/resource/screenShot/%E8%AF%84%E8%AE%BA%E5%88%97%E8%A1%A8.png" alt="Logo" width="200" height="400" />

+ 关注列表

<img src="https://raw.githubusercontent.com/product-manager-is-right/tiktok-simple/master/resource/screenShot/%E5%85%B3%E6%B3%A8%E5%88%97%E8%A1%A8.png" alt="Logo" width="200" height="400" />

+ 好友列表

<img src="https://raw.githubusercontent.com/product-manager-is-right/tiktok-simple/master/resource/screenShot/%E5%A5%BD%E5%8F%8B%E5%88%97%E8%A1%A8.png" alt="Logo" width="200" height="400" />

+ 聊天页面

<img src="https://raw.githubusercontent.com/product-manager-is-right/tiktok-simple/master/resource/screenShot/%E5%A5%BD%E5%8F%8B%E8%81%8A%E5%A4%A9.png" alt="Logo" width="200" height="400" />

#### 1.3 项目视频



#### 1.4 项目部署

项目采用docker进行部署，我们针对每个模块编写Dockerfile文件，基于Dockerfile文件可以构建每个微服务模块的镜像。
单个部署容器是比较繁琐的，因此我们采用docker-compose对容器进行编排部署，可以简化部署的工作。


### 2 项目设计

#### 2.1 文件目录

```txt
├─TIKTOK_Gateway//网关服务
│  ├─cache
│  │  └─naming
│  ├─cmd
│  ├─configs//配置文件
│  ├─log
│  ├─resolver//网关的均衡负载算法
│  ├─route//路由
│  └─test
├─TIKTOK_User//用户模块
│  ├─api
│  │  ├─handler//控制层
│  │  └─router//路由
│  ├─cache
│  │  └─naming
│  ├─cmd
│  ├─configs//User模块的配置文件
│  ├─dal
│  │  └─mysql//mysql连接
│  ├─log
│  ├─model//实体类
│  │  └─vo
│  ├─mw
│  │  ├─rabbitMQ//消息中间件
│  │  │  ├─consumer//消费者
│  │  │  └─producer//生产者
│  │  └─redis//redis启动
│  ├─resolver//服务发现
│  ├─service
│  │  └─serviceImpl//服务实现
│  ├─test//测试类
│  │  ├─cache
│  │  │  └─naming
│  │  └─log
├─TIKTOK_Video//视频模块，结构与User一致
│   ├─api
│   │  ├─handler
│   │  └─router
│   ├─cache
│   │  └─naming
│   ├─cmd
│   ├─configs
│   ├─dal
│   │  └─mysql
│   ├─log
│   ├─model
│   │  └─vo
│   ├─mw
│   │  └─rabbitMQ
│   │      ├─consumer
│   │      └─producer
│   ├─resolver
│   ├─resources
│   ├─service
│   │  └─ServiceImpl
│   ├─test
│   │  ├─cache
│   │  │  └─naming
│   │  └─log
│   └─util

```



#### 2.2 项目架构设计

***项目设计***

![项目设计图](https://raw.githubusercontent.com/product-manager-is-right/tiktok-simple/master/resource/%E5%90%8E%E7%AB%AF%E6%9E%B6%E6%9E%84%E5%9B%BE.png)

本项目采用hertz作为基础框架，我们通过业务的不同需求，将服务分为用户和视频两个模块，用户服务模块主要管理用户信息相关接口，如用户信息，关注，点赞，聊天等，视频模块主要管理视频相关信息等，如视频发布，评论等。我们再使用Nacos注册中心将不同的服务集群注册，再使用Gateway进行服务器进行统一的服务器接口管理。同样的数据库也被我们划分为两个部分以便服务的解耦，分别对应用户服务与视频服务。采用minio文件服务器进行视频和封面的存储，采用rabbitMQ对数据库的写入操作进行削峰，redis对读取数据进行缓存，减少数据库的查询压力。

#### 2.3 数据库安全以及设计

+ 阅读[数据库安全以及索引文档](https://zv0hnc8742.feishu.cn/docx/YtSsdvp7LoCFN3xp1OXckQYyn2c)获取详细信息


#### 2.4 服务模块

##### 用户模块的设计

+ 用户模块包括用户注册（/douyin/user/register/）、用户登录（/douyin/user/login/）和用户信息（/douyin/user/）三个接口。

+ 阅读[用户模块文档](https://nd8dqd1ncj.feishu.cn/docx/QeAVdIxa5oUpqzxIVV1cw4i8nUh)获取详细信息

##### 视频模块的设计

+ 视频模块包括视频Feed流获取（/douyin/feed/ ）、视频投稿（/douyin/publish/action/）和发布列表（/douyin/publish/list/ - 发布列表）。

+ 阅读[视频模块文档](https://nd8dqd1ncj.feishu.cn/docx/Qev3d1AceoBEw5xhmXrcRQ8rnEd)获取详细信息

##### 点赞模块的设计

+ 点赞模块包括点赞操作（/douyin/favorite/action/）和获取点赞列表（/douyin/favorite/list/）。
+ 阅读[点赞模块文档](https://zv0hnc8742.feishu.cn/docx/Cfk9daRdgo9tI9xZdddcDo5an2d)获取详细信息

##### 关注模块的设计

+ 关注模块包括关注（/douyin/relation/action/）、获取关注列表（/douyin/relatioin/follow/list/）、获取粉丝列表（/douyin/relation/follower/list/）。
+ 阅读[关注模块文档](https://mc8cteit07.feishu.cn/docx/OyFndy7MfoAEcwxfVv9cjUdXnig)获取详细设计

##### 评论模块的设计

+ 评论模块包括评论（/douyin/comment/action/）和查看评论（/douyin/comment/list/）。
+ 阅读[评论模块文档](https://hoyew2yzsw.feishu.cn/docx/LX3fdY3IdoB18DxO0OHcxcdRnUb) 获取详细设计。

##### 聊天模块的设计

+ 聊天模块获取好友列表，获取聊天记录，以及聊天消息的发送。
+ 阅读[聊天模块文档](https://nd8dqd1ncj.feishu.cn/docx/TYGldpnuMo9urFxc6h4cbzninOL) 获取详细设计。

#### 2.5 性能优化

##### 2.5.1 微服务架构

###### 网关

网关是介于客户端和服务器端之间的中间层，所有的外部请求都会先经过 网关这一层。在本项目中，我们使用网关对传入请求进行均衡负载，安全性的判断（限制上传视频大小）以及统一不同服务器的接口。

###### 服务注册、发现
我们使用nacos作为微服务注册，服务发现组件。为了提高系统的可用性，我们基于docker和docker-compose容器化部署了包含3个节点的nacos集群，为了给客户端提供统一的nacos访问地址，我们还部署了1个nginx节点，用作反向代理和负载均衡，客户端发送的请求会由nginx负载均衡到3个Nacos节点中的随机一个。

同时为了服务之间的能够相互通信我们新建了一下接口

```go
//查询video数据库的详细信息判断video数据库是否点赞
/douyin/favorite/IsFavor/
//将用户上传的视频存储在um数据库中
/douyin/publish/UserVideo/
//根据User端返回的videoid，查询video数据库的详细信息
/douyin/publish/GetVideos/
```

##### 2.5.2 Redis架构设计

![redis](https://raw.githubusercontent.com/product-manager-is-right/tiktok-simple/dev/resource/redis.png)

使用redis为喜欢视频列表，喜欢列表，关注列表以及评论列表缓存,程序会先查询redis当中是否有缓存，如果没有再从数据库中查询。在存入redis的时候我们会将列表json以便存储。

##### 2.5.3 消息中间件架构设计

***MQ架构设计***

##### ![rabbitMQ](https://raw.githubusercontent.com/product-manager-is-right/tiktok-simple/master/resource/rabbitMQ.png)

我们对关注模块，点赞模块和视频模块使用rabbitMQ进行了削峰操作。

+ 在关注模块中，我们采用了rabbitMQ的订阅模式队列，将关注操作和取关操作发布请求至follow交换机，交换机再通过传入的routingkey "add"和"delet"将消息推入不同的消息队列进行消费。
+ 在喜欢模块中我们采用了发布模式队列，将点赞与取消点赞操作请求发布至favor交换机，favor交换机再将队列中的消息同时发布给User Consumer和Video Consumer，分别存入User数据库中和Video数据库中
+ 在评论模块中，由于评论需要返回是否成功的信息，我们只能对删除操作进行削峰操作，所以这里采用简单队列模式，将发布请求直接发布至评论删除队列消费。

#### 2.6 技术总结

可以使用redis覆盖更多的列表比如发布列表，但是考虑到json化后，视频信息较多可能会造成大key的问题

minIO是一个分布式存储引擎，由于服务器数量限制，这里暂时以单机演示

### 3 未来展望

#### 3.1 完全分布式

我们在未来可以将每一块的服务彻底拆分开来，将其分为，喜欢，点赞，关注，再使用grpc进行服务之间的通信，实现服务之间的解耦。

可以更新补充网关部分的均衡负载算法，可以根据服务器的业务不同和请求不同，改变网关的均衡负载算法

#### 3.2 推送算法

本项目中我们并未实现推荐算法，但是根据用户刷视频的title和视频的标签我们可以实现视频推荐

### 4 贡献者以及分工

+ 宋孟欣（组长） ： 团队分工，架构设计，视频模块服务，服务发现，消息中间件设计。 联系方式：smxmorgab@gmail.com
+ 张志昊：微服务架构设计，数据库设计，用户模块，网关，服务注册，聊天模块，缓存设计
+ 秀耿：评论模块，服务发现，消息中间件设计
+ 刘厚飞：关注模块，网关，服务注册，聊天模块，项目部署
+ 曾帅淇：数据库安全以及索引优化，点赞模块，缓存实现
+ 杨光光：关注模块，缓存实现
+ 詹豪：关注模块，缓存实现

### 5 项目地址

需要先在安卓端安装app-release.apk软件，双击右下方“我”
软件地址：[⁣‌﻿⁤‌‍⁤‬‍⁣⁡‍‍⁣‍﻿‍‌⁣‬﻿⁤﻿⁣﻿⁢﻿⁣‌﻿⁤‍⁣‍﻿⁣﻿⁤‌⁡‬‌极简抖音App使用说明 - 青训营版 - 飞书云文档 (feishu.cn)](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)
输入一下地址：http://101.43.5.142:8070/

（只开放到2023.2.28日，服务器到期）

### 6 鸣谢

+ [字节跳动后端青训营](https://youthcamp.bytedance.com/)
