package kernel

type Body struct {
	// 原始请求
	Raws []*Raw `default:"[{'contains': 'raw'}]" json:"raws,omitempty"`
}

func (b *Body) Check(check string) (checked bool) {
	for _, _raw := range b.Raws {
		if checked = _raw.Check(check); checked {
			break
		}
	}

	return
}
