package vod

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

// DeleteMultipartUpload invokes the vod.DeleteMultipartUpload API synchronously
func (client *Client) DeleteMultipartUpload(request *DeleteMultipartUploadRequest) (response *DeleteMultipartUploadResponse, err error) {
	response = CreateDeleteMultipartUploadResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMultipartUploadWithChan invokes the vod.DeleteMultipartUpload API asynchronously
func (client *Client) DeleteMultipartUploadWithChan(request *DeleteMultipartUploadRequest) (<-chan *DeleteMultipartUploadResponse, <-chan error) {
	responseChan := make(chan *DeleteMultipartUploadResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMultipartUpload(request)
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

// DeleteMultipartUploadWithCallback invokes the vod.DeleteMultipartUpload API asynchronously
func (client *Client) DeleteMultipartUploadWithCallback(request *DeleteMultipartUploadRequest, callback func(response *DeleteMultipartUploadResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMultipartUploadResponse
		var err error
		defer close(result)
		response, err = client.DeleteMultipartUpload(request)
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

// DeleteMultipartUploadRequest is the request struct for api DeleteMultipartUpload
type DeleteMultipartUploadRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string           `position:"Query" name:"ResourceOwnerId"`
	ResourceRealOwnerId  requests.Integer `position:"Query" name:"ResourceRealOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              string           `position:"Query" name:"OwnerId"`
	MediaId              string           `position:"Query" name:"MediaId"`
	MediaType            string           `position:"Query" name:"MediaType"`
}

// DeleteMultipartUploadResponse is the response struct for api DeleteMultipartUpload
type DeleteMultipartUploadResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteMultipartUploadRequest creates a request to invoke DeleteMultipartUpload API
func CreateDeleteMultipartUploadRequest() (request *DeleteMultipartUploadRequest) {
	request = &DeleteMultipartUploadRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DeleteMultipartUpload", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteMultipartUploadResponse creates a response to parse from DeleteMultipartUpload response
func CreateDeleteMultipartUploadResponse() (response *DeleteMultipartUploadResponse) {
	response = &DeleteMultipartUploadResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
