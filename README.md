# bark-cli：一个用于 [Bark](https://github.com/Finb/Bark) 的命令行工具

基于 golang 实现的通知工具 bark 的命令行，目前支持两种子命令：

- `bark plain` 发送纯文本
- `bark url` 发送直达链接

可通过 `bark help` 进行详细查看

## 安装
- `git clone https://github.com/KushNee/bark-cli.git`
- `cd <path/to/bark-cli>`
- `go install`

## 配置
目前仅支持通过环境变量配置

- `barkApi` 本人的设备码 **必须**
- `barkServer` bark所在服务器，此为可选项，默认为 `https://api.day.app/`