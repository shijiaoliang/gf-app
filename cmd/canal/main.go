/**
 *******************************************51cto********************************************
 * Copyright (c)  www.51cto.com
 * Created by 51canal.
 * User: shijl@51cto.com
 * Date: 2020/08/27
 * Time: 11:18
 ********************************************************************************************
 */

package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/mysql"
)

type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	g.Log().Info(e)
	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}

func main() {
	//cfg
	cfg := canal.NewDefaultConfig()
	cfg.Addr = g.Cfg().GetString("source.Addr")
	cfg.User = g.Cfg().GetString("source.User")
	cfg.Password = g.Cfg().GetString("source.Password")
	cfg.Dump.ExecutionPath = ""
	cfg.IncludeTableRegex = g.Cfg().GetStrings("source.IncludeTableRegex")

	//canal
	c, _ := canal.NewCanal(cfg)
	c.SetEventHandler(&MyEventHandler{})

	//startPos
	var startPos mysql.Position
	var err error
	startPosName := g.Cfg().GetString("source.startPos.Name")
	startPosPos := g.Cfg().GetUint32("source.startPos.Pos")
	if startPosName != "" && startPosPos > 0 {
		startPos = mysql.Position{
			Name: startPosName,
			Pos:  startPosPos,
		}
	} else {
		startPos, err = c.GetMasterPos()
		if err != nil {
			g.Log().Printf("getMasterPos err %v", err)
			return
		}
	}

	//run
	err = c.RunFrom(startPos)
	if err != nil {
		g.Log().Printf("start 51canal err %v", err)
	}
}
