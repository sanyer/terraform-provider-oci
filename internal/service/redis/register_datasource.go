// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package redis

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_redis_oci_cache_user", RedisOciCacheUserDataSource())
	tfresource.RegisterDatasource("oci_redis_oci_cache_users", RedisOciCacheUsersDataSource())
	tfresource.RegisterDatasource("oci_redis_redis_cluster", RedisRedisClusterDataSource())
	tfresource.RegisterDatasource("oci_redis_redis_cluster_nodes", RedisRedisClusterNodesDataSource())
	tfresource.RegisterDatasource("oci_redis_redis_clusters", RedisRedisClustersDataSource())
}
