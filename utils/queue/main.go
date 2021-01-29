package queue

import (
	"fmt"
	"log"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/tamnk74/todolist-mysql-go/constants"
)

var redisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", ":6379")
	},
}

type Context struct {
	ID int64
}

var pool *work.WorkerPool

var enqueuer = work.NewEnqueuer(constants.APP_NAME, redisPool)

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: ", job.Name)
	return next()
}

func Init() *work.WorkerPool {
	if pool != nil {
		return pool
	}
	pool = work.NewWorkerPool(Context{}, 10, constants.APP_NAME, redisPool)
	pool.Middleware((*Context).Log)

	pool.Job(constants.SEND_EMAIL_Q, (*Context).SendEmail)

	pool.Start()
	return pool
}

func CreateJob(jobName string, payload work.Q) {
	_, err := enqueuer.Enqueue(jobName, payload)
	if err != nil {
		log.Fatal(err)
	}
}
