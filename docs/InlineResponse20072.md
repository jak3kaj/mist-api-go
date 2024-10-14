# InlineResponse20072

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUsermacOverride** | **bool** | can be set to true to allow the override by usermac result | [optional] [default to false]
**CreatedTime** | **float64** |  | [optional] [default to null]
**EgressVlanNames** | **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;egress_vlan_names&#x60;, list of egress vlans to return | [optional] [default to null]
**GbpTag** | **int32** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;gbp_tag&#x60; | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**Match** | [***Object**](.md) |  | [optional] [default to null]
**MatchAll** | **bool** | This field is applicable only when &#x60;type&#x60;&#x3D;&#x3D;&#x60;match&#x60;   * &#x60;false&#x60;: means it is sufficient to match any of the values (i.e., match-any behavior)   * &#x60;true&#x60;: means all values should be matched (i.e., match-all behavior)   Currently it makes sense to set this field to &#x60;true&#x60; only if the &#x60;match&#x60;&#x3D;&#x3D;&#x60;idp_role&#x60; or &#x60;match&#x60;&#x3D;&#x3D;&#x60;usermac_label&#x60; | [optional] [default to false]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**RadiusAttrs** | **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;radius_attrs&#x60;, user can specify a list of one or more standard attributes in the field \&quot;radius_attrs\&quot;.  It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected. Note that it is allowed to have more than one radius_attrs in the result of a given rule. | [optional] [default to null]
**RadiusGroup** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;radius_group&#x60; | [optional] [default to null]
**RadiusVendorAttrs** | **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;radius_vendor_attrs&#x60;, user can specify a list of one or more vendor-specific attributes in the field \&quot;radius_vendor_attrs\&quot;.  It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected. Note that it is allowed to have more than one radius_vendor_attrs in the result of a given rule. | [optional] [default to null]
**SessionTimeout** | **int32** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;session_timeout, in seconds | [optional] [default to null]
**Type_** | [***Object**](.md) |  | [default to null]
**UsernameAttr** | [***Object**](.md) |  | [optional] [default to null]
**Values** | **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;match&#x60; | [optional] [default to null]
**Vlan** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;vlan&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

