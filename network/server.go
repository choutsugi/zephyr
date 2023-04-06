package network

type (
	StartHandler      func()
	StopHandler       func()
	ConnectHandler    func(sess Session)
	DisconnectHandler func(sess Session)
	ReceiveHandler    func(sess Session, msg []byte, msgType int)
)

type Server interface {
	Addr() string
	Start() error
	Stop() error
	Protocol() string
	OnStart(handler StartHandler)
	OnStop(handler StopHandler)
	OnConnect(handler ConnectHandler)
	OnDisconnect(handler DisconnectHandler)
	OnReceive(handler ReceiveHandler)
}
