package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func InitConfig(path string) string {
	var s string
	f,err := os.Open(path)

	if err == io.EOF{
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		b,_,err := r.ReadLine()
		if err != nil {
			if err == io.EOF{
				break
			}
			panic(err)
		}
		//去除单行属性两端的空格
		//s := strings.TrimSpace(string(b))
		s = strings.TrimSpace(string(b))
/*		//fmt.Println(s)

		//判断等号=在该行的位置
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		//取得等号左边的key值，判断是否为空
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}

		//取得等号右边的value值，判断是否为空
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		//这样就成功吧配置文件里的属性key=value对，成功载入到内存中c对象里
		myMap[key] = value*/
	}
	return s
}

/*func main() {
	configure := InitConfig("./config/db_configuration.txt")
	fmt.Println(configure)
}*/