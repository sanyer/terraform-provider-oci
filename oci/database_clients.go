// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_database "github.com/oracle/oci-go-sdk/v41/database"

	oci_common "github.com/oracle/oci-go-sdk/v41/common"
)

func init() {
	RegisterOracleClient("oci_database.DatabaseClient", &OracleClient{initClientFn: initDatabaseDatabaseClient})
}

func initDatabaseDatabaseClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_database.NewDatabaseClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) databaseClient() *oci_database.DatabaseClient {
	return m.GetClient("oci_database.DatabaseClient").(*oci_database.DatabaseClient)
}
