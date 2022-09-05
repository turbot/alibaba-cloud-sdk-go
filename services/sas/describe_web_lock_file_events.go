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

// DescribeWebLockFileEvents invokes the sas.DescribeWebLockFileEvents API synchronously
func (client *Client) DescribeWebLockFileEvents(request *DescribeWebLockFileEventsRequest) (response *DescribeWebLockFileEventsResponse, err error) {
	response = CreateDescribeWebLockFileEventsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWebLockFileEventsWithChan invokes the sas.DescribeWebLockFileEvents API asynchronously
func (client *Client) DescribeWebLockFileEventsWithChan(request *DescribeWebLockFileEventsRequest) (<-chan *DescribeWebLockFileEventsResponse, <-chan error) {
	responseChan := make(chan *DescribeWebLockFileEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWebLockFileEvents(request)
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

// DescribeWebLockFileEventsWithCallback invokes the sas.DescribeWebLockFileEvents API asynchronously
func (client *Client) DescribeWebLockFileEventsWithCallback(request *DescribeWebLockFileEventsRequest, callback func(response *DescribeWebLockFileEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWebLockFileEventsResponse
		var err error
		defer close(result)
		response, err = client.DescribeWebLockFileEvents(request)
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

// DescribeWebLockFileEventsRequest is the request struct for api DescribeWebLockFileEvents
type DescribeWebLockFileEventsRequest struct {
	*requests.RpcRequest
	Remark      string           `position:"Query" name:"Remark"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	TsEnd       requests.Integer `position:"Query" name:"TsEnd"`
	ProcessName string           `position:"Query" name:"ProcessName"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	Dealed      string           `position:"Query" name:"Dealed"`
	TsBegin     requests.Integer `position:"Query" name:"TsBegin"`
}

// DescribeWebLockFileEventsResponse is the response struct for api DescribeWebLockFileEvents
type DescribeWebLockFileEventsResponse struct {
	*responses.BaseResponse
	CurrentPage int    `json:"CurrentPage" xml:"CurrentPage"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
	PageSize    int    `json:"PageSize" xml:"PageSize"`
	TotalCount  int    `json:"TotalCount" xml:"TotalCount"`
	List        []Info `json:"List" xml:"List"`
}

// CreateDescribeWebLockFileEventsRequest creates a request to invoke DescribeWebLockFileEvents API
func CreateDescribeWebLockFileEventsRequest() (request *DescribeWebLockFileEventsRequest) {
	request = &DescribeWebLockFileEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeWebLockFileEvents", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeWebLockFileEventsResponse creates a response to parse from DescribeWebLockFileEvents response
func CreateDescribeWebLockFileEventsResponse() (response *DescribeWebLockFileEventsResponse) {
	response = &DescribeWebLockFileEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
