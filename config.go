package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//SERVICE define which IP version we are usig (currently only dhcp4 supported)
var SERVICE []string = []string{"dhcp4"}

type configReq struct {
	Command string   `json:"command"`
	Service []string `json:"service"`
}
type configSet struct {
	Command   string    `json:"command"`
	Service   []string  `json:"service"`
	Arguments Arguments `json:"arguments"`
}

//NestedElem elements in root
type NestedElem struct {
	Arguments Arguments `json:"arguments"`
	Result    int       `json:"result"`
}

//Arguments structure
type Arguments struct {
	Dhcp4   Dhcp4       `json:"Dhcp4"`
	Logging interface{} `json:"Logging"` // If you dont need to modify this area just pass as interface{}
}

//Dhcp4 structure
type Dhcp4 struct {
	Subnet4                 []Subnet4   `json:"subnet4"`
	InterfacesConfig        interface{} `json:"interfaces-config"`
	ControlSocket           interface{} `json:"control-socket"`
	LeaseDatabase           interface{} `json:"lease-database"`
	ExpiredLeasesProcessing interface{} `json:"expired-leases-processing"`
	OptionData              interface{} `json:"option-data"`
}

//Subnet4 structure
type Subnet4 struct {
	Subnet        interface{}    `json:"subnet"`
	Pools         interface{}    `json:"pools"`
	OptionData    interface{}    `json:"option-data"`
	RenewTimer    int            `json:"renew-timer"`
	RebindTimer   int            `json:"rebind-timer"`
	ValidLifetime int            `json:"valid-lifetime"`
	Reservations  []Reservations `json:"reservations"`
}

//Reservations structure
type Reservations struct {
	Hostname  string `json:"hostname"`
	Hwaddress string `json:"hw-address"`
	Ipaddress string `json:"ip-address"`
}

// Config struct for provider
type Config struct {
	Server     string
	Username   string
	Password   string
	Configfile string
}

// Client for connections
type Client struct {
	Config        *Config
	currentConfig []NestedElem
	httpClient    *http.Client
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

//Client In go (c *Config) means this is method assosieted to struct Config
func (c *Config) Client() (*Client, error) {

	log.Println("[INFO] Configuring kea-api client")
	htclient := &http.Client{}
	var data []byte
	services := []string{"dhcp4"}
	jsonStr := configReq{Command: "config-get", Service: services}
	b, err := json.Marshal(jsonStr)
	check(err)
	req, err := http.NewRequest("POST", c.Server, bytes.NewBuffer(b))
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Content-Type", "application/json")
	resp, err := htclient.Do(req)
	if err != nil {
		log.Fatalf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ = ioutil.ReadAll(resp.Body)
		log.Printf("[INFO] %s\n", string(data))
	}
	log.Println("[INFO] Terminating GET_CONF application...")
	var m []NestedElem
	err = json.Unmarshal(data, &m)

	client := &Client{
		Config:        c,
		currentConfig: m,
		httpClient:    htclient,
	}

	return client, nil
}

//NewLease method to create new lease fot kea-dhcpd4
func (c *Client) NewLease(r Reservations) error {
	var data []byte
	c.currentConfig[0].Arguments.Dhcp4.Subnet4[0].Reservations = append(c.currentConfig[0].Arguments.Dhcp4.Subnet4[0].Reservations, r)
	jsonSet := configSet{Command: "config-set", Service: SERVICE, Arguments: c.currentConfig[0].Arguments}
	enc, err := json.Marshal(jsonSet)
	check(err)
	req, err := http.NewRequest("POST", c.Config.Server, bytes.NewBuffer(enc))
	req.SetBasicAuth(c.Config.Username, c.Config.Password)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("[Error] The HTTP request failed with error %s\n", err)
	} else {
		data, _ = ioutil.ReadAll(resp.Body)
		log.Printf("[INFO] New lease shoud be added.. %s\n", string(data))
	}

	return nil
}

//SaveConfig method to save a config file for kea-dhcpd4
func (c *Client) SaveConfig(r Reservations) error {
	var data []byte
	jsonWrite := "{ \"command\": \"config-write\", \"service\": [ \"dhcp4\" ], \"arguments\":{\"filename\":\"" + c.Config.Configfile + "\"} }"
	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(jsonWrite)
	req, err := http.NewRequest("POST", c.Config.Server, strings.NewReader(jsonWrite))
	check(err)
	req.SetBasicAuth(c.Config.Username, c.Config.Password)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("[ERROR] The HTTP request failed with error %s\n", err)
	} else {
		data, _ = ioutil.ReadAll(resp.Body)
		log.Printf("[INFO] Configuration file should be saved.. %s\n", string(data))
	}
	return nil
}
