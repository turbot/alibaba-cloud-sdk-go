package dytnsapi

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

// Data is a nested struct in dytnsapi response
type Data struct {
	IsNumberPortability bool   `json:"IsNumberPortability" xml:"IsNumberPortability"`
	NumberSegment       int64  `json:"NumberSegment" xml:"NumberSegment"`
	IsConsistent        int    `json:"IsConsistent" xml:"IsConsistent"`
	CarrierCode         string `json:"CarrierCode" xml:"CarrierCode"`
	Carrier             string `json:"Carrier" xml:"Carrier"`
	VerifyResult        string `json:"VerifyResult" xml:"VerifyResult"`
	BasicCarrier        string `json:"BasicCarrier" xml:"BasicCarrier"`
	IsPrivacyNumber     bool   `json:"IsPrivacyNumber" xml:"IsPrivacyNumber"`
	Status              string `json:"Status" xml:"Status"`
	City                string `json:"City" xml:"City"`
	Number              string `json:"Number" xml:"Number"`
	Province            string `json:"Province" xml:"Province"`
}
