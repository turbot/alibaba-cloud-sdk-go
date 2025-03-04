package config

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

// ConfigRulesItem is a nested struct in config response
type ConfigRulesItem struct {
	RiskLevel             int                        `json:"RiskLevel" xml:"RiskLevel"`
	ConfigRuleName        string                     `json:"ConfigRuleName" xml:"ConfigRuleName"`
	DeveloperType         string                     `json:"DeveloperType" xml:"DeveloperType"`
	ManagedRuleIdentifier string                     `json:"ManagedRuleIdentifier" xml:"ManagedRuleIdentifier"`
	Description           string                     `json:"Description" xml:"Description"`
	ManagedRuleName       string                     `json:"ManagedRuleName" xml:"ManagedRuleName"`
	DefaultEnable         bool                       `json:"DefaultEnable" xml:"DefaultEnable"`
	ConfigRuleId          string                     `json:"ConfigRuleId" xml:"ConfigRuleId"`
	ControlId             string                     `json:"ControlId" xml:"ControlId"`
	ControlDescription    string                     `json:"ControlDescription" xml:"ControlDescription"`
	ConfigRuleParameters  []ConfigRuleParametersItem `json:"ConfigRuleParameters" xml:"ConfigRuleParameters"`
}
