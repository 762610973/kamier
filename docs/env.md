- linux: ubuntu:22.04
- 安装 podman
  - sudo apt update
  - sudo apt install podman
  - 配置下载镜像源: 阿里云
- podman 安装 MongoDB
  - 配置镜像源
    - `vim /etc/containers/registries.conf`
    - ```toml
		ualified-search-registries = ["docker.io","registry.access.redhat.com"]
		[[registry]]
		prefix = "docker.io"
		location = "9sj33cdj.mirror.aliyuncs.com"
		insecure = true
		```
  - podman pull docker.io/library/mongo:latest
  - mongodb启动
    - ```shell
      podman run -d -p 27017:27017 --name=xl_mongo  \
      -e MONGO_INITDB_ROOT_USERNAME=xl \
      -e MONGO_INITDB_ROOT_PASSWORD=1112 \
      mongo:latest
      ```

- 不使用容器安装MongoDB,需要配置一下账号密码
```shell
use admin

db.createUser(
    {
        user:'xl',
        pwd:'1112',
        roles:[
            {
                role:'root',
                db: 'admin'
            }
        ]
    }
  
```

	
