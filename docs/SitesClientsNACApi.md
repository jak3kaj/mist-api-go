# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteNacClientEvents**](SitesClientsNACApi.md#CountSiteNacClientEvents) | **Get** /api/v1/sites/{site_id}/nac_clients/events/count | countSiteNacClientEvents
[**CountSiteNacClients**](SitesClientsNACApi.md#CountSiteNacClients) | **Get** /api/v1/sites/{site_id}/nac_clients/count | countSiteNacClients
[**SearchSiteNacClientEvents**](SitesClientsNACApi.md#SearchSiteNacClientEvents) | **Get** /api/v1/sites/{site_id}/nac_clients/events/search | searchSiteNacClientEvents
[**SearchSiteNacClients**](SitesClientsNACApi.md#SearchSiteNacClients) | **Get** /api/v1/sites/{site_id}/nac_clients/search | searchSiteNacClients

# **CountSiteNacClientEvents**
> InlineResponse20016 CountSiteNacClientEvents(ctx, siteId, optional)
countSiteNacClientEvents

Count by Distinct Attributes of NAC Client-Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesClientsNACApiCountSiteNacClientEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesClientsNACApiCountSiteNacClientEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of SiteNacClientEventsCountDistinct**](.md)|  | 
 **type_** | **optional.String**| see [listDeviceEventsDefinitions]($e/Constants%20Events/listNacEventsDefinitions) | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountSiteNacClients**
> InlineResponse20016 CountSiteNacClients(ctx, siteId, optional)
countSiteNacClients

Count by Distinct Attributes of NAC Clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesClientsNACApiCountSiteNacClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesClientsNACApiCountSiteNacClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of Distinct4**](.md)| NAC Policy Rule ID, if matched | 
 **lastNacruleId** | **optional.String**| NAC Policy Rule ID, if matched | 
 **nacruleMatched** | **optional.Bool**| NAC Policy Rule Matched | 
 **authType** | **optional.String**| authentication type, e.g. \&quot;eap-tls\&quot;, \&quot;peap-tls\&quot;, \&quot;eap-ttls\&quot;, \&quot;eap-teap\&quot;, \&quot;mab\&quot;, \&quot;psk\&quot;, \&quot;device-auth\&quot; | 
 **lastVlanId** | **optional.String**| Vlan ID | 
 **lastNasVendor** | **optional.String**| vendor of NAS device | 
 **idpId** | **optional.String**| SSO ID, if present and used | 
 **lastSsid** | **optional.String**| SSID | 
 **lastUsername** | **optional.String**| Username presented by the client | 
 **timestamp** | **optional.Float64**| start time, in epoch | 
 **lastAp** | **optional.String**| AP MAC connected to by client | 
 **mac** | **optional.String**| MAC address | 
 **lastStatus** | **optional.String**| Connection status of client i.e “permitted”, “denied, “session_ended” | 
 **type_** | **optional.String**| Client type i.e. “wireless”, “wired” etc. | 
 **mdmComplianceStatus** | **optional.String**| MDM compliancy of client i.e “compliant”, “not compliant” | 
 **mdmProvider** | **optional.String**| MDM provider of client’s organisation eg “intune”, “jamf” | 
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

# **SearchSiteNacClientEvents**
> InlineResponse20037 SearchSiteNacClientEvents(ctx, siteId, optional)
searchSiteNacClientEvents

Search NAC Client Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesClientsNACApiSearchSiteNacClientEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesClientsNACApiSearchSiteNacClientEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| see [listDeviceEventsDefinitions]($e/Constants%20Events/listNacEventsDefinitions) | 
 **nacruleId** | [**optional.Interface of string**](.md)| NAC Policy Rule ID, if matched | 
 **nacruleMatched** | **optional.Bool**| NAC Policy Rule Matched | 
 **dryrunNacruleId** | **optional.String**| NAC Policy Dry Run Rule ID, if present and matched | 
 **dryrunNacruleMatched** | **optional.Bool**| True - if dryrun rule present and matched with priority, False - if not matched or not present | 
 **authType** | **optional.String**| authentication type, e.g. \&quot;eap-tls\&quot;, \&quot;peap-tls\&quot;, \&quot;eap-ttls\&quot;, \&quot;eap-teap\&quot;, \&quot;mab\&quot;, \&quot;psk\&quot;, \&quot;device-auth\&quot; | 
 **vlan** | **optional.Int32**| Vlan ID | 
 **nasVendor** | **optional.String**| vendor of NAS device | 
 **bssid** | **optional.String**| BSSID | 
 **idpId** | [**optional.Interface of string**](.md)| SSO ID, if present and used | 
 **idpRole** | **optional.String**| IDP returned roles/groups for the user | 
 **idpUsername** | **optional.String**| Username presented to the Identity Provider | 
 **respAttrs** | [**optional.Interface of []string**](string.md)| Radius attributes returned by NAC to NAS Devive | 
 **ssid** | **optional.String**| SSID | 
 **username** | **optional.String**| Username presented by the client | 
 **ap** | **optional.String**| AP MAC | 
 **randomMac** | **optional.Bool**| AP random macMAC | 
 **mac** | **optional.String**| MAC address | 
 **timestamp** | **optional.Float64**| time, in epoch | 
 **usermacLabel** | **optional.String**| labels derived from usermac entry | 
 **text** | **optional.String**| partial / full MAC address, username, device_mac or ap | 
 **nasIp** | **optional.String**| IP address of NAS device | 
 **sort** | **optional.String**| sort options, ‘-‘ prefix represents DESC order, default is wcid in ASC order | 
 **ingressVlan** | **optional.String**| vendor specific Vlan ID in radius requests | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteNacClients**
> InlineResponse20038 SearchSiteNacClients(ctx, siteId, optional)
searchSiteNacClients

Search Site NAC Clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesClientsNACApiSearchSiteNacClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesClientsNACApiSearchSiteNacClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nacruleId** | **optional.String**| NAC Policy Rule ID, if matched | 
 **nacruleMatched** | **optional.Bool**| NAC Policy Rule Matched | 
 **authType** | **optional.String**| authentication type, e.g. \&quot;eap-tls\&quot;, \&quot;peap-tls\&quot;, \&quot;eap-ttls\&quot;, \&quot;eap-teap\&quot;, \&quot;mab\&quot;, \&quot;psk\&quot;, \&quot;device-auth\&quot; | 
 **vlan** | **optional.String**| Vlan name or ID assigned to the client | 
 **nasVendor** | **optional.String**| vendor of NAS device | 
 **idpId** | **optional.String**| SSO ID, if present and used | 
 **ssid** | **optional.String**| SSID | 
 **username** | **optional.String**| Username presented by the client | 
 **timestamp** | **optional.Float64**| start time, in epoch | 
 **ap** | **optional.String**| AP MAC connected to by client | 
 **mac** | **optional.String**| MAC address | 
 **mxedgeId** | **optional.String**| ID of Mist Edge that the client is connected through | 
 **nacruleName** | **optional.String**| NAC Policy Rule Name matched | 
 **status** | **optional.String**| Connection status of client i.e “permitted”, “denied, “session_ended” | 
 **type_** | **optional.String**| Client type i.e. “wireless”, “wired” etc. | 
 **mdmCompliance** | **optional.String**| MDM compliancy of client i.e “compliant”, “not compliant” | 
 **mdmProvider** | **optional.String**| MDM provider of client’s organisation eg “intune”, “jamf” | 
 **sort** | **optional.String**| sort options, ‘-‘ prefix represents DESC order, default is wcid in ASC order | 
 **ingressVlan** | **optional.String**| vendor specific Vlan ID in radius requests | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse20038**](inline_response_200_38.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

