package ftdClient

import "fmt"

func (c *Client) GetProtocolObject(ID string) (*ProtocolObject, error) {

	var protocolObject ProtocolObject
	err := doFTDRequest(&protocolObject, fmt.Sprintf("object/protocols/%s", ID), "GET", c)

	return &protocolObject, err
}

func (c *Client) CreateProtocolObject(protocolObject ProtocolObject) (*ProtocolObject, error) {
	err := doFTDRequest(&protocolObject, "object/protocols", "POST", c)
	return &protocolObject, err
}

func (c *Client) UpdateProtocolObject(protocolObject ProtocolObject) (*ProtocolObject, error) {

	err := doFTDRequest(&protocolObject, fmt.Sprintf("object/protocols/%s", protocolObject.ID), "PUT", c)
	return &protocolObject, err
}

func (c *Client) DeleteProtocolObject(protocolObject ProtocolObject) error {

	err := doFTDRequest(&protocolObject, fmt.Sprintf("object/protocols/%s", protocolObject.ID), "DELETE", c)
	return err
}
