// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementCloudListenerDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cloud_listener_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseManagementCloudListenerResource(), fieldMap, readSingularDatabaseManagementCloudListener)
}

func readSingularDatabaseManagementCloudListener(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudListenerDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementCloudListenerDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetCloudListenerResponse
}

func (s *DatabaseManagementCloudListenerDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementCloudListenerDataSourceCrud) Get() error {
	request := oci_database_management.GetCloudListenerRequest{}

	if cloudListenerId, ok := s.D.GetOkExists("cloud_listener_id"); ok {
		tmp := cloudListenerId.(string)
		request.CloudListenerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetCloudListener(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementCloudListenerDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.AdrHomeDirectory != nil {
		s.D.Set("adr_home_directory", *s.Res.AdrHomeDirectory)
	}

	if s.Res.CloudConnectorId != nil {
		s.D.Set("cloud_connector_id", *s.Res.CloudConnectorId)
	}

	if s.Res.CloudDbHomeId != nil {
		s.D.Set("cloud_db_home_id", *s.Res.CloudDbHomeId)
	}

	if s.Res.CloudDbNodeId != nil {
		s.D.Set("cloud_db_node_id", *s.Res.CloudDbNodeId)
	}

	if s.Res.CloudDbSystemId != nil {
		s.D.Set("cloud_db_system_id", *s.Res.CloudDbSystemId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComponentName != nil {
		s.D.Set("component_name", *s.Res.ComponentName)
	}

	if s.Res.DbaasId != nil {
		s.D.Set("dbaas_id", *s.Res.DbaasId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	endpoints := []interface{}{}
	for _, item := range s.Res.Endpoints {
		endpoints = append(endpoints, CloudListenerEndpointToMap(item))
	}
	s.D.Set("endpoints", endpoints)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HostName != nil {
		s.D.Set("host_name", *s.Res.HostName)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ListenerAlias != nil {
		s.D.Set("listener_alias", *s.Res.ListenerAlias)
	}

	if s.Res.ListenerOraLocation != nil {
		s.D.Set("listener_ora_location", *s.Res.ListenerOraLocation)
	}

	s.D.Set("listener_type", s.Res.ListenerType)

	if s.Res.LogDirectory != nil {
		s.D.Set("log_directory", *s.Res.LogDirectory)
	}

	if s.Res.OracleHome != nil {
		s.D.Set("oracle_home", *s.Res.OracleHome)
	}

	servicedAsms := []interface{}{}
	for _, item := range s.Res.ServicedAsms {
		servicedAsms = append(servicedAsms, CloudServicedAsmToMap(item))
	}
	s.D.Set("serviced_asms", servicedAsms)

	servicedDatabases := []interface{}{}
	for _, item := range s.Res.ServicedDatabases {
		servicedDatabases = append(servicedDatabases, CloudListenerServicedDatabaseToMap(item))
	}
	s.D.Set("serviced_databases", servicedDatabases)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TraceDirectory != nil {
		s.D.Set("trace_directory", *s.Res.TraceDirectory)
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
