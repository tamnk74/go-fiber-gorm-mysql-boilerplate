package queue

import (
	"github.com/gocraft/work"
	Email "github.com/tamnk74/todolist-mysql-go/utils/email"
)

func (c *Context) SendEmail(job *work.Job) error {
	subject := job.ArgString("subject")
	email := job.ArgString("email")
	if err := job.ArgError(); err != nil {
		return err
	}
	Email.Send([]string{email}, subject)
	return nil
}
