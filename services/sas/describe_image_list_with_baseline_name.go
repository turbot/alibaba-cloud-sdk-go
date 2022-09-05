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

// DescribeImageListWithBaselineName invokes the sas.DescribeImageListWithBaselineName API synchronously
func (client *Client) DescribeImageListWithBaselineName(request *DescribeImageListWithBaselineNameRequest) (response *DescribeImageListWithBaselineNameResponse, err error) {
	response = CreateDescribeImageListWithBaselineNameResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeImageListWithBaselineNameWithChan invokes the sas.DescribeImageListWithBaselineName API asynchronously
func (client *Client) DescribeImageListWithBaselineNameWithChan(request *DescribeImageListWithBaselineNameRequest) (<-chan *DescribeImageListWithBaselineNameResponse, <-chan error) {
	responseChan := make(chan *DescribeImageListWithBaselineNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeImageListWithBaselineName(request)
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

// DescribeImageListWithBaselineNameWithCallback invokes the sas.DescribeImageListWithBaselineName API asynchronously
func (client *Client) DescribeImageListWithBaselineNameWithCallback(request *DescribeImageListWithBaselineNameRequest, callback func(response *DescribeImageListWithBaselineNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeImageListWithBaselineNameResponse
		var err error
		defer close(result)
		response, err = client.DescribeImageListWithBaselineName(request)
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

// DescribeImageListWithBaselineNameRequest is the request struct for api DescribeImageListWithBaselineName
type DescribeImageListWithBaselineNameRequest struct {
	*requests.RpcRequest
	Criteria        string           `position:"Query" name:"Criteria"`
	RepoNamespace   string           `position:"Query" name:"RepoNamespace"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	ImageDigest     string           `position:"Query" name:"ImageDigest"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	CriteriaType    string           `position:"Query" name:"CriteriaType"`
	Lang            string           `position:"Query" name:"Lang"`
	CurrentPage     requests.Integer `position:"Query" name:"CurrentPage"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	RepoName        string           `position:"Query" name:"RepoName"`
	BaselineNameKey string           `position:"Query" name:"BaselineNameKey"`
	RepoInstanceId  string           `position:"Query" name:"RepoInstanceId"`
}

// DescribeImageListWithBaselineNameResponse is the response struct for api DescribeImageListWithBaselineName
type DescribeImageListWithBaselineNameResponse struct {
	*responses.BaseResponse
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	PageInfo   PageInfo         `json:"PageInfo" xml:"PageInfo"`
	ImageInfos []ImageInfosItem `json:"ImageInfos" xml:"ImageInfos"`
}

// CreateDescribeImageListWithBaselineNameRequest creates a request to invoke DescribeImageListWithBaselineName API
func CreateDescribeImageListWithBaselineNameRequest() (request *DescribeImageListWithBaselineNameRequest) {
	request = &DescribeImageListWithBaselineNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeImageListWithBaselineName", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeImageListWithBaselineNameResponse creates a response to parse from DescribeImageListWithBaselineName response
func CreateDescribeImageListWithBaselineNameResponse() (response *DescribeImageListWithBaselineNameResponse) {
	response = &DescribeImageListWithBaselineNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
