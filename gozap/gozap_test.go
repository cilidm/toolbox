package gozap

import "testing"

func TestZap(t *testing.T) {
	InitLogger("./zap.log", "debug")
	Info("aa", "bb")
}
