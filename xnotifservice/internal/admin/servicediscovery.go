package admin

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	sd "github.com/hashicorp/consul/api"
)

type client struct {
	consul       *sd.Client
	registration *sd.AgentServiceRegistration
}

// ServiceDiscovery is a wrapper around the Consul API.
type ServiceDiscovery interface {
	// Register a service with the local agent.
	Register() error

	// Deregister a service with the local agent.
	Deregister() error

	// Service
	//Service(service, tag string, passingOnly bool, queryOpts *api.QueryOptions) ([]*api.ServiceEntry, *api.QueryMeta, error)

	// Key Value
	KV() error
}

//NewSDClient ...
func NewSDClient(consulAddress string, consulPort string, advertiseAddress string, advertisePort string, logger *log.Logger) (ServiceDiscovery, error) {

	cClient, err := sd.NewClient(&sd.Config{Address: fmt.Sprintf("%s:%s", consulAddress, consulPort)})
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("error instantiating new consul client")
		return nil, err
	}

	check := sd.AgentServiceCheck{
		HTTP:     "http://" + advertiseAddress + ":" + advertisePort + "/notification/admin/health",
		Interval: "10s",
		Timeout:  "1s",
		Notes:    "Basic health checks",
	}

	port, _ := strconv.Atoi(advertisePort)
	id := uuid.New()

	asr := &sd.AgentServiceRegistration{
		ID:      id.String(), //unique service ID
		Name:    "notification",
		Address: advertiseAddress,
		Port:    port,
		Tags:    []string{"notification", "contextPath=/notification", "swaggerPath=/swagger", "instanceUuid=" + id.String()},
		Check:   &check,
	}

	return &client{consul: cClient, registration: asr}, nil
}

//register service
func (c *client) Register() error {
	return c.consul.Agent().ServiceRegister(c.registration)
}

//Derigster the Client using the registration ID
func (c *client) Deregister() error {
	return c.consul.Agent().ServiceDeregister(c.registration.ID)
}

func (c *client) Service(service, tag string, passingOnly bool, queryOpts *api.QueryOptions) ([]*api.ServiceEntry, *api.QueryMeta, error) {
	return c.consul.Health().Service(service, tag, passingOnly, queryOpts)
}

func (c *client) KV() error {
	return nil
}
