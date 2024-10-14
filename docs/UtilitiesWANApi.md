# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearSiteDeviceSession**](UtilitiesWANApi.md#ClearSiteDeviceSession) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/clear_session | clearSiteDeviceSession
[**ClearSiteSsrArpCache**](UtilitiesWANApi.md#ClearSiteSsrArpCache) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/clear_arp | clearSiteSsrArpCache
[**ClearSiteSsrBgpRoutes**](UtilitiesWANApi.md#ClearSiteSsrBgpRoutes) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/clear_bgp | clearSiteSsrBgpRoutes
[**ReleaseSiteSsrDhcpLease**](UtilitiesWANApi.md#ReleaseSiteSsrDhcpLease) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/release_dhcp | releaseSiteSsrDhcpLease
[**RunSiteSrxTopCommand**](UtilitiesWANApi.md#RunSiteSrxTopCommand) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/run_top | runSiteSrxTopCommand
[**ServicePingFromSsr**](UtilitiesWANApi.md#ServicePingFromSsr) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/service_ping | servicePingFromSsr
[**ShowSiteSsrAndSrxRoutes**](UtilitiesWANApi.md#ShowSiteSsrAndSrxRoutes) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_route | showSiteSsrAndSrxRoutes
[**ShowSiteSsrAndSrxSessions**](UtilitiesWANApi.md#ShowSiteSsrAndSrxSessions) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_session | showSiteSsrAndSrxSessions
[**ShowSiteSsrOspfDatabase**](UtilitiesWANApi.md#ShowSiteSsrOspfDatabase) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_ospf_database | showSiteSsrOspfDatabase
[**ShowSiteSsrOspfInterfaces**](UtilitiesWANApi.md#ShowSiteSsrOspfInterfaces) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_ospf_interfaces | showSiteSsrOspfInterfaces
[**ShowSiteSsrOspfNeighbors**](UtilitiesWANApi.md#ShowSiteSsrOspfNeighbors) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_ospf_neighbors | showSiteSsrOspfNeighbors
[**ShowSiteSsrOspfSummary**](UtilitiesWANApi.md#ShowSiteSsrOspfSummary) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_ospf_summary | showSiteSsrOspfSummary
[**ShowSiteSsrServicePath**](UtilitiesWANApi.md#ShowSiteSsrServicePath) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_service_path | showSiteSsrServicePath
[**TestSiteSsrDnsResolution**](UtilitiesWANApi.md#TestSiteSsrDnsResolution) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/resolve_dns | testSiteSsrDnsResolution

# **ClearSiteDeviceSession**
> ClearSiteDeviceSession(ctx, siteId, deviceId, optional)
clearSiteDeviceSession

Clear session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesWANApiClearSiteDeviceSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWANApiClearSiteDeviceSessionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdClearSessionBody**](DeviceIdClearSessionBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearSiteSsrArpCache**
> WebsocketSession ClearSiteSsrArpCache(ctx, siteId, deviceId, optional)
clearSiteSsrArpCache

Clear ARP cache for SSR, SRX and Switch  Clear the entire ARP cache or a subset if arguments are provided.  *Note*: port_id is optional if neither vlan nor ip is specified

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesWANApiClearSiteSsrArpCacheOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWANApiClearSiteSsrArpCacheOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UtilsClearArp**](UtilsClearArp.md)|  | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearSiteSsrBgpRoutes**
> WebsocketSession ClearSiteSsrBgpRoutes(ctx, siteId, deviceId, optional)
clearSiteSsrBgpRoutes

Clear routes associated with one or all BGP neighbors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesWANApiClearSiteSsrBgpRoutesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWANApiClearSiteSsrBgpRoutesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdClearBgpBody**](DeviceIdClearBgpBody.md)|  | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReleaseSiteSsrDhcpLease**
> WebsocketSession ReleaseSiteSsrDhcpLease(ctx, siteId, deviceId, optional)
releaseSiteSsrDhcpLease

Releases an active DHCP lease.  The output will be available through websocket. As there can be multiple command issued  against the same Device at the same time and the output all goes through the same websocket stream, session is introduced for demux.   #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {     \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" }```    #### Example output from ws stream  ```json {     \"channel\": \"/sites/d6fb4f96-xxxx-xxxx-xxxx-xxxxxxxxxxxx/devices/00000000-0000-0000-1000-xxxxxxxxxxxx/cmd\",   \"event\": \"data\",   \"data\": {     \"session\": \"9106e908-74dc-4a4f-9050-9c2adcaf44a5\",     \"raw\": \"Running traceroute...\\ntraceroute to 8.8.8.8, 64 hops max\\n 0  192.168.1.1 1 ms  192.168.1.1 1 ms  192.168.1.1 1 ms\\n 1  80.10.236.81 2 ms  80.10.236.81 4 ms  80.10.236.81 2 ms\\n 2  193.253.80.250 3 ms  193.253.80.250 2 ms  193.253.80.250 2 ms\\n 3  193.252.159.41 2 ms  193.252.159.41 1 ms  193.252.159.41 3 ms\\n\"   } } ``` \"

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesWANApiReleaseSiteSsrDhcpLeaseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWANApiReleaseSiteSsrDhcpLeaseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdReleaseDhcpBody**](DeviceIdReleaseDhcpBody.md)|  | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunSiteSrxTopCommand**
> WebsocketSessionWithUrl RunSiteSrxTopCommand(ctx, siteId, deviceId)
runSiteSrxTopCommand

Run top command on switches and SRX. The output will be available through websocket.   As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {   \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" } ```

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

# **ServicePingFromSsr**
> WebsocketSession ServicePingFromSsr(ctx, siteId, deviceId, optional)
servicePingFromSsr

Ping from SSR  Service Ping can be performed from the Device. The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, session is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {     \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" } ``` ##### Example output from ws stream ```json {     \"event\": \"data\",     \"channel\": \"/sites/4ac1dcf4-9d8b-7211-65c4-057819f0862b/devices/00000000-0000-0000-1000-5c5b350e0060/cmd\",     \"data\": {         \"session\": \"session_id\",         \"raw\": \"64 bytes from 23.211.0.110: seq=8 ttl=58 time=12.323 ms\\n\"     } } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesWANApiServicePingFromSsrOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWANApiServicePingFromSsrOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdServicePingBody**](DeviceIdServicePingBody.md)| Request Body | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSiteSsrAndSrxRoutes**
> WebsocketSession ShowSiteSsrAndSrxRoutes(ctx, siteId, deviceId, optional)
showSiteSsrAndSrxRoutes

Get routes from SSR, SRX and Switch.   The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {     \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" } ``` ##### Example output from ws stream ``` admin@labsystem1.fiedler# show bgp neighbors BGP neighbor is 192.168.4.1, remote AS 4200000001, local AS 4200000128, external link   BGP version 4, remote router ID 1.1.1.1   BGP state = Established, up for 00:27:25   Last read 00:00:25, hold time is 90, keepalive interval is 30 seconds   Configured hold time is 90, keepalive interval is 30 seconds   Neighbor capabilities:     4 Byte AS: advertised and received     Route refresh: advertised and received(old &amp; new)     Address family IPv4 Unicast: advertised and received     Graceful Restart Capabilty: advertised and received       Remote Restart timer is 120 seconds       Address families by peer:         none         ... ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesWANApiShowSiteSsrAndSrxRoutesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWANApiShowSiteSsrAndSrxRoutesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdShowRouteBody**](DeviceIdShowRouteBody.md)| all attributes are optional | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSiteSsrAndSrxSessions**
> WebsocketSession ShowSiteSsrAndSrxSessions(ctx, siteId, deviceId, optional)
showSiteSsrAndSrxSessions

Get active sessions passing through the Device.  The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, session is introduced for demux.   #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json { \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" }```  #### Example output from ws stream ```json { \"channel\": \"/sites/d6fb4f96-xxxx-xxxx-xxxx-xxxxxxxxxxxx/devices/00000000-0000-0000-1000-xxxxxxxxxxxx/cmd\", \"event\": \"data\", \"data\": { \"session\": \"f517bf29-1141-41ae-a084-17cacb0ccb57\", \"raw\": \"{\\\"status\\\":\\\"SUCCESS\\\",\\\"finished\\\":true,\\\"rows\\\":[{\\\"session_id\\\":\\\"a04b1cc7-dcc1-40a6-a010-0fe46ca38551\\\",\\\"direction\\\":\\\"forward\\\",\\\"service\\\":\\\"internet\\\",\\\"tenant\\\":\\\"SRV.PRD-Core\\\",\\\"device_interface\\\":\\\"ge-0/0/3\\\",\\\"network_interface\\\":\\\"ge-0/0/3.100\\\",\\\"protocol\\\":\\\"TCP\\\",\\\"source_ip\\\":\\\"10.3.20.101\\\",\\\"source_port\\\":45733,\\\"destination_ip\\\":\\\"13.38.46.35\\\",\\\"destination_port\\\":443,\\\"nat_ip\\\":\\\"192.168.1.115\\\",\\\"nat_port\\\":45256,\\\"payload_encrypted\\\":false,\\\"timeout\\\":1581,\\\"uptime\\\":319},{\\\"session_id\\\":\\\"a04b1cc7-dcc1-40a6-a010-0fe46ca38551\\\",\\\"direction\\\":\\\"reverse\\\",\\\"service\\\":\\\"internet\\\",\\\"tenant\\\":\\\"SRV.PRD-Core\\\",\\\"device_interface\\\":\\\"ge-0/0/0\\\",\\\"network_interface\\\":\\\"ge-0/0/0\\\",\\\"protocol\\\":\\\"TCP\\\",\\\"source_ip\\\":\\\"13.38.46.35\\\",\\\"source_port\\\":443,\\\"destination_ip\\\":\\\"192.168.1.115\\\",\\\"destination_port\\\":45256,\\\"nat_ip\\\":\\\"0.0.0.0\\\",\\\"nat_port\\\":0,\\\"payload_encrypted\\\":false,\\\"timeout\\\":1581,\\\"uptime\\\":319}]}\\n\" } } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesWANApiShowSiteSsrAndSrxSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWANApiShowSiteSsrAndSrxSessionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdShowSessionBody**](DeviceIdShowSessionBody.md)|  | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSiteSsrOspfDatabase**
> WebsocketSession ShowSiteSsrOspfDatabase(ctx, siteId, deviceId, optional)
showSiteSsrOspfDatabase

Get OSPF Database from the Device. The output will be available through websocket.   As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {   \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" } ```  #### Example output from ws stream ``` ===== ==================== ========== ======= ======== ================ =================== ================= Vrf   Neighbor Router ID   Priority   State   Uptime   Dead Timer Due   Interface Address   Interface State ===== ==================== ========== ======= ======== ================ =================== =================       1.0.0.3                     1   Full       852               38   172.16.3.2          Backup       1.0.0.4                     1   Full       811               33   172.16.3.2          DROther       1.0.0.3                     1   Full       852               38   172.16.4.2          Backup       1.0.0.4                     1   Full       811               34   172.16.4.2          DROther ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesWANApiShowSiteSsrOspfDatabaseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWANApiShowSiteSsrOspfDatabaseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdShowOspfDatabaseBody**](DeviceIdShowOspfDatabaseBody.md)| all attributes are optional | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSiteSsrOspfInterfaces**
> WebsocketSession ShowSiteSsrOspfInterfaces(ctx, siteId, deviceId, optional)
showSiteSsrOspfInterfaces

Get OSPF interfaces from the Device. The output will be available through websocket.   As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {   \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" } ```  #### Example output from ws stream ``` ===== ================== =================== ============== =============== =========== ========= =========== Vrf   Device Interface   Network Interface   Interface Up   IP Address      OSPF Type   Area ID   Area Type ===== ================== =================== ============== =============== =========== ========= ===========       net1               g1                          True   172.16.1.2/24   Broadcast   0.0.0.0   default       net3               g3                          True   172.16.3.2/24   Broadcast   0.0.0.0   default       net4               g4                          True   172.16.4.2/24   Broadcast   0.0.0.4   default ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesWANApiShowSiteSsrOspfInterfacesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWANApiShowSiteSsrOspfInterfacesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdShowOspfInterfacesBody**](DeviceIdShowOspfInterfacesBody.md)| all attributes are optional | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSiteSsrOspfNeighbors**
> WebsocketSession ShowSiteSsrOspfNeighbors(ctx, siteId, deviceId, optional)
showSiteSsrOspfNeighbors

Get OSPF Neighbors from the Device. The output will be available through websocket.   As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {   \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" } ```  #### Example output from ws stream ``` ===== ==================== ========== ======= ======== ================ =================== ================= Vrf   Neighbor Router ID   Priority   State   Uptime   Dead Timer Due   Interface Address   Interface State ===== ==================== ========== ======= ======== ================ =================== =================       1.0.0.3                     1   Full       852               38   172.16.3.2          Backup       1.0.0.4                     1   Full       811               33   172.16.3.2          DROther       1.0.0.3                     1   Full       852               38   172.16.4.2          Backup       1.0.0.4                     1   Full       811               34   172.16.4.2          DROther ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesWANApiShowSiteSsrOspfNeighborsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWANApiShowSiteSsrOspfNeighborsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdShowOspfNeighborsBody**](DeviceIdShowOspfNeighborsBody.md)| all attributes are optional | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSiteSsrOspfSummary**
> WebsocketSession ShowSiteSsrOspfSummary(ctx, siteId, deviceId, optional)
showSiteSsrOspfSummary

Get OSPF summary from the Device. The output will be available through websocket.   As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {   \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" } ```  #### Example output from ws stream ``` ===== =========== ========== ============= ==================== ========= =========== ============= Vrf   Router ID   ABR Type   ASBR Router   External LSA Count   Area ID   Area Type   Area Border                                                                                       Router ===== =========== ========== ============= ==================== ========= =========== =============       1.0.0.2     cisco            False                    0   0.0.0.0       1.0.0.2     cisco            False                    0   0.0.0.4   default ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesWANApiShowSiteSsrOspfSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWANApiShowSiteSsrOspfSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdShowOspfSummaryBody**](DeviceIdShowOspfSummaryBody.md)| all attributes are optional | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSiteSsrServicePath**
> WebsocketSession ShowSiteSsrServicePath(ctx, siteId, deviceId, optional)
showSiteSsrServicePath

Get service path information of the Device.  The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, session is introduced for demux.   #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json { \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" } ``` #### Example output from ws stream ```json { \"channel\": \"/sites/d6fb4f96-xxxx-xxxx-xxxx-xxxxxxxxxxxx/devices/00000000-0000-0000-1000-xxxxxxxxxxxx/cmd\", \"event\": \"data\", \"data\": { \"session\":\"5cb8a6db-d11a-42cd-bed7-19e9f29e637\", \"raw\":\"{\\\"status\\\":\\\"SUCCESS\\\",\\\"finished\\\":true,\\\"rows\\\":[{\\\"service\\\":\\\"management\\\",\\\"type\\\":\\\"service-agent\\\",\\\"network_interface\\\":\\\"ge-0/0/0\\\",\\\"destination\\\":\\\"\\\",\\\"gateway_ip\\\":\\\"192.168.1.1\\\",\\\"vector\\\":\\\"\\\",\\\"cost\\\":0,\\\"rate\\\":0,\\\"state\\\":\\\"Up\\\",\\\"capacity\\\":\\\"0/unlimited\\\",\\\"meetsSLA\\\":\\\"Yes\\\"},{\\\"service\\\":\\\"management\\\",\\\"type\\\":\\\"service-agent\\\",\\\"network_interface\\\":\\\"ge-0/0/1\\\",\\\"destination\\\":\\\"\\\",\\\"gateway_ip\\\":\\\"192.168.0.1\\\",\\\"vector\\\":\\\"\\\",\\\"cost\\\":0,\\\"rate\\\":0,\\\"state\\\":\\\"Up\\\",\\\"capacity\\\":\\\"0/unlimited\\\",\\\"meetsSLA\\\":\\\"Yes\\\"}]}\" } } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesWANApiShowSiteSsrServicePathOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWANApiShowSiteSsrServicePathOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdShowServicePathBody**](DeviceIdShowServicePathBody.md)|  | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestSiteSsrDnsResolution**
> WebsocketSession TestSiteSsrDnsResolution(ctx, siteId, deviceId)
testSiteSsrDnsResolution

DNS resolutions are performed on the Device.  The output will be available through websocket. As there can be multiple command issued against the same SSR at the same time and the output all goes through the same websocket stream, `session` is used for demux.    #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {     \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\" } ``` ##### Example output from ws stream ```  Router      | Hostname               | Resolved | Last Resolved        | Expiration -------------|------------------------|----------|----------------------|---------------------  test-device | xxx.yyy.net            | Y        | 2022-03-28T03:56:49Z | 2022-03-28T03:57:49Z ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

