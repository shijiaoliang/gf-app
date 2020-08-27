package main

import (
	"fmt"
	"github.com/siddontang/go-mysql/canal"
)

type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	fmt.Println(e)
	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}

func main() {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = "rm-2zejg15af6d9b2r9y.mysql.rds.aliyuncs.com:3306"
	cfg.User = "admin_root"
	cfg.Password = "Y0yeWjaAqkLZIb8xH9xt"
	cfg.Dump.ExecutionPath = ""

	cfg.IncludeTableRegex = []string{"edu_ceping\\.admin", "edu_ceping\\.app"}

	c, _ := canal.NewCanal(cfg)
	c.SetEventHandler(&MyEventHandler{})

	//startPos := mysql.Position{
	//	Name: "mysql-bin.002473",
	//	Pos:  uint32(1436044),
	//}
	startPos, err := c.GetMasterPos()

	err = c.RunFrom(startPos)
	if err != nil {
		fmt.Printf("start canal err %v", err)
	}
}
