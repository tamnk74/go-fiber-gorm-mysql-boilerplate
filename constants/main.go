package constants

const AUTH = "user:"

const APP_NAME = "go_app"

const SEND_EMAIL_Q = "send_email"

var STATUS = struct {
	ACTIVE   uint
	INACTIVE uint
}{
	ACTIVE:   1,
	INACTIVE: 0,
}
