package testing

import (
	"github.com/Microsoft/hcsshim/hcn"
)

const (
	guid = "123ABC"
)

type HCN interface {
	GetNetworkByName(networkName string) (*hcn.HostComputeNetwork, error)
	ListEndpointsOfNetwork(networkId string) ([]hcn.HostComputeEndpoint, error)
	GetEndpointByID(endpointId string) (*hcn.HostComputeEndpoint, error)
	ListEndpoints() ([]hcn.HostComputeEndpoint, error)
	GetEndpointByName(endpointName string) (*hcn.HostComputeEndpoint, error)
	ListLoadBalancers() ([]hcn.HostComputeLoadBalancer, error)
	GetLoadBalancerByID(loadBalancerId string) (*hcn.HostComputeLoadBalancer, error)
	CreateEndpoint(endpoint *hcn.HostComputeEndpoint) (*hcn.HostComputeEndpoint, error)
	CreateLoadBalancer(loadbalancer *hcn.HostComputeLoadBalancer) (*hcn.HostComputeLoadBalancer, error)
}

type FakeHCN struct {
	endpoints     []*hcn.HostComputeEndpoint
	loadbalancers []*hcn.HostComputeLoadBalancer
}

func NewFakeHCN() *FakeHCN {
	return &FakeHCN{}
}

func (HCN FakeHCN) GetNetworkByName(networkName string) (*hcn.HostComputeNetwork, error) {
	return &hcn.HostComputeNetwork{
		Id:   guid,
		Name: networkName,
		Type: "overlay",
		MacPool: hcn.MacPool{
			Ranges: []hcn.MacRange{
				{
					StartMacAddress: "00-15-5D-52-C0-00",
					EndMacAddress:   "00-15-5D-52-CF-FF",
				},
			},
		},
		Ipams: []hcn.Ipam{
			{
				Type: "Static",
				Subnets: []hcn.Subnet{
					{
						IpAddressPrefix: "192.168.1.0/24",
						Routes: []hcn.Route{
							{
								NextHop:           "192.168.1.1",
								DestinationPrefix: "0.0.0.0/0",
							},
						},
					},
				},
			},
		},
		SchemaVersion: hcn.SchemaVersion{
			Major: 2,
			Minor: 0,
		},
	}, nil
}

func (HCN FakeHCN) ListEndpointsOfNetwork(networkId string) ([]hcn.HostComputeEndpoint, error) {
	var endpoints []hcn.HostComputeEndpoint
	for _, ep := range HCN.endpoints {
		if ep.HostComputeNetwork == networkId {
			endpoints = append(endpoints, *ep)
		}
	}
	return endpoints, nil
}

func (HCN FakeHCN) GetEndpointByID(endpointId string) (*hcn.HostComputeEndpoint, error) {
	endpoint := &hcn.HostComputeEndpoint{}
	for _, ep := range HCN.endpoints {
		if ep.Id == endpointId {
			endpoint.Id = endpointId
			endpoint.Name = ep.Name
			endpoint.HostComputeNetwork = ep.HostComputeNetwork
			endpoint.Health = ep.Health
			endpoint.IpConfigurations = ep.IpConfigurations
		}
	}
	return endpoint, nil
}

func (HCN FakeHCN) ListEndpoints() ([]hcn.HostComputeEndpoint, error) {

	var endpoints []hcn.HostComputeEndpoint
	for _, ep := range HCN.endpoints {
		endpoints = append(endpoints, *ep)
	}
	return endpoints, nil
}

func (HCN FakeHCN) GetEndpointByName(endpointName string) (*hcn.HostComputeEndpoint, error) {
	endpoint := &hcn.HostComputeEndpoint{}
	for _, ep := range HCN.endpoints {
		if ep.Name == endpointName {
			endpoint.Id = ep.Id
			endpoint.Name = endpointName
			endpoint.HostComputeNetwork = ep.HostComputeNetwork
			endpoint.Health = ep.Health
			endpoint.IpConfigurations = ep.IpConfigurations
		}
	}
	return endpoint, nil
}

func (HCN FakeHCN) ListLoadBalancers() ([]hcn.HostComputeLoadBalancer, error) {
	var loadbalancers []hcn.HostComputeLoadBalancer
	for _, lb := range HCN.loadbalancers {
		loadbalancers = append(loadbalancers, *lb)
	}
	return loadbalancers, nil
}

func (HCN FakeHCN) GetLoadBalancerByID(loadBalancerId string) (*hcn.HostComputeLoadBalancer, error) {
	loadbalancer := &hcn.HostComputeLoadBalancer{}
	for _, lb := range HCN.loadbalancers {
		if lb.Id == loadBalancerId {
			loadbalancer.Id = loadBalancerId
			loadbalancer.Flags = lb.Flags
			loadbalancer.HostComputeEndpoints = lb.HostComputeEndpoints
			loadbalancer.SourceVIP = lb.SourceVIP
		}
	}
	return loadbalancer, nil
}

func (HCN FakeHCN) CreateEndpoint(endpoint *hcn.HostComputeEndpoint) (*hcn.HostComputeEndpoint, error) {
	newEndpoint := &hcn.HostComputeEndpoint{
		Id:                 endpoint.Id,
		Name:               endpoint.Name,
		HostComputeNetwork: guid,
		IpConfigurations:   endpoint.IpConfigurations,
		MacAddress:         endpoint.MacAddress,
		Flags:              hcn.EndpointFlagsNone,
		SchemaVersion:      endpoint.SchemaVersion,
		Health:             endpoint.Health,
	}

	HCN.endpoints = append(HCN.endpoints, newEndpoint)

	return newEndpoint, nil
}

func (HCN FakeHCN) CreateRemoteEndpoint(endpoint *hcn.HostComputeEndpoint) (*hcn.HostComputeEndpoint, error) {
	newEndpoint := &hcn.HostComputeEndpoint{
		Id:                 endpoint.Id,
		Name:               endpoint.Name,
		HostComputeNetwork: guid,
		IpConfigurations:   endpoint.IpConfigurations,
		MacAddress:         endpoint.MacAddress,
		Flags:              hcn.EndpointFlagsRemoteEndpoint | endpoint.Flags,
		SchemaVersion:      endpoint.SchemaVersion,
		Health:             endpoint.Health,
	}

	HCN.endpoints = append(HCN.endpoints, newEndpoint)

	return newEndpoint, nil
}

func (HCN FakeHCN) CreateLoadBalancer(loadbalancer *hcn.HostComputeLoadBalancer) (*hcn.HostComputeLoadBalancer, error) {
	newLoadBalancer := &hcn.HostComputeLoadBalancer{
		Id:                   loadbalancer.Id,
		HostComputeEndpoints: loadbalancer.HostComputeEndpoints,
		SourceVIP:            loadbalancer.SourceVIP,
		Flags:                loadbalancer.Flags,
		PortMappings:         loadbalancer.PortMappings,
		FrontendVIPs:         loadbalancer.FrontendVIPs,
		SchemaVersion:        loadbalancer.SchemaVersion,
	}

	HCN.loadbalancers = append(HCN.loadbalancers, newLoadBalancer)

	return newLoadBalancer, nil
}

func (HCN FakeHCN) DeleteLoadBalancer(loadbalancer *hcn.HostComputeLoadBalancer) {
	var i int

	for _, lb := range HCN.loadbalancers {
		i++
		if lb.Id == loadbalancer.Id {
			break
		}
	}

	HCN.loadbalancers = append(HCN.loadbalancers[:i], HCN.loadbalancers[i+1:]...)
}
