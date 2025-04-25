package kernel

import (
	"strings"
)

type Matcher struct {
	// 等于
	Equal string `json:"equal,omitempty"`
	// 前缀
	Prefix string `json:"prefix,omitempty"`
	// 后缀
	Suffix string `json:"suffix,omitempty"`
	// 包含
	Contains string `json:"contains,omitempty"`
}

func (m *Matcher) Test(key string) (new string, match bool) {
	key = strings.ToLower(key)
	new = key
	if "" != m.Equal && m.Equal == key {
		match = true
	}

	if "" != m.Prefix && strings.HasPrefix(key, m.Prefix) {
		match = true
	}

	if "" != m.Suffix && strings.HasSuffix(key, m.Suffix) {
		new = key
		match = true
	}

	if "" != m.Contains && strings.Contains(key, m.Contains) {
		new = key
		match = true
	}

	return
}
