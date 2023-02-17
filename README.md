#  Tiktok (Simple version) 抖音基础版

**简介:**

本项目采用微服务架构，将项目服务端拆分为User模块和Video模块，并使用网关来实现网络互连。实现了抖音基础功能，包括但不限于：用户登录，视频上传，点赞关注，评论，聊天等功能

使用技术栈：hertz框架，Mysql数据库，jwt，Nacos，minio文件系统，redis，rabbitMQ，ffmpeg

## 目录

+ 项目演示 

  + 配置要求
  + 演示界面
  + 演示视频

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

+ 鸣谢

### 1 项目演示

#### 1.1 项目环境

+ go 1.19

+ MySQL数据库

+ 服务端搭建Nacos、Redis、RabbitMQ、ffmpeg环境

+ 配置文件服务器minio

#### 1.2 项目页面

+ 登录页面



+ 视频页面



+ 视频上传



+ 评论列表



+ 关注列表



+ 好友列表



+ 聊天页面



#### 1.3 项目视频



### 2 项目设计

#### 2.1 文件目录

```txt
├─TIKTOK_Gateway
│  ├─cache
│  │  └─naming
│  ├─cmd
│  ├─configs
│  ├─log
│  ├─resolver
│  ├─route
│  └─test
├─TIKTOK_User
│  ├─api
│  │  ├─handler
│  │  └─router
│  ├─cache
│  │  └─naming
│  ├─cmd
│  ├─configs
│  ├─dal
│  │  └─mysql
│  ├─log
│  ├─model
│  │  └─vo
│  ├─mw
│  │  ├─rabbitMQ
│  │  │  ├─consumer
│  │  │  └─producer
│  │  └─redis
│  ├─resolver
│  ├─service
│  │  └─serviceImpl
│  ├─test
│  │  ├─cache
│  │  │  └─naming
│  │  └─log
    ├─resolver
    ├─resources
    ├─service
    │  └─ServiceImpl
    ├─test
    │  ├─cache
    │  │  └─naming
    │  └─log
    └─util

```



#### 2.2 项目架构设计

***项目设计***

![项目设计图](https://raw.githubusercontent.com/product-manager-is-right/tiktok-simple/master/resource/%E5%90%8E%E7%AB%AF%E6%9E%B6%E6%9E%84%E5%9B%BE.png)

本项目采用hertz作为基础框架，我们通过业务的不同需求，将服务分为用户和视频两个模块，用户服务模块主要管理用户信息相关接口，如用户信息，关注，点赞，聊天等，视频模块主要管理视频相关信息等，如视频发布，评论等。我们再使用Nacos注册中心将不同的服务集群注册，再使用Gateway进行服务器进行统一的服务器接口管理。同样的数据库也被我们划分为两个部分以便服务的解耦，分别对应用户服务与视频服务。采用minio文件服务器进行视频和封面的存储，采用rabbitMQ对数据库的写入操作进行削峰，redis对读取数据进行缓存，减少数据库的查询压力。

#### 2.3 数据库设计

***数据库设计图***



#### 2.4 服务模块

##### 用户模块的设计

+ 用户模块包括用户注册（/douyin/user/register/）、用户登录（/douyin/user/login/）和用户信息（/douyin/user/）三个接口。

+ 阅读[用户模块文档](**文档地址**)获取详细信息

##### 视频模块的设计

+ 视频模块包括视频Feed流获取（/douyin/feed/ ）、视频投稿（/douyin/publish/action/）和发布列表（/douyin/publish/list/ - 发布列表）。

+ 阅读[视频模块文档](https://nd8dqd1ncj.feishu.cn/docx/Qev3d1AceoBEw5xhmXrcRQ8rnEd)获取详细信息

##### 点赞模块的设计

+ 点赞模块包括点赞操作（/douyin/favorite/action/）和获取点赞列表（/douyin/favorite/list/）。
+ 阅读[点赞模块文档](**文档地址**)获取详细信息

##### 关注模块的设计

+ 关注模块包括关注（/douyin/relation/action/）、获取关注列表（/douyin/relatioin/follow/list/）、获取粉丝列表（/douyin/relation/follower/list/）。
+ 阅读[关注模块文档](**文档地址**)获取详细设计

##### 评论模块的设计

+ 评论模块包括评论（/douyin/comment/action/）和查看评论（/douyin/comment/list/）。
+ 阅读[评论模块文档]() 获取详细设计。

##### 聊天模块的设计

+ 聊天模块包括发表评论、删除评论和查看评论。
+ 阅读[聊天模块文档]() 获取详细设计。

#### 2.5 性能优化

##### 2.5.1 微服务架构

###### 网关

###### 服务发现

###### 服务注册

##### 2.5.2 Redis架构设计



##### 2.5.3 消息中间件架构设计

***MQ架构设计***

##### ![rabbitMQ](https://raw.githubusercontent.com/product-manager-is-right/tiktok-simple/master/resource/rabbitMQ.png)

我们对关注模块，点赞模块和视频模块使用rabbitMQ进行了削峰操作。

+ 在关注模块中，我们采用了rabbitMQ的订阅模式队列，将关注操作和取关操作发布请求至follow交换机，交换机再通过传入的routingkey "add"和"delet"将消息推入不同的消息队列进行消费。
+ 在喜欢模块中我们采用了发布模式队列，将点赞与取消点赞操作请求发布至favor交换机，favor交换机再将队列中的消息同时发布给User Consumer和Video Consumer，分别存入User数据库中和Video数据库中
+ 在评论模块中，由于评论需要返回是否成功的信息，我们只能对删除操作进行削峰操作，所以这里采用简单队列模式，将发布请求直接发布至评论删除队列消费。

##### 2.5.4 数据库索引优化



#### 2.6 技术总结



### 3 未来展望



### 4 贡献者以及分工

+ 宋孟欣（组长） ： 团队分工，视频模块服务，服务发现，消息中间件设计。 联系方式：smxmorgab@gmail.com
+ 张志昊：微服务架构设计，数据库设计，用户模块，网关，服务注册，聊天模块
+ 秀耿：评论模块，服务发现，消息中间件设计
+ 刘厚飞：关注模块，网关，服务注册，聊天模块
+ 曾帅淇：数据库安全以及索引优化，点赞模块，redis缓存设计
+ 杨光光：关注模块，redis缓存设计
+ 詹豪：关注模块，redis缓存设计

### 5 鸣谢

+ [字节跳动后端青训营](https://youthcamp.bytedance.com/)
