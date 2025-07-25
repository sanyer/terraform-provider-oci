// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCloudClustersRequest wrapper for the ListCloudClusters operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCloudClusters.go.html to see an example of how to use ListCloudClustersRequest.
type ListCloudClustersRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system.
	CloudDbSystemId *string `mandatory:"false" contributesTo:"query" name:"cloudDbSystemId"`

	// A filter to only return the resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for `TIMECREATED` is descending and the default sort order for `DISPLAYNAME` is ascending.
	// The `DISPLAYNAME` sort order is case-sensitive.
	SortBy ListCloudClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListCloudClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudClustersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudClustersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudClustersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudClustersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudClustersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudClustersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudClustersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudClustersResponse wrapper for the ListCloudClusters operation
type ListCloudClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CloudClusterCollection instances
	CloudClusterCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudClustersSortByEnum Enum with underlying type: string
type ListCloudClustersSortByEnum string

// Set of constants representing the allowable values for ListCloudClustersSortByEnum
const (
	ListCloudClustersSortByTimecreated ListCloudClustersSortByEnum = "TIMECREATED"
	ListCloudClustersSortByDisplayname ListCloudClustersSortByEnum = "DISPLAYNAME"
)

var mappingListCloudClustersSortByEnum = map[string]ListCloudClustersSortByEnum{
	"TIMECREATED": ListCloudClustersSortByTimecreated,
	"DISPLAYNAME": ListCloudClustersSortByDisplayname,
}

var mappingListCloudClustersSortByEnumLowerCase = map[string]ListCloudClustersSortByEnum{
	"timecreated": ListCloudClustersSortByTimecreated,
	"displayname": ListCloudClustersSortByDisplayname,
}

// GetListCloudClustersSortByEnumValues Enumerates the set of values for ListCloudClustersSortByEnum
func GetListCloudClustersSortByEnumValues() []ListCloudClustersSortByEnum {
	values := make([]ListCloudClustersSortByEnum, 0)
	for _, v := range mappingListCloudClustersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudClustersSortByEnumStringValues Enumerates the set of values in String for ListCloudClustersSortByEnum
func GetListCloudClustersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListCloudClustersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudClustersSortByEnum(val string) (ListCloudClustersSortByEnum, bool) {
	enum, ok := mappingListCloudClustersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudClustersSortOrderEnum Enum with underlying type: string
type ListCloudClustersSortOrderEnum string

// Set of constants representing the allowable values for ListCloudClustersSortOrderEnum
const (
	ListCloudClustersSortOrderAsc  ListCloudClustersSortOrderEnum = "ASC"
	ListCloudClustersSortOrderDesc ListCloudClustersSortOrderEnum = "DESC"
)

var mappingListCloudClustersSortOrderEnum = map[string]ListCloudClustersSortOrderEnum{
	"ASC":  ListCloudClustersSortOrderAsc,
	"DESC": ListCloudClustersSortOrderDesc,
}

var mappingListCloudClustersSortOrderEnumLowerCase = map[string]ListCloudClustersSortOrderEnum{
	"asc":  ListCloudClustersSortOrderAsc,
	"desc": ListCloudClustersSortOrderDesc,
}

// GetListCloudClustersSortOrderEnumValues Enumerates the set of values for ListCloudClustersSortOrderEnum
func GetListCloudClustersSortOrderEnumValues() []ListCloudClustersSortOrderEnum {
	values := make([]ListCloudClustersSortOrderEnum, 0)
	for _, v := range mappingListCloudClustersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudClustersSortOrderEnumStringValues Enumerates the set of values in String for ListCloudClustersSortOrderEnum
func GetListCloudClustersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudClustersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudClustersSortOrderEnum(val string) (ListCloudClustersSortOrderEnum, bool) {
	enum, ok := mappingListCloudClustersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
