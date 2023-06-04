package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"horus-watcher/storage"
	"horus-watcher/types"
	"net/http"
	"os/exec"
	"regexp"
)

func ManageServices() {
	data, err := GetServicesStatus()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Services", data)
	sub := storage.GetSubscribedUrls()
	if sub != nil {
		for _, url := range sub {
			fmt.Println("URL", url)
			PostServicesStatus(url, data)
		}
	}
	f := storage.GetFlaggedServices()
	if f != nil {
		for _, service := range f {
			fmt.Println("Service Flagg", service)
		}
	}
}

func GetServicesStatus() (types.Data, error) {
	sNames := storage.GetServicesNames()
	var d types.Data
	for _, sName := range sNames {

		cmd := exec.Command("sc", "query", sName)

		output, err := cmd.Output()
		if err != nil {
			return d, err
		}

		// fmt.Println(string(output))
		serviceInfo := string(output)
		re := regexp.MustCompile(`ESTADO\s+:\s+(\d+)\s+([A-Z_]+)`)
		match := re.FindStringSubmatch(serviceInfo)

		if len(match) > 2 {
			estadoValue := match[2]
			fmt.Printf("ESTADO da aplicação %s: %s\n", sName, estadoValue)
			d.Data = append(d.Data, types.Service{Name: sName, Status: estadoValue})
		} else {
			fmt.Printf("Could not find ESTADO field for %s ", sName)
		}
	}
	fmt.Println("\n--------------------------------------------------")
	return d, nil
}

func PostServicesStatus(url string, data types.Data) {
	body, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("Status code 200")
	} else {
		fmt.Println("Status code", resp.StatusCode)
	}
}
