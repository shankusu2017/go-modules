/* Package define 全局用的产量，避免到处定义常量,公共定义放在hippogames.com.cn/common/define中
** E 错误码
** U 常量
** COM 通用的模块定义
** C CONFIG
** D DB相关的定义，和相应的表密切相关
** P：PROTOCOL，在对应的protocol文件中
** I: INTERFACE, 接口
** M：(Macro definition)宏定义
**    MS: MacroState  状态
**    MA: MacroAction 指令
 */
package define

// 模块ID
const (
	// COMMODIDBROADCAST 广播
	COMMODIDBROADCAST = 0
	// COMMODIDHTTPMGR http服务
	COMMODIDHTTPMGR
	// COMMODIDDBMGR dbmgr
	COMMODIDIOMGR
)
