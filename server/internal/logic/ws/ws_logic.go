package ws

import (
	"fmt"
	"net/http"
	"server/internal/utils"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/tideland/golib/logger"
	"go.uber.org/zap"
)

type WsLogic interface {
	CreateWsConnection(w http.ResponseWriter, r *http.Request) error
	HeartBeatCheck()
	PushLoop()
}

type wsLogic struct {
	wsConn             *WsConnection
	AllWsUserConnInfos map[string]*WsConnection
	AddWsUserConnInfos chan *WsConnection
	DelWsUserConnInfos chan *WsConnection
}

// websocket结构体
type WsConnection struct {
	ConnectId   string
	Socket      *websocket.Conn
	Connection  bool
	Node        string
	ConnectTime int64
	Mx          sync.Mutex
	Messages    chan []byte
}

func NewWsLogic() WsLogic {
	return &wsLogic{
		AllWsUserConnInfos: make(map[string]*WsConnection),
		AddWsUserConnInfos: make(chan *WsConnection),
		DelWsUserConnInfos: make(chan *WsConnection),
	}
}

// 创建链接
func (w *wsLogic) CreateWsConnection(writer http.ResponseWriter, request *http.Request) error {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: time.Second * 5,
		CheckOrigin: func(r *http.Request) bool { // 允许ws跨域
			return true
		},
	}

	var connectId = uuid.New().String()
	c, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		fmt.Printf("CreateWsConnection error,err=%v", err)
		return err
	}

	ip, _ := utils.GetLocalIP()
	port := utils.GetPort()
	node := fmt.Sprintf("%s:%s", ip, port)

	wsConnction := &WsConnection{
		ConnectId:   connectId,
		Socket:      c,
		Connection:  true,
		Node:        node,
		ConnectTime: time.Now().Unix(),
		Mx:          sync.Mutex{},
		Messages:    make(chan []byte, 1024),
	}
	w.wsConn = wsConnction
	w.AddWsUserConnInfos <- wsConnction

	return nil
}

// 心跳检测
func (w *wsLogic) HeartBeatCheck() {
	defer func() {
		w.DelWsUserConnInfos <- w.wsConn
	}()
	for {
		select {
		// 1秒没有消息就检测是否断连
		case <-time.After(time.Millisecond * 10):
			if !w.IsConnect() {
				// logger.Logger.Warn("ping websocket user conn failed", zap.Int("user_id", w.UID), zap.String("user_conn_id", w.ID))
				return
			}
		}
	}
}

// 循环推送队列消息给用户
func (w *wsLogic) PushLoop() {
	// 监听推送消息
	for {
		// 判断是否断开链接
		if !w.IsConnect() {
			fmt.Printf("user websocket disconnect")
			break
		}

		msg, err := PopWsMsgFromQueue()
		if err != nil {
			time.Sleep(time.Second * 1)
			continue
		}

		msg.ConnId = w.ID

		// 向某个用户的所有链接同步推送消息
		go func(userId int, msg Msg) {
			var userConnIdList []WsUserConnInfo
			userConnIdList, err = GetAllUserInfoList(userId)
			if err != nil {
				logger.Logger.Warn("get user all websocket conn id failed", zap.Int("user_id", userId), zap.Any("msg", msg), zap.Error(err))
				return
			}

			for _, userConn := range userConnIdList {
				if userConn.Closed {
					continue
				}
				if _, ok := AllWsUserConnInfos[userConn.ID]; ok {
					go msg.PushMsg(userConn.ID)
				} else {
					go msg.PushMsgToOtherServer(userConn.Node, userConn.ID)
				}

				// 记录消息到延迟队列，判断用户是否收到消息ACK
				if msg.Retries > 0 {
					msg.ConnId = userConn.ID
					go msg.PushWsMsgToDelayQueue()
				}
			}

		}(w.UID, msg)

		logger.Logger.Info("send websocket msg success", zap.Int("user_id", w.UID), zap.Any("msg", msg), zap.String("user_conn_id", w.ID))
	}
}

func (w *wsLogic) IsConnect() bool {
	if w.wsConn.Connection == true {
		if err := w.ping(); err != nil {
			w.wsConn.Connection = false
		}
	}
	return w.wsConn.Connection
}

// 检查链接是否中断
func (w *wsLogic) ping() error {
	w.wsConn.Mx.Lock()
	defer w.wsConn.Mx.Unlock()
	return w.wsConn.Socket.WriteMessage(websocket.PingMessage, nil)
}
