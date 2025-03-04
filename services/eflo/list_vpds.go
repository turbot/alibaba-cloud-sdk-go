package eflo

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

// ListVpds invokes the eflo.ListVpds API synchronously
func (client *Client) ListVpds(request *ListVpdsRequest) (response *ListVpdsResponse, err error) {
	response = CreateListVpdsResponse()
	err = client.DoAction(request, response)
	return
}

// ListVpdsWithChan invokes the eflo.ListVpds API asynchronously
func (client *Client) ListVpdsWithChan(request *ListVpdsRequest) (<-chan *ListVpdsResponse, <-chan error) {
	responseChan := make(chan *ListVpdsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVpds(request)
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

// ListVpdsWithCallback invokes the eflo.ListVpds API asynchronously
func (client *Client) ListVpdsWithCallback(request *ListVpdsRequest, callback func(response *ListVpdsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVpdsResponse
		var err error
		defer close(result)
		response, err = client.ListVpds(request)
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

// ListVpdsRequest is the request struct for api ListVpds
type ListVpdsRequest struct {
	*requests.RpcRequest
	PageNumber     requests.Integer `position:"Body" name:"PageNumber"`
	WithDependence requests.Boolean `position:"Body" name:"WithDependence"`
	WithoutVcc     requests.Boolean `position:"Body" name:"WithoutVcc"`
	PageSize       requests.Integer `position:"Body" name:"PageSize"`
	ForSelect      requests.Boolean `position:"Body" name:"ForSelect"`
	FilterErId     string           `position:"Body" name:"FilterErId"`
	VpdId          string           `position:"Body" name:"VpdId"`
	EnablePage     requests.Boolean `position:"Body" name:"EnablePage"`
	Name           string           `position:"Body" name:"Name"`
	Status         string           `position:"Body" name:"Status"`
}

// ListVpdsResponse is the response struct for api ListVpds
type ListVpdsResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateListVpdsRequest creates a request to invoke ListVpds API
func CreateListVpdsRequest() (request *ListVpdsRequest) {
	request = &ListVpdsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "ListVpds", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListVpdsResponse creates a response to parse from ListVpds response
func CreateListVpdsResponse() (response *ListVpdsResponse) {
	response = &ListVpdsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
