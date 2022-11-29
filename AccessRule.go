package ftdClient

import "fmt"

func (c *Client) GetAccessRule(ID string) (*AccessRule, error) {
	ar := AccessRule{}
	ar.Type = "accessrule"
	err := doFTDRequest(&ar, fmt.Sprintf("policy/accesspolicies/%s", ID), "GET", c)
	return &ar, err
}

func (c *Client) CreateAccessRule(ar AccessRule) (*AccessRule, error) {
	ar.Type = "accessrule"
	err := doFTDRequest(&ar, "policy/accesspolicies", "POST", c)
	return &ar, err
}

func (c *Client) UpdateAccessRule(ar AccessRule) (*AccessRule, error) {
	ar.Type = "accessrule"
	err := doFTDRequest(&ar, fmt.Sprintf("policy/accesspolicies/%s", ar.ID), "PUT", c)
	return &ar, err
}

func (c *Client) DeleteAccessRule(ar AccessRule) (*AccessRule, error) {
	ar.Type = "accessrule"
	err := doFTDRequest(&ar, fmt.Sprintf("policy/accesspolicies/%s", ar.ID), "DELETE", c)
	return &ar, err
}
