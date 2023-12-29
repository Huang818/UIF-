---
sidebar_position: 4
---

# UIF 设置

## 设置密码

:::tip
UIF 默认不需要密码。如果你将 API 暴露给公网（即将地址设置为 `0.0.0.0`），则任何人都可以访问，UIF 要求**必须**使用密码。
:::

当你设置公网可见后，首次运行时会在 UIF 目录下生成一个随机密码。请到目录找到名为 `uif_key.txt` 的文件，打开即可看到生成的随机密码。

请复制密码，然后到 [首页](https://uiforfreedom.github.io/#/home) 处，粘贴密码，然后点击右上角 `连接后端` 即可。

如果你要修改密码（不建议，密码应该越随机越好），直接修改 `uif_key.txt` 文件内容即可。

## 修改默认 Web 和 API 地址和端口

`Web` 默认监听 `127.0.0.1:9527`; `API` 默认监听 `127.0.0.1:9413`.

如果你需要改变端口，或者暴露到公网使用，请在 UIF 目录下添加文件名为 `uif_api_address.txt` 或 `uif_web_address.txt`；

将 IP 地址修改为 `0.0.0.0`，则意味着公网可见。

### 例子

例如：将 `uif_api_address.txt` 使用 3333 端口，并且公网可见：

```bash
0.0.0.0:3333
```

将 `uif_web_address.txt` 使用 1234 端口，并且公网可见：

```bash
0.0.0.0:1234
```