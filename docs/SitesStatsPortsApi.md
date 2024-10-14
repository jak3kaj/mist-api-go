# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteSwOrGwPorts**](SitesStatsPortsApi.md#CountSiteSwOrGwPorts) | **Get** /api/v1/sites/{site_id}/stats/ports/count | countSiteSwOrGwPorts
[**CountSiteSwitchPorts**](SitesStatsPortsApi.md#CountSiteSwitchPorts) | **Get** /api/v1/sites/{site_id}/stats/switch_ports/count | countSiteSwitchPorts
[**SearchSiteSwOrGwPorts**](SitesStatsPortsApi.md#SearchSiteSwOrGwPorts) | **Get** /api/v1/sites/{site_id}/stats/ports/search | searchSiteSwOrGwPorts
[**SearchSiteSwitchPorts**](SitesStatsPortsApi.md#SearchSiteSwitchPorts) | **Get** /api/v1/sites/{site_id}/stats/switch_ports/search | searchSiteSwitchPorts

# **CountSiteSwOrGwPorts**
> InlineResponse20016 CountSiteSwOrGwPorts(ctx, siteId, optional)
countSiteSwOrGwPorts

Count by Distinct Attributes of Switch/Gateway Ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsPortsApiCountSiteSwOrGwPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsPortsApiCountSiteSwOrGwPortsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of SitePortsCountDistinct**](.md)|  | 
 **fullDuplex** | **optional.Bool**| indicates full or half duplex | 
 **mac** | **optional.String**| device identifier | 
 **neighborMac** | **optional.String**| Chassis identifier of the chassis type listed | 
 **neighborPortDesc** | **optional.String**| Description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” | 
 **neighborSystemName** | **optional.String**| Name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” | 
 **poeDisabled** | **optional.Bool**| is the POE configured not be disabled. | 
 **poeMode** | **optional.String**| poe mode depending on class E.g. “802.3at” | 
 **poeOn** | **optional.Bool**| is the device attached to POE | 
 **portId** | **optional.String**| interface name | 
 **portMac** | **optional.String**| interface mac address | 
 **powerDraw** | **optional.Float64**| Amount of power being used by the interface at the time the command is executed. Unit in watts. | 
 **txPkts** | **optional.Int32**| Output packets | 
 **rxPkts** | **optional.Int32**| Input packets | 
 **rxBytes** | **optional.Int32**| Input bytes | 
 **txBps** | **optional.Int32**| Output rate | 
 **rxBps** | **optional.Int32**| Input rate | 
 **txMcastPkts** | **optional.Int32**| Multicast output packets | 
 **txBcastPkts** | **optional.Int32**| Broadcast output packets | 
 **rxMcastPkts** | **optional.Int32**| Multicast input packets | 
 **rxBcastPkts** | **optional.Int32**| Broadcast input packets | 
 **speed** | **optional.Int32**| port speed | 
 **stpState** | [**optional.Interface of StpState2**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **stpRole** | [**optional.Interface of StpRole2**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **authState** | [**optional.Interface of AuthState2**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; &amp;&amp; has Authenticator role | 
 **up** | **optional.Bool**| indicates if interface is up | 
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

# **CountSiteSwitchPorts**
> InlineResponse20016 CountSiteSwitchPorts(ctx, siteId, optional)
countSiteSwitchPorts

Count by Distinct Attributes of Switch/Gateway Ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsPortsApiCountSiteSwitchPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsPortsApiCountSiteSwitchPortsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of SiteSwitchPortsCountDistinct**](.md)|  | 
 **fullDuplex** | **optional.Bool**| indicates full or half duplex | 
 **mac** | **optional.String**| device identifier | 
 **neighborMac** | **optional.String**| Chassis identifier of the chassis type listed | 
 **neighborPortDesc** | **optional.String**| Description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” | 
 **neighborSystemName** | **optional.String**| Name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” | 
 **poeDisabled** | **optional.Bool**| is the POE configured not be disabled. | 
 **poeMode** | **optional.String**| poe mode depending on class E.g. “802.3at” | 
 **poeOn** | **optional.Bool**| is the device attached to POE | 
 **portId** | **optional.String**| interface name | 
 **portMac** | **optional.String**| interface mac address | 
 **powerDraw** | **optional.Float64**| Amount of power being used by the interface at the time the command is executed. Unit in watts. | 
 **txPkts** | **optional.Int32**| Output packets | 
 **rxPkts** | **optional.Int32**| Input packets | 
 **rxBytes** | **optional.Int32**| Input bytes | 
 **txBps** | **optional.Int32**| Output rate | 
 **rxBps** | **optional.Int32**| Input rate | 
 **txMcastPkts** | **optional.Int32**| Multicast output packets | 
 **txBcastPkts** | **optional.Int32**| Broadcast output packets | 
 **rxMcastPkts** | **optional.Int32**| Multicast input packets | 
 **rxBcastPkts** | **optional.Int32**| Broadcast input packets | 
 **speed** | **optional.Int32**| port speed | 
 **stpState** | [**optional.Interface of StpState4**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **stpRole** | [**optional.Interface of StpRole4**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **authState** | [**optional.Interface of AuthState4**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **up** | **optional.Bool**| indicates if interface is up | 
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

# **SearchSiteSwOrGwPorts**
> InlineResponse200175 SearchSiteSwOrGwPorts(ctx, siteId, optional)
searchSiteSwOrGwPorts

Search Switch / Gateway Ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsPortsApiSearchSiteSwOrGwPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsPortsApiSearchSiteSwOrGwPortsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullDuplex** | **optional.Bool**| indicates full or half duplex | 
 **mac** | **optional.String**| device identifier | 
 **deviceType** | [**optional.Interface of DeviceType1**](.md)| device type | 
 **neighborMac** | **optional.String**| Chassis identifier of the chassis type listed | 
 **neighborPortDesc** | **optional.String**| Description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” | 
 **neighborSystemName** | **optional.String**| Name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” | 
 **poeDisabled** | **optional.Bool**| is the POE configured not be disabled. | 
 **poeMode** | **optional.String**| poe mode depending on class E.g. “802.3at” | 
 **poeOn** | **optional.Bool**| is the device attached to POE | 
 **portId** | **optional.String**| interface name | 
 **portMac** | **optional.String**| interface mac address | 
 **powerDraw** | **optional.Float64**| Amount of power being used by the interface at the time the command is executed. Unit in watts. | 
 **txPkts** | **optional.Int32**| Output packets | 
 **rxPkts** | **optional.Int32**| Input packets | 
 **rxBytes** | **optional.Int32**| Input bytes | 
 **txBps** | **optional.Int32**| Output rate | 
 **rxBps** | **optional.Int32**| Input rate | 
 **txErrors** | **optional.Int32**| Output errors | 
 **rxErrors** | **optional.Int32**| Input errors | 
 **txMcastPkts** | **optional.Int32**| Multicast output packets | 
 **txBcastPkts** | **optional.Int32**| Broadcast output packets | 
 **rxMcastPkts** | **optional.Int32**| Multicast input packets | 
 **rxBcastPkts** | **optional.Int32**| Broadcast input packets | 
 **speed** | **optional.Int32**| port speed | 
 **macLimit** | **optional.Int32**| Limit on number of dynamically learned macs | 
 **macCount** | **optional.Int32**| Number of mac addresses in the forwarding table | 
 **up** | **optional.Bool**| indicates if interface is up | 
 **active** | **optional.Bool**| indicates if interface is active/inactive | 
 **jitter** | **optional.Float64**| Last sampled jitter of the interface | 
 **loss** | **optional.Float64**| Last sampled loss of the interface | 
 **latency** | **optional.Float64**| Last sampled latency of the interface | 
 **stpState** | [**optional.Interface of StpState3**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **stpRole** | [**optional.Interface of StpRole3**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **xcvrPartNumber** | **optional.String**| Optic Slot Partnumber, Check for null/empty | 
 **authState** | [**optional.Interface of AuthState3**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; &amp;&amp; has Authenticator role | 
 **lteImsi** | **optional.String**| LTE IMSI value, Check for null/empty | 
 **lteIccid** | **optional.String**| LTE ICCID value, Check for null/empty | 
 **lteImei** | **optional.String**| LTE IMEI value, Check for null/empty | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse200175**](inline_response_200_175.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteSwitchPorts**
> InlineResponse200175 SearchSiteSwitchPorts(ctx, siteId, optional)
searchSiteSwitchPorts

Search Switch / Gateway Ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsPortsApiSearchSiteSwitchPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsPortsApiSearchSiteSwitchPortsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullDuplex** | **optional.Bool**| indicates full or half duplex | 
 **mac** | **optional.String**| device identifier | 
 **neighborMac** | **optional.String**| Chassis identifier of the chassis type listed | 
 **neighborPortDesc** | **optional.String**| Description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” | 
 **neighborSystemName** | **optional.String**| Name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” | 
 **poeDisabled** | **optional.Bool**| is the POE configured not be disabled. | 
 **poeMode** | **optional.String**| poe mode depending on class E.g. “802.3at” | 
 **poeOn** | **optional.Bool**| is the device attached to POE | 
 **portId** | **optional.String**| interface name | 
 **portMac** | **optional.String**| interface mac address | 
 **powerDraw** | **optional.Float64**| Amount of power being used by the interface at the time the command is executed. Unit in watts. | 
 **txPkts** | **optional.Int32**| Output packets | 
 **rxPkts** | **optional.Int32**| Input packets | 
 **rxBytes** | **optional.Int32**| Input bytes | 
 **txBps** | **optional.Int32**| Output rate | 
 **rxBps** | **optional.Int32**| Input rate | 
 **txMcastPkts** | **optional.Int32**| Multicast output packets | 
 **txBcastPkts** | **optional.Int32**| Broadcast output packets | 
 **rxMcastPkts** | **optional.Int32**| Multicast input packets | 
 **rxBcastPkts** | **optional.Int32**| Broadcast input packets | 
 **speed** | **optional.Int32**| port speed | 
 **stpState** | [**optional.Interface of StpState5**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **stpRole** | [**optional.Interface of StpRole5**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **authState** | [**optional.Interface of AuthState5**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; &amp;&amp; has Authenticator role | 
 **up** | **optional.Bool**| indicates if interface is up | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse200175**](inline_response_200_175.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

