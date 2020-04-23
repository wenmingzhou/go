package conf

//LogTransfer 全局变量
type LogTransferCfg struct {
	KafkaCfg `ini:"kafka"` //必须对应.ini的节(kafka)
	ESCfg    `ini:"es"`    //必须对应.ini的节(es)
}

type KafkaCfg struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type ESCfg struct {
	Address string `ini:"address"`
}
