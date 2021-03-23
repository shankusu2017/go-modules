/*Package snowflake 包
** TODOBETTER某个毫秒内序列号消耗完毕时的sleep函数的时间可以优化
 */
package snowflake

import (
	"sync"
	"time"

	log "hippogames.com.cn/common/note"
)

var (
	worker      *COMIDWorker // id单例句柄
	mutexLocker *sync.Mutex
)

// COMIDWorker 雪花算法中的worker类型
type COMIDWorker struct {
	JobNumber     int // 等效于snowflake中的(DatacenterID||WorkerID)
	LastTimestamp int64
	Sequence      uint32
}

// Lock 加锁
func (m *COMIDWorker) Lock() {
	mutexLocker.Lock()
}

// Unlock 解锁
func (m *COMIDWorker) Unlock() {
	mutexLocker.Unlock()
}

func workID() int {
	DatacenterID := 1 << 5
	threaderID := 1 << 0
	id := DatacenterID | threaderID
	return id
}

// init 初始化
func init() {
	// 初始化worker
	worker = new(COMIDWorker)
	worker.JobNumber = workID()
	worker.LastTimestamp = 0
	worker.Sequence = 0

	log.Log("uuid.first(%d)", Next())
}

// getCurMilliseconds 获取特定的时间戳(单位：毫秒)
func getCurMilliseconds() int64 {
	return (int64)((time.Now().UnixNano() / 1000000) - 1577836800000) // unix.timestamp.millisecond
}

// getCurMicroseconds 获取特定的时间戳(单位：微妙)
func getCurMicroseconds() int64 {
	return (int64)((time.Now().UnixNano() / 1000) - 1577836800000000) // unix.timestamp.microsecond
}

// getNext 生成UUID-内部实现
func (m *COMIDWorker) getNext() uint64 {
	m.Lock()
	defer m.Unlock()

	var cur int64
	for {
		cur = getCurMilliseconds()
		if cur < m.LastTimestamp { // 预防闰秒等时间回滚
			log.Err("Clock moved backwards.Refusing to generate id for %d milliseconds", cur-m.LastTimestamp)
			time.Sleep(time.Duration(100) * time.Microsecond)
		} else if cur == m.LastTimestamp {
			// 本毫秒内的序号已经消耗完毕，等待进入下一个timestamp区间
			if m.Sequence >= 4095 { // 2^12
				diff := 1000 - getCurMicroseconds()%1000 + 7 // 直接休眠到下一个timestamp
				time.Sleep(time.Duration(diff) * time.Microsecond)
			} else { // 继续使用本timestamp内的下一个seq
				break
			}
		} else { // 进入了新的可用timestamp
			m.Sequence = 0 // 重置seq
			m.LastTimestamp = cur
			break
		}
	}

	m.Sequence++ // 消耗一枚本timestamp内的seq

	// 组装成UUID
	uuid := (uint64)((uint64)(m.LastTimestamp<<(10+12)) |
		(uint64)((m.JobNumber)<<12) |
		(uint64)(m.Sequence))

	return uuid
}

// Next 生成UUID，对外接口
func Next() uint64 {
	return worker.getNext()
}
