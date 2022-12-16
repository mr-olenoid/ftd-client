package ftdClient

import "fmt"

func (c *Client) GetPortGroup(ID string) (*PortGroup, error) {
	var portGroup PortGroup
	err := doFTDRequest(&portGroup, fmt.Sprintf("object/icmpv4ports/%s", ID), "GET", c)
	return &portGroup, err
}

func (c *Client) CreatePortGroup(portGroup PortGroup) (*PortGroup, error) {
	err := doFTDRequest(&portGroup, "object/icmpv4ports", "POST", c)
	return &portGroup, err
}

func (c *Client) UpdatePortGroup(portGroup PortGroup) (*PortGroup, error) {
	err := doFTDRequest(&portGroup, fmt.Sprintf("object/icmpv4ports/%s", portGroup.ID), "PUT", c)
	return &portGroup, err
}

func (c *Client) DeletePortGroup(portGroup PortGroup) error {
	err := doFTDRequest(&portGroup, fmt.Sprintf("object/icmpv4ports/%s", portGroup.ID), "DELETE", c)
	return err
}
