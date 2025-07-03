// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dbmulticloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DbmulticloudOracleDbAzureConnectorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbmulticloudOracleDbAzureConnectors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_cluster_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_db_azure_connector_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_db_azure_connector_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DbmulticloudOracleDbAzureConnectorResource()),
						},
					},
				},
			},
		},
	}
}

func readDbmulticloudOracleDbAzureConnectors(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureConnectorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDBAzureConnectorClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudOracleDbAzureConnectorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.OracleDBAzureConnectorClient
	Res    *oci_dbmulticloud.ListOracleDbAzureConnectorsResponse
}

func (s *DbmulticloudOracleDbAzureConnectorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbAzureConnectorsDataSourceCrud) Get() error {
	request := oci_dbmulticloud.ListOracleDbAzureConnectorsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbClusterResourceId, ok := s.D.GetOkExists("db_cluster_resource_id"); ok {
		tmp := dbClusterResourceId.(string)
		request.DbClusterResourceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if oracleDbAzureConnectorId, ok := s.D.GetOkExists("id"); ok {
		tmp := oracleDbAzureConnectorId.(string)
		request.OracleDbAzureConnectorId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dbmulticloud.OracleDbAzureConnectorLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.ListOracleDbAzureConnectors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOracleDbAzureConnectors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DbmulticloudOracleDbAzureConnectorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DbmulticloudOracleDbAzureConnectorsDataSource-", DbmulticloudOracleDbAzureConnectorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	oracleDbAzureConnector := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OracleDbAzureConnectorSummaryToMap(item))
	}
	oracleDbAzureConnector["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DbmulticloudOracleDbAzureConnectorsDataSource().Schema["oracle_db_azure_connector_summary_collection"].Elem.(*schema.Resource).Schema)
		oracleDbAzureConnector["items"] = items
	}

	resources = append(resources, oracleDbAzureConnector)
	if err := s.D.Set("oracle_db_azure_connector_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
