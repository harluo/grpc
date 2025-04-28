package core

import (
	"strings"
)

type Raw struct {
	// 前缀
	Prefix string `json:"prefix,omitempty"`
	// 后缀
	Suffix string `json:"suffix,omitempty"`
	// 包含
	Contains string `default:"Raw" json:"contains,omitempty"`
}

func (r *Raw) Check(check string) (checked bool) {
	if "" != r.Contains && strings.Contains(check, r.Contains) {
		checked = true
	} else if "" != r.Prefix && strings.HasPrefix(check, r.Prefix) {
		checked = true
	} else if "" != r.Suffix && strings.HasSuffix(check, r.Suffix) {
		checked = true
	}

	return
}
