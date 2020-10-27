package cloudesl

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

// BatchInsertItems invokes the cloudesl.BatchInsertItems API synchronously
func (client *Client) BatchInsertItems(request *BatchInsertItemsRequest) (response *BatchInsertItemsResponse, err error) {
	response = CreateBatchInsertItemsResponse()
	err = client.DoAction(request, response)
	return
}

// BatchInsertItemsWithChan invokes the cloudesl.BatchInsertItems API asynchronously
func (client *Client) BatchInsertItemsWithChan(request *BatchInsertItemsRequest) (<-chan *BatchInsertItemsResponse, <-chan error) {
	responseChan := make(chan *BatchInsertItemsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchInsertItems(request)
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

// BatchInsertItemsWithCallback invokes the cloudesl.BatchInsertItems API asynchronously
func (client *Client) BatchInsertItemsWithCallback(request *BatchInsertItemsRequest, callback func(response *BatchInsertItemsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchInsertItemsResponse
		var err error
		defer close(result)
		response, err = client.BatchInsertItems(request)
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

// BatchInsertItemsRequest is the request struct for api BatchInsertItems
type BatchInsertItemsRequest struct {
	*requests.RpcRequest
	ExtraParams string                      `position:"Body" name:"ExtraParams"`
	StoreId     string                      `position:"Body" name:"StoreId"`
	ItemInfo    *[]BatchInsertItemsItemInfo `position:"Body" name:"ItemInfo"  type:"Repeated"`
}

// BatchInsertItemsItemInfo is a repeated param struct in BatchInsertItemsRequest
type BatchInsertItemsItemInfo struct {
	MemberPrice       string `name:"MemberPrice"`
	ActionPrice       string `name:"ActionPrice"`
	BeSourceCode      string `name:"BeSourceCode"`
	BrandName         string `name:"BrandName"`
	PromotionStart    string `name:"PromotionStart"`
	PriceUnit         string `name:"PriceUnit"`
	Rank              string `name:"Rank"`
	ItemInfoIndex     string `name:"ItemInfoIndex"`
	ItemBarCode       string `name:"ItemBarCode"`
	CustomizeFeatureK string `name:"CustomizeFeatureK"`
	CustomizeFeatureL string `name:"CustomizeFeatureL"`
	CustomizeFeatureM string `name:"CustomizeFeatureM"`
	BePromotion       string `name:"BePromotion"`
	CustomizeFeatureN string `name:"CustomizeFeatureN"`
	CustomizeFeatureO string `name:"CustomizeFeatureO"`
	PromotionEnd      string `name:"PromotionEnd"`
	ItemTitle         string `name:"ItemTitle"`
	CustomizeFeatureC string `name:"CustomizeFeatureC"`
	CustomizeFeatureD string `name:"CustomizeFeatureD"`
	ItemQrCode        string `name:"ItemQrCode"`
	CustomizeFeatureE string `name:"CustomizeFeatureE"`
	InventoryStatus   string `name:"InventoryStatus"`
	PromotionReason   string `name:"PromotionReason"`
	CustomizeFeatureF string `name:"CustomizeFeatureF"`
	CustomizeFeatureG string `name:"CustomizeFeatureG"`
	CustomizeFeatureH string `name:"CustomizeFeatureH"`
	CustomizeFeatureI string `name:"CustomizeFeatureI"`
	CustomizeFeatureJ string `name:"CustomizeFeatureJ"`
	CustomizeFeatureA string `name:"CustomizeFeatureA"`
	CustomizeFeatureB string `name:"CustomizeFeatureB"`
	SuggestPrice      string `name:"SuggestPrice"`
	ForestFirstId     string `name:"ForestFirstId"`
	ProductionPlace   string `name:"ProductionPlace"`
	Manufacturer      string `name:"Manufacturer"`
	SourceCode        string `name:"SourceCode"`
	ItemId            string `name:"ItemId"`
	SalesPrice        string `name:"SalesPrice"`
	OriginalPrice     string `name:"OriginalPrice"`
	ItemShortTitle    string `name:"ItemShortTitle"`
	ForestSecondId    string `name:"ForestSecondId"`
	ItemPicUrl        string `name:"ItemPicUrl"`
	SupplierName      string `name:"SupplierName"`
	Material          string `name:"Material"`
	ModelNumber       string `name:"ModelNumber"`
	SaleSpec          string `name:"SaleSpec"`
	CategoryName      string `name:"CategoryName"`
	TaxFee            string `name:"TaxFee"`
	EnergyEfficiency  string `name:"EnergyEfficiency"`
	PromotionText     string `name:"PromotionText"`
	SkuId             string `name:"SkuId"`
}

// BatchInsertItemsResponse is the response struct for api BatchInsertItems
type BatchInsertItemsResponse struct {
	*responses.BaseResponse
	ErrorMessage   string        `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string        `json:"ErrorCode" xml:"ErrorCode"`
	Message        string        `json:"Message" xml:"Message"`
	DynamicCode    string        `json:"DynamicCode" xml:"DynamicCode"`
	Code           string        `json:"Code" xml:"Code"`
	DynamicMessage string        `json:"DynamicMessage" xml:"DynamicMessage"`
	RequestId      string        `json:"RequestId" xml:"RequestId"`
	Success        bool          `json:"Success" xml:"Success"`
	BatchResults   []BatchResult `json:"BatchResults" xml:"BatchResults"`
}

// CreateBatchInsertItemsRequest creates a request to invoke BatchInsertItems API
func CreateBatchInsertItemsRequest() (request *BatchInsertItemsRequest) {
	request = &BatchInsertItemsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "BatchInsertItems", "cloudesl", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBatchInsertItemsResponse creates a response to parse from BatchInsertItems response
func CreateBatchInsertItemsResponse() (response *BatchInsertItemsResponse) {
	response = &BatchInsertItemsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
