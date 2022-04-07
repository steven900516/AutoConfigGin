package tool

import (
	"OwnWeb/controller"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

var RegisterIOC = make(map[string]interface{})

func RegistRouter(app *gin.Engine, functionName string) {
	RegistMap()
	path, err := os.Getwd()
	if err != nil {
		panic("读取controller包下路径错误……")
	}
	path = path + "/controller/"
	files, err := ioutil.ReadDir(path)
	for _, file := range files {
		realName := string(file.Name()[:len(file.Name())-3])
		if v, ok := RegisterIOC[realName]; !ok {
			log.Fatal(realName + "装配失败……请先在tool包下中的RegistMap方法先行添加映射")
		} else {
			ref := reflect.ValueOf(v)
			routerMethod := ref.MethodByName(functionName)
			args := []reflect.Value{reflect.ValueOf(app)}
			routerMethod.Call(args)
		}
	}
}

func RegistMap() {
	RegisterIOC["HelloController"] = &controller.HelloController{}
	RegisterIOC["IndexController"] = &controller.IndexController{}
}
