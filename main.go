package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var file string
var api string

/**
Export Routes to Kong. For Routes, we need to make extra calls to get a list of services each route corresponds to.
*/
func exportRoutesToKong(c *cli.Context) error {
	yamlFile, err := ioutil.ReadFile(file)
	routes := RouteFile{}

	if err != nil {
		log.Fatal(err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, &routes)

	if err != nil {
		log.Fatal(err)
		return err
	}

	//Get Services as Well since we will need them
	services := []Service{}
	serr := getServices(api, &services)
	if serr != nil {
		log.Fatal(serr)
	}

	//create a map of services by name to Id
	serviceMap := make(map[string]string)

	for _, val := range services {
		serviceMap[*val.Name] = *val.Id
	}

	//for each route, we will need to make a api call
	for _, val := range routes.Routes {
		//get the corresponding Id of a Service Name
		serviceId := serviceMap[val.ServiceName]
		val.Service.Id = serviceId
		log.Println(val)
	}

	return nil
}

/*
 Export Services to kong. Since services are the building block, we dont need to do anything extra
*/
func exportServicesToKong(c *cli.Context) error {
	yamlFile, err := ioutil.ReadFile(file)
	services := ServiceFile{}

	if err != nil {
		log.Fatal(err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, &services)

	if err != nil {
		log.Fatal(err)
		return err
	}
	//for each route, we will need to make a api call
	for _, val := range services.Services {
		y, _ := json.Marshal(val)
		log.Println(string(y))
		MakeServices(api, y)
	}

	return nil
}

/*

For the import Logic, We can Import Services as-is from Kong.

But for Routes however, we need to make a intermittent map of Services to Service.ID.

This is becuase when we save the Routes from Kong, It will get a ID of X. But we reimport the services, The IDs will change.

So if old ID X = service A, when we re-import services, service A will have a different ID. Thus importing Routes as is wont work.


*/
func importFromKong(c *cli.Context) error {
	//Get Services and save to service.yaml
	services := []Service{}
	serr := getServices(api, &services)
	if serr != nil {
		log.Fatal(serr)
		return serr
	}
	serviceFile := ServiceFile{}
	serviceFile.Services = services

	marshalYaml, err := yaml.Marshal(serviceFile)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("service.yaml", marshalYaml, 0644)
	log.Println("Kong Services Imported")

	//create a map of Services ID to name
	serviceMap := make(map[string]string)

	for _, val := range serviceFile.Services {
		serviceMap[*val.Id] = *val.Name
	}
	//Get Routes and save to route.yaml
	routes := []Routes{}
	err = getRoutes(api, &routes)
	if err != nil {
		log.Fatal(err)
		return err
	}

	//now update Routes based on Service id name
	for index, val := range routes {
		val.ServiceName = serviceMap[val.Service.Id]
		routes[index] = val
	}
	routefile := RouteFile{}
	routefile.Routes = routes
	marshalYaml, err = yaml.Marshal(routefile)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = ioutil.WriteFile("routes.yaml", marshalYaml, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("Kong Routes Imported")
	return err

}

func main() {
	app := cli.NewApp()

	// #App version
	app.Version = "0.0.1"
	app.Name = "Kong Config - Config Management for Kong"
	app.Usage = "Config management via Yamls"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file, f",
			Value:       "",
			Usage:       "Yaml File to use",
			Destination: &file,
		},
		cli.StringFlag{
			Name:        "api, a",
			Value:       "http://localhost:8001",
			Usage:       "Kong API location",
			Destination: &api,
		},
	}

	//import sub command TODO
	app.Commands = []cli.Command{
		{
			Name:    "import",
			Aliases: []string{"i"},
			Usage:   "Imports a yaml config from kong",
			Action:  importFromKong,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "api, a",
					Value:       "http://localhost:8001",
					Usage:       "Kong API location",
					Destination: &api,
				},
			},
		},
		{
			Name:    "routes",
			Aliases: []string{"r"},
			Usage:   "Creates Routes based on a file",
			Action:  exportRoutesToKong,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "file, f",
					Value:       "",
					Usage:       "Yaml File to use",
					Destination: &file,
				},
				cli.StringFlag{
					Name:        "api, a",
					Value:       "http://localhost:8001",
					Usage:       "Kong API location",
					Destination: &api,
				},
			},
		},
		{
			Name:    "services",
			Aliases: []string{"r"},
			Usage:   "Creates Services based on a file",
			Action:  exportServicesToKong,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "file, f",
					Value:       "",
					Usage:       "Yaml File to use",
					Destination: &file,
				},
				cli.StringFlag{
					Name:        "api, a",
					Value:       "http://localhost:8001",
					Usage:       "Kong API location",
					Destination: &api,
				},
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		yamlFile, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		service := ServiceFile{}
		err = yaml.Unmarshal(yamlFile, &service)
		if err != nil {
			return err
		}
		log.Println(service)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
