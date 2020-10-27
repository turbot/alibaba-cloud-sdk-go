package cloudesl

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

// DescribePlanogramRails invokes the cloudesl.DescribePlanogramRails API synchronously
func (client *Client) DescribePlanogramRails(request *DescribePlanogramRailsRequest) (response *DescribePlanogramRailsResponse, err error) {
	response = CreateDescribePlanogramRailsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePlanogramRailsWithChan invokes the cloudesl.DescribePlanogramRails API asynchronously
func (client *Client) DescribePlanogramRailsWithChan(request *DescribePlanogramRailsRequest) (<-chan *DescribePlanogramRailsResponse, <-chan error) {
	responseChan := make(chan *DescribePlanogramRailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePlanogramRails(request)
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

// DescribePlanogramRailsWithCallback invokes the cloudesl.DescribePlanogramRails API asynchronously
func (client *Client) DescribePlanogramRailsWithCallback(request *DescribePlanogramRailsRequest, callback func(response *DescribePlanogramRailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePlanogramRailsResponse
		var err error
		defer close(result)
		response, err = client.DescribePlanogramRails(request)
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

// DescribePlanogramRailsRequest is the request struct for api DescribePlanogramRails
type DescribePlanogramRailsRequest struct {
	*requests.RpcRequest
	ExtraParams string           `position:"Body" name:"ExtraParams"`
	StoreId     string           `position:"Body" name:"StoreId"`
	Layer       requests.Integer `position:"Body" name:"Layer"`
	PageNumber  requests.Integer `position:"Body" name:"PageNumber"`
	PageSize    requests.Integer `position:"Body" name:"PageSize"`
	Shelf       string           `position:"Body" name:"Shelf"`
	RailCode    string           `position:"Body" name:"RailCode"`
}

// DescribePlanogramRailsResponse is the response struct for api DescribePlanogramRails
type DescribePlanogramRailsResponse struct {
	*responses.BaseResponse
	DynamicMessage     string              `json:"DynamicMessage" xml:"DynamicMessage"`
	RequestId          string              `json:"RequestId" xml:"RequestId"`
	Success            bool                `json:"Success" xml:"Success"`
	ErrorMessage       string              `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode          string              `json:"ErrorCode" xml:"ErrorCode"`
	Message            string              `json:"Message" xml:"Message"`
	PageNumber         int                 `json:"PageNumber" xml:"PageNumber"`
	DynamicCode        string              `json:"DynamicCode" xml:"DynamicCode"`
	StoreId            string              `json:"StoreId" xml:"StoreId"`
	TotalCount         int                 `json:"TotalCount" xml:"TotalCount"`
	Code               string              `json:"Code" xml:"Code"`
	PageSize           int                 `json:"PageSize" xml:"PageSize"`
	PlanogramRailInfos []PlanogramRailInfo `json:"PlanogramRailInfos" xml:"PlanogramRailInfos"`
}

// CreateDescribePlanogramRailsRequest creates a request to invoke DescribePlanogramRails API
func CreateDescribePlanogramRailsRequest() (request *DescribePlanogramRailsRequest) {
	request = &DescribePlanogramRailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "DescribePlanogramRails", "cloudesl", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePlanogramRailsResponse creates a response to parse from DescribePlanogramRails response
func CreateDescribePlanogramRailsResponse() (response *DescribePlanogramRailsResponse) {
	response = &DescribePlanogramRailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
