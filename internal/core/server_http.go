package core

import (
	"net"
	"net/http"
	"strings"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/harluo/grpc/internal/internal/constant"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

func (s *Server) handler(grpc *grpc.Server, gateway http.Handler) (handler http.Handler) {
	// 增加原始数据解析
	handler = s.handle(gateway)
	// 处理跨域
	handler = gox.Ift(s.gateway.CorsEnabled(), s.cors(handler), handler)
	// 如果端口配置为一样，需要合并处理
	handler = gox.Ift(s.diffPort(), handler, s.combine(grpc, handler))

	return
}

func (s *Server) combine(grpc *grpc.Server, gateway http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(rsp http.ResponseWriter, req *http.Request) {
		fields := s.fields(req)
		s.logger.Debug("收到请求", fields[0], fields[1:]...)
		if req.ProtoMajor >= 2 && constant.GrpcHeaderValue == req.Header.Get(constant.HeaderContentType) {
			grpc.ServeHTTP(rsp, req)
		} else {
			gateway.ServeHTTP(rsp, req)
		}
	}), new(http2.Server))
}

func (s *Server) cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rsp http.ResponseWriter, req *http.Request) {
		// 设置允许跨域访问的源
		rsp.Header().Set(constant.HeaderAllowOrigin, strings.Join(s.gateway.Cors.Allows, constant.Comma))
		// 设置允许的请求方法
		rsp.Header().Set(constant.HeaderAllowMethods, strings.Join(s.gateway.Cors.Methods, constant.Comma))
		// 设置允许的请求头
		rsp.Header().Set(constant.HeaderAllowHeaders, strings.Join(s.gateway.Cors.Headers, constant.Comma))

		// 如果是预检请求，直接返回
		if req.Method == constant.MethodOptions {
			return
		}

		// 调用实际的处理器函数
		next.ServeHTTP(rsp, req)
	})
}

func (s *Server) handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rsp http.ResponseWriter, req *http.Request) {
		// 处理原始数据
		if nil != req && nil != s.gateway && s.gateway.Body.Check(req.URL.Path) {
			req.Header.Set(constant.HeaderContentType, constant.RawHeaderValue)
		}

		// 调用实际的处理器函数
		next.ServeHTTP(rsp, req)
	})
}

func (s *Server) ip(req *http.Request) (ip string) {
	ip = req.Header.Get(constant.HeaderXRealIp)
	if netIp := net.ParseIP(ip); nil != netIp {
		return
	}

	ips := req.Header.Get(constant.HeaderXForwardedFor)
	for _, _ip := range strings.Split(ips, constant.Comma) {
		ip = _ip
		if netIP := net.ParseIP(ip); nil != netIP {
			return
		}
	}

	if _ip, _, err := net.SplitHostPort(req.RemoteAddr); nil == err {
		if netIp := net.ParseIP(_ip); nil != netIp {
			ip = _ip
		}
	}

	return
}

func (s *Server) fields(request *http.Request) gox.Fields[any] {
	return gox.Fields[any]{
		field.New("method", request.Method),
		field.New("url", request.URL.String()),
		field.New("ip", s.ip(request)),
		field.New("useragent", request.UserAgent()),
		field.New("referer", request.Referer()),
	}
}
