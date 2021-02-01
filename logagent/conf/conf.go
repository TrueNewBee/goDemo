package conf

// 定义一个app的结构体

type AppConf struct{
	KafkaConf	`ini:"kafka"`  //要和ini中的节对应
	TaillogConf	`ini:"taillog"`
}

type KafkaConf struct {
	Address string	`ini:"address"`
	Topic string	`ini:"topic"`
}

type TaillogConf struct{
	FileName string	`ini:"filename"`
}