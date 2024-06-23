// pkg/raft/raft.go
package raft

import (
	"math/rand"
	"sync"
	"time"
)

type State int

const (
	Follower State = iota
	Candidate
	Leader
)

type LogEntry struct {
	Operation string
	Key       string
	Value     string
}

type Raft struct {
	mu         sync.Mutex
	state      State
	log        []LogEntry
	peers      []string
	me         int
	electionCh chan bool
}

func NewRaft(peers []string, me int) *Raft {
	r := &Raft{
		state:      Follower,
		peers:      peers,
		me:         me,
		electionCh: make(chan bool),
	}
	go r.electionTimeout()
	return r
}

func (r *Raft) electionTimeout() {
	timeout := time.Duration(150+rand.Intn(150)) * time.Millisecond
	for {
		select {
		case <-time.After(timeout):
			r.startElection()
		case <-r.electionCh:
			timeout = time.Duration(150+rand.Intn(150)) * time.Millisecond
		}
	}
}

func (r *Raft) startElection() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.state = Candidate
	r.broadcastRequestVote()
}

func (r *Raft) broadcastRequestVote() {
	// 逻辑：向其他节点发送投票请求
}

func (r *Raft) State() State {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.state
}

func (r *Raft) AppendEntries(entries []LogEntry) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(entries) == 0 {
		return true
	}

	for _, entry := range entries {
		r.log = append(r.log, entry)
		applyLogEntry(entry)
	}

	return true
}

func applyLogEntry(entry LogEntry) {
	if entry.Operation == "set" {
		// 调用KVStore的Set方法
	}
}
