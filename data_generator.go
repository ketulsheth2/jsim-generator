package main

import (
	"math/rand"
	"time"
)

func generateGhostData(startTime time.Time, users []Users, appData AppData) *GhostData {
	// Generate random list of pool and leased IPs
	rand.Seed(time.Now().UnixNano())
	poolIps := generateRandomIps(rand.Intn(10))
	leasedIps := generateRandomIps(rand.Intn(2) + 1)

	// Generate ghost data
	ghostData := &GhostData{
		Version: "1.2.1",
		Ghost: Ghost{
			Events: generateGhostEvents(startTime, appData),
			DhcpProxy: GhostDhcpProxy{
				GhostDhcpClient: GhostDhcpClient{
					Pool:  poolIps,
					Lease: leasedIps,
				},
			},
			Status: generateGhostStatus(startTime, users, appData, poolIps, leasedIps),
		},
	}
	return ghostData
}

func generateGhostStatus(startTime time.Time, users []Users, appData AppData, poolIps, leasedIps []string) GhostStatus {
	randIp := generateRandomIp()
	rand.Seed(time.Now().UnixNano())
	// Pick a random user from the users list
	randUser := users[rand.Intn(len(users))]
	geoLocation := randUser.Geolocation
	ghostStatus := &GhostStatus{
		Mac:             randUser.GhostID,
		Online:          true,
		GhostIP:         randUser.Geolocation.Query,
		Username:        randUser.Username,
		State:           appData.DhcpStatus[rand.Intn(len(appData.DhcpStatus))],
		NumberOfClients: len(leasedIps),
		Interface:       "en0",
		DhcpServer:      randIp,
		Subnet:          "255.255.255.240",
		Gateway:         randIp,
		Rules: GhostStatusRules{
			Enabled:      true,
			Interval:     ES_INTERVAL,
			GhostClients: generateGhostClients(len(leasedIps), randUser, appData, startTime),
		},
		NameServers:     []string{generateRandomIp()},
		NumberOfPoolIps: len(poolIps),
		AvailablePool:   poolIps,
		ManagedClients:  generateManagedClients(leasedIps, randUser, startTime, appData),
		PubIpAddr:       randUser.Geolocation.Query,
		WhoIs: WhoIs{
			Status:      "success",
			Asn:         geoLocation.As,
			City:        geoLocation.City,
			Country:     geoLocation.Country,
			CountryCode: geoLocation.CountryCode,
			Latitude:    geoLocation.Lat,
			Longitude:   geoLocation.Lon,
			Timezone:    geoLocation.Timezone,
			Isp:         geoLocation.Isp,
			Org:         geoLocation.Org,
			Region:      geoLocation.Region,
			RegionName:  geoLocation.RegionName,
			Zip:         geoLocation.Zip,
		},
	}
	return *ghostStatus
}

func generateManagedClients(leasedIps []string, user Users, startTime time.Time, appData AppData) []ManagedClients {
	// Make map for managed clients for devices
	managedClientMap := GetUserDeviceMap(user)
	// Create clients based on the leased IPs
	managedClients := make([]ManagedClients, 0, len(leasedIps))
	activeApps := getActiveApps(appData, startTime)
	for i := 0; i < len(leasedIps); i++ {
		managedClient := ManagedClients{
			// TODO: which mac address?
			Mac:        managedClientMap[i].Mac,
			LeasedIp:   leasedIps[i],
			LeaseTime:  rand.Intn(500),
			RenewedAt:  startTime.Format(time.RFC3339),
			Authorized: rand.Intn(2) == 0,
			UserAgent:  appData.UserAgents[rand.Intn(len(appData.UserAgents))],
			Rules: ManagedClientsRules{
				Apps: GhostApps{
					Active:   activeApps,
					Inactive: getInactiveApps(appData, activeApps, startTime),
				},
				Device: managedClientMap[i].DeviceInfo,
			},
		}
		managedClients = append(managedClients, managedClient)
	}
	return managedClients
}

func generateGhostClients(numClients int, user Users, appData AppData, startTime time.Time) []GhostClients {
	rand.Seed(time.Now().UnixNano())
	// Get devices that are available for the user
	devices := make([]string, 0, 3)
	if user.Devices.IPad != "" {
		devices = append(devices, "ipad")
	}
	if user.Devices.IPadPro != "" {
		devices = append(devices, "ipad pro")
	}
	if user.Devices.IPhone != "" {
		devices = append(devices, "iphone")
	}

	// Get the active apps
	activeApps := getActiveApps(appData, startTime)

	// Create clients equal to the number of leased IPs
	ghostClients := make([]GhostClients, 0, numClients)
	for i := 0; i < numClients; i++ {
		ghostClient := &GhostClients{
			Id: GhostId{
				Devices: devices,
			},
			Apps: GhostApps{
				Active: activeApps,
				// Get inactive apps
				Inactive: getInactiveApps(appData, activeApps, startTime),
			},
		}
		ghostClients = append(ghostClients, *ghostClient)
	}
	return ghostClients
}

