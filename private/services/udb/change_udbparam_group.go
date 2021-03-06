//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB ChangeUDBParamGroup

package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// ChangeUDBParamGroupRequest is request schema for ChangeUDBParamGroup action
type ChangeUDBParamGroupRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// DB实例Id
	DBId *string `required:"true"`

	// 参数组Id
	GroupId *string `required:"true"`
}

// ChangeUDBParamGroupResponse is response schema for ChangeUDBParamGroup action
type ChangeUDBParamGroupResponse struct {
	response.CommonBase
}

// NewChangeUDBParamGroupRequest will create request of ChangeUDBParamGroup action.
func (c *UDBClient) NewChangeUDBParamGroupRequest() *ChangeUDBParamGroupRequest {
	req := &ChangeUDBParamGroupRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ChangeUDBParamGroup - 修改配置文件
func (c *UDBClient) ChangeUDBParamGroup(req *ChangeUDBParamGroupRequest) (*ChangeUDBParamGroupResponse, error) {
	var err error
	var res ChangeUDBParamGroupResponse

	err = c.Client.InvokeAction("ChangeUDBParamGroup", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
