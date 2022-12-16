package ftdClient

import "fmt"

func (c *Client) GetIcmp4port(ID string) (*Icmpv4port, error) {

	var icmpv4port Icmpv4port
	err := doFTDRequest(&icmpv4port, fmt.Sprintf("object/icmpv4ports/%s", ID), "GET", c)

	return &icmpv4port, err
}

func (c *Client) CreateIcmp4port(icmpv4port Icmpv4port) (*Icmpv4port, error) {
	err := doFTDRequest(&icmpv4port, "object/icmpv4ports", "POST", c)
	return &icmpv4port, err
}

func (c *Client) UpdateIcmp4port(icmpv4port Icmpv4port) (*Icmpv4port, error) {

	err := doFTDRequest(&icmpv4port, fmt.Sprintf("object/icmpv4ports/%s", icmpv4port.ID), "PUT", c)
	return &icmpv4port, err
}

func (c *Client) DeleteIcmp4port(icmpv4port Icmpv4port) error {

	err := doFTDRequest(&icmpv4port, fmt.Sprintf("object/icmpv4ports/%s", icmpv4port.ID), "DELETE", c)
	return err
}
