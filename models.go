package ftdClient

type FtdModel interface {
	SecurityZone | NetworkObject | AccessRule | NetworkInterface
}

type SecurityZone struct {
	ID          string           `json:"id,omitempty"`
	Version     string           `json:"version,omitempty"`
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Interfaces  []ReferenceModel `json:"interfaces,omitempty"`
	Mode        string           `json:"mode"`
	Type        string           `json:"type"` //securityzone
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
	DnsResolution string `json:"dnsResolution,omitempty"` //['IPV4_ONLY', 'IPV6_ONLY', 'IPV4_AND_IPV6']
	Type          string `json:"type"`                    //networkobject
}

type ReferenceModel struct {
	ID      string `json:"id,omitempty"`
	Version string `json:"version,omitempty"`
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
}

type ItemTypes interface {
	SecurityZone | AccessPolicy | NetworkInterface
}

// List return wraper
type Items[T ItemTypes] struct {
	Items []T `json:"items"`
}

// add all fields later
type AccessPolicy struct {
	ID      string `json:"id"`
	Version string `json:"version,omitempty"`
	Name    string `json:"name,omitempty"`
}

type AccessRule struct {
	ID                  string            `json:"id,omitempty"`
	RuleID              int               `json:"ruleId,omitempty"`
	Version             string            `json:"version,omitempty"`
	Name                string            `json:"name"`
	Description         string            `json:"description,omitempty"`
	SourceZones         []ReferenceModel  `json:"sourceZones,omitempty"`
	DestinationZones    []ReferenceModel  `json:"destinationZones,omitempty"`
	SourceNetworks      []ReferenceModel  `json:"sourceNetworks,omitempty"`
	DestinationNetworks []ReferenceModel  `json:"destinationNetworks,omitempty"`
	SourcePorts         []ReferenceModel  `json:"sourcePorts,omitempty"`
	DestinationPorts    []ReferenceModel  `json:"destinationPorts,omitempty"`
	RulePosition        int               `json:"rulePosition,omitempty"`
	RuleAction          string            `json:"ruleAction"`
	EventLogAction      string            `json:"eventLogAction,omitempty"` //['LOG_FLOW_START', 'LOG_FLOW_END', 'LOG_BOTH', 'LOG_NONE']
	IdentitySources     []ReferenceModel  `json:"identitySources,omitempty"`
	Users               []TrafficEntry    `json:"users,omitempty"`
	EmbeddedAppFilter   EmbeddedAppFilter `json:"embeddedAppFilter,omitempty"`
	UrlFilter           EmbeddedURLFilter `json:"urlFilter,omitempty"`
	IntrusionPolicy     ReferenceModel    `json:"intrusionPolicy,omitempty"`
	FilePolicy          ReferenceModel    `json:"filePolicy,omitempty"`
	LogFiles            bool              `json:"logFiles,omitempty"`
	SyslogServer        ReferenceModel    `json:"syslogServer,omitempty"`
	//HitCount                  bool              `json:"hitCount,omitempty"`
	DestinationDynamicObjects []ReferenceModel `json:"destinationDynamicObjects,omitempty"`
	SourceDynamicObjects      []ReferenceModel `json:"sourceDynamicObjects,omitempty"`
	TimeRangeObjects          []ReferenceModel `json:"timeRangeObjects,omitempty"`
	Type                      string           `json:"type"` //accessrule
}

type TrafficEntry struct {
	Name           string         `json:"name"`
	IdentitySource ReferenceModel `json:"identitySource"`
	Type           string         `json:"type"` //trafficentry
}

type EmbeddedAppFilter struct {
	Applications       []ReferenceModel             `json:"applications,omitempty"`
	ApplicationFilters []ReferenceModel             `json:"applicationFilters,omitempty"`
	Conditions         []ApplicationFilterCondition `json:"ApplicationFilterCondition,omitempty"`
	Type               string                       `json:"type,omitempty"` //embeddedappfilter
}

type EmbeddedURLFilter struct {
	UrlObjects    []ReferenceModel     `json:"urlObjects,omitempty"`
	UrlCategories []URLCategoryMatcher `json:"urlCategories,omitempty"`
	Type          string               `json:"type,omitempty"` //embeddedurlfilter
}

type URLCategoryMatcher struct {
	UrlCategory                 ReferenceModel `json:"urlCategory"`
	UrlReputation               ReferenceModel `json:"urlReputation"`
	IncludeUnknownUrlReputation bool           `json:"includeUnknownUrlReputation,omitempty"`
	Type                        string         `json:"type"` //urlcategorymatcher
}

