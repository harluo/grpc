package core

type Cors struct {
	// 是否开启
	Enabled *bool `json:"enabled,omitempty"`
	// 允许跨域访问的源
	Allows []string `default:"['*']" json:"allows,omitempty"`
	// 允许的请求方法
	// nolint:lll
	Methods []string `default:"['GET', 'POST', 'PUT', 'DELETE']" json:"methods,omitempty"`
	// 允许的请求头
	// nolint:lll
	Headers []string `default:"['Content-Type', 'Authorization']" json:"headers,omitempty"`
}
