package finalize

import (
	"INServer/src/common/global"
	"INServer/src/common/util"
	"INServer/src/services/cluster"
	"INServer/src/services/world"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Wait 等待结束
func Wait() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	for {
		stopped := false
		select {
		case sig := <-sigs:
			if sig.String() == "interrupt" {
				stopped = true
				stopNode()
			}
			break
		}
		if stopped {
			break
		}
		time.Sleep(time.Millisecond * 10)
	}
}

func stopNode() {
	global.PendingExit = true
	stopServer()
}

func stopServer() {
	switch global.CurrentServerType {
	case global.WorldServer:
		world.Instance.Stop()
		break
	case global.DatabaseServer:
		util.Wait(func() bool {
			return cluster.RunningCount() == 2
		}, "等待其他服务器关服完成...", time.Second)
		break
	case global.CenterServer:
		util.Wait(func() bool {
			return cluster.RunningCount() == 1
		}, "等待其他服务器关服完成...", time.Second)
		break
	}
}

func init() {
	global.PendingExit = false
}
