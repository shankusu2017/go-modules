// Package cfg 加载json等格的配置文件
package cfg

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

/******************************** JSON ********************************/
/******************************** JSON ********************************/

type jsonStruct struct {
}

// newJSONStruct comment
func newJSONStruct() *jsonStruct {
	return &jsonStruct{}
}

// loadJSON 通用的解析函数：path对应的文件->结构体
func (jst *jsonStruct) loadJSON(path string, v interface{}) error {
	// ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("read file fail, path:", path)
		return err
	}

	// 读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		log.Println("Unmarshal file fail, path:", path)
		return err
	}
	return err
}

// LoadJSONCfg 加载JSON格式的配置文件
func LoadJSONCfg(path string, v interface{}) (interface{}, error) {
	log.Printf("load %s config....\n", path)

	jsonParse := newJSONStruct()
	err := jsonParse.loadJSON(path, v)
	return v, err
}