type ApplicationFilterCondition struct {
	Risks            []RiskCondition         `json:"risks,omitempty"`
	Productivities   []ProductivityCondition `json:"productivities,omitempty"`
	Tags             []ReferenceModel        `json:"tags,omitempty"`
	Categories       []ReferenceModel        `json:"categories,omitempty"`
	Filter           string                  `json:"filter,omitempty"`
	ApplicationTypes []TypeCondition         `json:"applicationTypes,omitempty"`
	Type             string                  `json:"type"` //applicationfiltercondition
}

type RiskCondition struct {
	Risk string `json:"risk,omitempty"` // ['UNKNOWN', 'VERY_LOW', 'LOW', 'MEDIUM', 'HIGH', 'CRITICAL'],
	Type string `json:"type"`           //riskcondition
}

type ProductivityCondition struct {
	Productivity string `json:"productivity,omitempty"`
	Type         string `json:"type"` //productivitycondition
}

type TypeCondition struct {
	ApplicationType string `json:"applicationType,omitempty"`
	Type            string `json:"type"` //typecondition
}

type NetworkInterface struct {
	ID               string        `json:"id,omitempty"`
	Version          string        `json:"version,omitempty"`
	Name             string        `json:"name,omitempty"`
	Description      string        `json:"description,omitempty"`
	HardwareName     string        `json:"hardwareName,omitempty"`
	MonitorInterface bool          `json:"monitorInterface,omitempty"`
	Ipv4             InterfaceIPv4 `json:"ipv4,omitempty"`
	//ipv6
	ManagementOnly      bool   `json:"managementOnly,omitempty"`
	ManagementInterface bool   `json:"managementInterface,omitempty"`
	Mode                string `json:"mode,omitempty"` //['PASSIVE', 'ROUTED', 'SWITCHPORT', 'BRIDGEGROUPMEMBER']
	Mtu                 int    `json:"mtu,omitempty"`
	Enabled             bool   `json:"enabled,omitempty"`
	MacAddress          string `json:"macAddress,omitempty"`
	StandbyMacAddress   string `json:"standbyMacAddress,omitempty"`
	//pppoe
	CtsEnabled bool   `json:"ctsEnabled,omitempty"`
	FecMode    string `json:"fecMode,omitempty"`    //['AUTO', 'CL108_RS', 'CL74_FC', 'CL91_RS', 'DISABLE']
	SpeedType  string `json:"speedType,omitempty"`  //['SFP_DETECT', 'AUTO', 'FORTY_THOUSAND', 'TWENTYFIVE_THOUSAND', 'TEN_THOUSAND', 'THOUSAND', 'HUNDRED', 'TEN', 'NO_NEGOTIATE', 'IGNORE']
	DuplexType string `json:"duplexType,omitempty"` //['AUTO', 'HALF', 'FULL', 'IGNORE']
	//switchPortConfig
	//powerOverEthernet
	AutoNeg             bool   `json:"autoNeg,omitempty"`
	BreakOutCapable     bool   `json:"breakOutCapable,omitempty"`
	Present             bool   `json:"present,omitempty"`
	SplitInterface      bool   `json:"splitInterface,omitempty"`
	TenGigabitInterface bool   `json:"tenGigabitInterface,omitempty"`
	GigabitInterface    bool   `json:"gigabitInterface,omitempty"`
	Type                string `json:"type"` //physicalinterface

}

type InterfaceIPv4 struct {
	IpType                string        `json:"ipType,omitempty"` //['DHCP', 'STATIC', 'PPPOE']
	DefaultRouteUsingDHCP bool          `json:"defaultRouteUsingDHCP,omitempty"`
	DhcpRouteMetric       int           `json:"dhcpRouteMetric,omitempty"`
	IpAddress             HAIPv4Address `json:"ipAddress,omitempty"`
	Dhcp                  bool          `json:"dhcp,omitempty"`
	AddressNull           bool          `json:"addressNull,omitempty"`
	Type                  string        `json:"type,omitempty"` //interfaceipv4
}

type HAIPv4Address struct {
	IpAddress        string `json:"ipAddress,omitempty"`
	Netmask          string `json:"netmask,omitempty"`
	StandbyIpAddress string `json:"standbyIpAddress,omitempty"`
	Type             string `json:"type,omitempty"` //haipv4address
}
