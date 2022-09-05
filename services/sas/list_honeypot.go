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

// ListHoneypot invokes the sas.ListHoneypot API synchronously
func (client *Client) ListHoneypot(request *ListHoneypotRequest) (response *ListHoneypotResponse, err error) {
	response = CreateListHoneypotResponse()
	err = client.DoAction(request, response)
	return
}

// ListHoneypotWithChan invokes the sas.ListHoneypot API asynchronously
func (client *Client) ListHoneypotWithChan(request *ListHoneypotRequest) (<-chan *ListHoneypotResponse, <-chan error) {
	responseChan := make(chan *ListHoneypotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListHoneypot(request)
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

// ListHoneypotWithCallback invokes the sas.ListHoneypot API asynchronously
func (client *Client) ListHoneypotWithCallback(request *ListHoneypotRequest, callback func(response *ListHoneypotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListHoneypotResponse
		var err error
		defer close(result)
		response, err = client.ListHoneypot(request)
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

// ListHoneypotRequest is the request struct for api ListHoneypot
type ListHoneypotRequest struct {
	*requests.RpcRequest
	HoneypotName string           `position:"Query" name:"HoneypotName"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Lang         string           `position:"Query" name:"Lang"`
	NodeId       string           `position:"Query" name:"NodeId"`
	CurrentPage  requests.Integer `position:"Query" name:"CurrentPage"`
	NodeName     string           `position:"Query" name:"NodeName"`
	HoneypotIds  *[]string        `position:"Query" name:"HoneypotIds"  type:"Repeated"`
}

// ListHoneypotResponse is the response struct for api ListHoneypot
type ListHoneypotResponse struct {
	*responses.BaseResponse
	Success        bool       `json:"Success" xml:"Success"`
	Code           string     `json:"Code" xml:"Code"`
	Message        string     `json:"Message" xml:"Message"`
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PageInfo       PageInfo   `json:"PageInfo" xml:"PageInfo"`
	List           []ListItem `json:"List" xml:"List"`
}

// CreateListHoneypotRequest creates a request to invoke ListHoneypot API
func CreateListHoneypotRequest() (request *ListHoneypotRequest) {
	request = &ListHoneypotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ListHoneypot", "", "")
	request.Method = requests.POST
	return
}

// CreateListHoneypotResponse creates a response to parse from ListHoneypot response
func CreateListHoneypotResponse() (response *ListHoneypotResponse) {
	response = &ListHoneypotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
