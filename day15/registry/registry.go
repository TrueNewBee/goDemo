package registry

import "context"

// 总的注册接口

// Name(): 插件名字,例如传etcd
// Init(opts ...Option): 初始化,里面用选项设计模式做初始化
// Regiter(): 服务注册
// Unregister(): 服务反注册, 例如服务端停了,注册列表销毁
// GetService: 服务发现(ip port[] string)

type Registry interface {
	// 插件名字
	Name() string
	// 初始化
	Init(ctx context.Context, opts ...Option)(err error)
	// 服务注册
	Regiter(ctx context.Context, service *Service )(err error)
	// 反服务注册
	Unregiter(ctx context.Context, service *Service)(err error)
	// 服务发现
	GetService(ctx context.Context, name string)(service *Service, err error)
}
