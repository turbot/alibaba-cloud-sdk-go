package resourcemanager

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

// DisableControlPolicy invokes the resourcemanager.DisableControlPolicy API synchronously
func (client *Client) DisableControlPolicy(request *DisableControlPolicyRequest) (response *DisableControlPolicyResponse, err error) {
	response = CreateDisableControlPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DisableControlPolicyWithChan invokes the resourcemanager.DisableControlPolicy API asynchronously
func (client *Client) DisableControlPolicyWithChan(request *DisableControlPolicyRequest) (<-chan *DisableControlPolicyResponse, <-chan error) {
	responseChan := make(chan *DisableControlPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableControlPolicy(request)
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

// DisableControlPolicyWithCallback invokes the resourcemanager.DisableControlPolicy API asynchronously
func (client *Client) DisableControlPolicyWithCallback(request *DisableControlPolicyRequest, callback func(response *DisableControlPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableControlPolicyResponse
		var err error
		defer close(result)
		response, err = client.DisableControlPolicy(request)
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

// DisableControlPolicyRequest is the request struct for api DisableControlPolicy
type DisableControlPolicyRequest struct {
	*requests.RpcRequest
}

// DisableControlPolicyResponse is the response struct for api DisableControlPolicy
type DisableControlPolicyResponse struct {
	*responses.BaseResponse
	EnablementStatus string `json:"EnablementStatus" xml:"EnablementStatus"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableControlPolicyRequest creates a request to invoke DisableControlPolicy API
func CreateDisableControlPolicyRequest() (request *DisableControlPolicyRequest) {
	request = &DisableControlPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "DisableControlPolicy", "", "")
	request.Method = requests.POST
	return
}

// CreateDisableControlPolicyResponse creates a response to parse from DisableControlPolicy response
func CreateDisableControlPolicyResponse() (response *DisableControlPolicyResponse) {
	response = &DisableControlPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
