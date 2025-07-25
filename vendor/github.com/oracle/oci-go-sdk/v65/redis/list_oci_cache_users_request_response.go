// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOciCacheUsersRequest wrapper for the ListOciCacheUsers operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ListOciCacheUsers.go.html to see an example of how to use ListOciCacheUsersRequest.
type ListOciCacheUsersRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return the resources that match with the given OCI cache user name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return the resources, whose lifecycleState matches with the given lifecycleState.
	LifecycleState OciCacheUserLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOciCacheUsersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListOciCacheUsersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOciCacheUsersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOciCacheUsersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOciCacheUsersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOciCacheUsersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOciCacheUsersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOciCacheUserLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOciCacheUserLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOciCacheUsersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOciCacheUsersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOciCacheUsersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOciCacheUsersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOciCacheUsersResponse wrapper for the ListOciCacheUsers operation
type ListOciCacheUsersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OciCacheUserCollection instances
	OciCacheUserCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOciCacheUsersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOciCacheUsersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOciCacheUsersSortOrderEnum Enum with underlying type: string
type ListOciCacheUsersSortOrderEnum string

// Set of constants representing the allowable values for ListOciCacheUsersSortOrderEnum
const (
	ListOciCacheUsersSortOrderAsc  ListOciCacheUsersSortOrderEnum = "ASC"
	ListOciCacheUsersSortOrderDesc ListOciCacheUsersSortOrderEnum = "DESC"
)

var mappingListOciCacheUsersSortOrderEnum = map[string]ListOciCacheUsersSortOrderEnum{
	"ASC":  ListOciCacheUsersSortOrderAsc,
	"DESC": ListOciCacheUsersSortOrderDesc,
}

var mappingListOciCacheUsersSortOrderEnumLowerCase = map[string]ListOciCacheUsersSortOrderEnum{
	"asc":  ListOciCacheUsersSortOrderAsc,
	"desc": ListOciCacheUsersSortOrderDesc,
}

// GetListOciCacheUsersSortOrderEnumValues Enumerates the set of values for ListOciCacheUsersSortOrderEnum
func GetListOciCacheUsersSortOrderEnumValues() []ListOciCacheUsersSortOrderEnum {
	values := make([]ListOciCacheUsersSortOrderEnum, 0)
	for _, v := range mappingListOciCacheUsersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOciCacheUsersSortOrderEnumStringValues Enumerates the set of values in String for ListOciCacheUsersSortOrderEnum
func GetListOciCacheUsersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOciCacheUsersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOciCacheUsersSortOrderEnum(val string) (ListOciCacheUsersSortOrderEnum, bool) {
	enum, ok := mappingListOciCacheUsersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOciCacheUsersSortByEnum Enum with underlying type: string
type ListOciCacheUsersSortByEnum string

// Set of constants representing the allowable values for ListOciCacheUsersSortByEnum
const (
	ListOciCacheUsersSortByTimecreated ListOciCacheUsersSortByEnum = "timeCreated"
	ListOciCacheUsersSortByDisplayname ListOciCacheUsersSortByEnum = "displayName"
)

var mappingListOciCacheUsersSortByEnum = map[string]ListOciCacheUsersSortByEnum{
	"timeCreated": ListOciCacheUsersSortByTimecreated,
	"displayName": ListOciCacheUsersSortByDisplayname,
}

var mappingListOciCacheUsersSortByEnumLowerCase = map[string]ListOciCacheUsersSortByEnum{
	"timecreated": ListOciCacheUsersSortByTimecreated,
	"displayname": ListOciCacheUsersSortByDisplayname,
}

// GetListOciCacheUsersSortByEnumValues Enumerates the set of values for ListOciCacheUsersSortByEnum
func GetListOciCacheUsersSortByEnumValues() []ListOciCacheUsersSortByEnum {
	values := make([]ListOciCacheUsersSortByEnum, 0)
	for _, v := range mappingListOciCacheUsersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOciCacheUsersSortByEnumStringValues Enumerates the set of values in String for ListOciCacheUsersSortByEnum
func GetListOciCacheUsersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOciCacheUsersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOciCacheUsersSortByEnum(val string) (ListOciCacheUsersSortByEnum, bool) {
	enum, ok := mappingListOciCacheUsersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
