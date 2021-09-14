package config

// Options 服务配置参数
type Options struct {
	Server struct {
		Port int `json:"port"` // 服务端口
	} `json:"server"`
}
