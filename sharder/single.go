package sharder

import (
	"github.com/honeycombio/malinois/logger"
)

// SingleShard implements the Shard interface
type SingleShard string

var selfShard SingleShard = "self"

func (s *SingleShard) Equals(other Shard) bool { return true }

// GetAddress will never be used because every shard is my shard
func (s *SingleShard) GetAddress() string { return "" }

type SingleServerSharder struct {
	Logger logger.Logger `inject:""`
}

func (s *SingleServerSharder) MyShard() Shard {
	return &selfShard
}

func (s *SingleServerSharder) WhichShard(traceID string) Shard {
	s.Logger.Debugf("single server sharder; choosing self for trace ID %s", traceID)
	return &selfShard
}