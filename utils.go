package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
)

type ManagedClientsInfo struct {
	Mac        string
	DeviceInfo GhostManagedClientDeviceInfo
}

// Difference returns the elements in `a` that aren't in `b`.
func Difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

// Append unique data from "from" to "to" till capacity is reached
func AppendData(from []string, to []string, capacity int) []string {
	for i := 0; i < capacity; i++ {
		index := rand.Intn(len(from))
		to = append(to, from[index])
		from = removeElemFromSlice(from, index)
	}
	return to
}

// Remove element from slice at index i
func removeElemFromSlice(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// Generate an array of size n of random IPs
func generateRandomIps(n int) []string {
	seedIp := generateRandomIp()
	randIps := make([]string, 0, n)
	for i := 0; i < n; i++ {
		randIps = append(randIps, seedIp)
		seedIp = getNextIpAddr(seedIp)
	}
	return randIps
}

// Generate a random IP address
func generateRandomIp() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

// Get the next IP based on the seed input
func getNextIpAddr(seedIp string) string {
	ip := net.ParseIP(seedIp)
	ip = ip.To4()
	if ip == nil {
		log.Fatal("non ipv4 address")
	}
	ip[3]++
	return ip.String()
}

// Get device info for a user based on the registered devices for that user.
func GetUserDeviceMap(user Users) map[int]ManagedClientsInfo {
	managedClientMap := make(map[int]ManagedClientsInfo)
	i := 0
	if user.Devices.IPhone != "" {
		managedClientMap[i] = ManagedClientsInfo{
			Mac: user.Devices.IPhone,
			DeviceInfo: GhostManagedClientDeviceInfo{
				Platform: GhostPlatform{
					Name:    "iOS",
					Version: "15.2.1",
				},
				Os: GhostOs{
					Name: "iOS",
				},
				Bot: rand.Intn(2) == 0,
				Dist: GhostPlatform{
					Name:    "iphone",
					Version: "15.2.1",
				},
			},
		}
		i++
	}

	if user.Devices.IPad != "" {
		managedClientMap[i] = ManagedClientsInfo{
			Mac: user.Devices.IPad,
			DeviceInfo: GhostManagedClientDeviceInfo{
				Platform: GhostPlatform{
					Name:    "iOS",
					Version: "15.2.0",
				},
				Os: GhostOs{
					Name: "iOS",
				},
				Bot: rand.Intn(2) == 0,
				Dist: GhostPlatform{
					Name:    "iPad",
					Version: "15.2.0",
				},
			},
		}
		i++
	}

	if user.Devices.IPadPro != "" {
		managedClientMap[i] = ManagedClientsInfo{
			Mac: user.Devices.IPadPro,
			DeviceInfo: GhostManagedClientDeviceInfo{
				Platform: GhostPlatform{
					Name:    "iOS",
					Version: "15.2.1",
				},
				Os: GhostOs{
					Name: "iOS",
				},
				Bot: rand.Intn(2) == 0,
				Dist: GhostPlatform{
					Name:    "iPad Pro",
					Version: "15.2.1",
				},
			},
		}
	}
	return managedClientMap
}
