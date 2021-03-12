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

// SubmitOperationCredentials invokes the domain.SubmitOperationCredentials API synchronously
func (client *Client) SubmitOperationCredentials(request *SubmitOperationCredentialsRequest) (response *SubmitOperationCredentialsResponse, err error) {
	response = CreateSubmitOperationCredentialsResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitOperationCredentialsWithChan invokes the domain.SubmitOperationCredentials API asynchronously
func (client *Client) SubmitOperationCredentialsWithChan(request *SubmitOperationCredentialsRequest) (<-chan *SubmitOperationCredentialsResponse, <-chan error) {
	responseChan := make(chan *SubmitOperationCredentialsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitOperationCredentials(request)
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

// SubmitOperationCredentialsWithCallback invokes the domain.SubmitOperationCredentials API asynchronously
func (client *Client) SubmitOperationCredentialsWithCallback(request *SubmitOperationCredentialsRequest, callback func(response *SubmitOperationCredentialsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitOperationCredentialsResponse
		var err error
		defer close(result)
		response, err = client.SubmitOperationCredentials(request)
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

// SubmitOperationCredentialsRequest is the request struct for api SubmitOperationCredentials
type SubmitOperationCredentialsRequest struct {
	*requests.RpcRequest
	Credentials   string           `position:"Query" name:"Credentials"`
	AuditRecordId requests.Integer `position:"Query" name:"AuditRecordId"`
	RegType       requests.Integer `position:"Query" name:"RegType"`
	AuditType     requests.Integer `position:"Query" name:"AuditType"`
	Lang          string           `position:"Query" name:"Lang"`
}

// SubmitOperationCredentialsResponse is the response struct for api SubmitOperationCredentials
type SubmitOperationCredentialsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSubmitOperationCredentialsRequest creates a request to invoke SubmitOperationCredentials API
func CreateSubmitOperationCredentialsRequest() (request *SubmitOperationCredentialsRequest) {
	request = &SubmitOperationCredentialsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SubmitOperationCredentials", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitOperationCredentialsResponse creates a response to parse from SubmitOperationCredentials response
func CreateSubmitOperationCredentialsResponse() (response *SubmitOperationCredentialsResponse) {
	response = &SubmitOperationCredentialsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
