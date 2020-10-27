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

// DescribeUserLog invokes the cloudesl.DescribeUserLog API synchronously
func (client *Client) DescribeUserLog(request *DescribeUserLogRequest) (response *DescribeUserLogResponse, err error) {
	response = CreateDescribeUserLogResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserLogWithChan invokes the cloudesl.DescribeUserLog API asynchronously
func (client *Client) DescribeUserLogWithChan(request *DescribeUserLogRequest) (<-chan *DescribeUserLogResponse, <-chan error) {
	responseChan := make(chan *DescribeUserLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserLog(request)
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

// DescribeUserLogWithCallback invokes the cloudesl.DescribeUserLog API asynchronously
func (client *Client) DescribeUserLogWithCallback(request *DescribeUserLogRequest, callback func(response *DescribeUserLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserLogResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserLog(request)
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

// DescribeUserLogRequest is the request struct for api DescribeUserLog
type DescribeUserLogRequest struct {
	*requests.RpcRequest
	ExtraParams     string           `position:"Body" name:"ExtraParams"`
	StoreId         string           `position:"Body" name:"StoreId"`
	UserId          string           `position:"Body" name:"UserId"`
	PageNumber      requests.Integer `position:"Body" name:"PageNumber"`
	FromDate        string           `position:"Body" name:"FromDate"`
	OperationStatus string           `position:"Body" name:"OperationStatus"`
	ToDate          string           `position:"Body" name:"ToDate"`
	EslBarCode      string           `position:"Body" name:"EslBarCode"`
	PageSize        requests.Integer `position:"Body" name:"PageSize"`
	ItemBarCode     string           `position:"Body" name:"ItemBarCode"`
	ItemShortTitle  string           `position:"Body" name:"ItemShortTitle"`
	OperationType   string           `position:"Body" name:"OperationType"`
	LogId           string           `position:"Body" name:"LogId"`
}

// DescribeUserLogResponse is the response struct for api DescribeUserLog
type DescribeUserLogResponse struct {
	*responses.BaseResponse
	ErrorMessage   string        `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string        `json:"ErrorCode" xml:"ErrorCode"`
	PageNumber     int           `json:"PageNumber" xml:"PageNumber"`
	Message        string        `json:"Message" xml:"Message"`
	DynamicCode    string        `json:"DynamicCode" xml:"DynamicCode"`
	Code           string        `json:"Code" xml:"Code"`
	PageSize       int           `json:"PageSize" xml:"PageSize"`
	DynamicMessage string        `json:"DynamicMessage" xml:"DynamicMessage"`
	RequestId      string        `json:"RequestId" xml:"RequestId"`
	Success        bool          `json:"Success" xml:"Success"`
	TotalCount     int           `json:"TotalCount" xml:"TotalCount"`
	UserLogs       []UserLogInfo `json:"UserLogs" xml:"UserLogs"`
}

// CreateDescribeUserLogRequest creates a request to invoke DescribeUserLog API
func CreateDescribeUserLogRequest() (request *DescribeUserLogRequest) {
	request = &DescribeUserLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "DescribeUserLog", "cloudesl", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeUserLogResponse creates a response to parse from DescribeUserLog response
func CreateDescribeUserLogResponse() (response *DescribeUserLogResponse) {
	response = &DescribeUserLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
