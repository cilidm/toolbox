package trace

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

var incrNum uint64

// NewTraceID New trace id
func NewTraceID() string {
	return fmt.Sprintf("trace-id-%d-%s-%d",
		os.Getpid(),
		time.Now().Format("2006.01.02.15.04.05.999"),
		atomic.AddUint64(&incrNum, 1))
}
