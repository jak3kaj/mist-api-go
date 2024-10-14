# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgWirelessClients**](OrgsClientsWirelessApi.md#CountOrgWirelessClients) | **Get** /api/v1/orgs/{org_id}/clients/count | countOrgWirelessClients
[**CountOrgWirelessClientsSessions**](OrgsClientsWirelessApi.md#CountOrgWirelessClientsSessions) | **Get** /api/v1/orgs/{org_id}/clients/sessions/count | countOrgWirelessClientsSessions
[**SearchOrgWirelessClientEvents**](OrgsClientsWirelessApi.md#SearchOrgWirelessClientEvents) | **Get** /api/v1/orgs/{org_id}/clients/events/search | searchOrgWirelessClientEvents
[**SearchOrgWirelessClientSessions**](OrgsClientsWirelessApi.md#SearchOrgWirelessClientSessions) | **Get** /api/v1/orgs/{org_id}/clients/sessions/search | searchOrgWirelessClientSessions
[**SearchOrgWirelessClients**](OrgsClientsWirelessApi.md#SearchOrgWirelessClients) | **Get** /api/v1/orgs/{org_id}/clients/search | searchOrgWirelessClients

# **CountOrgWirelessClients**
> InlineResponse20016 CountOrgWirelessClients(ctx, orgId, optional)
countOrgWirelessClients

Count Org Wireless Clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsClientsWirelessApiCountOrgWirelessClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsClientsWirelessApiCountOrgWirelessClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of OrgClientsCountDistinct**](.md)|  | 
 **mac** | **optional.String**| partial / full MAC address | 
 **hostname** | **optional.String**| partial / full hostname | 
 **device** | **optional.String**| device type, e.g. Mac, Nvidia, iPhone | 
 **os** | **optional.String**| os, e.g. Sierra, Yosemite, Windows 10 | 
 **model** | **optional.String**| model, e.g. “MBP 15 late 2013”, 6, 6s, “8+ GSM” | 
 **ap** | **optional.String**| AP mac where the client has connected to | 
 **vlan** | **optional.String**| vlan | 
 **ssid** | **optional.String**| SSID | 
 **ipAddress** | **optional.String**|  | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountOrgWirelessClientsSessions**
> InlineResponse20016 CountOrgWirelessClientsSessions(ctx, orgId, optional)
countOrgWirelessClientsSessions

Count Org Wireless Clients Sessions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsClientsWirelessApiCountOrgWirelessClientsSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsClientsWirelessApiCountOrgWirelessClientsSessionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of OrgClientSessionsCountDistinct**](.md)|  | 
 **ap** | **optional.String**| AP MAC | 
 **band** | [**optional.Interface of Band1**](.md)| 802.11 Band | 
 **clientFamily** | **optional.String**| E.g. “Mac”, “iPhone”, “Apple watch” | 
 **clientManufacture** | **optional.String**| E.g. “Apple” | 
 **clientModel** | **optional.String**| E.g. “8+”, “XS” | 
 **clientOs** | **optional.String**| E.g. “Mojave”, “Windows 10”, “Linux” | 
 **ssid** | **optional.String**| SSID | 
 **wlanId** | **optional.String**| wlan_id | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrgWirelessClientEvents**
> InlineResponse20042 SearchOrgWirelessClientEvents(ctx, orgId, optional)
searchOrgWirelessClientEvents

Get Org Clients Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsClientsWirelessApiSearchOrgWirelessClientEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsClientsWirelessApiSearchOrgWirelessClientEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **reasonCode** | **optional.Int32**| for assoc/disassoc events | 
 **ssid** | **optional.String**| SSID Name | 
 **ap** | **optional.String**| AP MAC | 
 **proto** | [**optional.Interface of Proto**](.md)| a / b / g / n / ac / ax | 
 **band** | [**optional.Interface of Band**](.md)| 802.11 Band | 
 **wlanId** | **optional.String**| wlan_id | 
 **nacruleId** | **optional.String**| nacrule_id | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20042**](inline_response_200_42.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrgWirelessClientSessions**
> InlineResponse20044 SearchOrgWirelessClientSessions(ctx, orgId, optional)
searchOrgWirelessClientSessions

Search Org Wireless Clients Sessions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsClientsWirelessApiSearchOrgWirelessClientSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsClientsWirelessApiSearchOrgWirelessClientSessionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ap** | **optional.String**| AP MAC | 
 **band** | [**optional.Interface of Band2**](.md)| 802.11 Band | 
 **clientFamily** | **optional.String**| E.g. “Mac”, “iPhone”, “Apple watch” | 
 **clientManufacture** | **optional.String**| E.g. “Apple” | 
 **clientModel** | **optional.String**| E.g. “8+”, “XS” | 
 **clientUsername** | **optional.String**| Username | 
 **clientOs** | **optional.String**| E.g. “Mojave”, “Windows 10”, “Linux” | 
 **ssid** | **optional.String**| SSID | 
 **wlanId** | **optional.String**| wlan_id | 
 **pskId** | **optional.String**| PSK ID | 
 **pskName** | **optional.String**| PSK Name | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20044**](inline_response_200_44.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrgWirelessClients**
> InlineResponse20043 SearchOrgWirelessClients(ctx, orgId, optional)
searchOrgWirelessClients

Search Org Wireless Clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsClientsWirelessApiSearchOrgWirelessClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsClientsWirelessApiSearchOrgWirelessClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **optional.String**| Site ID | 
 **mac** | **optional.String**| partial / full MAC address | 
 **ipAddress** | **optional.String**|  | 
 **hostname** | **optional.String**| partial / full hostname | 
 **device** | **optional.String**| device type, e.g. Mac, Nvidia, iPhone | 
 **os** | **optional.String**| os, e.g. Sierra, Yosemite, Windows 10 | 
 **model** | **optional.String**| model, e.g. “MBP 15 late 2013”, 6, 6s, “8+ GSM” | 
 **ap** | **optional.String**| AP mac where the client has connected to | 
 **pskId** | **optional.String**| PSK ID | 
 **pskName** | **optional.String**| PSK Name | 
 **vlan** | **optional.String**| vlan | 
 **ssid** | **optional.String**| SSID | 
 **text** | **optional.String**| partial / full MAC address, hostname, username, psk_name or ip | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20043**](inline_response_200_43.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

