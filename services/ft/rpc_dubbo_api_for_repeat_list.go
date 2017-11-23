
package ft

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

func (client *Client) RpcDubboApiForRepeatList(request *RpcDubboApiForRepeatListRequest) (response *RpcDubboApiForRepeatListResponse, err error) {
response = CreateRpcDubboApiForRepeatListResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) RpcDubboApiForRepeatListWithChan(request *RpcDubboApiForRepeatListRequest) (<-chan *RpcDubboApiForRepeatListResponse, <-chan error) {
responseChan := make(chan *RpcDubboApiForRepeatListResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.RpcDubboApiForRepeatList(request)
responseChan <- response
errChan <- err
})
if err != nil {
errChan <- err
close(responseChan)
close(errChan)
}
return responseChan, errChan
}

func (client *Client) RpcDubboApiForRepeatListWithCallback(request *RpcDubboApiForRepeatListRequest, callback func(response *RpcDubboApiForRepeatListResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *RpcDubboApiForRepeatListResponse
var err error
defer close(result)
response, err = client.RpcDubboApiForRepeatList(request)
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

type RpcDubboApiForRepeatListRequest struct {
*requests.RpcRequest
            SwitchValue  string `position:"Query" name:"SwitchValue"`
            MixedRepeatListDisk  *[]RpcDubboApiForRepeatListMixedRepeatListDisk `position:"Query" name:"MixedRepeatListDisk"  type:"Repeated"`
            StringValue  string `position:"Query" name:"StringValue"`
            EnumValue  string `position:"Query" name:"EnumValue"`
            StandardRepeatListDisk  *[]RpcDubboApiForRepeatListStandardRepeatListDisk `position:"Query" name:"StandardRepeatListDisk"  type:"Repeated"`
            OtherParam  string `position:"Query" name:"OtherParam"`
            ResultSwitchValue  string `position:"Query" name:"ResultSwitchValue"`
            Code  string `position:"Query" name:"Code"`
            IntValue  string `position:"Query" name:"IntValue"`
            NumberRange  string `position:"Query" name:"NumberRange"`
            JsonObject  string `position:"Query" name:"JsonObject"`
            JsonRepeatList  *[]RpcDubboApiForRepeatListJsonRepeatList `position:"Query" name:"JsonRepeatList"  type:"Repeated"`
            Message  string `position:"Query" name:"Message"`
            HttpStatusCode  string `position:"Query" name:"HttpStatusCode"`
            RegionId  string `position:"Query" name:"RegionId"`
            ListDisk  *[]string `position:"Query" name:"ListDisk"  type:"Repeated"`
            NestedRepeatListDisk  *[]RpcDubboApiForRepeatListNestedRepeatListDisk `position:"Query" name:"NestedRepeatListDisk"  type:"Repeated"`
            Success  string `position:"Query" name:"Success"`
            DefaultValue  string `position:"Query" name:"DefaultValue"`
            SimpleRepeatListDisk  *[]string `position:"Query" name:"SimpleRepeatListDisk"  type:"Repeated"`
            RequiredValue  string `position:"Query" name:"RequiredValue"`
}

type RpcDubboApiForRepeatListMixedRepeatListDisk struct{
        Size string `name:"Size"`
        Type *[]string `name:"Type" type:"Repeated"`
}
type RpcDubboApiForRepeatListStandardRepeatListDisk struct{
        Size string `name:"Size"`
        Type string `name:"Type"`
}
type RpcDubboApiForRepeatListJsonRepeatList struct{
        Size string `name:"Size"`
        Type string `name:"Type"`
}
type RpcDubboApiForRepeatListNestedRepeatListDisk struct{
        Size *[]string `name:"Size" type:"Repeated"`
        Type *[]string `name:"Type" type:"Repeated"`
}

type RpcDubboApiForRepeatListResponse struct {
*responses.BaseResponse
            IntValue     string `json:"IntValue"`
            NumberRange     string `json:"NumberRange"`
            StringValue     string `json:"StringValue"`
            SwitchValue     string `json:"SwitchValue"`
            EnumValue     string `json:"EnumValue"`
            RequiredValue     string `json:"RequiredValue"`
            DefaultValue     string `json:"DefaultValue"`
            Success     string `json:"Success"`
            Code     string `json:"Code"`
            Message     string `json:"Message"`
            HttpStatusCode     string `json:"HttpStatusCode"`
            NullToEmptyValue     string `json:"NullToEmptyValue"`
            ResultSwitchValue     string `json:"ResultSwitchValue"`
            RegionId     string `json:"RegionId"`
            Name     string `json:"Name"`
            Age     string `json:"Age"`
            JsonRepeatList     string `json:"JsonRepeatList"`
            ListDisks     string `json:"ListDisks"`
            SimpleRepeatListDisks     string `json:"SimpleRepeatListDisks"`
            StandardRepeatListDisks     string `json:"StandardRepeatListDisks"`
            MixedRepeatListDisks     string `json:"MixedRepeatListDisks"`
            NestedRepeatListDisks     string `json:"NestedRepeatListDisks"`
            NullToEmptyStructValue struct {
            NullToEmptyStructChildValue     string `json:"NullToEmptyStructChildValue"`
            }  `json:"NullToEmptyStructValue"`
            StructValue struct {
            StructChildValue     string `json:"StructChildValue"`
            }  `json:"StructValue"`
            ArrayValue  []struct {
            ArrayChildValue     string `json:"ArrayChildValue"`
            }  `json:"ArrayValue"`
}

func CreateRpcDubboApiForRepeatListRequest() (request *RpcDubboApiForRepeatListRequest) {
request = &RpcDubboApiForRepeatListRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ft", "2015-01-01", "RpcDubboApiForRepeatList", "", "")
return
}

func CreateRpcDubboApiForRepeatListResponse() (response *RpcDubboApiForRepeatListResponse) {
response = &RpcDubboApiForRepeatListResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

