package model

import "fmt"

var (
	ErrObjectNotExists = fmt.Errorf("no such object")
	ErrDb = fmt.Errorf("db error")

	ErrCreateConsumerGroup = fmt.Errorf("failed to create consumer group")
	ErrConsumerHandler = fmt.Errorf("failed to consume via handler")
)