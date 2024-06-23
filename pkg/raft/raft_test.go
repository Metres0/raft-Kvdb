// pkg/raft/raft_test.go
package raft

import (
	"testing"
	"time"
)

func TestRaftElection(t *testing.T) {
	peers := []string{"0", "1", "2"}
	r := NewRaft(peers, 0)

	time.Sleep(1 * time.Second)

	if r.State() != Candidate {
		t.Fatalf("Expected state to be Candidate, got %v", r.State())
	}
}
