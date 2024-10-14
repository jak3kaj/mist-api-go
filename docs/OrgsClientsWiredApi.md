# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgWiredClients**](OrgsClientsWiredApi.md#CountOrgWiredClients) | **Get** /api/v1/orgs/{org_id}/wired_clients/count | countOrgWiredClients
[**SearchOrgWiredClients**](OrgsClientsWiredApi.md#SearchOrgWiredClients) | **Get** /api/v1/orgs/{org_id}/wired_clients/search | searchOrgWiredClients

# **CountOrgWiredClients**
> InlineResponse20016 CountOrgWiredClients(ctx, orgId, optional)
countOrgWiredClients

Count by Distinct Attributes of Clients  Note: For list of avaialable `type` values, please refer to [listClientEventsDefinitions]($e/Constants%20Events/listClientEventsDefinitions)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsClientsWiredApiCountOrgWiredClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsClientsWiredApiCountOrgWiredClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of OrgWiredClientsCountDistinct**](.md)|  | 
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

# **SearchOrgWiredClients**
> InlineResponse20041 SearchOrgWiredClients(ctx, orgId, optional)
searchOrgWiredClients

Search for Wired Clients in org  Note: For list of avaialable `type` values, please refer to [listClientEventsDefinitions]($e/Constants%20Events/listClientEventsDefinitions)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsClientsWiredApiSearchOrgWiredClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsClientsWiredApiSearchOrgWiredClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **optional.String**| Site ID | 
 **deviceMac** | **optional.String**| device mac | 
 **mac** | **optional.String**| client mac | 
 **portId** | **optional.String**| port id | 
 **vlan** | **optional.Int32**| vlan | 
 **ip** | **optional.String**| ip | 
 **manufacture** | **optional.String**| client manufacturer | 
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

