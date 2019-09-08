package xlog

import (
	"strings"
	"testing"

	"github.com/spf13/viper"
)

var conf = `
{
    "logger":{
        "filename":"./log/server.log",
        "maxsize":20000,
        "maxage":7,
        "maxbackups":10,
        "compress":false
    },
    "interval":"24h",
    "level":"debug"
}
`

func TestA(t *testing.T) {
	c := viper.New()
	c.SetConfigType("json")
	c.ReadConfig(strings.NewReader(conf))
	Init(c)
	Debug("aaaa", 111211)
	Debug("aaaa", 111211)
	Debug("aaaa", 111211)
	Debug("aaaa", 111211)
	Debug("aaaa", 111211)
	Debug("aaaa", 111211)
}

func Benchmark_A(b *testing.B) {
	c := viper.New()
	c.SetConfigType("json")
	c.ReadConfig(strings.NewReader(conf))
	Init(c)
	for i := 0; i < b.N; i++ { //use b.N for looping
		Debug("aaaa", 111211)
	}
}
