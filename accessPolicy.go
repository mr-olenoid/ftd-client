package ftdClient

import "fmt"

func (c *Client) GetAccessPolicy(ID string) (*AccessPolicy, error) {
	ap := AccessPolicy{}
	ap.Type = "accesspolicy"
	err := doFTDRequest(&ap, fmt.Sprintf("policy/accesspolicies/%s", ID), "GET", c)
	return &ap, err
}

func (c *Client) UpdateAccessPolicy(ap AccessPolicy) (*AccessPolicy, error) {
	ap.Type = "accesspolicy"
	err := doFTDRequest(&ap, fmt.Sprintf("policy/accesspolicies/%s", ap.ID), "PUT", c)
	return &ap, err
}
