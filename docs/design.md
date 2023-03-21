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
  - 删除节点: `DELETE /node/delete/id?id=`
  - 更新节点: `POST /node/update/id?id=`
## compute
> 分布式的计算服务
- 接口
  - 同步计算
  - 异步计算
  - 查询计算结果
- grpc
  - processServer
  - httpServer
  - nodeServer
## container
> 容器内执行计算方法
- struct
  - 标明节点信息: Site
  - 请求Server的客户端: Client
  - 数据: Data
    - 公共数据: PublicData
    - 私有数据: PrivateData