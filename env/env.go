//Package env 读取程序配置
package env

import (
	"os"
	"path/filepath"

	"hippogames.com.cn/common/cfg"
	"hippogames.com.cn/common/define"
	log "hippogames.com.cn/common/note"
)

var (
	srvCfg *define.CLogSrvCfg // 整体配置
)

// init 初始化httpserver包，绑定指定的接口
// 一个URL有一个专门的goroutine,不同的URL的goroutine互不干扰-goroutine自行解决同步问题
func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Log("program.abs.dir:%s", dir)

	srvCfg = new(define.CLogSrvCfg)

	path := "../cfg/srv.json"
	_, err = cfg.LoadJSONCfg(path, srvCfg)
	if nil != err {
		log.Panic("load cfg fail err(%s)", err.Error())
	}
}

// GetCfg 获取配置
func GetCfg() define.CLogSrvCfg {
	return *srvCfg
}
