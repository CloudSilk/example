# 轻舟低代码平台DEMO

## 环境要求：

nodejs>=18
golang>=1.20

## 源码运行

```bash
# 编译web
cd web && WEB_BASE=/web yarn build
# 运行server
SERVICE_MODE=ALL  go run main.go --ui ./web/dist --port 48080
```