package model

type Status int8

const (
	IDLE = 1 << iota
	PROCESSING
	PAUSE
	DONE
	TERMINATED
)
