package core

type Json struct {
	// 是否允许多行
	Multiline bool `json:"multiline,omitempty"`
	// 前缀
	Indent string `json:"indent,omitempty"`
	// 允许部分
	Partial bool `json:"partial,omitempty"`
	// 选项列表
	// nolint: lll
	Options []string `default:"['enum_as_numbers', 'name_as_proto']" json:"options,omitempty"`
	// 是否允许返回零值
	Unpopulated bool `json:"unpopulated,omitempty"`
	// 是否允许丢弃
	Discard *bool `default:"true" json:"discard,omitempty"`
}
