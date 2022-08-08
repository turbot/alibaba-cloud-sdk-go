package nlb

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

// AttachCommonBandwidthPackageToLoadBalancer invokes the nlb.AttachCommonBandwidthPackageToLoadBalancer API synchronously
func (client *Client) AttachCommonBandwidthPackageToLoadBalancer(request *AttachCommonBandwidthPackageToLoadBalancerRequest) (response *AttachCommonBandwidthPackageToLoadBalancerResponse, err error) {
	response = CreateAttachCommonBandwidthPackageToLoadBalancerResponse()
	err = client.DoAction(request, response)
	return
}

// AttachCommonBandwidthPackageToLoadBalancerWithChan invokes the nlb.AttachCommonBandwidthPackageToLoadBalancer API asynchronously
func (client *Client) AttachCommonBandwidthPackageToLoadBalancerWithChan(request *AttachCommonBandwidthPackageToLoadBalancerRequest) (<-chan *AttachCommonBandwidthPackageToLoadBalancerResponse, <-chan error) {
	responseChan := make(chan *AttachCommonBandwidthPackageToLoadBalancerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachCommonBandwidthPackageToLoadBalancer(request)
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

// AttachCommonBandwidthPackageToLoadBalancerWithCallback invokes the nlb.AttachCommonBandwidthPackageToLoadBalancer API asynchronously
func (client *Client) AttachCommonBandwidthPackageToLoadBalancerWithCallback(request *AttachCommonBandwidthPackageToLoadBalancerRequest, callback func(response *AttachCommonBandwidthPackageToLoadBalancerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachCommonBandwidthPackageToLoadBalancerResponse
		var err error
		defer close(result)
		response, err = client.AttachCommonBandwidthPackageToLoadBalancer(request)
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

// AttachCommonBandwidthPackageToLoadBalancerRequest is the request struct for api AttachCommonBandwidthPackageToLoadBalancer
type AttachCommonBandwidthPackageToLoadBalancerRequest struct {
	*requests.RpcRequest
	ClientToken        string           `position:"Body" name:"ClientToken"`
	BandwidthPackageId string           `position:"Body" name:"BandwidthPackageId"`
	DryRun             requests.Boolean `position:"Body" name:"DryRun"`
	LoadBalancerId     string           `position:"Body" name:"LoadBalancerId"`
}

// AttachCommonBandwidthPackageToLoadBalancerResponse is the response struct for api AttachCommonBandwidthPackageToLoadBalancer
type AttachCommonBandwidthPackageToLoadBalancerResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	JobId          string `json:"JobId" xml:"JobId"`
}

// CreateAttachCommonBandwidthPackageToLoadBalancerRequest creates a request to invoke AttachCommonBandwidthPackageToLoadBalancer API
func CreateAttachCommonBandwidthPackageToLoadBalancerRequest() (request *AttachCommonBandwidthPackageToLoadBalancerRequest) {
	request = &AttachCommonBandwidthPackageToLoadBalancerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Nlb", "2022-04-30", "AttachCommonBandwidthPackageToLoadBalancer", "nlb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAttachCommonBandwidthPackageToLoadBalancerResponse creates a response to parse from AttachCommonBandwidthPackageToLoadBalancer response
func CreateAttachCommonBandwidthPackageToLoadBalancerResponse() (response *AttachCommonBandwidthPackageToLoadBalancerResponse) {
	response = &AttachCommonBandwidthPackageToLoadBalancerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
