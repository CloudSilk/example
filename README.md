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

# 打开浏览器，输入http://localhost:48080/web/#/web/
```
## 截图

![1](/images/screen1.png)
![2](/images/screen2.png)
![3](/images/screen3.png)
![4](/images/screen4.png)

## 社区

如果微信群二维码过期，请添加社区助手的微信，备注云梭。

![微信群](/images/wechat.jpg)

![社区助手](/images/assistant.jpg)