# kongconfig
Generate / Import  Kong api configs


## Download
`go get -u github.com/nfons/kongconfig`


## Usage
`$ kongconfig [subcommand] [OPTIONS]`  

### Sub Commands
`import` imports kong state (routes / services) to local files

`routes` Exports Routes to a local file: routes.yaml

`services` Exports Services to a local file: service.yaml


### Options:
* `-f, --file` : "File name to use during Export

* `-a, --api` : URL for the kong api server


### Sample Config File:
Kongconfig Yaml files are basically the same as kong json payloads:

for example, a config for a route is:

    - hosts:
      - YOUR_HOST.com
      methods: []
      paths: []
      preserver_host: false
      protocols:
      - http
      service:
        Id: 891a01b4-2fdc-4658-aaf5-f42b924ad2b8
      strip_path: true

### Example Usage

Assuming you have kong running on a k8s cluster:

`$ kubectl port-forward deployment/kong 8001:8001`

This will port-forward kong to run on localhost:8001


`$ kongconfig import --api http://localhost:8001`

This will Create 2 Files from your kong state:

`service.yaml` and `routes.yaml`
