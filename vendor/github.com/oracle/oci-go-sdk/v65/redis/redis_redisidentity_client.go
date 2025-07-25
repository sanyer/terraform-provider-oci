// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// RedisIdentityClient a client for RedisIdentity
type RedisIdentityClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewRedisIdentityClientWithConfigurationProvider Creates a new default RedisIdentity client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewRedisIdentityClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client RedisIdentityClient, err error) {
	if enabled := common.CheckForEnabledServices("redis"); !enabled {
		return client, fmt.Errorf("the Developer Tool configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local developer-tool-configuration.json file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newRedisIdentityClientFromBaseClient(baseClient, provider)
}

// NewRedisIdentityClientWithOboToken Creates a new default RedisIdentity client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewRedisIdentityClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client RedisIdentityClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newRedisIdentityClientFromBaseClient(baseClient, configProvider)
}

func newRedisIdentityClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client RedisIdentityClient, err error) {
	// RedisIdentity service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("RedisIdentity"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = RedisIdentityClient{BaseClient: baseClient}
	client.BasePath = "20220315"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *RedisIdentityClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("redis", "https://redis.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *RedisIdentityClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *RedisIdentityClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateIdentityToken Generates an identity token to sign in with the specified redis user for the redis cluster
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/CreateIdentityToken.go.html to see an example of how to use CreateIdentityToken API.
// A default retry strategy applies to this operation CreateIdentityToken()
func (client RedisIdentityClient) CreateIdentityToken(ctx context.Context, request CreateIdentityTokenRequest) (response CreateIdentityTokenResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createIdentityToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateIdentityTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateIdentityTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateIdentityTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateIdentityTokenResponse")
	}
	return
}

// createIdentityToken implements the OCIOperation interface (enables retrying operations)
func (client RedisIdentityClient) createIdentityToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/redisClusters/{redisClusterId}/actions/identityToken", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateIdentityTokenResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/CreateIdentityTokenDetails/CreateIdentityToken"
		err = common.PostProcessServiceError(err, "RedisIdentity", "CreateIdentityToken", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
