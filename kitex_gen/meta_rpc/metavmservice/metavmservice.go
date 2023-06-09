// Code generated by Kitex v0.5.2. DO NOT EDIT.

package metavmservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	meta_rpc "github.com/renjs2001/WebManager/kitex_gen/meta_rpc"
)

func serviceInfo() *kitex.ServiceInfo {
	return metaVMServiceServiceInfo
}

var metaVMServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "MetaVMService"
	handlerType := (*meta_rpc.MetaVMService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetVMInfo":    kitex.NewMethodInfo(getVMInfoHandler, newMetaVMServiceGetVMInfoArgs, newMetaVMServiceGetVMInfoResult, false),
		"GetAbversion": kitex.NewMethodInfo(getAbversionHandler, newMetaVMServiceGetAbversionArgs, newMetaVMServiceGetAbversionResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "meta_rpc",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func getVMInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*meta_rpc.MetaVMServiceGetVMInfoArgs)
	realResult := result.(*meta_rpc.MetaVMServiceGetVMInfoResult)
	success, err := handler.(meta_rpc.MetaVMService).GetVMInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMetaVMServiceGetVMInfoArgs() interface{} {
	return meta_rpc.NewMetaVMServiceGetVMInfoArgs()
}

func newMetaVMServiceGetVMInfoResult() interface{} {
	return meta_rpc.NewMetaVMServiceGetVMInfoResult()
}

func getAbversionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*meta_rpc.MetaVMServiceGetAbversionArgs)
	realResult := result.(*meta_rpc.MetaVMServiceGetAbversionResult)
	success, err := handler.(meta_rpc.MetaVMService).GetAbversion(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMetaVMServiceGetAbversionArgs() interface{} {
	return meta_rpc.NewMetaVMServiceGetAbversionArgs()
}

func newMetaVMServiceGetAbversionResult() interface{} {
	return meta_rpc.NewMetaVMServiceGetAbversionResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetVMInfo(ctx context.Context, req *meta_rpc.VMInfoReq) (r *meta_rpc.VMInfoResp, err error) {
	var _args meta_rpc.MetaVMServiceGetVMInfoArgs
	_args.Req = req
	var _result meta_rpc.MetaVMServiceGetVMInfoResult
	if err = p.c.Call(ctx, "GetVMInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAbversion(ctx context.Context, req *meta_rpc.VersionReq) (r *meta_rpc.VersionResp, err error) {
	var _args meta_rpc.MetaVMServiceGetAbversionArgs
	_args.Req = req
	var _result meta_rpc.MetaVMServiceGetAbversionResult
	if err = p.c.Call(ctx, "GetAbversion", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
