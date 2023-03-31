package main

import "github.com/dtherhtun/Learning-go/rest-servers/basic/internal/taskstore"

type taskServer struct {
	store *taskstore.TaskStore
}
