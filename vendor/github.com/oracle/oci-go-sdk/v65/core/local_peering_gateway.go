// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LocalPeeringGateway A local peering gateway (LPG) is an object on a VCN that lets that VCN peer
// with another VCN in the same region. *Peering* means that the two VCNs can
// communicate using private IP addresses, but without the traffic traversing the
// internet or routing through your on-premises network. For more information,
// see VCN Peering (https://docs.oracle.com/iaas/Content/Network/Tasks/VCNpeering.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
type LocalPeeringGateway struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the LPG.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The LPG's Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
	Id *string `mandatory:"true" json:"id"`

	// Whether the VCN at the other end of the peering is in a different tenancy.
	// Example: `false`
	IsCrossTenancyPeering *bool `mandatory:"true" json:"isCrossTenancyPeering"`

	// The LPG's current lifecycle state.
	LifecycleState LocalPeeringGatewayLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Whether the LPG is peered with another LPG. `NEW` means the LPG has not yet been
	// peered. `PENDING` means the peering is being established. `REVOKED` means the
	// LPG at the other end of the peering has been deleted.
	PeeringStatus LocalPeeringGatewayPeeringStatusEnum `mandatory:"true" json:"peeringStatus"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peered LPG.
	PeerId *string `mandatory:"true" json:"peerId"`

	// The date and time the LPG was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN that uses the LPG.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Security attributes (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels
	// for a resource that can be referenced in a Zero Trust Packet Routing (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm)
	// (ZPR) policy to control access to ZPR-supported resources.
	// Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// The smallest aggregate CIDR that contains all the CIDR routes advertised by the VCN
	// at the other end of the peering from this LPG. See `peerAdvertisedCidrDetails` for
	// the individual CIDRs. The value is `null` if the LPG is not peered.
	// Example: `192.168.0.0/16`, or if aggregated with `172.16.0.0/24` then `128.0.0.0/1`
	PeerAdvertisedCidr *string `mandatory:"false" json:"peerAdvertisedCidr"`

	// The specific ranges of IP addresses available on or via the VCN at the other
	// end of the peering from this LPG. The value is `null` if the LPG is not peered.
	// You can use these as destination CIDRs for route rules to route a subnet's
	// traffic to this LPG.
	// Example: [`192.168.0.0/16`, `172.16.0.0/24`]
	PeerAdvertisedCidrDetails []string `mandatory:"false" json:"peerAdvertisedCidrDetails"`

	// Additional information regarding the peering status, if applicable.
	PeeringStatusDetails *string `mandatory:"false" json:"peeringStatusDetails"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the LPG is using.
	// For information about why you would associate a route table with an LPG, see
	// Transit Routing: Access to Multiple VCNs in Same Region (https://docs.oracle.com/iaas/Content/Network/Tasks/transitrouting.htm).
	RouteTableId *string `mandatory:"false" json:"routeTableId"`
}

func (m LocalPeeringGateway) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LocalPeeringGateway) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLocalPeeringGatewayLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLocalPeeringGatewayLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLocalPeeringGatewayPeeringStatusEnum(string(m.PeeringStatus)); !ok && m.PeeringStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PeeringStatus: %s. Supported values are: %s.", m.PeeringStatus, strings.Join(GetLocalPeeringGatewayPeeringStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LocalPeeringGatewayLifecycleStateEnum Enum with underlying type: string
type LocalPeeringGatewayLifecycleStateEnum string

// Set of constants representing the allowable values for LocalPeeringGatewayLifecycleStateEnum
const (
	LocalPeeringGatewayLifecycleStateProvisioning LocalPeeringGatewayLifecycleStateEnum = "PROVISIONING"
	LocalPeeringGatewayLifecycleStateAvailable    LocalPeeringGatewayLifecycleStateEnum = "AVAILABLE"
	LocalPeeringGatewayLifecycleStateTerminating  LocalPeeringGatewayLifecycleStateEnum = "TERMINATING"
	LocalPeeringGatewayLifecycleStateTerminated   LocalPeeringGatewayLifecycleStateEnum = "TERMINATED"
)

var mappingLocalPeeringGatewayLifecycleStateEnum = map[string]LocalPeeringGatewayLifecycleStateEnum{
	"PROVISIONING": LocalPeeringGatewayLifecycleStateProvisioning,
	"AVAILABLE":    LocalPeeringGatewayLifecycleStateAvailable,
	"TERMINATING":  LocalPeeringGatewayLifecycleStateTerminating,
	"TERMINATED":   LocalPeeringGatewayLifecycleStateTerminated,
}

