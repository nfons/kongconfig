package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
	"github.com/urfave/cli"
)

var file string
var api string

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

	//Get Routes and save to route.yaml
	routes := []Routes{}
	err = getRoutes(api, &routes)
	if err != nil {
		log.Fatal(err)
		return err
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
