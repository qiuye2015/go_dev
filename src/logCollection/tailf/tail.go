package tailf

import (
	"fmt"
	"time"

	"github.com/beego/beego/core/logs"
	"github.com/hpcloud/tail"
)

type CollectConf struct {
	LogPath string
	Topic   string
}

type TailObj struct {
	tail *tail.Tail
	conf CollectConf
}
type TextMsg struct {
	Msg   string
	Topic string
}
type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan  chan *TextMsg
}

var (
	tailObjMgr *TailObjMgr
)

func InitTail(confs []CollectConf, chanSize int) (err error) {
	if len(confs) == 0 {
		err = fmt.Errorf("invalid config for log collect,conf:%v", confs)
		return
	}
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize),
	}
	for _, v := range confs {
		obj := &TailObj{
			conf: v,
		}

		tails, errT := tail.TailFile(v.LogPath, tail.Config{
			ReOpen: true,
			Follow: true,
			// Location: &SeekInfo{0, os.SEEK_END},
			MustExist: false,
			Poll:      true,
		})
		if errT != nil {
			logs.Error("tail file err:", errT)
			return
		}
		obj.tail = tails
		tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj)

		go readFromTail(obj)
	}
	return
}

func readFromTail(tailObj *TailObj) {
	for true {
		line, ok := <-tailObj.tail.Lines
		if !ok {
			logs.Warn("tail file close reopen, filename %s", tailObj.tail.Filename)
			time.Sleep(100 * time.Millisecond)
		}
		textMsg := &TextMsg{
			Msg:   line.Text,
			Topic: tailObj.conf.Topic,
		}
		tailObjMgr.msgChan <- textMsg
	}
}

func GetOneLine() (msg *TextMsg) {
	msg = <-tailObjMgr.msgChan
	return
}
