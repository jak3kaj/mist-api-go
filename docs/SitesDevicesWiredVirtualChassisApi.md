# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteVirtualChassis**](SitesDevicesWiredVirtualChassisApi.md#CreateSiteVirtualChassis) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/vc | createSiteVirtualChassis
[**DeleteSiteVirtualChassis**](SitesDevicesWiredVirtualChassisApi.md#DeleteSiteVirtualChassis) | **Delete** /api/v1/sites/{site_id}/devices/{device_id}/vc | deleteSiteVirtualChassis
[**GetSiteDeviceVirtualChassis**](SitesDevicesWiredVirtualChassisApi.md#GetSiteDeviceVirtualChassis) | **Get** /api/v1/sites/{site_id}/devices/{device_id}/vc | getSiteDeviceVirtualChassis
[**SetSiteVcPort**](SitesDevicesWiredVirtualChassisApi.md#SetSiteVcPort) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/vc/vc_port | setSiteVcPort
[**UpdateSiteVirtualChassisMember**](SitesDevicesWiredVirtualChassisApi.md#UpdateSiteVirtualChassisMember) | **Put** /api/v1/sites/{site_id}/devices/{device_id}/vc | updateSiteVirtualChassisMember

# **CreateSiteVirtualChassis**
> InlineResponse2007 CreateSiteVirtualChassis(ctx, siteId, deviceId, optional)
createSiteVirtualChassis

For models (e.g. EX3400 and up) having dedicated VC ports, it is easier to form a VC by just connecting cables with the dedicated VC ports. Cloud will detect the new VC and update the inventory.  In case that the user would like to choose the dedicated switch as a VC master. Or for EX2300-C-12P and EX2300-C-12T which doesn’t have dedicated VC ports, below are procedures to automate the VC creation:  1. Power on the switch that is choosen as the VC master first. And the powering on the other member switches. 2. Claim or adopt all these switches under the same organization’s Inventory 3. Assign these switches into the same Site 4. Invoke vc command on the switch choosen to be the VC master. For EX2300-C-12P, VC ports will be created automatically. 5. Connect the cables to the VC ports for these switches 6. Wait for the VC to be formed. The Org’s inventory will be updated for the new VC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesWiredVirtualChassisApiCreateSiteVirtualChassisOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesWiredVirtualChassisApiCreateSiteVirtualChassisOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdVcBody1**](DeviceIdVcBody1.md)| Request Body | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteVirtualChassis**
> DeleteSiteVirtualChassis(ctx, siteId, deviceId)
deleteSiteVirtualChassis

When all the member switches of VC are removed and only member ID 0 is left, the cloud would detect this situation and automatically changes the single switch to non-VC role.  For some unexpected cases that the VC is gone and disconncted, the API below could be used to change the state of VC’s switches to be standalone. After it is executed, all the switches will be shown as standalone switches under Inventory.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteDeviceVirtualChassis**
> InlineResponse2007 GetSiteDeviceVirtualChassis(ctx, siteId, deviceId)
getSiteDeviceVirtualChassis

Get VC Status  The API returns a combined view of the VC status which includes topology and stats_

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSiteVcPort**
> SetSiteVcPort(ctx, siteId, deviceId, optional)
setSiteVcPort

Set VC port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesWiredVirtualChassisApiSetSiteVcPortOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesWiredVirtualChassisApiSetSiteVcPortOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of VcVcPortBody**](VcVcPortBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSiteVirtualChassisMember**
> InlineResponse2007 UpdateSiteVirtualChassisMember(ctx, siteId, deviceId, optional)
updateSiteVirtualChassisMember

The VC creation and adding member switch API will update the device’ s virtual chassis config which is applied after VC is formed to create JUNOS pre-provisioned virtual chassis configuration.  ## Change to use preprovisioned VC To switch the VC to use preprovisioned VC, enable preprovisioned in virtual_chassis config. Both vc_role master and backup will be matched to routing-engine role in Junos preprovisioned VC config.  In this config, fpc0 has to be the same as the mac of device_id. Use renumber if you want to replace fpc0 which involves device_id change.  Notice: to configure preprovisioned VC, every member of the VC must be in the inventory.  ## Add new members For models (e.g. EX4300 and up) having dedicated VC ports, it is easier to add new member switches into a VC by just connecting cables with the dedicated VC ports. Cloud will detect the new members and update the inventory.  For EX2300 VC, adding new members requires to follow the procedures below: 1. Powering on the new member switches and ensuring cables are not connected to any VC ports. 2. Claim or adopt all new member switches under the VC’s organization Inventory 3. Assign all new member switches to the same Site as the VC 4. Invoke vc command to add switches to the VC. 5. Connect the cables to the VC ports for these switches 6. After a while, the Org’s Inventory shows this new switches has been added into the VC.  ## Removing member switch To remove a member switch from the VC, following the procedures below:  1. Ensuring the VC is connected to the cloud first 2. Unplug the cable from the VC port of the switch 3. Waiting for the VC state (vc_state) of this switch is changed to not-present 4. Invoke update_vc with remove to remove this switch from the VC 5. The Org’s Inventory shows the switch is removed.  Please notice that member ID 0 (fpc0) cannot be removed. When a VC has two switches left, unpluging the cable may result in the situation that fpc0 becomes a line card (LC). When this situation is happened, please re-plug in the cable, wait for both switches becoming present (show virtual-chassis) and then removing the cable again.  ## Renumber a member switch When a member switch doesn’ t work properly and needed to be replaced, the renumber API could be used. The following two types of renumber are supported:  1. Replace a non-fpc0 member switch 2. Replace fpc0. When fpc0 is relaced, PAPI device config and JUNOS config will be both updated.  For renumber to work, the following procedures are needed:  1. Ensuring the VC is connected to the cloud and the state of the member switch to be replaced must be non present.  2. Adding the new member switch to the VC  3. Waiting for the VC state (vc_state) of this VC to be updated to API server  4. Invoke vc with renumber to replace the new member switch from fpc X to  ## Perprovision VC members By specifying “preprovision” op, you can convert the current VC to pre-provisioned mode, update VC members as well as specify vc_ports when adding new members for device models without dedicated vc ports. Use renumber for fpc0 replacement which involves device_id change.  Note:  1. vc_ports is used for adding new members and not needed if * the device model has dedicated vc ports, or * no new member is added  2. New VC members to be added should exist in the same Site as the VC  Update Device’s VC config can achieve similar purpose by directly modifying current virtual_chassis config. However, it cannot fulfill requests to enabling vc_ports on new members that are yet to belong to current VC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesWiredVirtualChassisApiUpdateSiteVirtualChassisMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesWiredVirtualChassisApiUpdateSiteVirtualChassisMemberOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdVcBody**](DeviceIdVcBody.md)| Request Body | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

