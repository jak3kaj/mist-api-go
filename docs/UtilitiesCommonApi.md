# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArpFromDevice**](UtilitiesCommonApi.md#ArpFromDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/arp | arpFromDevice
[**ClearSiteDeviceMacTable**](UtilitiesCommonApi.md#ClearSiteDeviceMacTable) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/clear_mac_table | clearSiteDeviceMacTable
[**ClearSiteDevicePolicyHitCount**](UtilitiesCommonApi.md#ClearSiteDevicePolicyHitCount) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/clear_policy_hit_count | clearSiteDevicePolicyHitCount
[**CreateSiteDeviceShellSession**](UtilitiesCommonApi.md#CreateSiteDeviceShellSession) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/shell | createSiteDeviceShellSession
[**GetSiteDeviceConfigCmd**](UtilitiesCommonApi.md#GetSiteDeviceConfigCmd) | **Get** /api/v1/sites/{site_id}/devices/{device_id}/config_cmd | getSiteDeviceConfigCmd
[**GetSiteDeviceZtpPassword**](UtilitiesCommonApi.md#GetSiteDeviceZtpPassword) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/request_ztp_password | getSiteDeviceZtpPassword
[**MonitorSiteDeviceTraffic**](UtilitiesCommonApi.md#MonitorSiteDeviceTraffic) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/monitor_traffic | monitorSiteDeviceTraffic
[**PingFromDevice**](UtilitiesCommonApi.md#PingFromDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/ping | pingFromDevice
[**ReadoptSiteOctermDevice**](UtilitiesCommonApi.md#ReadoptSiteOctermDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/readopt | readoptSiteOctermDevice
[**ReleaseSiteDeviceDhcpLease**](UtilitiesCommonApi.md#ReleaseSiteDeviceDhcpLease) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/release_dhcp_leases | releaseSiteDeviceDhcpLease
[**ReprovisionSiteOctermDevice**](UtilitiesCommonApi.md#ReprovisionSiteOctermDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/reprovision | reprovisionSiteOctermDevice
[**RestartSiteDevice**](UtilitiesCommonApi.md#RestartSiteDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/restart | restartSiteDevice
[**ShowSiteDeviceArpTable**](UtilitiesCommonApi.md#ShowSiteDeviceArpTable) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_arp | showSiteDeviceArpTable
[**ShowSiteDeviceBgpSummary**](UtilitiesCommonApi.md#ShowSiteDeviceBgpSummary) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_bgp_rummary | showSiteDeviceBgpSummary
[**ShowSiteDeviceDhcpLeases**](UtilitiesCommonApi.md#ShowSiteDeviceDhcpLeases) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_dhcp_leases | showSiteDeviceDhcpLeases
[**ShowSiteDeviceEvpnDatabase**](UtilitiesCommonApi.md#ShowSiteDeviceEvpnDatabase) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_evpn_database | showSiteDeviceEvpnDatabase
[**ShowSiteDeviceForwardingTable**](UtilitiesCommonApi.md#ShowSiteDeviceForwardingTable) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_forwarding_table | showSiteDeviceForwardingTable
[**ShowSiteDeviceMacTable**](UtilitiesCommonApi.md#ShowSiteDeviceMacTable) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_mac_table | showSiteDeviceMacTable
[**StartSiteLocateDevice**](UtilitiesCommonApi.md#StartSiteLocateDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/locate | startSiteLocateDevice
[**StopSiteLocateDevice**](UtilitiesCommonApi.md#StopSiteLocateDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/unlocate | stopSiteLocateDevice
[**TracerouteFromDevice**](UtilitiesCommonApi.md#TracerouteFromDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/traceroute | tracerouteFromDevice
[**UploadSiteDeviceSupportFile**](UtilitiesCommonApi.md#UploadSiteDeviceSupportFile) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/support | uploadSiteDeviceSupportFile

# **ArpFromDevice**
> WebsocketSession ArpFromDevice(ctx, siteId, deviceId, optional)
arpFromDevice

ARP can be performed on the Device. The output will be available through websocket. As there can be multiple command issued against the same AP at the same time and the output all goes through the same websocket stream, session is introduced for demux.   #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {     \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" } ``` ##### Example output from ws stream ```json {   \"event\": \"data\",   \"channel\": \"/sites/4ac1dcf4-9d8b-7211-65c4-057819f0862b/devices/00000000-0000-0000-1000-5c5b350e0060/cmd\",   \"data\": {     \"session\": \"session_id\",     \"raw\":     \"Output\": \"\\tMAC\\t\\tDEV\\tVLAN\\tRx Packets\\t\\t Rx Bytes\\t\\tTx Packets\\t\\t Tx Bytes\\tFlows\\tIdle sec\\n-----------------------------------------------------------------------------------------------------------------------\"   }  } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesCommonApiArpFromDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesCommonApiArpFromDeviceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdArpBody**](DeviceIdArpBody.md)|  | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearSiteDeviceMacTable**
> WebsocketSession ClearSiteDeviceMacTable(ctx, siteId, deviceId, optional)
clearSiteDeviceMacTable

Clear MAC Table from the Device.  The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {     \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" } ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesCommonApiClearSiteDeviceMacTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesCommonApiClearSiteDeviceMacTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdClearMacTableBody**](DeviceIdClearMacTableBody.md)| all attributes are optional | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearSiteDevicePolicyHitCount**
> WebsocketSessionWithUrl ClearSiteDevicePolicyHitCount(ctx, siteId, deviceId)
clearSiteDevicePolicyHitCount

Clear application policy hit counts for all the policies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 

### Return type

[**WebsocketSessionWithUrl**](websocket_session_with_url.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSiteDeviceShellSession**
> WebsocketSessionWithUrl CreateSiteDeviceShellSession(ctx, siteId, deviceId)
createSiteDeviceShellSession

Create Shell Session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 

### Return type

[**WebsocketSessionWithUrl**](websocket_session_with_url.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteDeviceConfigCmd**
> InlineResponse200186 GetSiteDeviceConfigCmd(ctx, siteId, deviceId, optional)
getSiteDeviceConfigCmd

Get Config CLI Commands For a brown-field switch deployment where we adopted the switch through Adoption Command, we do not wipe out / overwrite the existing config automatically. Instead, we generate CLI commands that we would have generated. The user can inspect, modify, and incorporate this into their existing config manually.  Once they feel comfortable about the config we generate, they can enable allow_mist_config where we will take full control of their config like a claimed switch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesCommonApiGetSiteDeviceConfigCmdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesCommonApiGetSiteDeviceConfigCmdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sort** | **optional.Bool**| Make output cmds sorted (for better readability) or not. | [default to false]

### Return type

[**InlineResponse200186**](inline_response_200_186.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteDeviceZtpPassword**
> InlineResponse200187 GetSiteDeviceZtpPassword(ctx, siteId, deviceId)
getSiteDeviceZtpPassword

In the case where soemthing happens during/after ZTP, the root-password is modified (required for ZTP to set up outbound-ssh) but the user-defined password config has not be configured. This API can be used to retrieve the temporary password.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200187**](inline_response_200_187.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorSiteDeviceTraffic**
> WebsocketSessionWithUrl MonitorSiteDeviceTraffic(ctx, siteId, deviceId, optional)
monitorSiteDeviceTraffic

Monitor traffic on switches and SRX.   * JUNOS uses cmd \"monitor interface <port>\" to monitor traffic on particular <port>   * JUNOS uses cmd \"monitor interface traffic\" to monitor traffic on all ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesCommonApiMonitorSiteDeviceTrafficOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesCommonApiMonitorSiteDeviceTrafficOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UtilsMonitorTraffic**](UtilsMonitorTraffic.md)|  | 

### Return type

[**WebsocketSessionWithUrl**](websocket_session_with_url.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PingFromDevice**
> WebsocketSession PingFromDevice(ctx, siteId, deviceId, optional)
pingFromDevice

Ping from AP, Switch and SSR  Ping can be performed from the Device. The output will be available through websocket. As there can be multiple command issued against the same AP at the same time and the output all goes through the same websocket stream, session is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {     \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" } ``` ##### Example output from ws stream ```json {     \"event\": \"data\",     \"channel\": \"/sites/4ac1dcf4-9d8b-7211-65c4-057819f0862b/devices/00000000-0000-0000-1000-5c5b350e0060/cmd\",     \"data\": {         \"session\": \"session_id\",         \"raw\": \"64 bytes from 23.211.0.110: seq=8 ttl=58 time=12.323 ms\\n\"     } } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesCommonApiPingFromDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesCommonApiPingFromDeviceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdPingBody**](DeviceIdPingBody.md)| Request Body | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadoptSiteOctermDevice**
> ReadoptSiteOctermDevice(ctx, siteId, deviceId)
readoptSiteOctermDevice

For the octerm devices, the device ID must come from fpc0. However, for a VC, the users may change the original fpc0 from CLI. To fix the issue, the readopt API could be used to trigger the readopt process so the device would get the corret device ID to connect the cloud.

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

# **ReleaseSiteDeviceDhcpLease**
> ReleaseSiteDeviceDhcpLease(ctx, siteId, deviceId, optional)
releaseSiteDeviceDhcpLease

Releases an active DHCP lease.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesCommonApiReleaseSiteDeviceDhcpLeaseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesCommonApiReleaseSiteDeviceDhcpLeaseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdReleaseDhcpLeasesBody**](DeviceIdReleaseDhcpLeasesBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReprovisionSiteOctermDevice**
> ReprovisionSiteOctermDevice(ctx, siteId, deviceId)
reprovisionSiteOctermDevice

To force one device to reprovision itself again.

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

# **RestartSiteDevice**
> RestartSiteDevice(ctx, siteId, deviceId, optional)
restartSiteDevice

Restart / Reboot a device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesCommonApiRestartSiteDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesCommonApiRestartSiteDeviceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UtilsDevicesRestart**](UtilsDevicesRestart.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSiteDeviceArpTable**
> WebsocketSession ShowSiteDeviceArpTable(ctx, siteId, deviceId, optional)
showSiteDeviceArpTable

Get ARP Table from the Device.  The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {     \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" } ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesCommonApiShowSiteDeviceArpTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesCommonApiShowSiteDeviceArpTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdShowArpBody**](DeviceIdShowArpBody.md)| all attributes are optional | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSiteDeviceBgpSummary**
> WebsocketSession ShowSiteDeviceBgpSummary(ctx, siteId, deviceId, optional)
showSiteDeviceBgpSummary

Get BGP Summary from SSR, SRX and Switch.   The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {     \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\"  } ```  ##### Example output from ws stream ``` Tue 2024-04-23 16:36:06 UTC Retrieving bgp entries... BGP table version is 354, local router ID is 10.224.8.16, vrf id 0 Default local pref 100, local AS 65000 Status codes:  s suppressed, d damped, h history, * valid, > best, = multipath,               i internal, r RIB_failure, S Stale, R Removed Nexthop codes: @NNN nexthop's vrf id, < announce-nh-self Origin codes:  i - IGP, e - EGP, ? - incomplete RPKI validation codes: V valid, I invalid, N Not found    Network                                      Next Hop                                  Metric LocPrf Weight Path *> 161.161.161.0/24 ```\"

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesCommonApiShowSiteDeviceBgpSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesCommonApiShowSiteDeviceBgpSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdShowBgpRummaryBody**](DeviceIdShowBgpRummaryBody.md)| all attributes are optional | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSiteDeviceDhcpLeases**
> WebsocketSession ShowSiteDeviceDhcpLeases(ctx, siteId, deviceId, optional)
showSiteDeviceDhcpLeases

Shows DHCP leases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesCommonApiShowSiteDeviceDhcpLeasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesCommonApiShowSiteDeviceDhcpLeasesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UtilsShowDhcpLeases**](UtilsShowDhcpLeases.md)|  | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSiteDeviceEvpnDatabase**
> WebsocketSession ShowSiteDeviceEvpnDatabase(ctx, siteId, deviceId, optional)
showSiteDeviceEvpnDatabase

Get EVPN Database from the Device. The output will be available through websocket.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesCommonApiShowSiteDeviceEvpnDatabaseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesCommonApiShowSiteDeviceEvpnDatabaseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdShowEvpnDatabaseBody**](DeviceIdShowEvpnDatabaseBody.md)| all attributes are optional | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSiteDeviceForwardingTable**
> WebsocketSession ShowSiteDeviceForwardingTable(ctx, siteId, deviceId, optional)
showSiteDeviceForwardingTable

Get forwarding table from the Device. The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {     \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" } ```  ##### Example output from ws stream ``` Mon 2024-05-20 16:47:30 UTC Retrieving fib entries… Entry Count: 3268 Capacity:    22668 ==================== ====== ======= ================== ===== ====================== =========== =========== ====== IP Prefix            Port   Proto   Tenant             VRF   Service                Next Hops   Vector      Cost ==================== ====== ======= ================== ===== ====================== =========== =========== ====== 0.0.0.0/0               0   None    Old_Mgmt           -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10 branch1-Kiosk      -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10 branch1-MGT        -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10 3.1.1.0/24              0   None    Old_Mgmt           -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10 branch1-Kiosk      -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10 branch1-MGT        -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10  ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesCommonApiShowSiteDeviceForwardingTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesCommonApiShowSiteDeviceForwardingTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdShowForwardingTableBody**](DeviceIdShowForwardingTableBody.md)| all attributes are optional | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSiteDeviceMacTable**
> WebsocketSession ShowSiteDeviceMacTable(ctx, siteId, deviceId, optional)
showSiteDeviceMacTable

Get MAC Table from the Device.  The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {     \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" } ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesCommonApiShowSiteDeviceMacTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesCommonApiShowSiteDeviceMacTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdShowMacTableBody**](DeviceIdShowMacTableBody.md)| all attributes are optional | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartSiteLocateDevice**
> StartSiteLocateDevice(ctx, siteId, deviceId, optional)
startSiteLocateDevice

### Access Points Locate an Access Point by blinking it's LED. It is a persisted state that has to be stopped by calling Stop Locating API  ### Switches Locate a Switch by blinking all port LEDs.  By default, request is sent to `master` switch and LEDs will keep flashing for 5 minutes. In case of virtual chassis (VC) the desired member mac has to be passed in the request payload.  At anypoint, only one VC member can be requested to flash the LED.  To stop LED flashing before the duration ends /unlocate API request can be made.  If /unlocate API is not called LED will continue to flash on device for the given duration.  Default duration is 5 minutes and 120 minutes is the maximum.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesCommonApiStartSiteLocateDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesCommonApiStartSiteLocateDeviceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of LocateSwitch**](LocateSwitch.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopSiteLocateDevice**
> StopSiteLocateDevice(ctx, siteId, deviceId)
stopSiteLocateDevice

Stop Locate a Device

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

# **TracerouteFromDevice**
> WebsocketSession TracerouteFromDevice(ctx, siteId, deviceId, optional)
tracerouteFromDevice

Traceroute can be performed from the Device.  The output will be available through websocket. As there can be multiple command issued against the same Device at the same time and the output all goes through the same websocket stream, session is introduced for demux.   #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {     \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" }```    #### Example output from ws stream  ```json {     \"channel\": \"/sites/d6fb4f96-xxxx-xxxx-xxxx-xxxxxxxxxxxx/devices/00000000-0000-0000-1000-xxxxxxxxxxxx/cmd\",   \"event\": \"data\",   \"data\": {     \"session\": \"9106e908-74dc-4a4f-9050-9c2adcaf44a5\",     \"raw\": \"Running traceroute...\\ntraceroute to 8.8.8.8, 64 hops max\\n 0  192.168.1.1 1 ms  192.168.1.1 1 ms  192.168.1.1 1 ms\\n 1  80.10.236.81 2 ms  80.10.236.81 4 ms  80.10.236.81 2 ms\\n 2  193.253.80.250 3 ms  193.253.80.250 2 ms  193.253.80.250 2 ms\\n 3  193.252.159.41 2 ms  193.252.159.41 1 ms  193.252.159.41 3 ms\\n\"   } } ``` \"

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesCommonApiTracerouteFromDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesCommonApiTracerouteFromDeviceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdTracerouteBody**](DeviceIdTracerouteBody.md)| Request Body | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadSiteDeviceSupportFile**
> UploadSiteDeviceSupportFile(ctx, siteId, deviceId, optional)
uploadSiteDeviceSupportFile

Support / Upload device support files  #### Info Param | Name | Type | Description | | --- | --- | --- | | process | string | Upload 1 file with output of show system processes extensive | | outbound-ssh | string | Upload 1 file that concatenates all /var/log/outbound-ssh.log* files | | messages | string | Upload 1 to 10 /var/log/messages* files | | core-dumps | string | Upload all core dump files, if any | | full | string | Upload 1 file with output of request support information, 1 file that concatenates all /var/log/outbound-ssh.log files, all core dump files, the 3 most recent /var/log/messages files, and Mist agent logs (for Junos devices running the Mist agent) | | var-logs | string | Upload all non-empty files in the /var/log/ directory | | jma-logs | string | Upload Mist agent logs (for Junos devices running the Mist agent only) | \"

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesCommonApiUploadSiteDeviceSupportFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesCommonApiUploadSiteDeviceSupportFileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdSupportBody**](DeviceIdSupportBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

