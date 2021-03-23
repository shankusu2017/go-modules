/*Package router 消息分发器
**
** Warn: 1.这里均是分发的指针，故而单播可以修改单播消息，但不能修改广播的消息
**         若后续要修改广播的消息则考虑不再传指针，而是传结构体的拷贝
**       2.由于同一个房间的消息具有前后的逻辑顺序，这里不用简单的启用多goroutinue来分发
**          否则会造成消息顺序的混乱
 */

package router

import (
	"hippogames.com.cn/common/api"
	"hippogames.com.cn/common/define"
	log "hippogames.com.cn/common/note"
	"hippogames.com.cn/common/protocol"
)

var (
	// 消息订阅者
	subMap = make(map[uint64]api.MsgHandler)
	msgCH  = make(chan api.IProtoHead, define.M1)
)

func init() {
	// 开启私有的goroutinue
	go loop()
}

// loop 转发消息专用goroutinue
func loop() {
	var msg api.IProtoHead
	var dstModID uint64
	for {
		msg = <-msgCH
		dstModID = msg.GetDstMod()
		if dstModID == define.COMMODIDBROADCAST { // 广播协议则循环转发
			for _, hdl := range subMap {
				if msg.CheckFresh() {
					hdl(msg)
				} else {
					log.Err("proto(%s) fresh is invalid for dstModID(%d)", protocol.ID2Name(msg.GetProtoID()), dstModID)
				}
			}
		} else { // 找到目的模块：转发一次即可
			hdl := subMap[dstModID]
			if nil == hdl {
				log.Err("can't find hdl for dstModID(%d)", dstModID)
			}
			if msg.CheckFresh() {
				hdl(msg)
			} else {
				log.Err("fresh is invalid for dstModID(%d)", dstModID)
			}
		}
	}
}

// Sub 请求订阅消息
func Sub(modID uint64, hdl api.MsgHandler) {
	if nil != subMap[modID] {
		log.Fatal("[router sub.fail] can't sub twices(%v)", modID)
	}
	subMap[modID] = hdl
}

// Router 请求转发消息
func Router(msg api.IProtoHead) {
	ttl := len(msgCH)
	if ttl >= 100 {
		log.Err("router'load too heavy(%d)", ttl)
	}
	msgCH <- msg
}
