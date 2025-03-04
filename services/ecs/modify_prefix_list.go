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

// ModifyPrefixList invokes the ecs.ModifyPrefixList API synchronously
func (client *Client) ModifyPrefixList(request *ModifyPrefixListRequest) (response *ModifyPrefixListResponse, err error) {
	response = CreateModifyPrefixListResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyPrefixListWithChan invokes the ecs.ModifyPrefixList API asynchronously
func (client *Client) ModifyPrefixListWithChan(request *ModifyPrefixListRequest) (<-chan *ModifyPrefixListResponse, <-chan error) {
	responseChan := make(chan *ModifyPrefixListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyPrefixList(request)
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

// ModifyPrefixListWithCallback invokes the ecs.ModifyPrefixList API asynchronously
func (client *Client) ModifyPrefixListWithCallback(request *ModifyPrefixListRequest, callback func(response *ModifyPrefixListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyPrefixListResponse
		var err error
		defer close(result)
		response, err = client.ModifyPrefixList(request)
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

// ModifyPrefixListRequest is the request struct for api ModifyPrefixList
type ModifyPrefixListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer               `position:"Query" name:"ResourceOwnerId"`
	Description          string                         `position:"Query" name:"Description"`
	PrefixListId         string                         `position:"Query" name:"PrefixListId"`
	AddEntry             *[]ModifyPrefixListAddEntry    `position:"Query" name:"AddEntry"  type:"Repeated"`
	ResourceOwnerAccount string                         `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                         `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer               `position:"Query" name:"OwnerId"`
	PrefixListName       string                         `position:"Query" name:"PrefixListName"`
	RemoveEntry          *[]ModifyPrefixListRemoveEntry `position:"Query" name:"RemoveEntry"  type:"Repeated"`
}

// ModifyPrefixListAddEntry is a repeated param struct in ModifyPrefixListRequest
type ModifyPrefixListAddEntry struct {
	Description string `name:"Description"`
	Cidr        string `name:"Cidr"`
}

// ModifyPrefixListRemoveEntry is a repeated param struct in ModifyPrefixListRequest
type ModifyPrefixListRemoveEntry struct {
	Cidr string `name:"Cidr"`
}

// ModifyPrefixListResponse is the response struct for api ModifyPrefixList
type ModifyPrefixListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyPrefixListRequest creates a request to invoke ModifyPrefixList API
func CreateModifyPrefixListRequest() (request *ModifyPrefixListRequest) {
	request = &ModifyPrefixListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyPrefixList", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyPrefixListResponse creates a response to parse from ModifyPrefixList response
func CreateModifyPrefixListResponse() (response *ModifyPrefixListResponse) {
	response = &ModifyPrefixListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
