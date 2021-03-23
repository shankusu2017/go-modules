/*Package protocol 项目私有协议
** version:
**			v0.1
** note:
**   json.Unmarshal 只能将[]byte解析成map[string]interface{},mapstructure.Decode再将map[string]interface{}
**		解析成具体的结构体,目前还不支持结构体嵌套，原来id->struct的方案还达不到预期，先放着后面再处理
**/

package protocol

import (
	"encoding/json"
	"fmt"

	"hippogames.com.cn/common/api"
	log "hippogames.com.cn/common/note"
)

// New 构造一个结构体并返回其"指针"
func New(id int) api.IProtoHead {
	var msg api.IProtoHead
	switch id {
	case PIDTICKSECEVENT:
		msg = new(PTickSecEvent)
	case PIDTICKMILSECEVENT:
		msg = new(PTickMilsecEvent)
	default:
		log.Panic("inavlid proto(%d)", id)
	}
	msg.SetProtoID(id)
	msg.SetFresh(1)
	return msg
}

// Parse 根据协议ID解析到具体的协议结构体
func (m *PCOMHead) Parse(msg []byte) api.IProtoHead {
	id := m.GetProtoID()
	switch id {
	case PIDTICKSECEVENT:
		var detailMsg PTickSecEvent
		json.Unmarshal(msg, &detailMsg)
		return &detailMsg
	case PIDTICKMILSECEVENT:
		var detailMsg PTickMilsecEvent
		json.Unmarshal(msg, &detailMsg)
		return &detailMsg
	default:
		log.Err("Err invalid protoid(%d)", id)
		return nil
	}
}

// ID2Name 协议ID->协议名称
func ID2Name(id int) string {
	switch id {
	case PIDTICKSECEVENT:
		return "tick-sec"
	case PIDTICKMILSECEVENT:
		return "tick.milsec"
	default:
		return fmt.Sprintf("uknow.proto(%d)", id)
	}
}
