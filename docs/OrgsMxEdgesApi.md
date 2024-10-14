# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrgMxEdgeImage**](OrgsMxEdgesApi.md#AddOrgMxEdgeImage) | **Post** /api/v1/orgs/{org_id}/mxedges/{mxedge_id}/image/{image_number} | addOrgMxEdgeImage
[**AssignOrgMxEdgeToSite**](OrgsMxEdgesApi.md#AssignOrgMxEdgeToSite) | **Post** /api/v1/orgs/{org_id}/mxedges/assign | assignOrgMxEdgeToSite
[**BounceOrgMxEdgeDataPorts**](OrgsMxEdgesApi.md#BounceOrgMxEdgeDataPorts) | **Post** /api/v1/orgs/{org_id}/mxedges/{mxedge_id}/services/tunterm/bounce_port | bounceOrgMxEdgeDataPorts
[**ClaimOrgMxEdge**](OrgsMxEdgesApi.md#ClaimOrgMxEdge) | **Post** /api/v1/orgs/{org_id}/mxedges/claim | claimOrgMxEdge
[**ControlOrgMxEdgeServices**](OrgsMxEdgesApi.md#ControlOrgMxEdgeServices) | **Post** /api/v1/orgs/{org_id}/mxedges/{mxedge_id}/services/{name}/{action} | controlOrgMxEdgeServices
[**CountOrgMxEdges**](OrgsMxEdgesApi.md#CountOrgMxEdges) | **Get** /api/v1/orgs/{org_id}/mxedges/count | countOrgMxEdges
[**CountOrgSiteMxEdgeEvents**](OrgsMxEdgesApi.md#CountOrgSiteMxEdgeEvents) | **Get** /api/v1/orgs/{org_id}/mxedges/events/count | countOrgSiteMxEdgeEvents
[**CreateOrgMxEdge**](OrgsMxEdgesApi.md#CreateOrgMxEdge) | **Post** /api/v1/orgs/{org_id}/mxedges | createOrgMxEdge
[**DeleteOrgMxEdge**](OrgsMxEdgesApi.md#DeleteOrgMxEdge) | **Delete** /api/v1/orgs/{org_id}/mxedges/{mxedge_id} | deleteOrgMxEdge
[**DeleteOrgMxEdgeImage**](OrgsMxEdgesApi.md#DeleteOrgMxEdgeImage) | **Delete** /api/v1/orgs/{org_id}/mxedges/{mxedge_id}/image/{image_number} | deleteOrgMxEdgeImage
[**GetOrgMxEdge**](OrgsMxEdgesApi.md#GetOrgMxEdge) | **Get** /api/v1/orgs/{org_id}/mxedges/{mxedge_id} | getOrgMxEdge
[**GetOrgMxEdgeUpgradeInfo**](OrgsMxEdgesApi.md#GetOrgMxEdgeUpgradeInfo) | **Get** /api/v1/orgs/{org_id}/mxedges/version | getOrgMxEdgeUpgradeInfo
[**ListOrgMxEdges**](OrgsMxEdgesApi.md#ListOrgMxEdges) | **Get** /api/v1/orgs/{org_id}/mxedges | listOrgMxEdges
[**RestartOrgMxEdge**](OrgsMxEdgesApi.md#RestartOrgMxEdge) | **Post** /api/v1/orgs/{org_id}/mxedges/{mxedge_id}/restart | restartOrgMxEdge
[**SearchOrgMistEdgeEvents**](OrgsMxEdgesApi.md#SearchOrgMistEdgeEvents) | **Get** /api/v1/orgs/{org_id}/mxedges/events/search | searchOrgMistEdgeEvents
[**SearchOrgMxEdges**](OrgsMxEdgesApi.md#SearchOrgMxEdges) | **Get** /api/v1/orgs/{org_id}/mxedges/search | searchOrgMxEdges
[**UnassignOrgMxEdgeFromSite**](OrgsMxEdgesApi.md#UnassignOrgMxEdgeFromSite) | **Post** /api/v1/orgs/{org_id}/mxedges/unassign | unassignOrgMxEdgeFromSite
[**UnregisterOrgMxEdge**](OrgsMxEdgesApi.md#UnregisterOrgMxEdge) | **Post** /api/v1/orgs/{org_id}/mxedges/{mxedge_id}/unregister | unregisterOrgMxEdge
[**UpdateOrgMxEdge**](OrgsMxEdgesApi.md#UpdateOrgMxEdge) | **Put** /api/v1/orgs/{org_id}/mxedges/{mxedge_id} | updateOrgMxEdge
[**UploadOrgMxEdgeSupportFiles**](OrgsMxEdgesApi.md#UploadOrgMxEdgeSupportFiles) | **Post** /api/v1/orgs/{org_id}/mxedges/{mxedge_id}/support | uploadOrgMxEdgeSupportFiles

# **AddOrgMxEdgeImage**
> AddOrgMxEdgeImage(ctx, orgId, mxedgeId, imageNumber, optional)
addOrgMxEdgeImage

Attach up to 3 images to a mxedge

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **mxedgeId** | [**string**](.md)|  | 
  **imageNumber** | **int32**|  | 
 **optional** | ***OrgsMxEdgesApiAddOrgMxEdgeImageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxEdgesApiAddOrgMxEdgeImageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of ImageImport**](ImageImport.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignOrgMxEdgeToSite**
> InlineResponse20046 AssignOrgMxEdgeToSite(ctx, orgId, optional)
assignOrgMxEdgeToSite

Assign Org MxEdge to Site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxEdgesApiAssignOrgMxEdgeToSiteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxEdgesApiAssignOrgMxEdgeToSiteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MxedgesAssignBody**](MxedgesAssignBody.md)| Request Body | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BounceOrgMxEdgeDataPorts**
> BounceOrgMxEdgeDataPorts(ctx, orgId, mxedgeId, optional)
bounceOrgMxEdgeDataPorts

Bounce TunTerm Data Ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **mxedgeId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxEdgesApiBounceOrgMxEdgeDataPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxEdgesApiBounceOrgMxEdgeDataPortsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of TuntermBouncePortBody**](TuntermBouncePortBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClaimOrgMxEdge**
> InlineResponse20068 ClaimOrgMxEdge(ctx, orgId, optional)
claimOrgMxEdge

For a Mist Edge in default state, it will show a random claim code like `135-546-673` which you can “claim” it into your Org

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxEdgesApiClaimOrgMxEdgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxEdgesApiClaimOrgMxEdgeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MxedgesClaimBody**](MxedgesClaimBody.md)| Request Body | 

### Return type

[**InlineResponse20068**](inline_response_200_68.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ControlOrgMxEdgeServices**
> ControlOrgMxEdgeServices(ctx, orgId, mxedgeId, name, action)
controlOrgMxEdgeServices

Control Services on a Mist Edge

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **mxedgeId** | [**string**](.md)|  | 
  **name** | [**Name**](.md)| enum: &#x60;mxagent&#x60;, &#x60;mxdas&#x60;, &#x60;mxnacedge&#x60;, &#x60;mxocproxy&#x60;, &#x60;radsecproxy&#x60;, &#x60;tunterm&#x60; | 
  **action** | [**Action**](.md)| restart or start or stop | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountOrgMxEdges**
> InlineResponse20016 CountOrgMxEdges(ctx, orgId, optional)
countOrgMxEdges

Count Org Mist Edges

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxEdgesApiCountOrgMxEdgesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxEdgesApiCountOrgMxEdgesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of OrgMxedgeCountDistinct**](.md)|  | 
 **mxedgeId** | **optional.String**| mist edge id | 
 **siteId** | **optional.String**| mist edge site id | 
 **mxclusterId** | **optional.String**| mist edge cluster id | 
 **model** | **optional.String**| model name | 
 **distro** | **optional.String**| debian code name(buster, bullseye) | 
 **tuntermVersion** | **optional.String**| tunterm version | 
 **sort** | **optional.String**| sort options, -prefix represents DESC order, default is -last_seen | 
 **stats** | **optional.Bool**| whether to return device stats, default is false | 
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

# **CountOrgSiteMxEdgeEvents**
> InlineResponse20016 CountOrgSiteMxEdgeEvents(ctx, orgId, optional)
countOrgSiteMxEdgeEvents

Count Org Mist Edge Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxEdgesApiCountOrgSiteMxEdgeEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxEdgesApiCountOrgSiteMxEdgeEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of OrgMxedgeEventsCountDistinct**](.md)|  | 
 **mxedgeId** | **optional.String**| mist edge id | 
 **mxclusterId** | **optional.String**| mist edge cluster id | 
 **type_** | **optional.String**| see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **service** | **optional.String**| service running on mist edge(mxagent, tunterm etc) | 
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

# **CreateOrgMxEdge**
> InlineResponse20067 CreateOrgMxEdge(ctx, orgId, optional)
createOrgMxEdge

Create MxEdge

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxEdgesApiCreateOrgMxEdgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxEdgesApiCreateOrgMxEdgeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdMxedgesBody**](OrgIdMxedgesBody.md)| Request Body | 

### Return type

[**InlineResponse20067**](inline_response_200_67.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgMxEdge**
> DeleteOrgMxEdge(ctx, orgId, mxedgeId)
deleteOrgMxEdge

Delete Org MxEdge

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **mxedgeId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgMxEdgeImage**
> DeleteOrgMxEdgeImage(ctx, orgId, mxedgeId, imageNumber)
deleteOrgMxEdgeImage

Remove MxEdge Image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **mxedgeId** | [**string**](.md)|  | 
  **imageNumber** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgMxEdge**
> InlineResponse20067 GetOrgMxEdge(ctx, orgId, mxedgeId)
getOrgMxEdge

Get Org MxEdge details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **mxedgeId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20067**](inline_response_200_67.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgMxEdgeUpgradeInfo**
> []MxedgeUpgradeInfoItems GetOrgMxEdgeUpgradeInfo(ctx, orgId, optional)
getOrgMxEdgeUpgradeInfo

Get Mist Edge Upgrade Information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxEdgesApiGetOrgMxEdgeUpgradeInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxEdgesApiGetOrgMxEdgeUpgradeInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channel** | [**optional.Interface of Channel**](.md)| upgrade channel to follow, stable (default) / beta / alpha | 

### Return type

[**[]MxedgeUpgradeInfoItems**](mxedge_upgrade_info_items.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgMxEdges**
> []Mxedge ListOrgMxEdges(ctx, orgId, optional)
listOrgMxEdges

Get List of Org MxEdges

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxEdgesApiListOrgMxEdgesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxEdgesApiListOrgMxEdgesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forSites** | [**optional.Interface of ForSites**](.md)| filter for site level mist edges | 
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Mxedge**](mxedge.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartOrgMxEdge**
> RestartOrgMxEdge(ctx, orgId, mxedgeId)
restartOrgMxEdge

In the case where a Mist Edge is replaced, you would need to unregister it. Which disconnects the currently the connected Mist Edge and allow another to register.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **mxedgeId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrgMistEdgeEvents**
> InlineResponse20069 SearchOrgMistEdgeEvents(ctx, orgId, optional)
searchOrgMistEdgeEvents

Search Org Mist Edge Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxEdgesApiSearchOrgMistEdgeEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxEdgesApiSearchOrgMistEdgeEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mxedgeId** | **optional.String**| mist edge id | 
 **mxclusterId** | **optional.String**| mist edge cluster id | 
 **type_** | **optional.String**| see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **service** | **optional.String**| service running on mist edge(mxagent, tunterm etc) | 
 **component** | **optional.String**| component like PS1, PS2 | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]

### Return type

[**InlineResponse20069**](inline_response_200_69.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrgMxEdges**
> InlineResponse20070 SearchOrgMxEdges(ctx, orgId, optional)
searchOrgMxEdges

Search Org Mist Edges

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxEdgesApiSearchOrgMxEdgesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxEdgesApiSearchOrgMxEdgesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mxedgeId** | **optional.String**| mist edge id | 
 **siteId** | **optional.String**| mist edge site id | 
 **mxclusterId** | **optional.String**| mist edge cluster id | 
 **model** | **optional.String**| model name | 
 **distro** | **optional.String**| debian code name(buster, bullseye) | 
 **tuntermVersion** | **optional.String**| tunterm version | 
 **sort** | **optional.String**| sort options, -prefix represents DESC order, default is -last_seen | 
 **stats** | **optional.Bool**| whether to return device stats, default is false | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse20070**](inline_response_200_70.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnassignOrgMxEdgeFromSite**
> InlineResponse20046 UnassignOrgMxEdgeFromSite(ctx, orgId, optional)
unassignOrgMxEdgeFromSite

Unassign Org MxEdge from Site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxEdgesApiUnassignOrgMxEdgeFromSiteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxEdgesApiUnassignOrgMxEdgeFromSiteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MxedgesUnassignBody**](MxedgesUnassignBody.md)| Request Body | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnregisterOrgMxEdge**
> UnregisterOrgMxEdge(ctx, orgId, mxedgeId)
unregisterOrgMxEdge

In the case where a Mist Edge is replaced, you would need to unregister it. Which disconnects the currently the connected Mist Edge and allow another to register.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **mxedgeId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgMxEdge**
> InlineResponse20067 UpdateOrgMxEdge(ctx, orgId, mxedgeId, optional)
updateOrgMxEdge

Update Org MxEdge

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **mxedgeId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxEdgesApiUpdateOrgMxEdgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxEdgesApiUpdateOrgMxEdgeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of MxedgesMxedgeIdBody**](MxedgesMxedgeIdBody.md)| Request Body | 

### Return type

[**InlineResponse20067**](inline_response_200_67.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadOrgMxEdgeSupportFiles**
> UploadOrgMxEdgeSupportFiles(ctx, orgId, mxedgeId)
uploadOrgMxEdgeSupportFiles

Support / Upload Mist Edge support files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **mxedgeId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

