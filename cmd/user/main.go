package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/EdwardShen125/bytedance-simple-server/cmd/user/dal"
	user "github.com/EdwardShen125/bytedance-simple-server/kitex_gen/user/userservice"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/bound"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/constants"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/jwt"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/middleware"
	tracer2 "github.com/EdwardShen125/bytedance-simple-server/pkg/tracer"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"net"
)

var Jwt *jwt.JWT

func Init() {
	tracer2.InitJaeger(constants.UserServiceName)
	dal.Init()
	Jwt = jwt.NewJWT([]byte(constants.SecretKey))
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress}) //r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.UserAddress)
	if err != nil {
		panic(err)
	}
	Init()
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}), //server name
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       //address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), //limit
		server.WithMuxTransport(),                                          //Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    //tracer
		server.WithBoundHandler(bound.NewCpuLimitHandler()),                //BoundHandler
		server.WithRegistry(r),                                             //registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
