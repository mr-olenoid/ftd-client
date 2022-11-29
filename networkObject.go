package ftdClient

import "fmt"

func (c *Client) GetNetworkObject(ID string) (*NetworkObject, error) {
	no := NetworkObject{}
	err := doFTDRequest(&no, fmt.Sprintf("object/networks/%s", ID), "GET", c)
	return &no, err
}

func (c *Client) CreateNetworkObject(no NetworkObject) (*NetworkObject, error) {
	err := doFTDRequest(&no, "object/networks", "POST", c)
	return &no, err
}

func (c *Client) UpdateNetworkObject(no NetworkObject) (*NetworkObject, error) {
	err := doFTDRequest(&no, fmt.Sprintf("object/networks/%s", no.ID), "PUT", c)
	return &no, err
}

func (c *Client) DeleteNetworkObject(no NetworkObject) error {
	err := doFTDRequest(&no, fmt.Sprintf("object/networks/%s", no.ID), "DELETE", c)
	return err
}
