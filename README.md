# Unofficial API for Blizzard Games

[English](README.md) | [中文](README_CN.md)

This project is an **Unofficial Go SDK and Proxy Server** for Blizzard games. It provides comprehensive type-safe structs and methods to interact with Blizzard APIs, and also includes a Gin-based web server that acts as a proxy/gateway.

> **Dark side**: This project was initially intended to integrate the unpublicized web API of the China region and eventually convert it into the same data structure.

> **Disclaimer**: This project is an unofficial API wrapper and is not affiliated with, endorsed, sponsored, or specifically approved by Blizzard Entertainment, Inc.


## Known issues
- All Search API is not working properly due to implementation.
- All China region APIs are temporarily removed.
- '/heartstone/cards' in routers will register duplicate route error.(I know how to fix this issue, but so far just use
  hand to delete them)
- '/heartstone/decks' in routers will register duplicate route error.(I know how to fix this issue, but so far just use
  hand to delete them)
- No Response Error handling system yet.
- Why some apis don't have model struct? Sadly, I'm living in China now, and many games I don't have account to test, so
  I just skip them.( like: wow/profileService, scII, etc.)

## 🚀 Features

- **Multi-Game Support**: World of Warcraft, WoW Classic, Hearthstone, Diablo 3, and StarCraft II.
- **Type-Safe SDK**: Auto-generated Go structs from official API documentation.
- **Proxy Server**: Built-in Gin server (`:80`) exposing REST endpoints.
- **Smart Routing**: Unified interface for API calls.
- **Extensible**: Customizable Request and Logger interfaces.

## 🎮 Supported Games

- **World of Warcraft** (Retail)
- **World of Warcraft Classic**
- **Hearthstone**
- **Diablo 3**
- **StarCraft II**

## ⚠️ Important Notice

1. **Cookie Usage**: Some endpoints (especially player profile data) may require a cookie setup (e.g. `StringCharacterProfileSummary` in `CharacterProfile.go`) to bypass official API restrictions. Please ensure compliance with Terms of Service.
2. **Beta Status**: The project is in early development. API methods may vary across games.

## 🛠 Usage

### 1. Run as Proxy Server
The project includes a web server that wraps the API calls.

```bash
go run main.go
# Server starts on port 80
```

### 2. Use as Go Library (SDK)

You can import the generated packages directly to make API calls in your Go application.

#### Code Snippet

```go
package main

import (
	"context"
	"fmt"
	// Import the specific game/service package
	"github.com/Thenecromance/BlizzardAPI/api/wow/DataService/Achievement"
)

func main() {
	ctx := context.Background()

	// 1. Define Request Parameters
	req := &wow_Achievement.AchievementFields{
		AchievementId: 6,
	}

	// 2. Call the API
	// Returns (any, error)
	resp, err := wow_Achievement.Achievement(ctx, req)
	if err != nil {
		panic(err)
	}

	// 3. Type Assert to access specific fields
	// The response is automatically unmarshalled into the correct model
	model := resp.(*wow_Achievement.AchievementModel)
	fmt.Printf("Achievement Name: %s\n", model.Name)
}
```

## 🏗 Project Structure

```
/github.com/Thenecromance/BlizzardAPI
├───api             // Auto-generated API client code (SDK)
│   ├───wow
│   ├───D3
│   ├───...
├───app             // Gin Server application logic
├───routers         // Http Routes definitions for the Proxy Server
├───bridge          // Core infrastructure (HTTP Client implementation)
├───global          // Global constants and configurations
├───Interface       // Interfaces (Logger, Request)
├───internal        // Internal utilities (Token management, etc.)
├───tools           // Code generation tools
│   ├───updater     // Logic to fetch docs and generate 'api/' code
└───utils           // General utilities
```

## method Convention

| Method Prefix | Description | Return Type |
|---|---|---|
| **(No Prefix)** | **Recommended**. Main entry point. Handles request & unmarshalling. | `interface{}` (Ptr to Model) |
| `String**` | Fetch raw JSON string. | `string` |
| `bridge**` | Internal logic (Decodes JSON to Struct). | `interface{}` |
| `CNHook**` | Extension point for China region logic. | `func` |

## 📦 Dependencies

- **Web Framework**: `github.com/gin-gonic/gin`
- **JSON Handling**: `github.com/bytedance/sonic` (High performance)
- **Logging**: `github.com/sirupsen/logrus`
- **Config**: `github.com/spf13/viper`
- **Code Gen**: `github.com/twpayne/go-jsonstruct`

## License

[MIT](LICENSE)
