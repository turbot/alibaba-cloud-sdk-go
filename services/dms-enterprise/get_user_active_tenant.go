package dms_enterprise

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

// GetUserActiveTenant invokes the dms_enterprise.GetUserActiveTenant API synchronously
func (client *Client) GetUserActiveTenant(request *GetUserActiveTenantRequest) (response *GetUserActiveTenantResponse, err error) {
	response = CreateGetUserActiveTenantResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserActiveTenantWithChan invokes the dms_enterprise.GetUserActiveTenant API asynchronously
func (client *Client) GetUserActiveTenantWithChan(request *GetUserActiveTenantRequest) (<-chan *GetUserActiveTenantResponse, <-chan error) {
	responseChan := make(chan *GetUserActiveTenantResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUserActiveTenant(request)
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

// GetUserActiveTenantWithCallback invokes the dms_enterprise.GetUserActiveTenant API asynchronously
func (client *Client) GetUserActiveTenantWithCallback(request *GetUserActiveTenantRequest, callback func(response *GetUserActiveTenantResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserActiveTenantResponse
		var err error
		defer close(result)
		response, err = client.GetUserActiveTenant(request)
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

// GetUserActiveTenantRequest is the request struct for api GetUserActiveTenant
type GetUserActiveTenantRequest struct {
	*requests.RpcRequest
	Tid requests.Integer `position:"Query" name:"Tid"`
}

// GetUserActiveTenantResponse is the response struct for api GetUserActiveTenant
type GetUserActiveTenantResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	Tenant       Tenant `json:"Tenant" xml:"Tenant"`
}

// CreateGetUserActiveTenantRequest creates a request to invoke GetUserActiveTenant API
func CreateGetUserActiveTenantRequest() (request *GetUserActiveTenantRequest) {
	request = &GetUserActiveTenantRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetUserActiveTenant", "", "")
	request.Method = requests.POST
	return
}

// CreateGetUserActiveTenantResponse creates a response to parse from GetUserActiveTenant response
func CreateGetUserActiveTenantResponse() (response *GetUserActiveTenantResponse) {
	response = &GetUserActiveTenantResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
