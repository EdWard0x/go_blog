package initialize

import (
	"context"
	"os"
	"server/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func ConnectRedis() redis.Client {
	redisCfg := global.Config.Redis

	// 使用单节点配置创建 Redis 客户端
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Address,  // 设置 Redis 服务器地址
		Password: redisCfg.Password, // 设置 Redis 密码
		DB:       redisCfg.DB,       // 设置使用的数据库索引
	})

	// Ping Redis 服务器以检查连接是否正常
	_, err := client.Ping(ctx).Result()
	if err != nil {
		global.Log.Error("Failed to connect to Redis:", zap.Error(err))
		os.Exit(1)
	}

	return *client
}
