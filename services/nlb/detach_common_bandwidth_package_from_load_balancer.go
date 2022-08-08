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

// DetachCommonBandwidthPackageFromLoadBalancer invokes the nlb.DetachCommonBandwidthPackageFromLoadBalancer API synchronously
func (client *Client) DetachCommonBandwidthPackageFromLoadBalancer(request *DetachCommonBandwidthPackageFromLoadBalancerRequest) (response *DetachCommonBandwidthPackageFromLoadBalancerResponse, err error) {
	response = CreateDetachCommonBandwidthPackageFromLoadBalancerResponse()
	err = client.DoAction(request, response)
	return
}

// DetachCommonBandwidthPackageFromLoadBalancerWithChan invokes the nlb.DetachCommonBandwidthPackageFromLoadBalancer API asynchronously
func (client *Client) DetachCommonBandwidthPackageFromLoadBalancerWithChan(request *DetachCommonBandwidthPackageFromLoadBalancerRequest) (<-chan *DetachCommonBandwidthPackageFromLoadBalancerResponse, <-chan error) {
	responseChan := make(chan *DetachCommonBandwidthPackageFromLoadBalancerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachCommonBandwidthPackageFromLoadBalancer(request)
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

// DetachCommonBandwidthPackageFromLoadBalancerWithCallback invokes the nlb.DetachCommonBandwidthPackageFromLoadBalancer API asynchronously
func (client *Client) DetachCommonBandwidthPackageFromLoadBalancerWithCallback(request *DetachCommonBandwidthPackageFromLoadBalancerRequest, callback func(response *DetachCommonBandwidthPackageFromLoadBalancerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachCommonBandwidthPackageFromLoadBalancerResponse
		var err error
		defer close(result)
		response, err = client.DetachCommonBandwidthPackageFromLoadBalancer(request)
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

// DetachCommonBandwidthPackageFromLoadBalancerRequest is the request struct for api DetachCommonBandwidthPackageFromLoadBalancer
type DetachCommonBandwidthPackageFromLoadBalancerRequest struct {
	*requests.RpcRequest
	ClientToken        string           `position:"Body" name:"ClientToken"`
	BandwidthPackageId string           `position:"Body" name:"BandwidthPackageId"`
	DryRun             requests.Boolean `position:"Body" name:"DryRun"`
	LoadBalancerId     string           `position:"Body" name:"LoadBalancerId"`
}

// DetachCommonBandwidthPackageFromLoadBalancerResponse is the response struct for api DetachCommonBandwidthPackageFromLoadBalancer
type DetachCommonBandwidthPackageFromLoadBalancerResponse struct {
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

// CreateDetachCommonBandwidthPackageFromLoadBalancerRequest creates a request to invoke DetachCommonBandwidthPackageFromLoadBalancer API
func CreateDetachCommonBandwidthPackageFromLoadBalancerRequest() (request *DetachCommonBandwidthPackageFromLoadBalancerRequest) {
	request = &DetachCommonBandwidthPackageFromLoadBalancerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Nlb", "2022-04-30", "DetachCommonBandwidthPackageFromLoadBalancer", "nlb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetachCommonBandwidthPackageFromLoadBalancerResponse creates a response to parse from DetachCommonBandwidthPackageFromLoadBalancer response
func CreateDetachCommonBandwidthPackageFromLoadBalancerResponse() (response *DetachCommonBandwidthPackageFromLoadBalancerResponse) {
	response = &DetachCommonBandwidthPackageFromLoadBalancerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
