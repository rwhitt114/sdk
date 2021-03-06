package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// WafPolicy waf policy
// swagger:model WafPolicy
type WafPolicy struct {

	// Creator name. Field introduced in 17.2.4.
	CreatedBy string `json:"created_by,omitempty"`

	// WAF Rules are categorized in to groups based on their characterization. These groups are system created with CRS groups. Field introduced in 17.2.1.
	CrsGroups []*WafRuleGroup `json:"crs_groups,omitempty"`

	//  Field introduced in 17.2.1.
	Description string `json:"description,omitempty"`

	// WAF Policy mode. This can be detection or enforcement. Enum options - WAF_MODE_DETECTION_ONLY, WAF_MODE_ENFORCEMENT. Field introduced in 17.2.1.
	// Required: true
	Mode string `json:"mode"`

	//  Field introduced in 17.2.1.
	// Required: true
	Name string `json:"name"`

	// WAF Ruleset paranoia  mode. This is used to select Rules based on the paranoia-level tag. Enum options - WAF_PARANOIA_LEVEL_LOW, WAF_PARANOIA_LEVEL_MEDIUM, WAF_PARANOIA_LEVEL_HIGH, WAF_PARANOIA_LEVEL_EXTREME. Field introduced in 17.2.1.
	ParanoiaLevel string `json:"paranoia_level,omitempty"`

	// WAF Rules are categorized in to groups based on their characterization. These groups are created by the user and will be enforced after the CRS groups. Field introduced in 17.2.1.
	PostCrsGroups []*WafRuleGroup `json:"post_crs_groups,omitempty"`

	// WAF Rules are categorized in to groups based on their characterization. These groups are created by the user and will be  enforced before the CRS groups. Field introduced in 17.2.1.
	PreCrsGroups []*WafRuleGroup `json:"pre_crs_groups,omitempty"`

	//  It is a reference to an object of type Tenant. Field introduced in 17.2.1.
	TenantRef string `json:"tenant_ref,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	//  Field introduced in 17.2.1.
	UUID string `json:"uuid,omitempty"`

	// WAF Profile for WAF policy. It is a reference to an object of type WafProfile. Field introduced in 17.2.1.
	// Required: true
	WafProfileRef string `json:"waf_profile_ref"`
}
