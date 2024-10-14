# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrgInventory**](OrgsInventoryApi.md#AddOrgInventory) | **Post** /api/v1/orgs/{org_id}/inventory | addOrgInventory
[**CreateOrgGatewayHaCluster**](OrgsInventoryApi.md#CreateOrgGatewayHaCluster) | **Post** /api/v1/orgs/{org_id}/inventory/create_ha_cluster | createOrgGatewayHaCluster
[**DeleteOrgGatewayHaCluster**](OrgsInventoryApi.md#DeleteOrgGatewayHaCluster) | **Post** /api/v1/orgs/{org_id}/inventory/delete_ha_cluster | deleteOrgGatewayHaCluster
[**GetOrgInventory**](OrgsInventoryApi.md#GetOrgInventory) | **Get** /api/v1/orgs/{org_id}/inventory | getOrgInventory
[**ReevaluateOrgAutoAssignment**](OrgsInventoryApi.md#ReevaluateOrgAutoAssignment) | **Post** /api/v1/orgs/{org_id}/inventory/reevaluate_auto_assignment | reevaluateOrgAutoAssignment
[**ReplaceOrgDevices**](OrgsInventoryApi.md#ReplaceOrgDevices) | **Post** /api/v1/orgs/{org_id}/inventory/replace | replaceOrgDevices
[**UpdateOrgInventoryAssignment**](OrgsInventoryApi.md#UpdateOrgInventoryAssignment) | **Put** /api/v1/orgs/{org_id}/inventory | updateOrgInventoryAssignment

# **AddOrgInventory**
> InlineResponse2006 AddOrgInventory(ctx, orgId, optional)
addOrgInventory

Add Device to Org Inventory with the device claim codes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsInventoryApiAddOrgInventoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsInventoryApiAddOrgInventoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []string**](string.md)| Request Body | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrgGatewayHaCluster**
> CreateOrgGatewayHaCluster(ctx, orgId, optional)
createOrgGatewayHaCluster

Create HA Cluster from unassigned Gateways

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsInventoryApiCreateOrgGatewayHaClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsInventoryApiCreateOrgGatewayHaClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of InventoryCreateHaClusterBody**](InventoryCreateHaClusterBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgGatewayHaCluster**
> DeleteOrgGatewayHaCluster(ctx, orgId, optional)
deleteOrgGatewayHaCluster

Delete HA Cluster  After HA cluster deleted, both of the nodes will be unassigned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsInventoryApiDeleteOrgGatewayHaClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsInventoryApiDeleteOrgGatewayHaClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of HaClusterDelete**](HaClusterDelete.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgInventory**
> []Inventory GetOrgInventory(ctx, orgId, optional)
getOrgInventory

Get Org Inventory  ### VC (Virtual-Chassis) Management Ideally VC should be managed as a single device - where - one managed entity where config / monitoring is anchored against (with a stable identify MAC) - all members appears in the inventory  In our implementation, we strive to achieve that without manual user configurations by  1. during claim or adoption a VC, we require FPC0 to exist and will use its MAC as identify for the entire chassis 2. other VC members will be automatically populated when they’re all present  The perceivable result is  1. from `/sites/:site_id/stats/devices/:fpc0_mac` API, you’d see the VC where module_stat contains the VC members  2. from `/orgs/:org_id/inventory?vc=true` API, you’d see all VC members with vc_mac pointing to the FPC0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsInventoryApiGetOrgInventoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsInventoryApiGetOrgInventoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serial** | **optional.String**| device serial | 
 **model** | **optional.String**| device model | 
 **type_** | [**optional.Interface of DeviceType**](.md)|  | 
 **mac** | **optional.String**| MAC address | 
 **siteId** | **optional.String**| site id if assigned, null if not assigned | 
 **vcMac** | **optional.String**| Virtual Chassis MAC Address | 
 **vc** | **optional.Bool**| To display Virtual Chassis members | [default to false]
 **unassigned** | **optional.Bool**| to display Unassigned devices | [default to true]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Inventory**](inventory.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReevaluateOrgAutoAssignment**
> ReevaluateOrgAutoAssignment(ctx, orgId)
reevaluateOrgAutoAssignment

Reevaluate Auto Assignment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceOrgDevices**
> InlineResponse20062 ReplaceOrgDevices(ctx, orgId, optional)
replaceOrgDevices

It’s a common request we get from the customers. When a AP HW has problem and need a replacement, they would want to copy the existing attributes (Device Config) of this old AP to the new one. It can be done by providing the MAC of a device that’s currently in the inventory but not assigned. The Device replaced will become unassigned.  This API also supports replacement of Mist Edges. This API copies device agnostic attributes from old Mist edge to new one. Mist manufactured Mist Edges will be reset to factory settings but will still be in Inventory.Brownfield or VM’s will be deleted from Inventory  **Note:** For Gateway devices only like-for-like replacements (can only replace a SRX320 with a SRX320 and not some otehr model) are allowed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsInventoryApiReplaceOrgDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsInventoryApiReplaceOrgDevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of InventoryReplaceBody**](InventoryReplaceBody.md)| Request Body | 

### Return type

[**InlineResponse20062**](inline_response_200_62.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgInventoryAssignment**
> InlineResponse20062 UpdateOrgInventoryAssignment(ctx, orgId, optional)
updateOrgInventoryAssignment

Update Org Inventory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsInventoryApiUpdateOrgInventoryAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsInventoryApiUpdateOrgInventoryAssignmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdInventoryBody**](OrgIdInventoryBody.md)|  | 

### Return type

[**InlineResponse20062**](inline_response_200_62.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

