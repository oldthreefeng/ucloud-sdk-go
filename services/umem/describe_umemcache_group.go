//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem DescribeUMemcacheGroup

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUMemcacheGroupRequest is request schema for DescribeUMemcacheGroup action
type DescribeUMemcacheGroupRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 组的ID,如果指定则获取描述，否则为列表操 作,需指定Offset/Limit
	GroupId *string `required:"false"`

	// 分页显示的起始偏移, 默认值为0
	Offset *int `required:"false"`

	// 分页显示的条目数, 默认值为20
	Limit *int `required:"false"`
}

// DescribeUMemcacheGroupResponse is response schema for DescribeUMemcacheGroup action
type DescribeUMemcacheGroupResponse struct {
	response.CommonBase

	// 组的总的节点个数
	TotalCount int

	// 组列表,参见 UMemcacheGroupSet
	DataSet []UMemcacheGroupSet
}

// NewDescribeUMemcacheGroupRequest will create request of DescribeUMemcacheGroup action.
func (c *UMemClient) NewDescribeUMemcacheGroupRequest() *DescribeUMemcacheGroupRequest {
	req := &DescribeUMemcacheGroupRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUMemcacheGroup - 显示Memcache
func (c *UMemClient) DescribeUMemcacheGroup(req *DescribeUMemcacheGroupRequest) (*DescribeUMemcacheGroupResponse, error) {
	var err error
	var res DescribeUMemcacheGroupResponse

	err = c.client.InvokeAction("DescribeUMemcacheGroup", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}