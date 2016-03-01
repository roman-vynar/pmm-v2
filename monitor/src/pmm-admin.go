package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

const (
	WORKDIR = "./prometheus"
	USAGE   = `
	The tool for managing Prometheus endpoints for auto-discovery.

	Usage:

	ppm-admin -add <alias> <address>	- Add host to monitoring.
	ppm-admin -del <alias>			- Remove host from monitoring.
	ppm-admin -ls				- List hosts being monitored.
	`
)

var PORTS = []string{"9100", "9104", "9105", "9106"}

var (
	addHost = flag.Bool("add", false, "Add host to monitoring.")
	delHost = flag.Bool("del", false, "Remove host from monitoring.")
	lsHosts = flag.Bool("ls", false, "List hosts being monitored.")
)

type Host struct {
	Alias   string
	Address string
}

type Endpoint struct {
	Targets []string
	Labels  map[string]string
}

func writeFiles(config []Host) {
	// Write the tool config
	yamlData, _ := yaml.Marshal(&config)
	if err := ioutil.WriteFile(path.Join(WORKDIR, "targets.yml"), yamlData, 0644); err != nil {
		log.Fatal(err)
	}

	// Write Prometheus configs
	for _, port := range PORTS {
		var configTarget []Endpoint
		for _, host := range config {
			target := Endpoint{
				Targets: []string{host.Address + ":" + port},
				Labels:  map[string]string{"alias": host.Alias},
			}
			configTarget = append(configTarget, target)
		}
		yamlData, _ = yaml.Marshal(&configTarget)
		if err := ioutil.WriteFile(path.Join(WORKDIR, "targets_"+port+".yml"), yamlData, 0644); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	flag.Usage = func() {
		fmt.Println(USAGE)
	}
	flag.Parse()

	if _, err := os.Stat(WORKDIR); err != nil {
		log.Fatal(err)
	}

	filename := path.Join(WORKDIR, "targets.yml")
	if _, err := os.Stat(filename); err != nil {
		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}

	yamlData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var config []Host
	if err := yaml.Unmarshal(yamlData, &config); err != nil {
		log.Fatal(err)
	}

	switch {
	case *addHost:
		if len(flag.Arg(0)) == 0 || len(flag.Arg(1)) == 0 {
			flag.Usage()
			log.Fatal("No enough arguments given.")
		}
		config = append(config, Host{Alias: flag.Arg(0), Address: flag.Arg(1)})
		writeFiles(config)
		fmt.Println("Added.")
	case *delHost:
		if len(flag.Arg(0)) == 0 {
			flag.Usage()
			log.Fatal("No enough arguments given.")
		}
		for i, host := range config {
			if host.Alias == flag.Arg(0) {
				config = append(config[:i], config[i+1:]...)
				writeFiles(config)
				fmt.Println("Deleted.")
				return
			}
		}
		fmt.Println("No host matched.")
	case *lsHosts:
		if len(config) == 0 {
			fmt.Println("No hosts configured.")
			return
		}
		fmt.Printf("There are %d host(s) configured.\n\n", len(config))
		for _, host := range config {
			fmt.Printf("%10s | %s\n", host.Alias, host.Address)
		}
	default:
		flag.Usage()
	}
}
