package json

import (
	"encoding/json"

	"github.com/inoth/ino-toybox/components/logger"
	"github.com/inoth/ino-toybox/servers/wssvc/accumulator"
	"github.com/inoth/ino-toybox/servers/wssvc/models"
	"github.com/inoth/ino-toybox/servers/wssvc/msg_pipeline/parsers"
)

// 解析json字符串
type JsonParser struct{}

func (JsonParser) Parser(msgbody []byte, acc accumulator.Accumulator) {
	logger.Log.Info(string(msgbody))

	var body models.SourceMessage
	err := json.Unmarshal(msgbody, &body)
	if err != nil {
		logger.Log.Error(err.Error())
		acc.Err(err)
		return
	}
	acc.Next(body.GenNextBody())
}

func init() {
	parsers.AddParsers("json", &JsonParser{})
}
