package core

type Size struct {
	// 消息
	Msg Msg `json:"msg,omitempty" validate:"required"`
	// 窗口
	Window Window `json:"window,omitempty" validate:"required"`
}
