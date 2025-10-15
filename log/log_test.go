package log_test

import (
	"github.com/dobyte/due/v2/utils/xtime"
	"math/rand"
	"testing"
	"time"

	"github.com/dobyte/due/v2/log"
)

func TestLog(t *testing.T) {
	logger := log.NewLogger()
	id := rand.Int31n(1000)
	//logger.Debug("welcome to due-framework")
	//logger.Info("welcome to due-framework")
	//logger.Warn("welcome to due-framework")
	//logger.Error("welcome to due-framework")
	for {
		time.Sleep(time.Second)
		logger.Debugf("id: %d %s", id, xtime.Now().String())
	}
}
