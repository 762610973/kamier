# 设计说明

## storage
> 中心化的存储服务
> 数据存储在MongoDB
- 存储计算函数
  - `name string`
  - `id   int64`
  - `description string`
- 存储公共数据
  - `name string`
  - `id int64`
  - `data []bytes`
  - `description string`
- 接口
  - GET ping
  - 创建计算函数: `POST /function/add`
  - 获取计算函数: `GET /function/get/id?id=`
  - 查询所有函数: `GET /function/getAllFunc`
  - 删除计算函数: `DELETE /function/delete/id?id=`
  - 更新计算函数: `PUT /function/update/id?id=`
  - ============================
  - 创建公共数据: `POST /data/add`
  - 查询公共数据: `GET /data/get/id?id=`
  - 查询所有数据: `GET /data/getAllData`
  - 删除公共数据: `DELETE /function/delete/id?id=`
  - 更新公共数据: `PUT /function/update/id?id=`
  - =============================
  - 注册节点: `POST /node/add`
  - 查询节点: `GET /node/get/id?id=`
  - 查询所有节点: `GET /node/getAllNode`
## compute
> 分布式的计算服务
- 接口
  - 同步计算
  - 异步计算
  - 查询计算结果
- grpc
  - containerServer
  - httpServer
  - nodeServer
- 执行过程分析
  - 参数校验: 参与计算的节点数量, 获取的计算函数是否存在等
  - 准备执行: 异步阻塞请求所有节点开始准备,其中一个节点准备失败,则整个计算失败
    - 参数校验,序列号准备
  - 启动执行: 异步非阻塞请求所有节点开始执行
  - 新建共识节点
    - 从端口池子中获取端口
  - 启动共识
  - 异步启动容器执行计算方法
  - 超时机制
## container
> 容器内执行计算方法
- struct
  - 标明节点信息: Site
  - 请求Server的客户端: Client
  - 数据: Data
    - 公共数据: PublicData
    - 私有数据: PrivateData
- 简要流程分析
  - 如果函数不发生在本节点，等待执行节点的complete
  - 如果函数发生在本节点,本节点放置完成标记(pid),获取完成标记(本节点)
- 