var mappingLocalPeeringGatewayLifecycleStateEnumLowerCase = map[string]LocalPeeringGatewayLifecycleStateEnum{
	"provisioning": LocalPeeringGatewayLifecycleStateProvisioning,
	"available":    LocalPeeringGatewayLifecycleStateAvailable,
	"terminating":  LocalPeeringGatewayLifecycleStateTerminating,
	"terminated":   LocalPeeringGatewayLifecycleStateTerminated,
}

// GetLocalPeeringGatewayLifecycleStateEnumValues Enumerates the set of values for LocalPeeringGatewayLifecycleStateEnum
func GetLocalPeeringGatewayLifecycleStateEnumValues() []LocalPeeringGatewayLifecycleStateEnum {
	values := make([]LocalPeeringGatewayLifecycleStateEnum, 0)
	for _, v := range mappingLocalPeeringGatewayLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLocalPeeringGatewayLifecycleStateEnumStringValues Enumerates the set of values in String for LocalPeeringGatewayLifecycleStateEnum
func GetLocalPeeringGatewayLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
	}
}

// GetMappingLocalPeeringGatewayLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLocalPeeringGatewayLifecycleStateEnum(val string) (LocalPeeringGatewayLifecycleStateEnum, bool) {
	enum, ok := mappingLocalPeeringGatewayLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LocalPeeringGatewayPeeringStatusEnum Enum with underlying type: string
type LocalPeeringGatewayPeeringStatusEnum string

// Set of constants representing the allowable values for LocalPeeringGatewayPeeringStatusEnum
const (
	LocalPeeringGatewayPeeringStatusInvalid LocalPeeringGatewayPeeringStatusEnum = "INVALID"
	LocalPeeringGatewayPeeringStatusNew     LocalPeeringGatewayPeeringStatusEnum = "NEW"
	LocalPeeringGatewayPeeringStatusPeered  LocalPeeringGatewayPeeringStatusEnum = "PEERED"
	LocalPeeringGatewayPeeringStatusPending LocalPeeringGatewayPeeringStatusEnum = "PENDING"
	LocalPeeringGatewayPeeringStatusRevoked LocalPeeringGatewayPeeringStatusEnum = "REVOKED"
)

var mappingLocalPeeringGatewayPeeringStatusEnum = map[string]LocalPeeringGatewayPeeringStatusEnum{
	"INVALID": LocalPeeringGatewayPeeringStatusInvalid,
	"NEW":     LocalPeeringGatewayPeeringStatusNew,
	"PEERED":  LocalPeeringGatewayPeeringStatusPeered,
	"PENDING": LocalPeeringGatewayPeeringStatusPending,
	"REVOKED": LocalPeeringGatewayPeeringStatusRevoked,
}

var mappingLocalPeeringGatewayPeeringStatusEnumLowerCase = map[string]LocalPeeringGatewayPeeringStatusEnum{
	"invalid": LocalPeeringGatewayPeeringStatusInvalid,
	"new":     LocalPeeringGatewayPeeringStatusNew,
	"peered":  LocalPeeringGatewayPeeringStatusPeered,
	"pending": LocalPeeringGatewayPeeringStatusPending,
	"revoked": LocalPeeringGatewayPeeringStatusRevoked,
}

// GetLocalPeeringGatewayPeeringStatusEnumValues Enumerates the set of values for LocalPeeringGatewayPeeringStatusEnum
func GetLocalPeeringGatewayPeeringStatusEnumValues() []LocalPeeringGatewayPeeringStatusEnum {
	values := make([]LocalPeeringGatewayPeeringStatusEnum, 0)
	for _, v := range mappingLocalPeeringGatewayPeeringStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLocalPeeringGatewayPeeringStatusEnumStringValues Enumerates the set of values in String for LocalPeeringGatewayPeeringStatusEnum
func GetLocalPeeringGatewayPeeringStatusEnumStringValues() []string {
	return []string{
		"INVALID",
		"NEW",
		"PEERED",
		"PENDING",
		"REVOKED",
	}
}

// GetMappingLocalPeeringGatewayPeeringStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLocalPeeringGatewayPeeringStatusEnum(val string) (LocalPeeringGatewayPeeringStatusEnum, bool) {
	enum, ok := mappingLocalPeeringGatewayPeeringStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
