package emr

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

// ListBackupPlans invokes the emr.ListBackupPlans API synchronously
// api document: https://help.aliyun.com/api/emr/listbackupplans.html
func (client *Client) ListBackupPlans(request *ListBackupPlansRequest) (response *ListBackupPlansResponse, err error) {
	response = CreateListBackupPlansResponse()
	err = client.DoAction(request, response)
	return
}

// ListBackupPlansWithChan invokes the emr.ListBackupPlans API asynchronously
// api document: https://help.aliyun.com/api/emr/listbackupplans.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListBackupPlansWithChan(request *ListBackupPlansRequest) (<-chan *ListBackupPlansResponse, <-chan error) {
	responseChan := make(chan *ListBackupPlansResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListBackupPlans(request)
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

// ListBackupPlansWithCallback invokes the emr.ListBackupPlans API asynchronously
// api document: https://help.aliyun.com/api/emr/listbackupplans.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListBackupPlansWithCallback(request *ListBackupPlansRequest, callback func(response *ListBackupPlansResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListBackupPlansResponse
		var err error
		defer close(result)
		response, err = client.ListBackupPlans(request)
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

// ListBackupPlansRequest is the request struct for api ListBackupPlans
type ListBackupPlansRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageCount       requests.Integer `position:"Query" name:"PageCount"`
	OrderMode       string           `position:"Query" name:"OrderMode"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	Limit           requests.Integer `position:"Query" name:"Limit"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	Id              requests.Integer `position:"Query" name:"Id"`
	CurrentSize     requests.Integer `position:"Query" name:"CurrentSize"`
	OrderField      string           `position:"Query" name:"OrderField"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	BizId           string           `position:"Query" name:"BizId"`
	Status          string           `position:"Query" name:"Status"`
}

// ListBackupPlansResponse is the response struct for api ListBackupPlans
type ListBackupPlansResponse struct {
	*responses.BaseResponse
	RequestId  string                 `json:"RequestId" xml:"RequestId"`
	PageNumber int                    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                    `json:"PageSize" xml:"PageSize"`
	TotalCount int                    `json:"TotalCount" xml:"TotalCount"`
	Items      ItemsInListBackupPlans `json:"Items" xml:"Items"`
}

// CreateListBackupPlansRequest creates a request to invoke ListBackupPlans API
func CreateListBackupPlansRequest() (request *ListBackupPlansRequest) {
	request = &ListBackupPlansRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListBackupPlans", "emr", "openAPI")
	return
}

// CreateListBackupPlansResponse creates a response to parse from ListBackupPlans response
func CreateListBackupPlansResponse() (response *ListBackupPlansResponse) {
	response = &ListBackupPlansResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}