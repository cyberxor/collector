package lib
import (
	"fmt"
	"log"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/shirou/gopsutil/v3/host"
)

// struct to hold host information
type HostSchema struct {
	OS string `json:"os"`
	Platform string `json:"platform"`
	PlatformFamily string `json:"platform_family"`
	PlatformVersion string `json:"platform_version"`
	KernelVersion string `json:"kernel_version"`
	KernelArch string `json:"kernel_arch"`
	Hostname string `json:"hostname"`
	VirtualizationSystem string `json:"virtualization_system"`
	VirtualizationRole string `json:"virtualization_role"`
}

// GetHostInfo function to get host information and post it to the server
func GetHostInfo(){
	info,err := host.Info()
	if err != nil {
		log.Fatalf("Error getting host info: %v", err)
	}
	host := HostSchema{
		OS: info.OS,
		Platform: info.Platform,
		PlatformFamily: info.PlatformFamily,
		PlatformVersion: info.PlatformVersion,
		KernelVersion: info.KernelVersion,
		KernelArch: info.KernelArch,
		Hostname: info.Hostname,
		VirtualizationSystem: info.VirtualizationSystem,
		VirtualizationRole: info.VirtualizationRole,
	}
	hostJson, err := json.Marshal(host)
	if err != nil {
		log.Fatalf("Error marshalling host info: %v", err)
	}
	fmt.Println(string(hostJson))
	client := resty.New()

	resp, err := client.R().
	SetHeader("Content-Type", "application/json").
	SetBody(hostJson).
	Post("http://localhost:8080/host")

	if err != nil {
		log.Fatalf("Error posting host info: %v", err)
	}
	fmt.Println(resp)
}

