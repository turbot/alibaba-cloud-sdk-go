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

// DescribeCriteria invokes the sas.DescribeCriteria API synchronously
func (client *Client) DescribeCriteria(request *DescribeCriteriaRequest) (response *DescribeCriteriaResponse, err error) {
	response = CreateDescribeCriteriaResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCriteriaWithChan invokes the sas.DescribeCriteria API asynchronously
func (client *Client) DescribeCriteriaWithChan(request *DescribeCriteriaRequest) (<-chan *DescribeCriteriaResponse, <-chan error) {
	responseChan := make(chan *DescribeCriteriaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCriteria(request)
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

// DescribeCriteriaWithCallback invokes the sas.DescribeCriteria API asynchronously
func (client *Client) DescribeCriteriaWithCallback(request *DescribeCriteriaRequest, callback func(response *DescribeCriteriaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCriteriaResponse
		var err error
		defer close(result)
		response, err = client.DescribeCriteria(request)
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

// DescribeCriteriaRequest is the request struct for api DescribeCriteria
type DescribeCriteriaRequest struct {
	*requests.RpcRequest
	SupportAutoTag requests.Boolean `position:"Query" name:"SupportAutoTag"`
	SourceIp       string           `position:"Query" name:"SourceIp"`
	Value          string           `position:"Query" name:"Value"`
	MachineTypes   string           `position:"Query" name:"MachineTypes"`
}

// DescribeCriteriaResponse is the response struct for api DescribeCriteria
type DescribeCriteriaResponse struct {
	*responses.BaseResponse
	RequestId    string     `json:"RequestId" xml:"RequestId"`
	CriteriaList []Criteria `json:"CriteriaList" xml:"CriteriaList"`
}

// CreateDescribeCriteriaRequest creates a request to invoke DescribeCriteria API
func CreateDescribeCriteriaRequest() (request *DescribeCriteriaRequest) {
	request = &DescribeCriteriaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeCriteria", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeCriteriaResponse creates a response to parse from DescribeCriteria response
func CreateDescribeCriteriaResponse() (response *DescribeCriteriaResponse) {
	response = &DescribeCriteriaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
