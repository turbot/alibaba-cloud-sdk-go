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

// DescribeLogstoreStorage invokes the sas.DescribeLogstoreStorage API synchronously
func (client *Client) DescribeLogstoreStorage(request *DescribeLogstoreStorageRequest) (response *DescribeLogstoreStorageResponse, err error) {
	response = CreateDescribeLogstoreStorageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLogstoreStorageWithChan invokes the sas.DescribeLogstoreStorage API asynchronously
func (client *Client) DescribeLogstoreStorageWithChan(request *DescribeLogstoreStorageRequest) (<-chan *DescribeLogstoreStorageResponse, <-chan error) {
	responseChan := make(chan *DescribeLogstoreStorageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLogstoreStorage(request)
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

// DescribeLogstoreStorageWithCallback invokes the sas.DescribeLogstoreStorage API asynchronously
func (client *Client) DescribeLogstoreStorageWithCallback(request *DescribeLogstoreStorageRequest, callback func(response *DescribeLogstoreStorageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLogstoreStorageResponse
		var err error
		defer close(result)
		response, err = client.DescribeLogstoreStorage(request)
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

// DescribeLogstoreStorageRequest is the request struct for api DescribeLogstoreStorage
type DescribeLogstoreStorageRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	From     string `position:"Query" name:"From"`
	Lang     string `position:"Query" name:"Lang"`
}

// DescribeLogstoreStorageResponse is the response struct for api DescribeLogstoreStorage
type DescribeLogstoreStorageResponse struct {
	*responses.BaseResponse
	Used      int64  `json:"Used" xml:"Used"`
	Ttl       int    `json:"Ttl" xml:"Ttl"`
	Logstore  string `json:"Logstore" xml:"Logstore"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Preserve  int64  `json:"Preserve" xml:"Preserve"`
}

// CreateDescribeLogstoreStorageRequest creates a request to invoke DescribeLogstoreStorage API
func CreateDescribeLogstoreStorageRequest() (request *DescribeLogstoreStorageRequest) {
	request = &DescribeLogstoreStorageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeLogstoreStorage", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeLogstoreStorageResponse creates a response to parse from DescribeLogstoreStorage response
func CreateDescribeLogstoreStorageResponse() (response *DescribeLogstoreStorageResponse) {
	response = &DescribeLogstoreStorageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
