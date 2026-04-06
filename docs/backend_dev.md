
API 接口文档：https://openbridge.apifox.cn

卢宇扬、郑源羽

#### 2026.04.06

按照项目规范搭建了后端的基本框架，详细如下。

```
backend/
├── main.go             # 程序入口文件
├── go.mod              
├── go.sum              
├── openbridge.db       # SQLite 数据库文件
├── internal/           # 内部业务代码
│   ├── domain/         # 领域层 - 定义核心业务模型和接口
│   ├── handler/        # 表现层 - HTTP 处理器，处理请求和响应
│   ├── repository/     # 数据访问层 - 数据库操作实现
│   ├── usecase/        # 应用层 - 业务逻辑实现
│   └── tool/           # 工具层 - 通用工具函数
└── pkg/                # 可被外部引用的公共包
    └── db/             # 数据库相关公共包
```

实现了以下两个提供给前端的接口，这里只做简要介绍，具体查看 API 接口文档。

- 上传AccessToken

由于大多数网盘都需要access token进行身份验证，因此需要将用户的access token上传到服务器，以便服务器可以代表用户进行操作。服务器将access token存储在数据库中，以便在需要时使用。

- 上传RefreshToken

由于access token通常具有较短的有效期，因此需要定期刷新access token。服务器将refresh token存储在数据库中，以便在需要时使用。

下一次开发任务：实现获取容量相关的接口。