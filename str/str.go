/*Package string
 */

package str

import (
	"encoding/json"
	"strconv"
)

// Str int->string
func Str(num int) string {
	return strconv.Itoa(num)
}

// ToInt string->int
func ToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

// ToInt string->int
func Int(str string) int {
	ret, err := strconv.Atoi(str)
	if nil != err {
		ret = 0
	}
	return ret
}

// JSON2Str 带json标签的结构体的相关域转string
func JSON2Str(js interface{}) string {
	by, _ := json.Marshal(js)
	content := string(by)
	return content
}
