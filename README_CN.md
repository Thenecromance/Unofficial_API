# Blizzard 游戏非官方 API

[English](README.md) | [中文](README_CN.md)

本项目是一个针对暴雪游戏的 **非官方 Go SDK 和代理服务器**。它提供了全面的类型安全结构体和方法来与暴雪 API 交互，并且包含一个基于 Gin 的 Web 服务器作为代理/网关。
> **悄悄的说**: 本项目最初旨在整合中国区未公开的 Web API，最终将其转换为相同的数据结构。

> **免责声明**: 本项目是一个非官方的 API 封装，与暴雪娱乐有限公司（Blizzard Entertainment, Inc.）无关，也未获得其认可、赞助或特别批准。

## 🚀 特性

- **多游戏支持**: 魔兽世界 (World of Warcraft), 魔兽世界怀旧服 (WoW Classic), 炉石传说 (Hearthstone), 暗黑破坏神 3 (Diablo 3), 以及 星际争霸 2 (StarCraft II).
- **类型安全 SDK**: 根据官方 API 文档自动生成的 Go 结构体。
- **代理服务器**: 内置 Gin 服务器 (`:80`)，暴露 REST 接口。
- **智能路由**: 统一的 API 调用接口。
- **可扩展**: 可自定义 Request 和 Logger 接口。

## 🎮 支持的游戏

- **魔兽世界** (正式服)
- **魔兽世界怀旧服**
- **炉石传说**
- **暗黑破坏神 3**
- **星际争霸 2**

## ⚠️ 重要提示

1. **Cookie 使用**: 某些端点（特别是玩家个人资料数据）可能需要设置 Cookie (例如 `CharacterProfile.go` 中的 `StringCharacterProfileSummary`) 才能绕过官方 API 限制。请确保遵守使用条款。
2. **Beta 状态**: 项目处于早期开发阶段。不同游戏的 API 方法可能会有所不同。

## 🛠 使用方法

### 1. 作为代理服务器运行
本项目包含一个包装了 API 调用的 Web 服务器。

```bash
go run main.go
# 服务器将在 80 端口启动
```

### 2. 作为 Go 库 (SDK) 使用

您可以直接导入生成的包，在您的 Go 应用程序中进行 API 调用。

#### 代码示例

```go
package main

import (
	"context"
	"fmt"
	// 导入特定游戏/服务的包
	"github.com/Thenecromance/BlizzardAPI/api/wow/DataService/Achievement"
)

func main() {
	ctx := context.Background()

	// 1. 定义请求参数
	req := &wow_Achievement.AchievementFields{
		AchievementId: 6,
	}

	// 2. 调用 API
	// 返回 (any, error)
	resp, err := wow_Achievement.Achievement(ctx, req)
	if err != nil {
		panic(err)
	}

	// 3. 类型断言以访问特定字段
	// 响应会自动反序列化为正确的模型
	model := resp.(*wow_Achievement.AchievementModel)
	fmt.Printf("成就名称: %s\n", model.Name)
}
```

## 🏗 项目结构

```
/github.com/Thenecromance/BlizzardAPI
├───api             // 自动生成的 API 客户端代码 (SDK)
│   ├───wow
│   ├───D3
│   ├───...
├───app             // Gin 服务器应用逻辑
├───routers         // 代理服务器的 HTTP 路由定义
├───bridge          // 核心基础设施 (HTTP 客户端实现)
├───global          // 全局常量和配置
├───Interface       // 接口定义 (Logger, Request)
├───internal        // 内部工具 (Token 管理等)
├───tools           // 代码生成工具
│   ├───updater     // 获取文档并生成 'api/' 代码的逻辑
└───utils           // 通用工具
```

## 方法命名约定

| 方法前缀 | 描述 | 返回类型 |
|---|---|---|
| **(无前缀)** | **推荐**。主要入口点。处理请求和反序列化。 | `interface{}` (指向模型的指针) |
| `String**` | 获取原始 JSON 字符串。 | `string` |
| `bridge**` | 内部逻辑（将 JSON 解码为结构体）。 | `interface{}` |
| `CNHook**` | 中国区逻辑的扩展点。 | `func` |

## 📦 依赖项

- **Web 框架**: `github.com/gin-gonic/gin`
- **JSON 处理**: `github.com/bytedance/sonic` (高性能)
- **日志记录**: `github.com/sirupsen/logrus`
- **配置**: `github.com/spf13/viper`
- **代码生成**: `github.com/twpayne/go-jsonstruct`

## 许可证

[MIT](LICENSE)
