package ftdClient

type FtdModel interface {
	SecurityZone | NetworkObject | AccessPolicyItems |
		AccessRule
}

type SecurityZone struct {
	ID          string `json:"id,omitempty"`
	Version     string `json:"version,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	//Interfaces
	Mode string `json:"mode"`
	Type string `json:"type"` //securityzone
	//links
}

type NetworkObject struct {
	ID          string `json:"id,omitempty"`
	Version     string `json:"version,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	SubType     string `json:"subType"` //HOST NETWORK FQDN RANGE
	Value       string `json:"value"`
	//isSystemDefined - bool
	//dnsResolution - string
	Type string `json:"type"` //networkobject
}

type ReferenceModel struct {
	ID      string `json:"id"`
	Version string `json:"version,omitempty"`
	Name    string `json:"name,omitempty"`
	Type    string `json:"type"`
}

//Access policy wraper
type AccessPolicyItems struct {
	Items []AccessPolicy `json:"items"`
}

//add all fields later
type AccessPolicy struct {
	ID      string `json:"id"`
	Version string `json:"version,omitempty"`
	Name    string `json:"name,omitempty"`
}

type AccessRule struct {
	ID               string           `json:"id,omitempty"`
	RuleID           int              `json:"ruleId,omitempty"`
	Version          string           `json:"version,omitempty"`
	Name             string           `json:"name"`
	Description      string           `json:"description,omitempty"`
	SourceZones      []ReferenceModel `json:"sourceZones,omitempty"`
	DestinationZones []ReferenceModel `json:"destinationZones,omitempty"`

	RuleAction string `json:"ruleAction"`

	Type string `json:"type"` //accessrule
}
