# common-go

通用 Go 基础库，适用于微服务/多项目场景，封装常用的响应结构体、错误处理、日志、配置等模块，便于各后端服务复用。

## 目录结构示例

```
common-go/
├── go.mod
├── pkg/
│   ├── config/
│   ├── logger/
│   ├── middleware/
│   ├── db/
│   ├── utils/
│   └── response/
└── README.md
```

> 所有通用模块均放在 `pkg/` 目录下，便于统一管理和扩展。

## 已有模块
- `response`：统一 API 响应结构体与通用 Success/Error 方法（位于 pkg/response）

## 如何在其他 Go 项目中引用

1. 在你的项目目录下执行：
   ```sh
   go get github.com/francislyj/common-go@latest
   # 或指定 tag
   # go get github.com/francislyj/common-go@v0.1.0
   ```
2. 在代码中引用：
   ```go
   import "github.com/francislyj/common-go/pkg/response"
   ```

## 贡献
欢迎补充更多通用模块，如日志、配置、工具函数等。 