package conf

// 定义一个app的结构体

type AppConf struct{
	KafkaConf	`ini:"kafka"`  //要和ini中的节对应
	EtcdConf	`ini:"etcdconf"`
}

type KafkaConf struct {
	Address string	`ini:"address"`
	ChanMaxSize int  `ini:chan_max_size`
}

type EtcdConf struct {
	Address string	`ini:"address"`
	Key string `ini:"collect_log_key"`
	Timeout int  `ini:"timeout"`
}


// ------ unused ↓-------

type TaillogConf struct{
	FileName string	`ini:"filename"`
}