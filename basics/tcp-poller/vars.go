package main

const (
	BACKLOG_NUM = 1024
	KEVENT_NUM  = 4
	BUFF_SIZE   = 32
)

var (
	_ NetPoller = (*KqueuePoller)(nil)
)
