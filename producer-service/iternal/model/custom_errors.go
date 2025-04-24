package model

import "fmt"

var (
	ErrObjectNotExists = fmt.Errorf("no such object")
	ErrDb = fmt.Errorf("db error")

	ErrStartProducer = fmt.Errorf("failed to start producer")

	ErrProcessMessage = fmt.Errorf("process message error")
)