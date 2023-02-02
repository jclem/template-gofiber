# jclem/template-gofiber

This is a template for a [GoFiber](https://gofiber.io/) application.

## Prerequisites

- [Go](https://golang.org/)
- [Air](https://github.com/cosmtrek/air)

## Usage

```shell
> git clone https://github.com/jclem/template-gofiber.git
> cd template-gofiber
> script/rename user\\/mymodule
> air
```

## Features

- [Viper](https://github.com/spf13/viper) for configuration
- [Zap](https://github.com/uber-go/zap) for logging
- Basic error handling
- Request logging
- Per-request context, including access to:
  - Fiber request context
  - Application configuration
  - Logger