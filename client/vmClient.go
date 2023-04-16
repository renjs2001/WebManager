package client

import (
	"github.com/cloudwego/kitex/client"
	"github.com/renjs2001/WebManager/kitex_gen/meta_rpc/metavmservice"
)

var VmClient metavmservice.Client

func Init() {
	var opts []client.Option

	opts = append(opts, client.WithHostPorts("localhost:8888"))

	VmClient = metavmservice.MustNewClient("meta", opts...)
}
