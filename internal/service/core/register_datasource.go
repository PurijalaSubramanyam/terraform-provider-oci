// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_core_app_catalog_listing", CoreAppCatalogListingDataSource())
	tfresource.RegisterDatasource("oci_core_app_catalog_listing_resource_version", CoreAppCatalogListingResourceVersionDataSource())
	tfresource.RegisterDatasource("oci_core_app_catalog_listing_resource_versions", CoreAppCatalogListingResourceVersionsDataSource())
	tfresource.RegisterDatasource("oci_core_app_catalog_listings", CoreAppCatalogListingsDataSource())
	tfresource.RegisterDatasource("oci_core_app_catalog_subscriptions", CoreAppCatalogSubscriptionsDataSource())
	tfresource.RegisterDatasource("oci_core_block_volume_replica", CoreBlockVolumeReplicaDataSource())
	tfresource.RegisterDatasource("oci_core_block_volume_replicas", CoreBlockVolumeReplicasDataSource())
	tfresource.RegisterDatasource("oci_core_boot_volume", CoreBootVolumeDataSource())
	tfresource.RegisterDatasource("oci_core_boot_volume_attachments", CoreBootVolumeAttachmentsDataSource())
	tfresource.RegisterDatasource("oci_core_boot_volume_backup", CoreBootVolumeBackupDataSource())
	tfresource.RegisterDatasource("oci_core_boot_volume_backups", CoreBootVolumeBackupsDataSource())
	tfresource.RegisterDatasource("oci_core_boot_volume_replica", CoreBootVolumeReplicaDataSource())
	tfresource.RegisterDatasource("oci_core_boot_volume_replicas", CoreBootVolumeReplicasDataSource())
	tfresource.RegisterDatasource("oci_core_boot_volumes", CoreBootVolumesDataSource())
	tfresource.RegisterDatasource("oci_core_byoip_allocated_ranges", CoreByoipAllocatedRangesDataSource())
	tfresource.RegisterDatasource("oci_core_byoip_range", CoreByoipRangeDataSource())
	tfresource.RegisterDatasource("oci_core_byoip_ranges", CoreByoipRangesDataSource())
	tfresource.RegisterDatasource("oci_core_capture_filter", CoreCaptureFilterDataSource())
	tfresource.RegisterDatasource("oci_core_capture_filters", CoreCaptureFiltersDataSource())
	tfresource.RegisterDatasource("oci_core_cluster_network", CoreClusterNetworkDataSource())
	tfresource.RegisterDatasource("oci_core_cluster_network_instances", CoreClusterNetworkInstancesDataSource())
	tfresource.RegisterDatasource("oci_core_cluster_networks", CoreClusterNetworksDataSource())
	tfresource.RegisterDatasource("oci_core_compute_capacity_reservation", CoreComputeCapacityReservationDataSource())
	tfresource.RegisterDatasource("oci_core_compute_capacity_reservation_instance_shapes", CoreComputeCapacityReservationInstanceShapesDataSource())
	tfresource.RegisterDatasource("oci_core_compute_capacity_reservation_instances", CoreComputeCapacityReservationInstancesDataSource())
	tfresource.RegisterDatasource("oci_core_compute_capacity_reservations", CoreComputeCapacityReservationsDataSource())
	tfresource.RegisterDatasource("oci_core_compute_capacity_topologies", CoreComputeCapacityTopologiesDataSource())
	tfresource.RegisterDatasource("oci_core_compute_capacity_topology", CoreComputeCapacityTopologyDataSource())
	tfresource.RegisterDatasource("oci_core_compute_capacity_topology_compute_bare_metal_hosts", CoreComputeCapacityTopologyComputeBareMetalHostsDataSource())
	tfresource.RegisterDatasource("oci_core_compute_capacity_topology_compute_hpc_islands", CoreComputeCapacityTopologyComputeHpcIslandsDataSource())
	tfresource.RegisterDatasource("oci_core_compute_capacity_topology_compute_network_blocks", CoreComputeCapacityTopologyComputeNetworkBlocksDataSource())
	tfresource.RegisterDatasource("oci_core_compute_cluster", CoreComputeClusterDataSource())
	tfresource.RegisterDatasource("oci_core_compute_clusters", CoreComputeClustersDataSource())
	tfresource.RegisterDatasource("oci_core_compute_global_image_capability_schema", CoreComputeGlobalImageCapabilitySchemaDataSource())
	tfresource.RegisterDatasource("oci_core_compute_global_image_capability_schemas", CoreComputeGlobalImageCapabilitySchemasDataSource())
	tfresource.RegisterDatasource("oci_core_compute_global_image_capability_schemas_version", CoreComputeGlobalImageCapabilitySchemasVersionDataSource())
	tfresource.RegisterDatasource("oci_core_compute_global_image_capability_schemas_versions", CoreComputeGlobalImageCapabilitySchemasVersionsDataSource())
	tfresource.RegisterDatasource("oci_core_compute_image_capability_schema", CoreComputeImageCapabilitySchemaDataSource())
	tfresource.RegisterDatasource("oci_core_compute_image_capability_schemas", CoreComputeImageCapabilitySchemasDataSource())
	tfresource.RegisterDatasource("oci_core_console_histories", CoreConsoleHistoriesDataSource())
	tfresource.RegisterDatasource("oci_core_console_history_data", CoreConsoleHistoryContentDataSource())
	tfresource.RegisterDatasource("oci_core_cpe_device_shape", CoreCpeDeviceShapeDataSource())
	tfresource.RegisterDatasource("oci_core_cpe_device_shapes", CoreCpeDeviceShapesDataSource())
	tfresource.RegisterDatasource("oci_core_cpes", CoreCpesDataSource())
	tfresource.RegisterDatasource("oci_core_cross_connect", CoreCrossConnectDataSource())
	tfresource.RegisterDatasource("oci_core_cross_connect_group", CoreCrossConnectGroupDataSource())
	tfresource.RegisterDatasource("oci_core_cross_connect_groups", CoreCrossConnectGroupsDataSource())
	tfresource.RegisterDatasource("oci_core_cross_connect_locations", CoreCrossConnectLocationsDataSource())
	tfresource.RegisterDatasource("oci_core_cross_connect_port_speed_shapes", CoreCrossConnectPortSpeedShapesDataSource())
	tfresource.RegisterDatasource("oci_core_cross_connect_status", CoreCrossConnectStatusDataSource())
	tfresource.RegisterDatasource("oci_core_cross_connects", CoreCrossConnectsDataSource())
	tfresource.RegisterDatasource("oci_core_dedicated_vm_host", CoreDedicatedVmHostDataSource())
	tfresource.RegisterDatasource("oci_core_dedicated_vm_host_instance_shapes", CoreDedicatedVmHostInstanceShapesDataSource())
	tfresource.RegisterDatasource("oci_core_dedicated_vm_host_shapes", CoreDedicatedVmHostShapesDataSource())
	tfresource.RegisterDatasource("oci_core_dedicated_vm_hosts", CoreDedicatedVmHostsDataSource())
	tfresource.RegisterDatasource("oci_core_dedicated_vm_hosts_instances", CoreDedicatedVmHostsInstancesDataSource())
	tfresource.RegisterDatasource("oci_core_dhcp_options", CoreDhcpOptionsDataSource())
	tfresource.RegisterDatasource("oci_core_drg_attachments", CoreDrgAttachmentsDataSource())
	tfresource.RegisterDatasource("oci_core_drg_route_distribution", CoreDrgRouteDistributionDataSource())
	tfresource.RegisterDatasource("oci_core_drg_route_distribution_statements", CoreDrgRouteDistributionStatementsDataSource())
	tfresource.RegisterDatasource("oci_core_drg_route_distributions", CoreDrgRouteDistributionsDataSource())
	tfresource.RegisterDatasource("oci_core_drg_route_table", CoreDrgRouteTableDataSource())
	tfresource.RegisterDatasource("oci_core_drg_route_table_route_rules", CoreDrgRouteTableRouteRulesDataSource())
	tfresource.RegisterDatasource("oci_core_drg_route_tables", CoreDrgRouteTablesDataSource())
	tfresource.RegisterDatasource("oci_core_drgs", CoreDrgsDataSource())
	tfresource.RegisterDatasource("oci_core_fast_connect_provider_service", CoreFastConnectProviderServiceDataSource())
	tfresource.RegisterDatasource("oci_core_fast_connect_provider_service_key", CoreFastConnectProviderServiceKeyDataSource())
	tfresource.RegisterDatasource("oci_core_fast_connect_provider_services", CoreFastConnectProviderServicesDataSource())
	tfresource.RegisterDatasource("oci_core_image", CoreImageDataSource())
	tfresource.RegisterDatasource("oci_core_image_shape", CoreImageShapeDataSource())
	tfresource.RegisterDatasource("oci_core_image_shapes", CoreImageShapesDataSource())
	tfresource.RegisterDatasource("oci_core_images", CoreImagesDataSource())
	tfresource.RegisterDatasource("oci_core_instance", CoreInstanceDataSource())
	tfresource.RegisterDatasource("oci_core_instance_configuration", CoreInstanceConfigurationDataSource())
	tfresource.RegisterDatasource("oci_core_instance_configurations", CoreInstanceConfigurationsDataSource())
	tfresource.RegisterDatasource("oci_core_instance_console_connections", CoreInstanceConsoleConnectionsDataSource())
	tfresource.RegisterDatasource("oci_core_instance_credentials", CoreInstanceCredentialDataSource())
	tfresource.RegisterDatasource("oci_core_instance_devices", CoreInstanceDevicesDataSource())
	tfresource.RegisterDatasource("oci_core_instance_maintenance_reboot", CoreInstanceMaintenanceRebootDataSource())
	tfresource.RegisterDatasource("oci_core_instance_measured_boot_report", CoreInstanceMeasuredBootReportDataSource())
	tfresource.RegisterDatasource("oci_core_instance_pool", CoreInstancePoolDataSource())
	tfresource.RegisterDatasource("oci_core_instance_pool_instances", CoreInstancePoolInstancesDataSource())
	tfresource.RegisterDatasource("oci_core_instance_pool_load_balancer_attachment", CoreInstancePoolLoadBalancerAttachmentDataSource())
	tfresource.RegisterDatasource("oci_core_instance_pools", CoreInstancePoolsDataSource())
	tfresource.RegisterDatasource("oci_core_instances", CoreInstancesDataSource())
	tfresource.RegisterDatasource("oci_core_internet_gateways", CoreInternetGatewaysDataSource())
	tfresource.RegisterDatasource("oci_core_ip_inventory_subnet", CoreIpInventorySubnetDataSource())
	tfresource.RegisterDatasource("oci_core_ip_inventory_subnet_cidr", CoreIpInventorySubnetCidrDataSource())
	tfresource.RegisterDatasource("oci_core_ip_inventory_vcn_overlaps", CoreIpInventoryVcnOverlapsDataSource())
	tfresource.RegisterDatasource("oci_core_ipsec_config", CoreIpSecConnectionDeviceConfigDataSource())
	tfresource.RegisterDatasource("oci_core_ipsec_status", CoreIpSecConnectionDeviceStatusDataSource())
	tfresource.RegisterDatasource("oci_core_ipsec_connection_tunnels", CoreIpSecConnectionTunnelsDataSource())
	tfresource.RegisterDatasource("oci_core_ipsec_connections", CoreIpSecConnectionsDataSource())
	tfresource.RegisterDatasource("oci_core_ipsec_algorithm", CoreIpsecAlgorithmDataSource())
	tfresource.RegisterDatasource("oci_core_ipsec_connection_tunnel", CoreIpSecConnectionTunnelDataSource())
	tfresource.RegisterDatasource("oci_core_ipsec_connection_tunnel_error", CoreIpsecConnectionTunnelErrorDataSource())
	tfresource.RegisterDatasource("oci_core_ipsec_connection_tunnel_routes", CoreIpsecConnectionTunnelRoutesDataSource())
	tfresource.RegisterDatasource("oci_core_ipv6", CoreIpv6DataSource())
	tfresource.RegisterDatasource("oci_core_ipv6s", CoreIpv6sDataSource())
	tfresource.RegisterDatasource("oci_core_letter_of_authority", CoreLetterOfAuthorityDataSource())
	tfresource.RegisterDatasource("oci_core_local_peering_gateways", CoreLocalPeeringGatewaysDataSource())
	tfresource.RegisterDatasource("oci_core_nat_gateway", CoreNatGatewayDataSource())
	tfresource.RegisterDatasource("oci_core_nat_gateways", CoreNatGatewaysDataSource())
	tfresource.RegisterDatasource("oci_core_network_security_group", CoreNetworkSecurityGroupDataSource())
	tfresource.RegisterDatasource("oci_core_network_security_group_security_rules", CoreNetworkSecurityGroupSecurityRulesDataSource())
	tfresource.RegisterDatasource("oci_core_network_security_group_vnics", CoreNetworkSecurityGroupVnicsDataSource())
	tfresource.RegisterDatasource("oci_core_network_security_groups", CoreNetworkSecurityGroupsDataSource())
	tfresource.RegisterDatasource("oci_core_peer_region_for_remote_peerings", CorePeerRegionForRemotePeeringsDataSource())
	tfresource.RegisterDatasource("oci_core_private_ip", CorePrivateIpDataSource())
	tfresource.RegisterDatasource("oci_core_private_ips", CorePrivateIpsDataSource())
	tfresource.RegisterDatasource("oci_core_public_ip", CorePublicIpDataSource())
	tfresource.RegisterDatasource("oci_core_public_ip_pool", CorePublicIpPoolDataSource())
	tfresource.RegisterDatasource("oci_core_public_ip_pools", CorePublicIpPoolsDataSource())
	tfresource.RegisterDatasource("oci_core_public_ips", CorePublicIpsDataSource())
	tfresource.RegisterDatasource("oci_core_remote_peering_connections", CoreRemotePeeringConnectionsDataSource())
	tfresource.RegisterDatasource("oci_core_route_tables", CoreRouteTablesDataSource())
	tfresource.RegisterDatasource("oci_core_security_lists", CoreSecurityListsDataSource())
	tfresource.RegisterDatasource("oci_core_service_gateways", CoreServiceGatewaysDataSource())
	tfresource.RegisterDatasource("oci_core_services", CoreServicesDataSource())
	tfresource.RegisterDatasource("oci_core_shapes", CoreShapesDataSource())
	tfresource.RegisterDatasource("oci_core_subnet", CoreSubnetDataSource())
	tfresource.RegisterDatasource("oci_core_subnets", CoreSubnetsDataSource())
	tfresource.RegisterDatasource("oci_core_tunnel_security_associations", CoreTunnelSecurityAssociationsDataSource())
	tfresource.RegisterDatasource("oci_core_vcn", CoreVcnDataSource())
	tfresource.RegisterDatasource("oci_core_vcn_dns_resolver_association", CoreVcnDnsResolverAssociationDataSource())
	tfresource.RegisterDatasource("oci_core_vcns", CoreVcnsDataSource())
	tfresource.RegisterDatasource("oci_core_virtual_circuit", CoreVirtualCircuitDataSource())
	tfresource.RegisterDatasource("oci_core_virtual_circuit_associated_tunnels", CoreVirtualCircuitAssociatedTunnelsDataSource())
	tfresource.RegisterDatasource("oci_core_virtual_circuit_bandwidth_shapes", CoreVirtualCircuitBandwidthShapesDataSource())
	tfresource.RegisterDatasource("oci_core_virtual_circuit_public_prefixes", CoreVirtualCircuitPublicPrefixesDataSource())
	tfresource.RegisterDatasource("oci_core_virtual_circuits", CoreVirtualCircuitsDataSource())
	tfresource.RegisterDatasource("oci_core_vlan", CoreVlanDataSource())
	tfresource.RegisterDatasource("oci_core_vlans", CoreVlansDataSource())
	tfresource.RegisterDatasource("oci_core_vnic", CoreVnicDataSource())
	tfresource.RegisterDatasource("oci_core_vnic_attachments", CoreVnicAttachmentsDataSource())
	tfresource.RegisterDatasource("oci_core_volume", CoreVolumeDataSource())
	tfresource.RegisterDatasource("oci_core_volume_attachments", CoreVolumeAttachmentsDataSource())
	tfresource.RegisterDatasource("oci_core_volume_backup_policies", CoreVolumeBackupPoliciesDataSource())
	tfresource.RegisterDatasource("oci_core_volume_backup_policy_assignments", CoreVolumeBackupPolicyAssignmentsDataSource())
	tfresource.RegisterDatasource("oci_core_volume_backups", CoreVolumeBackupsDataSource())
	tfresource.RegisterDatasource("oci_core_volume_group_backups", CoreVolumeGroupBackupsDataSource())
	tfresource.RegisterDatasource("oci_core_volume_group_replica", CoreVolumeGroupReplicaDataSource())
	tfresource.RegisterDatasource("oci_core_volume_group_replicas", CoreVolumeGroupReplicasDataSource())
	tfresource.RegisterDatasource("oci_core_volume_groups", CoreVolumeGroupsDataSource())
	tfresource.RegisterDatasource("oci_core_volumes", CoreVolumesDataSource())
	tfresource.RegisterDatasource("oci_core_vtap", CoreVtapDataSource())
	tfresource.RegisterDatasource("oci_core_vtaps", CoreVtapsDataSource())
}
