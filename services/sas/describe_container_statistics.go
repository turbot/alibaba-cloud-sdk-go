package sas

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

// DescribeContainerStatistics invokes the sas.DescribeContainerStatistics API synchronously
func (client *Client) DescribeContainerStatistics(request *DescribeContainerStatisticsRequest) (response *DescribeContainerStatisticsResponse, err error) {
	response = CreateDescribeContainerStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeContainerStatisticsWithChan invokes the sas.DescribeContainerStatistics API asynchronously
func (client *Client) DescribeContainerStatisticsWithChan(request *DescribeContainerStatisticsRequest) (<-chan *DescribeContainerStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeContainerStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeContainerStatistics(request)
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

// DescribeContainerStatisticsWithCallback invokes the sas.DescribeContainerStatistics API asynchronously
func (client *Client) DescribeContainerStatisticsWithCallback(request *DescribeContainerStatisticsRequest, callback func(response *DescribeContainerStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeContainerStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeContainerStatistics(request)
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

// DescribeContainerStatisticsRequest is the request struct for api DescribeContainerStatistics
type DescribeContainerStatisticsRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
	SourceIp  string `position:"Query" name:"SourceIp"`
}

// DescribeContainerStatisticsResponse is the response struct for api DescribeContainerStatistics
type DescribeContainerStatisticsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeContainerStatisticsRequest creates a request to invoke DescribeContainerStatistics API
func CreateDescribeContainerStatisticsRequest() (request *DescribeContainerStatisticsRequest) {
	request = &DescribeContainerStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeContainerStatistics", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeContainerStatisticsResponse creates a response to parse from DescribeContainerStatistics response
func CreateDescribeContainerStatisticsResponse() (response *DescribeContainerStatisticsResponse) {
	response = &DescribeContainerStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
