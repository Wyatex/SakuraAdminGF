# GoFrame Template For SingleRepo

# 目录说明

- app 主要的业务代码
  - system 系统领域代码
    - controller 控制器
    - route 路由
    - model 系统领域模型
    - service 系统领域逻辑处理
  - business 业务领域代码
    - 结构可以和系统领域类似 
  - router 路由管理器
  - model 公共模型
  - middleware 公共中间件
  - code 公告业务码
  - cmd 服务器启动入口
- api 接口定义
- lib 与业务无关的一些工具

business层如果子系统非常多，可以再按照子系统创建各自的目录，比如：

- business
  - customer
    - controller
    - service
    - route
    - model
  - xxxxxx 如果子系统内容非常少也可以直接放一个文件夹
    - controller.go
    - service.go
    - route.go
    - model.go

# 中间件
目前有：CORS（跨域）、HandleResponse（统一结构返回）、GToken三个中间件，

gtoken用于登录校验，默认绑定全部路由，不需要校验的路由添加到：
app-middleware-gtoken.go:authExcludePaths

规则说明："/xxx/a":指精确指定路由/xxx/a，"/xxx/*":指/xxx下所有路由

后续添加日志记录（OperationLog）、角色权限校验中间件（Casbin）中间件。

# 分支管理

- master 主分支
- develop 开发分支，用于合并特性分支
- feat_xxx 特性分支，用于特性开发

开发流程：如果有新的系统需要开发，从develop分支切换出新的分支，比如：
develop_customer，开发完成后提交pull request，请求合并到develop分支
然后管理者审核，再由管理者合并到master分支

修复分支：如果发现develop分支有bug，则切出新的fix分支，比如：fix_某处代码异常
修复完成后提交pull request，请求合并到develop

# 约定俗成
一些开上的约束，务必遵循！

## 后端
- HTTP方法：统一使用POST方法，前端调用接口更加方便，数据传输格式统一采用JSON
- modelName统一用大写开头的驼峰命名法比如：SysUser
- 路由统一使用小写开头的驼峰命名法比如：/sysUser
- 业务码，一个模块一般留100的范围，比如第一个模块是10001开始，第二个从10101开始
- 待添加

## 前端
- 待添加

## 数据库
- 用tiny(1)存放布尔，1表示true， 0表示false
- 待添加

# 注意事项
- 项目中的casbin在启动和修改用户、角色、权限时可能加载一次，但是服务器不能水平拓展