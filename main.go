package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/vancholen/autocheck/utils"
)

func main() {
	fmt.Println("program start running......")
	c := cron.New(cron.WithSeconds())
	// 每天定时8：30定时签到
	c.AddFunc("* 30 8 */1 * *", utils.LoginAndSign)
	c.Start()
	select {}
}
