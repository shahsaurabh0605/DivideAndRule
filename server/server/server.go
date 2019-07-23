package server

import (
	"../job"
	"../node"
	"../program"
	"github.com/rsms/gotalk"
	"sync"
)

type Server struct {
	Programs   map[string]program.Program
	RWPrograms sync.RWMutex
	Jobs       map[string]job.Job
	RWJobs     sync.RWMutex
	Socks      map[*gotalk.Sock]node.Node
	RWSocks    sync.RWMutex
	// X		   map[string]string
}

func (server *Server) Init() {
	server.Programs = make(map[string]program.Program)
	server.Jobs = make(map[string]job.Job)
	server.Socks = make(map[*gotalk.Sock]node.Node)
	// server.X = make(map[string]string)
}
