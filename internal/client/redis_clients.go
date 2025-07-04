// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_redis "github.com/oracle/oci-go-sdk/v65/redis"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_redis.OciCacheUserClient", &OracleClient{InitClientFn: initRedisOciCacheUserClient})
	RegisterOracleClient("oci_redis.RedisClusterClient", &OracleClient{InitClientFn: initRedisRedisClusterClient})
	RegisterOracleClient("oci_redis.RedisIdentityClient", &OracleClient{InitClientFn: initRedisRedisIdentityClient})
}

func initRedisOciCacheUserClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_redis.NewOciCacheUserClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) OciCacheUserClient() *oci_redis.OciCacheUserClient {
	return m.GetClient("oci_redis.OciCacheUserClient").(*oci_redis.OciCacheUserClient)
}

func initRedisRedisClusterClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_redis.NewRedisClusterClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) RedisClusterClient() *oci_redis.RedisClusterClient {
	return m.GetClient("oci_redis.RedisClusterClient").(*oci_redis.RedisClusterClient)
}

func initRedisRedisIdentityClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_redis.NewRedisIdentityClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) RedisIdentityClient() *oci_redis.RedisIdentityClient {
	return m.GetClient("oci_redis.RedisIdentityClient").(*oci_redis.RedisIdentityClient)
}
