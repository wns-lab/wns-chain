package types

import (
	"time"
)

type Node [32]byte

func (n Node) IsValid() bool {
	return true
}

type Record struct {
	Owner    string
	Resolver string
	TTL      time.Time
}

func (r Record) IsValid() bool {
	return true
}
