package ddoscoo

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

// CreateSceneDefensePolicy invokes the ddoscoo.CreateSceneDefensePolicy API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/createscenedefensepolicy.html
func (client *Client) CreateSceneDefensePolicy(request *CreateSceneDefensePolicyRequest) (response *CreateSceneDefensePolicyResponse, err error) {
	response = CreateCreateSceneDefensePolicyResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSceneDefensePolicyWithChan invokes the ddoscoo.CreateSceneDefensePolicy API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/createscenedefensepolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSceneDefensePolicyWithChan(request *CreateSceneDefensePolicyRequest) (<-chan *CreateSceneDefensePolicyResponse, <-chan error) {
	responseChan := make(chan *CreateSceneDefensePolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSceneDefensePolicy(request)
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

// CreateSceneDefensePolicyWithCallback invokes the ddoscoo.CreateSceneDefensePolicy API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/createscenedefensepolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSceneDefensePolicyWithCallback(request *CreateSceneDefensePolicyRequest, callback func(response *CreateSceneDefensePolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSceneDefensePolicyResponse
		var err error
		defer close(result)
		response, err = client.CreateSceneDefensePolicy(request)
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

// CreateSceneDefensePolicyRequest is the request struct for api CreateSceneDefensePolicy
type CreateSceneDefensePolicyRequest struct {
	*requests.RpcRequest
	Template  string           `position:"Query" name:"Template"`
	EndTime   requests.Integer `position:"Query" name:"EndTime"`
	StartTime requests.Integer `position:"Query" name:"StartTime"`
	SourceIp  string           `position:"Query" name:"SourceIp"`
	Name      string           `position:"Query" name:"Name"`
}

// CreateSceneDefensePolicyResponse is the response struct for api CreateSceneDefensePolicy
type CreateSceneDefensePolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateSceneDefensePolicyRequest creates a request to invoke CreateSceneDefensePolicy API
func CreateCreateSceneDefensePolicyRequest() (request *CreateSceneDefensePolicyRequest) {
	request = &CreateSceneDefensePolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "CreateSceneDefensePolicy", "ddoscoo", "openAPI")
	return
}

// CreateCreateSceneDefensePolicyResponse creates a response to parse from CreateSceneDefensePolicy response
func CreateCreateSceneDefensePolicyResponse() (response *CreateSceneDefensePolicyResponse) {
	response = &CreateSceneDefensePolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}