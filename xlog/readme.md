# 基于zap 的日志

## 1.最佳实践
```
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
	c := viper.New()
	c.SetConfigType("json")
	c.ReadConfig(strings.NewReader(conf))
    xlog.Init(c)
    Debug("aaaa", 111211)
```