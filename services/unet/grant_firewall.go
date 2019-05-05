//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UNet GrantFirewall

package unet

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// GrantFirewallRequest is request schema for GrantFirewall action
type GrantFirewallRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 防火墙资源ID
	FWId *string `required:"true"`

	// 绑定防火墙组的资源类型，默认为全部资源类型。枚举值为："unatgw"，NAT网关； "uhost"，云主机； "upm"，物理云主机； "hadoophost"，hadoop节点； "fortresshost"，堡垒机； "udhost"，私有专区主机；"udockhost"，容器；"dbaudit"，数据库审计.
	ResourceType *string `required:"true"`

	// 所应用资源ID
	ResourceId *string `required:"true"`
}

// GrantFirewallResponse is response schema for GrantFirewall action
type GrantFirewallResponse struct {
	response.CommonBase
}

// NewGrantFirewallRequest will create request of GrantFirewall action.
func (c *UNetClient) NewGrantFirewallRequest() *GrantFirewallRequest {
	req := &GrantFirewallRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// GrantFirewall - 将防火墙应用到资源上
func (c *UNetClient) GrantFirewall(req *GrantFirewallRequest) (*GrantFirewallResponse, error) {
	var err error
	var res GrantFirewallResponse

	err = c.Client.InvokeAction("GrantFirewall", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
