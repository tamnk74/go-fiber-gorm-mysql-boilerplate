package schedulers

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func Init() {
	c := cron.New(cron.WithSeconds())
	fmt.Println("Inited cronjob")
	c.AddFunc("*/5 * * * *", func() { fmt.Println("Every second ", time.Now()) })
	c.Start()
}
