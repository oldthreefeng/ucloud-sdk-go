//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UFS RemoveUFSVolume

package ufs

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// RemoveUFSVolumeRequest is request schema for RemoveUFSVolume action
type RemoveUFSVolumeRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 文件系统ID
	VolumeId *string `required:"true"`
}

// RemoveUFSVolumeResponse is response schema for RemoveUFSVolume action
type RemoveUFSVolumeResponse struct {
	response.CommonBase
}

// NewRemoveUFSVolumeRequest will create request of RemoveUFSVolume action.
func (c *UFSClient) NewRemoveUFSVolumeRequest() *RemoveUFSVolumeRequest {
	req := &RemoveUFSVolumeRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// RemoveUFSVolume - 删除UFS文件系统
func (c *UFSClient) RemoveUFSVolume(req *RemoveUFSVolumeRequest) (*RemoveUFSVolumeResponse, error) {
	var err error
	var res RemoveUFSVolumeResponse

	err = c.Client.InvokeAction("RemoveUFSVolume", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
