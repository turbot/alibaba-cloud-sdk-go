package domain

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

// QueryQualificationDetail invokes the domain.QueryQualificationDetail API synchronously
func (client *Client) QueryQualificationDetail(request *QueryQualificationDetailRequest) (response *QueryQualificationDetailResponse, err error) {
	response = CreateQueryQualificationDetailResponse()
	err = client.DoAction(request, response)
	return
}

// QueryQualificationDetailWithChan invokes the domain.QueryQualificationDetail API asynchronously
func (client *Client) QueryQualificationDetailWithChan(request *QueryQualificationDetailRequest) (<-chan *QueryQualificationDetailResponse, <-chan error) {
	responseChan := make(chan *QueryQualificationDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryQualificationDetail(request)
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

// QueryQualificationDetailWithCallback invokes the domain.QueryQualificationDetail API asynchronously
func (client *Client) QueryQualificationDetailWithCallback(request *QueryQualificationDetailRequest, callback func(response *QueryQualificationDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryQualificationDetailResponse
		var err error
		defer close(result)
		response, err = client.QueryQualificationDetail(request)
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

// QueryQualificationDetailRequest is the request struct for api QueryQualificationDetail
type QueryQualificationDetailRequest struct {
	*requests.RpcRequest
	QualificationType string `position:"Query" name:"QualificationType"`
	InstanceId        string `position:"Query" name:"InstanceId"`
	UserClientIp      string `position:"Query" name:"UserClientIp"`
	Lang              string `position:"Query" name:"Lang"`
}

// QueryQualificationDetailResponse is the response struct for api QueryQualificationDetail
type QueryQualificationDetailResponse struct {
	*responses.BaseResponse
	TrackId     string      `json:"TrackId" xml:"TrackId"`
	AuditStatus int         `json:"AuditStatus" xml:"AuditStatus"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	Credentials Credentials `json:"Credentials" xml:"Credentials"`
}

// CreateQueryQualificationDetailRequest creates a request to invoke QueryQualificationDetail API
func CreateQueryQualificationDetailRequest() (request *QueryQualificationDetailRequest) {
	request = &QueryQualificationDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryQualificationDetail", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryQualificationDetailResponse creates a response to parse from QueryQualificationDetail response
func CreateQueryQualificationDetailResponse() (response *QueryQualificationDetailResponse) {
	response = &QueryQualificationDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
