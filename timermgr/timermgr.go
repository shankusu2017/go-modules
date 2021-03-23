/*Package timermgr time管理模块
 */
package timermgr

import (
	"time"

	"hippogames.com.cn/common/api"
	"hippogames.com.cn/common/define"
	"hippogames.com.cn/common/protocol"
	"hippogames.com.cn/common/router"
	comTime "hippogames.com.cn/common/time"
)

func init() {
	go ticker()
}

// 简易的定时器
func ticker() {
	tSec := time.NewTimer(time.Duration(time.Second * 1))
	tMilSec := time.NewTimer(time.Duration(time.Millisecond * 100))
	defer tSec.Stop()
	defer tMilSec.Stop()
	var msg api.IProtoHead

	for {
		select { // 生成滴答信息
		case <-tSec.C:
			tSec.Reset(time.Second * 1)
			data := protocol.New(protocol.PIDTICKSECEVENT).(*protocol.PTickSecEvent)
			data.SetSec(comTime.GetSec())
			msg = data
		case <-tMilSec.C:
			tMilSec.Reset(time.Millisecond * 100)
			data := protocol.New(protocol.PIDTICKMILSECEVENT).(*protocol.PTickMilsecEvent)
			data.SetMilsec(comTime.GetMilSec())
			msg = data
		}

		// 广播
		msg.SetFresh(define.G1)
		router.Router(msg)
	}
}
