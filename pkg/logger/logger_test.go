package logger

import (
	"go-api-base/pkg/util/timeutil"
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger, err := NewLogger(
		WithTimeLayout(timeutil.CSTLayout),
	)
	if err != nil {
		t.Fatal(err)
	}

	defer logger.Sync()

	logger.Error("aaaaa")
	logger.Debug("bbbbb")
	logger.Info("ccccc")
}
