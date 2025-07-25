// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementOccmDemandSignalDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["occm_demand_signal_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CapacityManagementOccmDemandSignalResource(), fieldMap, readSingularCapacityManagementOccmDemandSignal)
}

func readSingularCapacityManagementOccmDemandSignal(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccmDemandSignalDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DemandSignalClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementOccmDemandSignalDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.DemandSignalClient
	Res    *oci_capacity_management.GetOccmDemandSignalResponse
}

func (s *CapacityManagementOccmDemandSignalDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementOccmDemandSignalDataSourceCrud) Get() error {
	request := oci_capacity_management.GetOccmDemandSignalRequest{}

	if occmDemandSignalId, ok := s.D.GetOkExists("occm_demand_signal_id"); ok {
		tmp := occmDemandSignalId.(string)
		request.OccmDemandSignalId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.GetOccmDemandSignal(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CapacityManagementOccmDemandSignalDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

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

	return nil
}
