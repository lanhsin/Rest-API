package main

import (
	"fmt"
	"net/http"
	. "pemgr/cmd/pemgr-server/schema"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/hpcloud/tail"
)

func handleVersion(c *gin.Context) {
	res := CustomVersion{
		Version: pInfo.Server.Version,
		Commit:  pInfo.Server.Commit,
		Date:    pInfo.Server.Date,
	}
	c.JSON(http.StatusOK, res)
}

func logHandler(c *gin.Context) {
}

/*
func handleDb(c *gin.Context) {
	result := make(map[string]string)

	keys, err := db.Keys("*")
	if err != nil {
		c.JSON(http.StatusInternalServerError, keys)
	}

	sort.Strings(keys)

	for _, key := range keys {
		val, err := db.Get(key)
		if err != nil {
			continue
		}
		result[key] = val
	}
	c.JSON(http.StatusOK, result)
}
*/

func pingHandler(c *gin.Context) {
}

var upWebsocket = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func getLog(filename string, ws *websocket.Conn, mt int) {
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,
		Follow:    true,
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	var msg *tail.Line
	var ok bool

	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Printf("msg(%T): %v\n", msg, msg)
		fmt.Printf("msg(%T): %v\n", msg.Text, msg.Text)

		log := []byte(msg.Text)

		err = ws.WriteMessage(mt, log)
		if err != nil {
			break
		}
		//time.Sleep(5 * time.Millisecond)
	}
}

func logsocket(c *gin.Context) {
	ws, err := upWebsocket.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	mt, cmd, err := ws.ReadMessage()
	if err != nil {
	}

	if string(cmd) == "plat-stdio" {
		getLog(pConfig.Plat.Stdout, ws, mt)
	} else if string(cmd) == "plat-stderr" {
		getLog(pConfig.Plat.Stderr, ws, mt)
	} else if string(cmd) == "server-stdio" {
		getLog(pConfig.Server.Stdout, ws, mt)
	} else if string(cmd) == "server-stderr" {
		getLog(pConfig.Server.Stderr, ws, mt)
	}
}
