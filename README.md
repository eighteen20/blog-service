# blog-service

### 基于gin框架实现的一个简单博客后端服务

#### Go version： v1.16

### 目录结构

├─configs ---------- 配置文件  
├─docs ------------- 文档<br>
├─global ----------- 全局变量<br>
├─internal --------- 内部模块<br>
│  ├─dao ----------- 数据访问层<br>
│  ├─middleware ---- HTTP中间件<br>
│  ├─model --------- 模型层<br>
│  ├─routers ------- 路由相关的逻辑<br>
│  └─service ------- 业务逻辑<br>
├─pkg -------------- 项目相关模块包<br>
├─scripts ---------- 项目生成的临时文件<br>
├─storage ---------- 各类脚本<br>
└─third_party ------ 第三方资源工具