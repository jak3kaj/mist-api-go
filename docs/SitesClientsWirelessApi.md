# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteWirelessClientEvents**](SitesClientsWirelessApi.md#CountSiteWirelessClientEvents) | **Get** /api/v1/sites/{site_id}/clients/events/count | countSiteWirelessClientEvents
[**CountSiteWirelessClientSessions**](SitesClientsWirelessApi.md#CountSiteWirelessClientSessions) | **Get** /api/v1/sites/{site_id}/clients/sessions/count | countSiteWirelessClientSessions
[**CountSiteWirelessClients**](SitesClientsWirelessApi.md#CountSiteWirelessClients) | **Get** /api/v1/sites/{site_id}/clients/count | countSiteWirelessClients
[**GetSiteEventsForClient**](SitesClientsWirelessApi.md#GetSiteEventsForClient) | **Get** /api/v1/sites/{site_id}/clients/{client_mac}/events | getSiteEventsForClient
[**SearchSiteWirelessClientEvents**](SitesClientsWirelessApi.md#SearchSiteWirelessClientEvents) | **Get** /api/v1/sites/{site_id}/clients/events/search | searchSiteWirelessClientEvents
[**SearchSiteWirelessClientSessions**](SitesClientsWirelessApi.md#SearchSiteWirelessClientSessions) | **Get** /api/v1/sites/{site_id}/clients/sessions/search | searchSiteWirelessClientSessions
[**SearchSiteWirelessClients**](SitesClientsWirelessApi.md#SearchSiteWirelessClients) | **Get** /api/v1/sites/{site_id}/clients/search | searchSiteWirelessClients

# **CountSiteWirelessClientEvents**
> InlineResponse20016 CountSiteWirelessClientEvents(ctx, siteId, optional)
countSiteWirelessClientEvents

Count by Distinct Attributes of Client-Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesClientsWirelessApiCountSiteWirelessClientEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesClientsWirelessApiCountSiteWirelessClientEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of Distinct6**](.md)|  | 
 **type_** | **optional.String**| see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **reasonCode** | **optional.Int32**| for assoc/disassoc events | 
 **ssid** | **optional.String**| SSID Name | 
 **ap** | **optional.String**| AP MAC | 
 **proto** | [**optional.Interface of Proto1**](.md)| a / b / g / n / ac / ax | 
 **band** | [**optional.Interface of Band3**](.md)| 802.11 Band | 
 **wlanId** | **optional.String**| wlan_id | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountSiteWirelessClientSessions**
> InlineResponse20016 CountSiteWirelessClientSessions(ctx, siteId, optional)
countSiteWirelessClientSessions

Count by Distinct Attributes of Client Sessions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesClientsWirelessApiCountSiteWirelessClientSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesClientsWirelessApiCountSiteWirelessClientSessionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of SiteClientSessionsCountDistinct**](.md)|  | 
 **ap** | **optional.String**| AP MAC | 
 **band** | [**optional.Interface of Band5**](.md)| 802.11 Band | 
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

# **CountSiteWirelessClients**
> InlineResponse20016 CountSiteWirelessClients(ctx, siteId, optional)
countSiteWirelessClients

Count by Distinct Attributes of Clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesClientsWirelessApiCountSiteWirelessClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesClientsWirelessApiCountSiteWirelessClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of Distinct5**](.md)|  | 
 **ssid** | **optional.String**|  | 
 **ap** | **optional.String**|  | 
 **ipAddress** | **optional.String**|  | 
 **vlan** | **optional.String**|  | 
 **hostname** | **optional.String**|  | 
 **os** | **optional.String**|  | 
 **model** | **optional.String**|  | 
 **device** | **optional.String**|  | 
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

# **GetSiteEventsForClient**
> InlineResponse200120 GetSiteEventsForClient(ctx, siteId, clientMac, optional)
getSiteEventsForClient

Get the list of events for a specific client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **clientMac** | **string**|  | 
 **optional** | ***SitesClientsWirelessApiGetSiteEventsForClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesClientsWirelessApiGetSiteEventsForClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **optional.String**| e.g. MARVIS_EVENT_CLIENT_DHCP_STUCK | 
 **proto** | [**optional.Interface of Proto3**](.md)| a / b / g / n / ac / ax | 
 **band** | [**optional.Interface of Band7**](.md)| 802.11 Band | 
 **channel** | **optional.String**|  | 
 **wlanId** | **optional.String**|  | 
 **ssid** | **optional.String**|  | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse200120**](inline_response_200_120.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteWirelessClientEvents**
> InlineResponse20042 SearchSiteWirelessClientEvents(ctx, siteId, optional)
searchSiteWirelessClientEvents

Get Site Clients Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesClientsWirelessApiSearchSiteWirelessClientEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesClientsWirelessApiSearchSiteWirelessClientEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **reasonCode** | **optional.Int32**| for assoc/disassoc events | 
 **ssid** | **optional.String**| SSID Name | 
 **ap** | **optional.String**| AP MAC | 
 **proto** | [**optional.Interface of Proto2**](.md)| a / b / g / n / ac / ax | 
 **band** | [**optional.Interface of Band4**](.md)| 802.11 Band | 
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

# **SearchSiteWirelessClientSessions**
> InlineResponse200119 SearchSiteWirelessClientSessions(ctx, siteId, optional)
searchSiteWirelessClientSessions

Search Client Sessions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesClientsWirelessApiSearchSiteWirelessClientSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesClientsWirelessApiSearchSiteWirelessClientSessionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ap** | **optional.String**| AP MAC | 
 **band** | [**optional.Interface of Band6**](.md)| 802.11 Band | 
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

[**InlineResponse200119**](inline_response_200_119.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteWirelessClients**
> InlineResponse20043 SearchSiteWirelessClients(ctx, siteId, optional)
searchSiteWirelessClients

Search Wireless Clients  **NOTE**: fuzzy logic can be used with ‘*’, supported filters: mac, hostname, device, os, model. E.g. /clients/search?device=Mac*&hostname=jerry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesClientsWirelessApiSearchSiteWirelessClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesClientsWirelessApiSearchSiteWirelessClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **optional.String**| partial / full MAC address | 
 **ipAddress** | **optional.String**|  | 
 **hostname** | **optional.String**| partial / full hostname | 
 **device** | **optional.String**| device type, e.g. Mac, Nvidia, iPhone | 
 **os** | **optional.String**| os, e.g. Sierra, Yosemite, Windows 10 | 
 **model** | **optional.String**| model, e.g. “MBP 15 late 2013”, 6, 6s, “8+ GSM” | 
 **ap** | **optional.String**| AP mac where the client has connected to | 
 **ssid** | **optional.String**|  | 
 **text** | **optional.String**| partial / full MAC address, hostname, username, psk_name or ip | 
 **nacruleId** | **optional.String**| nacrule_id | 
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

