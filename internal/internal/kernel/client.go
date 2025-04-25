package kernel

type Client struct {
	// 名称
	Name string `json:"name,omitempty" validate:"required_without=Names"`
	// 名称列表
	Names []string `json:"names,omitempty" validate:"required_without=Name"`
	// 连接地址
	Addr string `json:"addr,omitempty" validate:"required,url"`
}
