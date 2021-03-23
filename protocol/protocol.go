/*
** Package protocol 项目私有协议
 */
package protocol

import (
	//"github.com/goinggo/mapstructure" 使用json自带的库即可

	"encoding/json"
	"sync"
)

/*
P开头的是前后端协议用到的结构体proto
使用了gin自带的json库解析，不要随意改动元信息的标签字段
*/

// COMST 公共的Struct类型，包含非常基础的函数
type COMST struct {
}

// String 输出json格式的字符串
// 必须传入子类结构体p,不然编码出来的是父类(COMST)的“空”字符串{}
func (m *COMST) String(p interface{}) string {
	bytes, _ := json.Marshal(p)
	return string(bytes)
}

// P开头的常量
const (
	/******** 协议号 ********
	*sum=main*10000+sub
	*main[1,99]服务器内部使用（系统级）,客户端那边也私用了这个头
	*100：前后端约定的系统级消息
	*200: 具体游戏指令
	 */

	/****************************************************************/
	/****************************************************************/
	//main:1 系统级-组件的消息
	// PIDTICKSECEVENT 一秒的ticker
	PIDTICKSECEVENT = 10010
	// PIDTICKMILSECEVENT 100毫秒的ticker
	PIDTICKMILSECEVENT = 10011
)

// PCOMHead 公共头
// 需满足 api.IProtoHead 接口要求
type PCOMHead struct {
	// 公用结构体头
	COMST

	// 协议自带的protoID
	MainID int `json:"main"`
	SubID  int `json:"sub"`

	// 自定义结构体 [[----------------
	// 新鲜度-确保一条消息被有效处理的次数
	freshLock sync.RWMutex
	fresh     int // init:单播消息为1,群播或广播则为对应值

	// // 用于对广播类型的消息进行写保护
	// // 填充好广播消息后关闭锁,之后再广播出去，在后续的流程中，禁止修改此结构体的负载域
	// // 起到所有其它收到广播的mgr使用的负载数据是相同的效果
	// writeLocked bool

	// 来源
	SrcMod uint64 `json:"ssm"` // json字段仅供服务器内部使用
	SrcSub uint64 `json:"ssb"`

	// 指定目的地
	DstMod uint64 `json:"sdm"`
	DstSub uint64 `json:"sds"`

	// 拓展字段-通用
	extend string

	// 附加信息
	account string // 客户端账号
	version string // 客户端版本号-某些逻辑对于客户端版本是有要求的(客户端目前不能热更，所以服务器必须服务不同版本的客户端)

	// ----------------]]

	// 具体协议对应负载内容
	// Data {} `json:"data"`
}

// // AddWriteLock 给携带的负载域加上写锁
// func (m *PCOMHead) AddWriteLock() {
// 	m.writeLocked = true
// }

// // OpenWriteLock 尝试打开写保护，写入数据
// func (m *PCOMHead) OpenWriteLock() {
// 	if false == m.writeLocked {
// 		return
// 	}

// 	// 数据被保护起来了，不能被修改了，这里直接报错，便于早点发现问题
// 	panic(fmt.Sprintf("invalid write data for bc'msg(%d)", m.GetProtoID()))
// }

// SetProtoID 设置总协议号
func (m *PCOMHead) SetProtoID(id int) {
	m.MainID, m.SubID = id/10000, id%10000
}

// GetProtoID 获取协议号
func (m *PCOMHead) GetProtoID() int {
	return (m.MainID*10000 + m.SubID)
}

// SetFresh 设置初始新鲜度
func (m *PCOMHead) SetFresh(ttl int) {
	m.freshLock.Lock()
	defer m.freshLock.Unlock()
	m.fresh = ttl
}

// CostFresh 尝试扣除一次新鲜度
func (m *PCOMHead) CostFresh() bool {
	m.freshLock.Lock()
	defer m.freshLock.Unlock()
	// 判断失败，返回false
	if m.fresh < 1 {
		return false
	}

	// 扣除一次次数，返回成功
	m.fresh--
	return true
}

// CheckFresh 检测新鲜度
func (m *PCOMHead) CheckFresh() bool {
	m.freshLock.Lock()
	defer m.freshLock.Unlock()
	return (m.fresh > 0)
}

// SetSrc 设置消息来源
func (m *PCOMHead) SetSrc(mod, sub uint64) {
	m.SrcMod, m.SrcSub = mod, sub
}

// SetSrcMod 设置mod
func (m *PCOMHead) SetSrcMod(mod uint64) {
	m.SrcMod = mod
}

// SetSrcSub 设置sub
func (m *PCOMHead) SetSrcSub(sub uint64) {
	m.SrcSub = sub
}

// GetSrc 获取消息源-模块
func (m *PCOMHead) GetSrc() (uint64, uint64) {
	return m.SrcMod, m.SrcSub
}

// GetSrcMod 获取消息源
func (m *PCOMHead) GetSrcMod() uint64 {
	return m.SrcMod
}

// GetSrcSub 获取消息源
func (m *PCOMHead) GetSrcSub() uint64 {
	return m.SrcSub
}

// SetDst 设置消息目的地
func (m *PCOMHead) SetDst(mod, sub uint64) {
	m.DstMod, m.DstSub = mod, sub
}

// SetDstMod 设置消息目的地
func (m *PCOMHead) SetDstMod(mod uint64) {
	m.DstMod = mod
}

// SetDstSub 设置消息目的地
func (m *PCOMHead) SetDstSub(sub uint64) {
	m.DstSub = sub
}

// GetDst 设置消息目的地
func (m *PCOMHead) GetDst() (uint64, uint64) {
	return m.DstMod, m.DstSub
}

// GetDstMod 获取消息目的地
func (m *PCOMHead) GetDstMod() uint64 {
	return m.DstMod
}

// GetDstSub 设置消息目的地
func (m *PCOMHead) GetDstSub() uint64 {
	return m.DstSub
}

// PCOMLoad 格式化的Load字段
type PCOMLoad struct {
}
