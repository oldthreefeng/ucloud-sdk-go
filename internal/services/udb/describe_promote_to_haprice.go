//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB DescribePromoteToHAPrice

package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribePromoteToHAPriceRequest is request schema for DescribePromoteToHAPrice action
type DescribePromoteToHAPriceRequest struct {
	request.CommonBase

	// 实例所在的可用区
	Zone *string `required:"false"`

	// DB实例ID
	DBId *string `required:"true"`
}

// DescribePromoteToHAPriceResponse is response schema for DescribePromoteToHAPrice action
type DescribePromoteToHAPriceResponse struct {
	response.CommonBase

	// 如果返回错误，错误消息
	Message string

	// 普通升级高可用的差价，单位为分
	Price int
}

// NewDescribePromoteToHAPriceRequest will create request of DescribePromoteToHAPrice action.
func (c *UDBClient) NewDescribePromoteToHAPriceRequest() *DescribePromoteToHAPriceRequest {
	req := &DescribePromoteToHAPriceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribePromoteToHAPrice - 获取普通单点实例升级高可用实例的差价
func (c *UDBClient) DescribePromoteToHAPrice(req *DescribePromoteToHAPriceRequest) (*DescribePromoteToHAPriceResponse, error) {
	var err error
	var res DescribePromoteToHAPriceResponse

	err = c.Client.InvokeAction("DescribePromoteToHAPrice", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
