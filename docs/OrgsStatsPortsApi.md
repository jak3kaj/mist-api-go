# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgSwitchPorts**](OrgsStatsPortsApi.md#CountOrgSwitchPorts) | **Get** /api/v1/orgs/{org_id}/stats/switch_ports/count | countOrgSwitchPorts
[**SearchOrgSwOrGwPorts**](OrgsStatsPortsApi.md#SearchOrgSwOrGwPorts) | **Get** /api/v1/orgs/{org_id}/stats/ports/search | searchOrgSwOrGwPorts

# **CountOrgSwitchPorts**
> InlineResponse20016 CountOrgSwitchPorts(ctx, orgId, optional)
countOrgSwitchPorts

Count by Distinct Attributes of Switch/Gateway Ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsStatsPortsApiCountOrgSwitchPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsStatsPortsApiCountOrgSwitchPortsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of OrgSwitchPortCountDistinct**](.md)|  | 
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
 **stpState** | [**optional.Interface of StpState1**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **stpRole** | [**optional.Interface of StpRole1**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **authState** | [**optional.Interface of AuthState1**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
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

# **SearchOrgSwOrGwPorts**
> ResponsePortStatsSearch SearchOrgSwOrGwPorts(ctx, orgId, optional)
searchOrgSwOrGwPorts

Search Switch / Gateway Ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsStatsPortsApiSearchOrgSwOrGwPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsStatsPortsApiSearchOrgSwOrGwPortsOpts struct
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
 **stpState** | [**optional.Interface of StpState**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **stpRole** | [**optional.Interface of StpRole**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **authState** | [**optional.Interface of AuthState**](.md)| if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; &amp;&amp; has Authenticator role | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**ResponsePortStatsSearch**](response_port_stats_search.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

