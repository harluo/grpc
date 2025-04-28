package core

import (
	"fmt"

	"github.com/goexl/gox"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/harluo/grpc/internal/internal/constant"
	"google.golang.org/protobuf/encoding/protojson"
)

type Gateway struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled,omitempty"`
	// 名字
	Name string `default:"网关" json:"name,omitempty"`
	// 绑定监听主机
	Host string `json:"host,omitempty"`
	// 绑定监听端口
	Port int `default:"9001" json:"port,omitempty" validate:"required,min=1,max=65535"`
	// 路径
	Path string `json:"path,omitempty" validate:"omitempty,startswith=/,endsnotwith=/"`
	// 跨域
	Cors *Cors `json:"cors,omitempty"`
	// 超时
	Timeout Timeout `json:"timeout,omitempty"`
	// 序列化
	Json Json `json:"json,omitempty"`
	// 头
	Header Header `json:"header,omitempty"`
	// 消息体
	Body Body `json:"body,omitempty"`
	// 模式
	Unescape *Unescape `json:"unescape,omitempty"`
}

func (g *Gateway) Options() (options []runtime.ServeMuxOption) {
	options = make([]runtime.ServeMuxOption, 0, 1)
	json := &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			Multiline:       g.Json.Multiline,
			Indent:          g.Json.Indent,
			AllowPartial:    g.Json.Partial,
			UseProtoNames:   gox.Contains(&g.Json.Options, constant.NameAsProto),
			UseEnumNumbers:  gox.Contains(&g.Json.Options, constant.EnumAsNumbers),
			EmitUnpopulated: g.Json.Unpopulated,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			AllowPartial:   g.Json.Partial,
			DiscardUnknown: *g.Json.Discard,
		},
	}
	options = append(options, runtime.WithMarshalerOption(runtime.MIMEWildcard, json))

	return
}

func (g *Gateway) Addr() string {
	return fmt.Sprintf("%s:%d", g.Host, g.Port)
}

func (g *Gateway) Enable() bool {
	return nil != g.Enabled && *g.Enabled
}

func (g *Gateway) CorsEnabled() bool {
	return nil != g.Cors && *g.Cors.Enabled
}
