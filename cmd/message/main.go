package main

import (
	"net"

	"github.com/EdwardShen125/bytedance-simple-server/cmd/message/dal"
	message "github.com/EdwardShen125/bytedance-simple-server/kitex_gen/message/messageservice"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/constants"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/middleware"
	tracer2 "github.com/EdwardShen125/bytedance-simple-server/pkg/tracer"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	tracer2.InitJaeger(constants.MessageServiceName)
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress}) //r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.MessageAddress)
	if err != nil {
		panic(err)
	}
	Init()
	svr := message.NewServer(new(MessageServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.MessageServiceName}), //server name
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       //address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), //limit
		server.WithMuxTransport(),                                          //Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    //tracer
		server.WithRegistry(r),                                             //registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
