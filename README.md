# nacos-cli

`nacos-cli`是一个命令行工具，用来代替nacos的图形界面操作。

![carbon](https://github.com/szpinc/nacos-cli/assets/19821378/2899922a-e7c7-402d-80d4-a6bb27912efc)



## 安装
### linux

**AMD64**

`curl -L -o /usr/local/bin/nacos-cli https://github.com/szpinc/nacos-cli/releases/download/v1.2/nacos-cli_linux_amd64`

**ARM64**

`curl -L -o /usr/local/bin/nacos-cli https://github.com/szpinc/nacos-cli/releases/download/v1.2/nacos-cli_linux_arm64`

## 使用

### 环境变量

以下是支持的环境变量列表：

| 环境变量 | 说明 | 是否必须 | 默认值 |
|---------|------|---------|--------|
| NACOS_ADDR | Nacos 服务器地址 | 是 | http://127.0.0.1:8848/nacos |
| NACOS_API_VERSION | API 版本 | 否 | v1 |
| NACOS_USER | 用户名 | 否 | - |
| NACOS_PASSWD | 密码 | 否 | - |
| NACOS_NAMESPACE | 命名空间 | 否 | - |
| NACOS_GROUP | 配置组 | 否 | DEFAULT_GROUP |

### 命令行参数与环境变量优先级

命令行参数的优先级高于环境变量。具体规则如下：

1. 用户名和密码：
   - 优先使用命令行参数 `-u/--username` 和 `-p/--password`
   - 如果命令行参数未提供，则使用环境变量 `NACOS_USER` 和 `NACOS_PASSWD`
   - 如果环境变量也未设置，则不使用认证

2. 命名空间和分组：
   - 优先使用命令行参数 `-n/--namespace` 和 `-g/--group`
   - 如果未提供，则使用环境变量 `NACOS_NAMESPACE` 和 `NACOS_GROUP`
     - 命名空间默认值: ''
     - 分组默认值：`DEFAULT_GROUP`

### 基本操作

**获取所有配置列表**

``` bash
nacos-cli get config -A
```

**获取指定配置**

``` bash
nacos-cli get config common.yaml -n PUBLIC -g DEFAULT_GROUP
```

**编辑配置**

``` bash
nacos-cli edit config common.yaml -n PUBLIC -g DEFAULT_GROUP
```

**从文件更新配置**

``` bash
nacos-cli apply -f common.yaml -n public -g DEFAULT_GROUP --id common.yaml
```

## 使用认证

如果您的 Nacos 服务器配置了认证，您可以通过以下方式提供凭据：

**通过命令行参数：**

```bash
nacos-cli get config common.yaml -n PUBLIC -g DEFAULT_GROUP -u your_username -p your_password
```
> **安全性**：直接在命令行中传递密码或者将密码存储在环境变量中存在安全风险，建议采用环境变量的方式。

**通过环境变量：**

```bash
export NACOS_USER="your_username"
export NACOS_PASSWD="your_password"

nacos-cli get config common.yaml -n PUBLIC -g DEFAULT_GROUP
```
