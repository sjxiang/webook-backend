

# 代码组织

data - biz     - controller
repo - usecase - service


cache storage model

model - service - handler


po - dto



ddd 就像我的裤衩子，都知道有，但谁也没见过






分层
    handler/controller
        接收请求发送响应，中间负责json 和对象之间转换
    
    service
        侧重业务逻辑，贫血

    domain
        侧重业务逻辑，充血
        
    repository



不同 xxo 之间拷贝

    copier


尽量不要搞很多 xxo，一般 DTO（数据传输对象）+ Entity（数据存储实体）就够了；xxo 不在于多，而在于保持风格一致性。


三层两个实体


go-zero 





Review 步骤
    1. 服务 API 客户端
    2. 服务实现端
        
        go.mod
        configuration
        DB Schema
        Model + Repository
        Service
        Controller
        Main App
        一个接口样例
