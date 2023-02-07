package ftdClient

import "fmt"

func (c *Client) GetSamlServer(ID string) (*SamlServer, error) {
	samlServ := SamlServer{}
	err := doFTDRequest(&samlServ, fmt.Sprintf("object/samlservers/%s", ID), "GET", c)
	return &samlServ, err
}

func (c *Client) CreateSamlServer(samlServ SamlServer) (*SamlServer, error) {
	err := doFTDRequest(&samlServ, fmt.Sprintf("object/samlservers/%s", samlServ.ID), "POST", c)
	return &samlServ, err
}

func (c *Client) UpdateSamlServer(samlServ SamlServer) (*SamlServer, error) {
	err := doFTDRequest(&samlServ, fmt.Sprintf("object/samlservers/%s", samlServ.ID), "PUT", c)
	return &samlServ, err
}

func (c *Client) DeleteSamlServer(samlServ SamlServer) error {
	err := doFTDRequest(&samlServ, fmt.Sprintf("object/samlservers/%s", samlServ.ID), "DELETE", c)
	return err
}
