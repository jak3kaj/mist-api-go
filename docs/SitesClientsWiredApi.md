# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteWiredClients**](SitesClientsWiredApi.md#CountSiteWiredClients) | **Get** /api/v1/sites/{site_id}/wired_clients/count | countSiteWiredClients
[**SearchSiteWiredClients**](SitesClientsWiredApi.md#SearchSiteWiredClients) | **Get** /api/v1/sites/{site_id}/wired_clients/search | searchSiteWiredClients

# **CountSiteWiredClients**
> InlineResponse20016 CountSiteWiredClients(ctx, siteId, optional)
countSiteWiredClients

Count by Distinct Attributes of Clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesClientsWiredApiCountSiteWiredClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesClientsWiredApiCountSiteWiredClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of SiteWiredClientsCountDistinct**](.md)|  | 
 **mac** | **optional.String**| client mac | 
 **deviceMac** | **optional.String**| device mac | 
 **portId** | **optional.String**| port id | 
 **vlan** | **optional.String**| vlan | 
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

# **SearchSiteWiredClients**
> InlineResponse20041 SearchSiteWiredClients(ctx, siteId, optional)
searchSiteWiredClients

Search Wired Clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesClientsWiredApiSearchSiteWiredClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesClientsWiredApiSearchSiteWiredClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceMac** | **optional.String**| device mac | 
 **mac** | **optional.String**| client mac | 
 **ip** | **optional.String**| client ip | 
 **portId** | **optional.String**| port id | 
 **vlan** | **optional.String**| vlan | 
 **manufacture** | **optional.String**| manufacture | 
 **text** | **optional.String**| single entry of hostname/mac | 
 **nacruleId** | **optional.String**| nacrule_id | 
 **dhcpHostname** | **optional.String**| DHCP Hostname | 
 **dhcpFqdn** | **optional.String**| DHCP FQDN | 
 **dhcpClientIdentifier** | **optional.String**| DHCP Client Identifier | 
 **dhcpVendorClassIdentifier** | **optional.String**| DHCP Vendor Class Identifier | 
 **dhcpRequestParams** | **optional.String**| DHCP Request Parameters | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