func getActiveApps(appData AppData, timeOfDay time.Time) []string {

	// Initialize random counts for business and social apps
	rand.Seed(time.Now().UnixNano())
	businessApps := rand.Intn(len(appData.Applications.Business))
	socialApps := rand.Intn(len(appData.Applications.Social))

	// If the time of day is between 9am to 5pm, there should be more business apps.
	// Outside that time range, there should be more social apps,
	// If there are more social apps than business between 9-5, swap the counts
	// If there are more business apps than social outside 9-5, swap the counts
	// NOTE: this only works if both business and social have equal elements.
	hourOfDay := timeOfDay.UTC().Hour()
	if hourOfDay >= 9 && hourOfDay <= 17 {
		if socialApps >= businessApps {
			businessApps = rand.Intn((len(appData.Applications.Business)/2)+1) + (len(appData.Applications.Business) / 2) - 1
			socialApps = socialApps / (1 * (rand.Intn(2) + 1))
			// businessApps = len(appData.Applications.Business) - 1
			// socialApps = businessApps
			// businessApps = len(appData.Applications.Business) - 1
			// businessApps, socialApps = socialApps, businessApps
		}
	} else {
		if businessApps >= socialApps {
			socialApps = rand.Intn((len(appData.Applications.Social)/2)+1) + (len(appData.Applications.Social) / 2) - 1
			businessApps = businessApps / (1 * (rand.Intn(2) + 1))
			// businessApps = socialApps
			// socialApps = len(appData.Applications.Social) - 1
			// businessApps = businessApps - rand.Intn(len(appData.Applications.Business))
			// businessApps, socialApps = socialApps, businessApps
		}
	}

	activeAppsLen := businessApps + socialApps
	activeApps := make([]string, 0, activeAppsLen)

	// Copy the original business list and append random business apps to active apps
	businessAppsList := make([]string, len(appData.Applications.Business))
	copy(businessAppsList, appData.Applications.Business)
	activeApps = AppendData(businessAppsList, activeApps, businessApps)

	// Copy the original social list and append random social apps to active apps
	socialAppsList := make([]string, len(appData.Applications.Social))
	copy(socialAppsList, appData.Applications.Social)
	activeApps = AppendData(socialAppsList, activeApps, socialApps)

	return activeApps
}

func getInactiveApps(appData AppData, activeApps []string, timeOfDay time.Time) []string {

	// Remove all active apps from the business and social list to avoid overlap
	businessAppsList := Difference(appData.Applications.Business, activeApps)
	socialAppsList := Difference(appData.Applications.Social, activeApps)

	// Get randomized business and social apps
	rand.Seed(time.Now().UnixNano())
	businessApps := rand.Intn(len(businessAppsList))
	socialApps := rand.Intn(len(socialAppsList))
	hourOfDay := timeOfDay.UTC().Hour()

	// If the time of day is between 9am to 5pm, there should be more inactive social apps.
	// Outside that time range, there should be more inactive business apps,
	// If there are more inactive business apps than social between 9-5, swap the counts
	// If there are more inactive social apps than business outside 9-5, swap the counts
	// NOTE: this only works if both business and social have equal elements.
	if hourOfDay >= 9 && hourOfDay <= 17 {
		if businessApps >= socialApps {
			socialApps = rand.Intn((len(socialAppsList)/2)+1) + (len(socialAppsList) / 2) - 1
			businessApps = businessApps / (1 * (rand.Intn(2) + 1))
		}
	} else {
		if socialApps >= businessApps {
			businessApps = rand.Intn((len(businessAppsList)/2)+1) + (len(businessAppsList) / 2) - 1
			socialApps = socialApps / (1 * (rand.Intn(2) + 1))
		}
	}

	if businessApps < 0 {
		businessApps = 0
	}

	if socialApps < 0 {
		socialApps = 0
	}

	inactiveAppsLen := businessApps + socialApps
	inactiveApps := make([]string, 0, inactiveAppsLen)

	// Append data from business and social list to inactive apps list
	inactiveApps = AppendData(businessAppsList, inactiveApps, businessApps)
	inactiveApps = AppendData(socialAppsList, inactiveApps, socialApps)
	return inactiveApps
}

func generateGhostEvents(startTime time.Time, appData AppData) []GhostEvents {
	// Create random events (upto 10)
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	// Populate ghost events based on randomized values from the app data
	randGhostEvents := make([]GhostEvents, 0, n)
	for i := 0; i < n; i++ {
		ghostEvent := &GhostEvents{
			Anomaly:   appData.Anomalies[rand.Intn(len(appData.Anomalies))],
			Severity:  appData.LogSeverity[rand.Intn(len(appData.LogSeverity))],
			Timestamp: startTime.Format(time.RFC3339),
		}
		randGhostEvents = append(randGhostEvents, *ghostEvent)
	}

	return randGhostEvents
}
