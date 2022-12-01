package ftdClient

import "fmt"

func (c *Client) GetSecurityZone(ID string) (*SecurityZone, error) {
	sz := SecurityZone{}
	err := doFTDRequest(&sz, fmt.Sprintf("object/securityzones/%s", ID), "GET", c)
	return &sz, err
}

func (c *Client) CreateSecurityZone(sz SecurityZone) (*SecurityZone, error) {
	err := doFTDRequest(&sz, "object/securityzones", "POST", c)
	return &sz, err
}

func (c *Client) UpdateSecurityZone(sz SecurityZone) (*SecurityZone, error) {
	err := doFTDRequest(&sz, fmt.Sprintf("object/securityzones/%s", sz.ID), "PUT", c)
	return &sz, err
}

func (c *Client) DeleteSecurityZone(sz SecurityZone) error {
	err := doFTDRequest(&sz, fmt.Sprintf("object/securityzones/%s", sz.ID), "DELETE", c)
	return err
}
