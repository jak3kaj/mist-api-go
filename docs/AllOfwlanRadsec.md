# AllOfwlanRadsec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoaEnabled** | **bool** |  | [optional] [default to false]
**Enabled** | **bool** |  | [optional] [default to null]
**IdleTimeout** | **int32** |  | [optional] [default to null]
**MxclusterIds** | **[]string** | To use Org mxedges when this WLAN does not use mxtunnel, specify their mxcluster_ids. Org mxedge(s) identified by mxcluster_ids | [optional] [default to null]
**ProxyHosts** | **[]string** | default is site.mxedge.radsec.proxy_hosts which must be a superset of all wlans[*].radsec.proxy_hosts when radsec.proxy_hosts are not used, tunnel peers (org or site mxedges) are used irrespective of use_site_mxedge | [optional] [default to null]
**ServerName** | **string** | name of the server to verify (against the cacerts in Org Setting). Only if not Mist Edge. | [optional] [default to null]
**Servers** | [**[]RadsecServer**](radsec_server.md) | List of Radsec Servers. Only if not Mist Edge. | [optional] [default to null]
**UseMxedge** | **bool** | use mxedge(s) as radsecproxy | [optional] [default to null]
**UseSiteMxedge** | **bool** | To use Site mxedges when this WLAN does not use mxtunnel | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

