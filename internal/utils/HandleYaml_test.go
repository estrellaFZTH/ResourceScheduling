package utils

import (
	"testing"
)

func TestConfigEngine_Load(t *testing.T) {
	config := &ConfigEngine{}
	err := config.Load("config.yaml")
	if err != nil {
		t.Fatalf("config load error:%v", err)
	}
	t.Logf("%v", config)
	res := config.Get("mysql.username.data")
	t.Log(res)
}

func TestConfigEngine_GetString(t *testing.T) {
	config := &ConfigEngine{}
	err := config.Load("config.yaml")
	if err != nil {
		t.Fatalf("config load error:%v", err)
	}
	t.Logf("%v", config)
	res := config.GetString("app")
	t.Log(res)
}

func TestConfigEngine_GetBool(t *testing.T) {
	config := &ConfigEngine{}
	err := config.Load("config.yaml")
	if err != nil {
		t.Fatalf("config load error:%v", err)
	}
	t.Logf("%v", config)
	res := config.GetString("enable")
	t.Log(res)
}

func TestConfigEngine_GetInt(t *testing.T) {
	config := &ConfigEngine{}
	err := config.Load("config.yaml")
	if err != nil {
		t.Fatalf("config load error:%v", err)
	}
	t.Logf("%v", config)
	res := config.GetInt("port")
	t.Log(res)
}

func TestConfigEngine_GetFloat64(t *testing.T) {
	config := &ConfigEngine{}
	err := config.Load("config.yaml")
	if err != nil {
		t.Fatalf("config load error:%v", err)
	}
	t.Logf("%v", config)
	res := config.GetString("timeout")
	t.Log(res)
}

// 关于获取结构体的时候需要注意的是，结构体的字段必须和配置文件中的字段名称是一致的
type MysqlConfig struct {
	UserName string
	PassWord string
	DataBase string
	Host     string
	Port     int
	TimeOut  float64
}

func TestConfigEngine_GetStruct(t *testing.T) {
	config := &ConfigEngine{}
	err := config.Load("config.yaml")
	if err != nil {
		t.Fatalf("config load error:%v", err)
	}
	t.Logf("%v", config)
	res := config.GetStruct("mysql", &MysqlConfig{})
	t.Log(res)
}

func TestConfigEngine_Get(t *testing.T) {
	config := &ConfigEngine{}
	err := config.Load("config.yaml")
	if err != nil {
		t.Fatalf("config load error:%v", err)
	}
	t.Logf("%v", config)
	res := config.Get("")
	t.Log(res)
}

func TestConfigEngine_SetField(t *testing.T) {
	config := &ConfigEngine{}
	err := config.Load("config.yaml")
	if err != nil {
		t.Fatalf("config load error:%v", err)
	}
	t.Logf("%v", config)

	res := config.setField(&Config{}, "mysql", "0.5")
	t.Log(res)
}

//mysql:
//	UserName:
//		Data: xy
//	PassWord: 123456
type Config struct {
	Mysql string
}

type User struct {
	Data string
}

type MySQL struct {
	UserName User
	PassWord string
}
