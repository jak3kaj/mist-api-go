# OrgIdWxtagsBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | **float64** |  | [optional] [default to null]
**ForSite** | **bool** |  | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**LastIps** | **[]string** |  | [optional] [default to null]
**Mac** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;client&#x60;, Client MAC Address | [optional] [default to null]
**Match** | [***Object**](.md) |  | [optional] [default to null]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**Name** | **string** | The name | [default to null]
**Op** | [***Object**](.md) |  | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**ResourceMac** | **string** |  | [optional] [default to null]
**Services** | **[]string** |  | [optional] [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**Specs** | [**[]WxlanTagSpec**](wxlan_tag_spec.md) | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;spec&#x60; | [optional] [default to null]
**Subnet** | **string** |  | [optional] [default to null]
**Type_** | [***Object**](.md) |  | [default to null]
**Values** | **[]string** | required if &#x60;type&#x60;&#x3D;&#x3D;&#x60;match&#x60; and   * &#x60;match&#x60;&#x3D;&#x3D;&#x60;ap_id&#x60;: list of AP IDs   * &#x60;match&#x60;&#x3D;&#x3D;&#x60;app&#x60;: list of Application Names   * &#x60;match&#x60;&#x3D;&#x3D;&#x60;asset_mac&#x60;: list of Asset MAC Addresses   * &#x60;match&#x60;&#x3D;&#x3D;&#x60;client_mac&#x60;: list of Client MAC Addresses   * &#x60;match&#x60;&#x3D;&#x3D;&#x60;hostname&#x60;: list of Resources Hostnames   * &#x60;match&#x60;&#x3D;&#x3D;&#x60;ip_range_subnet&#x60;: list of IP Addresses and/or CIDRs   * &#x60;match&#x60;&#x3D;&#x3D;&#x60;psk_name&#x60;: list of PSK Names   * &#x60;match&#x60;&#x3D;&#x3D;&#x60;psk_role&#x60;: list of PSK Roles   * &#x60;match&#x60;&#x3D;&#x3D;&#x60;port&#x60;: list of Ports or Port Ranges   * &#x60;match&#x60;&#x3D;&#x3D;&#x60;radius_attr&#x60;: list of RADIUS Attributes. The values are [ “6&#x3D;1”, “26&#x3D;10.2.3.4” ], this support other RADIUS attributes where we know the type   * &#x60;match&#x60;&#x3D;&#x3D;&#x60;radius_class&#x60;: list of RADIUS Classes. This matches the ATTR-Class(25)   * &#x60;match&#x60;&#x3D;&#x3D;&#x60;radius_group&#x60;: list of RADIUS Groups. This is a smart tag that matches RADIUS-Filter-ID, Airespace-ACL-Name (VendorID&#x3D;14179, VendorType&#x3D;6) / Aruba-User-Role (VendorID&#x3D;14823, VendorType&#x3D;1)   * &#x60;match&#x60;&#x3D;&#x3D;&#x60;radius_username&#x60;: list of RADIUS Usernames. This matches the ATTR-User-Name(1)   * &#x60;match&#x60;&#x3D;&#x3D;&#x60;sdkclient_uuid&#x60;: list of SDK UUIDs   * &#x60;match&#x60;&#x3D;&#x3D;&#x60;wlan_id&#x60;: list of WLAN IDs  **Notes**: Variables are not allowed | [optional] [default to null]
**VlanId** | [***Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

