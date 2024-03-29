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
	SecurityZone | AccessPolicy | NetworkInterface | TcpUdpPort | Application | ApplicationCategory
}

// List return wraper
type Items[T ItemTypes] struct {
	Items  []T    `json:"items"`
	Paging Paging `json:"paging"`
}

type Paging struct {
	Prev []string `json:"prev"`
	Next []string `json:"next"`
}

// add all fields later
type AccessPolicy struct {
	ID                    string              `json:"id"`
	Version               string              `json:"version,omitempty"`
	Name                  string              `json:"name,omitempty"`
	DefaultAction         AccessDefaultAction `json:"defaultAction,omitempty"`
	SslPolicy             ReferenceModel      `json:"sslPolicy,omitempty"`
	CertVisibilityEnabled bool                `json:"certVisibilityEnabled,omitempty"`
	NetworkAnalysisPolicy ReferenceModel      `json:"networkAnalysisPolicy,omitempty"`
	AdvancedSettings      AdvancedSettings    `json:"advancedSettings,omitempty"`
	IdentityPolicySetting ReferenceModel      `json:"identityPolicySetting,omitempty"`
	SecurityIntelligence  ReferenceModel      `json:"securityIntelligence,omitempty"`
	Type                  string              `json:"type"` //accesspolicy
}

type AdvancedSettings struct {
	DnsReputationEnforcementEnabled bool   `json:"dnsReputationEnforcementEnabled,omitempty"`
	Type                            string `json:"type,omitempty"` //advancedsettings
}

type AccessDefaultAction struct {
	Action          string         `json:"action,omitempty"`         // ['PERMIT', 'TRUST', 'DENY']
	EventLogAction  string         `json:"eventLogAction,omitempty"` //[ 'LOG_FLOW_END', 'LOG_BOTH', 'LOG_NONE']
	IntrusionPolicy ReferenceModel `json:"intrusionPolicy,omitempty"`
	SyslogServer    ReferenceModel `json:"syslogServer,omitempty"`
	//hitCount
	Type string `json:"type"` //accessdefaultaction
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
	Conditions         []ApplicationFilterCondition `json:"conditions,omitempty"`
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
	Type                        string         `json:"type,omitempty"` //urlcategorymatcher
}

type ApplicationFilterCondition struct {
	Risks            []RiskCondition         `json:"risks,omitempty"`
	Productivities   []ProductivityCondition `json:"productivities,omitempty"`
	Tags             []ReferenceModel        `json:"tags,omitempty"`
	Categories       []ReferenceModel        `json:"categories,omitempty"`
	Filter           string                  `json:"filter,omitempty"`
	ApplicationTypes []TypeCondition         `json:"applicationTypes,omitempty"`
	Type             string                  `json:"type,omitempty"` //applicationfiltercondition
}

type RiskCondition struct {
	Risk string `json:"risk,omitempty"` // ['UNKNOWN', 'VERY_LOW', 'LOW', 'MEDIUM', 'HIGH', 'CRITICAL'],
	Type string `json:"type,omitempty"` //riskcondition
}

type ProductivityCondition struct {
	Productivity string `json:"productivity,omitempty"`
	Type         string `json:"type,omitempty"` //productivitycondition
}

type TypeCondition struct {
	ApplicationType string `json:"applicationType,omitempty"`
	Type            string `json:"type,omitempty"` //typecondition
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

type TcpUdpPort struct {
	Version         string `json:"version,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	IsSystemDefined bool   `json:"isSystemDefined,omitempty"`
	Port            string `json:"port,omitempty"`
	ID              string `json:"id,omitempty"`
	Type            string `json:"type,omitempty"` //tcpportobject or udpportobject
}

type ProtocolObject struct {
	Version         string `json:"version,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	IsSystemDefined bool   `json:"isSystemDefined,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	ID              string `json:"id,omitempty"`
	Type            string `json:"type,omitempty"` //protocolobject
}

type Icmpv4port struct {
	Version         string `json:"version,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	IsSystemDefined bool   `json:"isSystemDefined,omitempty"`
	Icmpv4Type      string `json:"icmpv4Type,omitempty"`
	Icmpv4Code      string `json:"icmpv4Code,omitempty"`
	ID              string `json:"id,omitempty"`
	Type            string `json:"type,omitempty"` //icmpv4portobject
}

type PortGroup struct {
	Version         string         `json:"version,omitempty"`
	Name            string         `json:"name,omitempty"`
	Description     string         `json:"description,omitempty"`
	IsSystemDefined bool           `json:"isSystemDefined,omitempty"`
	ID              string         `json:"id,omitempty"`
	Objects         ReferenceModel `json:"objects,omitempty"`
	Type            string         `json:"type,omitempty"` //portobjectgroup
}

type Application struct {
	Name             string           `json:"name,omitempty"`
	AppId            int              `json:"appId,omitempty"`
	Description      string           `json:"description,omitempty"`
	Tags             []ReferenceModel `json:"tags,omitempty"`
	Categories       []ReferenceModel `json:"categories,omitempty"`
	Deprecated       bool             `json:"deprecated,omitempty"`
	TagIds           []int            `json:"tagIds,omitempty"`
	CategoryIds      []int            `json:"categoryIds,omitempty"`
	ID               string           `json:"id,omitempty"`
	ApplicationTypes []string         `json:"applicationTypes,omitempty"`
	Productivity     int              `json:"productivity,omitempty"`
	Risk             int              `json:"risk,omitempty"`
	Type             string           `json:"type,omitempty"` //application
}

type ApplicationFilter struct {
	Name            string                       `json:"name,omitempty"`
	ID              string                       `json:"id,omitempty"`
	Version         string                       `json:"version,omitempty"`
	Applications    []ReferenceModel             `json:"applications,omitempty"`
	IsSystemDefined bool                         `json:"isSystemDefined,omitempty"`
	Conditions      []ApplicationFilterCondition `json:"conditions,omitempty"`
	Type            string                       `json:"type,omitempty"` //applicationfilter
}

type ApplicationCategory struct {
	Name        string `json:"name,omitempty"`
	ID          string `json:"id,omitempty"`
	Version     string `json:"version,omitempty"`
	AppId       int    `json:"appId,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"` //applicationcategory
}

type SamlServer struct {
	Version                 string         `json:"version,omitempty"`
	Name                    string         `json:"name,omitempty"`
	Description             string         `json:"description,omitempty"`
	SamlIssuerURL           string         `json:"samlIssuerURL,omitempty"`
	SignInURL               string         `json:"signInURL,omitempty"`
	SignOutURL              string         `json:"signOutURL,omitempty"`
	SamlIssuerCert          ReferenceModel `json:"samlIssuerCert,omitempty"`
	FtdCert                 ReferenceModel `json:"ftdCert,omitempty"`
	RequestTimeout          int            `json:"requestTimeout,omitempty"`
	ServerOnInternalNetwork bool           `json:"serverOnInternalNetwork,omitempty"`
	ReAuthAtLogin           bool           `json:"reAuthAtLogin,omitempty"`
	SignatureType           string         `json:"signatureType,omitempty"`
	ID                      string         `json:"id,omitempty"`
	Type                    string         `json:"type,omitempty"` //samlserver
}
