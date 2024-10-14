# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CableTestFromSwitch**](UtilitiesLANApi.md#CableTestFromSwitch) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/cable_test | cableTestFromSwitch
[**ClearAllLearnedMacsFromPortOnSwitch**](UtilitiesLANApi.md#ClearAllLearnedMacsFromPortOnSwitch) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/clear_macs | clearAllLearnedMacsFromPortOnSwitch
[**ClearBpduErrosFromPortsOnSwitch**](UtilitiesLANApi.md#ClearBpduErrosFromPortsOnSwitch) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/clear_bpdu_error | clearBpduErrosFromPortsOnSwitch
[**CreateSiteDeviceSnapshot**](UtilitiesLANApi.md#CreateSiteDeviceSnapshot) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/snapshot | createSiteDeviceSnapshot
[**PollSiteSwitchStats**](UtilitiesLANApi.md#PollSiteSwitchStats) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/poll_stats | pollSiteSwitchStats
[**PortsBounceFromSwitch**](UtilitiesLANApi.md#PortsBounceFromSwitch) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/bounce_port | portsBounceFromSwitch
[**ReauthOrgDot1xWiredClient**](UtilitiesLANApi.md#ReauthOrgDot1xWiredClient) | **Post** /api/v1/orgs/{org_id}/wired_clients/{client_mac}/coa | reauthOrgDot1xWiredClient
[**ReauthSiteDot1xWiredClient**](UtilitiesLANApi.md#ReauthSiteDot1xWiredClient) | **Post** /api/v1/sites/{site_id}/wired_clients/{client_mac}/coa | reauthSiteDot1xWiredClient
[**UpgradeDeviceBios**](UtilitiesLANApi.md#UpgradeDeviceBios) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/upgrade_bios | upgradeDeviceBios
[**UpgradeDeviceFPGA**](UtilitiesLANApi.md#UpgradeDeviceFPGA) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/upgrade_fpga | upgradeDeviceFPGA
[**UpgradeSiteDevicesBios**](UtilitiesLANApi.md#UpgradeSiteDevicesBios) | **Post** /api/v1/sites/{site_id}/devices/upgrade_bios | upgradeSiteDevicesBios
[**UpgradeSiteDevicesFpga**](UtilitiesLANApi.md#UpgradeSiteDevicesFpga) | **Post** /api/v1/sites/{site_id}/devices/upgrade_fpga | upgradeSiteDevicesFpga

# **CableTestFromSwitch**
> WebsocketSession CableTestFromSwitch(ctx, siteId, deviceId, optional)
cableTestFromSwitch

TDR can be performed from the Switch. The output will be available through websocket. As there can be multiple command issued against the same Switch at the same time and the output all goes through the same websocket stream, session is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {     \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" } ``` ##### Example output from ws stream ```json {     \"event\": \"data\",     \"channel\": \"/sites/4ac1dcf4-9d8b-7211-65c4-057819f0862b/devices/00000000-0000-0000-1000-5c5b350e0060/cmd\",     \"data\": {         \"session\": \"session_id\",         \"raw\": \"Interface TDR detail:\\nTest status : Test successfully executed  ge-0/0/0\\n\"     } } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesLANApiCableTestFromSwitchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesLANApiCableTestFromSwitchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdCableTestBody**](DeviceIdCableTestBody.md)|  | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearAllLearnedMacsFromPortOnSwitch**
> ClearAllLearnedMacsFromPortOnSwitch(ctx, siteId, deviceId, optional)
clearAllLearnedMacsFromPortOnSwitch

Clear all learned MAC addresses, including persistent MAC addresses, on a port.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesLANApiClearAllLearnedMacsFromPortOnSwitchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesLANApiClearAllLearnedMacsFromPortOnSwitchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdClearMacsBody**](DeviceIdClearMacsBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearBpduErrosFromPortsOnSwitch**
> ClearBpduErrosFromPortsOnSwitch(ctx, siteId, deviceId, optional)
clearBpduErrosFromPortsOnSwitch

Clear bridge protocol data unit (BPDU) error condition caused by the detection of a possible bridging loop from Spanning Tree Protocol (STP) operation that renders the port unoperational.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesLANApiClearBpduErrosFromPortsOnSwitchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesLANApiClearBpduErrosFromPortsOnSwitchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UtilsClearBpdu**](UtilsClearBpdu.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSiteDeviceSnapshot**
> InlineResponse200189 CreateSiteDeviceSnapshot(ctx, siteId, deviceId)
createSiteDeviceSnapshot

Create recovery device snapshot (Available on Junos OS EX2300-, EX3400-, EX4400- devices)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200189**](inline_response_200_189.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PollSiteSwitchStats**
> PollSiteSwitchStats(ctx, siteId, deviceId)
pollSiteSwitchStats

This API can be used to poll statistics from the Switch proactively once. After it is called, the statistics will be pushed back to the cloud within the statistics interval.

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

# **PortsBounceFromSwitch**
> PortsBounceFromSwitch(ctx, siteId, deviceId, optional)
portsBounceFromSwitch

Port Bounce can be performed from the Switch.The output will be available through websocket. As there can be multiple command issued against the same AP at the same time and the output all goes through the same websocket stream, session is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {     \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" } ``` ##### Example output from ws stream ```json {     \"event\": \"data\",     \"channel\": \"/sites/4ac1dcf4-9d8b-7211-65c4-057819f0862b/devices/00000000-0000-0000-1000-5c5b350e0060/cmd\",     \"data\": {         \"session\": \"session_id\",         \"raw\": \"Port bounce complete.\"     } } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesLANApiPortsBounceFromSwitchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesLANApiPortsBounceFromSwitchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdBouncePortBody**](DeviceIdBouncePortBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReauthOrgDot1xWiredClient**
> InlineResponse200188 ReauthOrgDot1xWiredClient(ctx, orgId, clientMac)
reauthOrgDot1xWiredClient

Trigger a CoA (change of authorization) against a Wiired client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **clientMac** | **string**|  | 

### Return type

[**InlineResponse200188**](inline_response_200_188.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReauthSiteDot1xWiredClient**
> InlineResponse200188 ReauthSiteDot1xWiredClient(ctx, siteId, clientMac)
reauthSiteDot1xWiredClient

Trigger a CoA (change of authorization) against a Wiired client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **clientMac** | **string**|  | 

### Return type

[**InlineResponse200188**](inline_response_200_188.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeDeviceBios**
> InlineResponse200190 UpgradeDeviceBios(ctx, siteId, deviceId, optional)
upgradeDeviceBios

Upgrade device bios

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesLANApiUpgradeDeviceBiosOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesLANApiUpgradeDeviceBiosOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpgradeBios**](UpgradeBios.md)|  | 

### Return type

[**InlineResponse200190**](inline_response_200_190.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeDeviceFPGA**
> InlineResponse200190 UpgradeDeviceFPGA(ctx, siteId, deviceId, optional)
upgradeDeviceFPGA

Upgrade device fpga

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesLANApiUpgradeDeviceFPGAOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesLANApiUpgradeDeviceFPGAOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpgradeFpga**](UpgradeFpga.md)|  | 

### Return type

[**InlineResponse200190**](inline_response_200_190.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeSiteDevicesBios**
> UpgradeSiteDevicesBios(ctx, siteId, optional)
upgradeSiteDevicesBios

Upgrade Bios on Multiple Device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesLANApiUpgradeSiteDevicesBiosOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesLANApiUpgradeSiteDevicesBiosOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpgradeBiosMulti**](UpgradeBiosMulti.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeSiteDevicesFpga**
> UpgradeSiteDevicesFpga(ctx, siteId, optional)
upgradeSiteDevicesFpga

Upgrade Bios on Multiple Device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesLANApiUpgradeSiteDevicesFpgaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesLANApiUpgradeSiteDevicesFpgaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpgradeFpgaMulti**](UpgradeFpgaMulti.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

