package protocol

// PTickSecEvent PIDTICKSECEVENT 秒级的滴答
type PTickSecEvent struct {
	PCOMHead              // 通用头
	Data     PTickSecLoad `json:"data"`
}

// PTickSecLoad 具体信息
type PTickSecLoad struct {
	sec int64 // 当前UNIX时间戳
}

// SetSec 设置秒
func (m *PTickSecEvent) SetSec(sec int64) {
	m.Data.sec = sec
}

// GetSec 获取
func (m *PTickSecEvent) GetSec() int64 {
	return m.Data.sec
}

// PTickMilsecEvent PIDTICKMILSECEVENT 秒级的滴答
type PTickMilsecEvent struct {
	PCOMHead                   // 通用头
	Data     PTickerMilsecLoad `json:"data"`
}

// PTickerMilsecLoad 具体信息
type PTickerMilsecLoad struct {
	milsec int64 // 当前UNIX时间戳(毫秒)
}

// SetMilsec 设置毫秒
func (m *PTickMilsecEvent) SetMilsec(milsec int64) {
	m.Data.milsec = milsec
}

// GetMilsec 获取毫秒
func (m *PTickMilsecEvent) GetMilsec() int64 {
	return m.Data.milsec
}
