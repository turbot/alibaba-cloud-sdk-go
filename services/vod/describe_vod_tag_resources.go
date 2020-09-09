package vod

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

// DescribeVodTagResources invokes the vod.DescribeVodTagResources API synchronously
func (client *Client) DescribeVodTagResources(request *DescribeVodTagResourcesRequest) (response *DescribeVodTagResourcesResponse, err error) {
	response = CreateDescribeVodTagResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVodTagResourcesWithChan invokes the vod.DescribeVodTagResources API asynchronously
func (client *Client) DescribeVodTagResourcesWithChan(request *DescribeVodTagResourcesRequest) (<-chan *DescribeVodTagResourcesResponse, <-chan error) {
	responseChan := make(chan *DescribeVodTagResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVodTagResources(request)
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

// DescribeVodTagResourcesWithCallback invokes the vod.DescribeVodTagResources API asynchronously
func (client *Client) DescribeVodTagResourcesWithCallback(request *DescribeVodTagResourcesRequest, callback func(response *DescribeVodTagResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVodTagResourcesResponse
		var err error
		defer close(result)
		response, err = client.DescribeVodTagResources(request)
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

// DescribeVodTagResourcesRequest is the request struct for api DescribeVodTagResources
type DescribeVodTagResourcesRequest struct {
	*requests.RpcRequest
	Scope        string                        `position:"Query" name:"Scope"`
	Tag          *[]DescribeVodTagResourcesTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceId   *[]string                     `position:"Query" name:"ResourceId"  type:"Repeated"`
	OwnerId      requests.Integer              `position:"Query" name:"OwnerId"`
	ResourceType string                        `position:"Query" name:"ResourceType"`
}

// DescribeVodTagResourcesTag is a repeated param struct in DescribeVodTagResourcesRequest
type DescribeVodTagResourcesTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// DescribeVodTagResourcesResponse is the response struct for api DescribeVodTagResources
type DescribeVodTagResourcesResponse struct {
	*responses.BaseResponse
	RequestId    string        `json:"RequestId" xml:"RequestId"`
	TagResources []TagResource `json:"TagResources" xml:"TagResources"`
}

// CreateDescribeVodTagResourcesRequest creates a request to invoke DescribeVodTagResources API
func CreateDescribeVodTagResourcesRequest() (request *DescribeVodTagResourcesRequest) {
	request = &DescribeVodTagResourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DescribeVodTagResources", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVodTagResourcesResponse creates a response to parse from DescribeVodTagResources response
func CreateDescribeVodTagResourcesResponse() (response *DescribeVodTagResourcesResponse) {
	response = &DescribeVodTagResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
