package ws

import "github.com/go-redis/redis"

// 消息结构体
type Msg struct {
	ConnectId string      `json:"connection_id"`
	Content   interface{} `json:"content"`
}

// 获取消息事件内容
func popWsMsgFromQueue(r *redis.Client) (msg Msg, err error) {

	return
}
