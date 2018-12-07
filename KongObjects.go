package main

type Service struct {
	Host          *string `json:"host"`
	Protocol      *string `json:"protocol"`
	Name          *string `json:"name"`
	Port          int     `json:"port"`
	Path          *string `json:"path"`
	Retries       int     `json:"retries"`
	Id            *string `json:"id"`
	Write_timeout int     `json:"write_timeout"`
	Read_timeout  int     `json:"read_timeout"`
}
type serviceField struct {
	Id string `json:"id"`
}
type Routes struct {
	Hosts         []string     `json:"hosts"`
	Preserve_host bool         `json:"preserve_host"`
	Service       serviceField `json:"service"`
	Paths         []string     `json:"paths"`
	Methods       []string     `json:"methods"`
	Strip_path    bool         `json:"strip_path"`
	Protocols     []string     `json:"protocols"`
	ServiceName   string       `json:"-"` //this is not an actual json field
}

type Certificates struct {
	Cert string `json:"cert"`
	Key  string `json:"key"`
}

type Sni struct {
	Name               string `json:"name"`
	Ssl_certificate_id string `json:"ssl_certificate_id"`
}

//
// Used By End Users for their yaml files
//

type RouteFile struct {
	Routes []Routes `json:"Routes"`
}

type ServiceFile struct {
	Services []Service `json:"Services"`
}
