// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagementStationEventData Provides additional information for a management station event.
type ManagementStationEventData struct {

	// Type of management station operation.
	OperationType ManagementStationEventDataOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Status of the management station operation.
	Status EventStatusEnum `mandatory:"true" json:"status"`

	AdditionalDetails *WorkRequestEventDataAdditionalDetails `mandatory:"false" json:"additionalDetails"`
}

func (m ManagementStationEventData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagementStationEventData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingManagementStationEventDataOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetManagementStationEventDataOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEventStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetEventStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManagementStationEventDataOperationTypeEnum Enum with underlying type: string
type ManagementStationEventDataOperationTypeEnum string

// Set of constants representing the allowable values for ManagementStationEventDataOperationTypeEnum
const (
	ManagementStationEventDataOperationTypeSetManagementStationConfig      ManagementStationEventDataOperationTypeEnum = "SET_MANAGEMENT_STATION_CONFIG"
	ManagementStationEventDataOperationTypeSyncManagementStationMirror     ManagementStationEventDataOperationTypeEnum = "SYNC_MANAGEMENT_STATION_MIRROR"
	ManagementStationEventDataOperationTypeUpdateManagementStationSoftware ManagementStationEventDataOperationTypeEnum = "UPDATE_MANAGEMENT_STATION_SOFTWARE"
)

var mappingManagementStationEventDataOperationTypeEnum = map[string]ManagementStationEventDataOperationTypeEnum{
	"SET_MANAGEMENT_STATION_CONFIG":      ManagementStationEventDataOperationTypeSetManagementStationConfig,
	"SYNC_MANAGEMENT_STATION_MIRROR":     ManagementStationEventDataOperationTypeSyncManagementStationMirror,
	"UPDATE_MANAGEMENT_STATION_SOFTWARE": ManagementStationEventDataOperationTypeUpdateManagementStationSoftware,
}

var mappingManagementStationEventDataOperationTypeEnumLowerCase = map[string]ManagementStationEventDataOperationTypeEnum{
	"set_management_station_config":      ManagementStationEventDataOperationTypeSetManagementStationConfig,
	"sync_management_station_mirror":     ManagementStationEventDataOperationTypeSyncManagementStationMirror,
	"update_management_station_software": ManagementStationEventDataOperationTypeUpdateManagementStationSoftware,
}

// GetManagementStationEventDataOperationTypeEnumValues Enumerates the set of values for ManagementStationEventDataOperationTypeEnum
func GetManagementStationEventDataOperationTypeEnumValues() []ManagementStationEventDataOperationTypeEnum {
	values := make([]ManagementStationEventDataOperationTypeEnum, 0)
	for _, v := range mappingManagementStationEventDataOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagementStationEventDataOperationTypeEnumStringValues Enumerates the set of values in String for ManagementStationEventDataOperationTypeEnum
func GetManagementStationEventDataOperationTypeEnumStringValues() []string {
	return []string{
		"SET_MANAGEMENT_STATION_CONFIG",
		"SYNC_MANAGEMENT_STATION_MIRROR",
		"UPDATE_MANAGEMENT_STATION_SOFTWARE",
	}
}

// GetMappingManagementStationEventDataOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagementStationEventDataOperationTypeEnum(val string) (ManagementStationEventDataOperationTypeEnum, bool) {
	enum, ok := mappingManagementStationEventDataOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
