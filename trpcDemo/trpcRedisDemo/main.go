// 声明当前文件属于 main 包
// Go 中可执行程序的入口包必须是 main 包
package main

// 导入依赖包
import (
	// fmt 是 Go 标准库，提供格式化 I/O 功能（如 Println、Printf 等）
	"fmt"

	// trpc-database/goredis/v3/redisex 是 tRPC 框架封装的 Redis 扩展客户端
	// redisex 提供了对 go-redis 的二次封装，支持 tRPC 的服务发现、配置管理等
	"git.woa.com/trpc-go/trpc-database/goredis/v3/redisex"

	// trpc-go 是腾讯 tRPC-Go 框架核心包，提供服务端、客户端、上下文等基础能力
	"git.code.oa.com/trpc-go/trpc-go"
	// log 是 tRPC 框架的日志组件，支持配置化的日志输出
	"git.code.oa.com/trpc-go/trpc-go/log"
)

// main 函数是程序执行的入口
func main() {
	// 创建一个 tRPC Server 实例
	// NewServer() 内部会读取 ./trpc_go.yaml 配置文件，初始化框架（包括日志、插件、客户端配置等）
	// 这里用 _ 丢弃返回值，是因为本示例不需要启动服务端，只需要触发框架初始化（加载配置）
	// 注意：如果当前目录下没有 trpc_go.yaml 文件，这里会 panic
	_ = trpc.NewServer()

	// 创建一个 redisex Redis 客户端实例
	// 参数 "trpc.example.test.redis" 是 service name（服务名）
	// 框架会根据这个 name 在 trpc_go.yaml 的 client.service 配置中查找对应的 Redis 连接配置（target、password 等）
	redisExclient, err := redisex.New("trpc.example.test.redis")
	// 如果创建客户端失败（如配置缺失、连接失败），打印 Fatal 日志并退出程序
	if err != nil {
		log.Fatalf("redis init failed: %v", err)
	}

	// 获取底层的原生 go-redis Client
	// redisex 包装了 go-redis，通过 .Client 字段可以拿到原生的 *redis.Client 来调用所有 Redis 命令
	client := redisExclient.Client

	// 定义要写入 Redis 的字符串值
	v1 := "v1"

	// 获取一个空的后台 Context，类似 context.Background()，但带有 tRPC 框架的上下文信息
	// 用于传递请求生命周期内的元数据（trace、超时等）
	ctx := trpc.BackgroundContext()

	// 定义 Redis Key
	key := "key_redis_universal_v1"

	// 执行 SET 命令：把 key 设置为 v1，过期时间 0 表示永不过期
	// .Err() 用于获取该命令的错误（如果有）
	err = client.Set(ctx, key, v1, 0).Err()
	// 如果 SET 失败，打印 Fatal 日志并退出
	if err != nil {
		log.Fatalf("set failed: %v", err)
	}

	// 执行 GET 命令：根据 key 获取对应的 value
	// .Result() 同时返回 value 和 error
	v, err := client.Get(ctx, key).Result()
	// 如果 GET 失败（如 key 不存在或网络错误），打印 Fatal 日志并退出
	if err != nil {
		log.Fatalf("get failed: %v", err)
	}

	// 打印从 Redis 读取出的 value，预期输出 "v1"
	fmt.Println(v)
}

// 以下是注释占位（TODO 标记），表示后续可能要补充的示例代码

// 原生 client 用例
// （未实现）：直接使用 go-redis 原生 client 操作 Redis

// cas client 用例
// （未实现）：使用 Compare-And-Swap（乐观锁）操作 Redis

// 分布式锁 client 用例
// （未实现）：使用 Redis 实现分布式锁（如基于 SETNX）

// 本地定时任务例子
// （未实现）：单机定时任务示例

// 分布式定时任务用例
// （未实现）：基于 Redis 的多实例分布式定时任务
