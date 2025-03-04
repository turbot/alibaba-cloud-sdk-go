package config

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

// GetAggregateResourceCountsGroupByResourceType invokes the config.GetAggregateResourceCountsGroupByResourceType API synchronously
func (client *Client) GetAggregateResourceCountsGroupByResourceType(request *GetAggregateResourceCountsGroupByResourceTypeRequest) (response *GetAggregateResourceCountsGroupByResourceTypeResponse, err error) {
	response = CreateGetAggregateResourceCountsGroupByResourceTypeResponse()
	err = client.DoAction(request, response)
	return
}

// GetAggregateResourceCountsGroupByResourceTypeWithChan invokes the config.GetAggregateResourceCountsGroupByResourceType API asynchronously
func (client *Client) GetAggregateResourceCountsGroupByResourceTypeWithChan(request *GetAggregateResourceCountsGroupByResourceTypeRequest) (<-chan *GetAggregateResourceCountsGroupByResourceTypeResponse, <-chan error) {
	responseChan := make(chan *GetAggregateResourceCountsGroupByResourceTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAggregateResourceCountsGroupByResourceType(request)
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

// GetAggregateResourceCountsGroupByResourceTypeWithCallback invokes the config.GetAggregateResourceCountsGroupByResourceType API asynchronously
func (client *Client) GetAggregateResourceCountsGroupByResourceTypeWithCallback(request *GetAggregateResourceCountsGroupByResourceTypeRequest, callback func(response *GetAggregateResourceCountsGroupByResourceTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAggregateResourceCountsGroupByResourceTypeResponse
		var err error
		defer close(result)
		response, err = client.GetAggregateResourceCountsGroupByResourceType(request)
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

// GetAggregateResourceCountsGroupByResourceTypeRequest is the request struct for api GetAggregateResourceCountsGroupByResourceType
type GetAggregateResourceCountsGroupByResourceTypeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AggregatorId    string           `position:"Query" name:"AggregatorId"`
	FolderId        string           `position:"Query" name:"FolderId"`
	Region          string           `position:"Query" name:"Region"`
}

// GetAggregateResourceCountsGroupByResourceTypeResponse is the response struct for api GetAggregateResourceCountsGroupByResourceType
type GetAggregateResourceCountsGroupByResourceTypeResponse struct {
	*responses.BaseResponse
	RequestId                       string                 `json:"RequestId" xml:"RequestId"`
	DiscoveredResourceCountsSummary []GroupedResourceCount `json:"DiscoveredResourceCountsSummary" xml:"DiscoveredResourceCountsSummary"`
}

// CreateGetAggregateResourceCountsGroupByResourceTypeRequest creates a request to invoke GetAggregateResourceCountsGroupByResourceType API
func CreateGetAggregateResourceCountsGroupByResourceTypeRequest() (request *GetAggregateResourceCountsGroupByResourceTypeRequest) {
	request = &GetAggregateResourceCountsGroupByResourceTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "GetAggregateResourceCountsGroupByResourceType", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAggregateResourceCountsGroupByResourceTypeResponse creates a response to parse from GetAggregateResourceCountsGroupByResourceType response
func CreateGetAggregateResourceCountsGroupByResourceTypeResponse() (response *GetAggregateResourceCountsGroupByResourceTypeResponse) {
	response = &GetAggregateResourceCountsGroupByResourceTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
