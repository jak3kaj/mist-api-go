# AclPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**[]AclPolicyAction**](acl_policy_action.md) | - for GBP-based policy, all src_tags and dst_tags have to be gbp-based - for ACL-based policy, &#x60;network&#x60; is required in either the source or destination so that we know where to attach the policy to | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**SrcTags** | **[]string** | - for GBP-based policy, all src_tags and dst_tags have to be gbp-based - for ACL-based policy, &#x60;network&#x60; is required in either the source or destination so that we know where to attach the policy to | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

