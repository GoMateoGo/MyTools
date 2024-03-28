package global

import "go.uber.org/zap"

var (
	Logger  *zap.SugaredLogger
	RunPath map[string]string
)

func init() {
	RunPath = make(map[string]string)
}
