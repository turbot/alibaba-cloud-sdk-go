package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeNetworkInterfacePermissions invokes the ecs.DescribeNetworkInterfacePermissions API synchronously
func (client *Client) DescribeNetworkInterfacePermissions(request *DescribeNetworkInterfacePermissionsRequest) (response *DescribeNetworkInterfacePermissionsResponse, err error) {
	response = CreateDescribeNetworkInterfacePermissionsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNetworkInterfacePermissionsWithChan invokes the ecs.DescribeNetworkInterfacePermissions API asynchronously
func (client *Client) DescribeNetworkInterfacePermissionsWithChan(request *DescribeNetworkInterfacePermissionsRequest) (<-chan *DescribeNetworkInterfacePermissionsResponse, <-chan error) {
	responseChan := make(chan *DescribeNetworkInterfacePermissionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNetworkInterfacePermissions(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeNetworkInterfacePermissionsWithCallback invokes the ecs.DescribeNetworkInterfacePermissions API asynchronously
func (client *Client) DescribeNetworkInterfacePermissionsWithCallback(request *DescribeNetworkInterfacePermissionsRequest, callback func(response *DescribeNetworkInterfacePermissionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNetworkInterfacePermissionsResponse
		var err error
		defer close(result)
		response, err = client.DescribeNetworkInterfacePermissions(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeNetworkInterfacePermissionsRequest is the request struct for api DescribeNetworkInterfacePermissions
type DescribeNetworkInterfacePermissionsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId              requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber                   requests.Integer `position:"Query" name:"PageNumber"`
	PageSize                     requests.Integer `position:"Query" name:"PageSize"`
	NetworkInterfacePermissionId *[]string        `position:"Query" name:"NetworkInterfacePermissionId"  type:"Repeated"`
	ResourceOwnerAccount         string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string           `position:"Query" name:"OwnerAccount"`
	OwnerId                      requests.Integer `position:"Query" name:"OwnerId"`
	NetworkInterfaceId           string           `position:"Query" name:"NetworkInterfaceId"`
}

// DescribeNetworkInterfacePermissionsResponse is the response struct for api DescribeNetworkInterfacePermissions
type DescribeNetworkInterfacePermissionsResponse struct {
	*responses.BaseResponse
	PageSize                    int                         `json:"PageSize" xml:"PageSize"`
	RequestId                   string                      `json:"RequestId" xml:"RequestId"`
	PageNumber                  int                         `json:"PageNumber" xml:"PageNumber"`
	TotalCount                  int                         `json:"TotalCount" xml:"TotalCount"`
	NetworkInterfacePermissions NetworkInterfacePermissions `json:"NetworkInterfacePermissions" xml:"NetworkInterfacePermissions"`
}

// CreateDescribeNetworkInterfacePermissionsRequest creates a request to invoke DescribeNetworkInterfacePermissions API
func CreateDescribeNetworkInterfacePermissionsRequest() (request *DescribeNetworkInterfacePermissionsRequest) {
	request = &DescribeNetworkInterfacePermissionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeNetworkInterfacePermissions", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeNetworkInterfacePermissionsResponse creates a response to parse from DescribeNetworkInterfacePermissions response
func CreateDescribeNetworkInterfacePermissionsResponse() (response *DescribeNetworkInterfacePermissionsResponse) {
	response = &DescribeNetworkInterfacePermissionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
