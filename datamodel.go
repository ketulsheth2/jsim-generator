package main

type UserEvents string
type DhcpStatus string
type LogSeverity string

// const (
// 	USER_AUTH_SUCCESSFUL   UserEvents = "User Authorization Successful"
// 	USER_AUTH_UNSUCCESSFUL            = "User Authorization Unsuccessful"
// )

// const (
// 	DHCP_READY                DhcpStatus = "DHCP_READY"
// 	DHCP_TAKEOVER_READY                  = "DHCP_TAKEOVER_READY"
// 	DHCP_TAKEOVER_IN_PROGRESS            = "DHCP_TAKEOVER_IN_PROGRESS"
// )

// const (
// 	INFO     LogSeverity = "Info"
// 	ERROR                = "Error"
// 	DEBUG                = "Debug"
// 	CRITICAL             = "Critical"
// )

type AppData struct {
	Applications struct {
		Business []string `json:"business"`
		Social   []string `json:"social"`
	} `json:"applications"`
	UserEvents  []string `json:"user_events"`
	DhcpStatus  []string `json:"dhcp_status"`
	LogSeverity []string `json:"log_severity"`
	Anomalies   []string `json:"anomalies"`
	UserAgents  []string `json:"user_agent"`
}

type Users struct {
	GhostID string `json:"ghost_id"`
	Devices struct {
		IPadPro string `json:"iPad Pro"`
		IPad    string `json:"iPad"`
		IPhone  string `json:"iPhone"`
	} `json:"devices"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Geolocation struct {
		Query       string  `json:"query"`
		Status      string  `json:"status"`
		Country     string  `json:"country"`
		CountryCode string  `json:"countryCode"`
		Region      string  `json:"region"`
		RegionName  string  `json:"regionName"`
		City        string  `json:"city"`
		Zip         string  `json:"zip"`
		Lat         float64 `json:"lat"`
		Lon         float64 `json:"lon"`
		Timezone    string  `json:"timezone"`
		Isp         string  `json:"isp"`
		Org         string  `json:"org"`
		As          string  `json:"as"`
	} `json:"geolocation"`
}

type Applications struct {
	Business []string `json:"business"`
	Social   []string `json:"social"`
}

type GhostId struct {
	Mac     string   `json:"mac"`
	Devices []string `json:"devices"`
}

type GhostApps struct {
	Active   []string `json:"active"`
	Inactive []string `json:"inactive"`
}

type GhostClients struct {
	Id   GhostId   `json:"id"`
	Apps GhostApps `json:"apps"`
}

type GhostStatusRules struct {
	Enabled      bool           `json:"enabled"`
	Interval     int            `json:"interval"`
	GhostClients []GhostClients `json:"clients"`
}

type GhostPlatform struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type GhostOs struct {
	/* TODO: can add version too here and reuse ghostplatform */
	Name string `json:"name"`
}

type GhostManagedClientDeviceInfo struct {
	Platform GhostPlatform `json:"platform"`
	Os       GhostOs       `json:"os"`
	Bot      bool          `json:"bot"`
	Dist     GhostPlatform `json:"dist"`
}

type ManagedClientsRules struct {
	Apps   GhostApps                    `json:"apps"`
	Device GhostManagedClientDeviceInfo `json:"device"`
}

type ManagedClients struct {
	Mac        string              `json:"mac"`
	LeasedIp   string              `json:"leased_ip"`
	LeaseTime  int                 `json:"lease_time"`
	RenewedAt  string              `json:"renewed_at"`
	Authorized bool                `json:"authorized"`
	UserAgent  string              `json:"user_agent"`
	Rules      ManagedClientsRules `json:"rules"`
}

type WhoIs struct {
	Status string `json:"status"`
	/* TODO: check if geolocation also needs to be renamed to asn */
	Asn         string  `json:"asn"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	Zip         string  `json:"zip"`
}

type GhostStatus struct {
	Mac             string           `json:"mac"`
	Online          bool             `json:"online"`
	GhostIP         string           `json:"ghost_ip"`
	Username        string           `json:"username"`
	State           string           `json:"state"`
	NumberOfClients int              `json:"no_clients"`
	Interface       string           `json:"interface"`
	DhcpServer      string           `json:"dhcp_server"`
	Subnet          string           `json:"subnet"`
	Gateway         string           `json:"gateway"`
	Rules           GhostStatusRules `json:"rules"`
	NameServers     []string         `json:"nameservers"`
	NumberOfPoolIps int              `json:"no_pool_ips"`
	AvailablePool   []string         `json:"available_pool"`
	ManagedClients  []ManagedClients `json:"managed_clients"`
	PubIpAddr       string           `json:"pub_ip_addr"`
	WhoIs           WhoIs            `json:"whois"`
}

type GhostEvents struct {
	Anomaly   string `json:"anomaly"`
	Severity  string `json:"severity"`
	Timestamp string `json:"timestamp"`
}

type GhostDhcpClient struct {
	Pool  []string `json:"pool"`
	Lease []string `json:"lease"`
}

type GhostDhcpProxy struct {
	GhostDhcpClient GhostDhcpClient `json:"dhcp_client"`
}

type Ghost struct {
	Status    GhostStatus    `json:"status"`
	DhcpProxy GhostDhcpProxy `json:"dhcp_proxy"`
	Events    []GhostEvents  `json:"events"`
}

type GhostData struct {
	Version string `json:"version"`
	Ghost   Ghost  `json:"ghost"`
}
