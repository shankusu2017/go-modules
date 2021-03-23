//Package api 接口定义
package api

// IStringer 将自己的内容格式化成stirng
type IStringer interface {
	// 返回符合自己格式的string
	String(v interface{}) (str string)
}

// IProtoHead 协议消息必备的接口
type IProtoHead interface {
	SetProtoID(id int) // 消息号
	GetProtoID() int

	// 新鲜度：一条消息能被消费的次数
	SetFresh(int)
	CostFresh() bool
	CheckFresh() bool // 检测是否还新鲜

	SetSrc(mod, sub uint64) // 消息来源
	SetSrcMod(uint64)
	SetSrcSub(uint64)
	GetSrc() (mod, sub uint64)
	GetSrcMod() uint64
	GetSrcSub() uint64

	SetDst(mod, sub uint64) // 消息来源
	SetDstMod(uint64)
	SetDstSub(uint64)
	GetDst() (mod, sub uint64)
	GetDstSub() uint64
	GetDstMod() uint64
}

// MsgHandler 处理请求的函数句柄
type MsgHandler func(msg IProtoHead)
