package ftdClient

import (
	"fmt"
)

func (c *Client) GetNetworkInterface(ID string) (*NetworkInterface, error) {
	i := NetworkInterface{}
	err := doFTDRequest(&i, fmt.Sprintf("/devices/default/interfaces/%s", ID), "GET", c)
	return &i, err
}

func (c *Client) GetNetworkInterfaceSecurityZone(interfaceID string) (*ReferenceModel, error) {
	secz := ReferenceModel{}
	securityZones := Items[SecurityZone]{}
	doFTDRequest(&securityZones, "/object/securityzones", "GET", c)
	for _, sz := range securityZones.Items {
		for _, iface := range sz.Interfaces {
			if iface.ID == interfaceID {
				secz.ID = sz.ID
				secz.Name = sz.Name
				secz.Type = sz.Type
				secz.Version = sz.Version
			}
		}
	}
	return &secz, nil
}

// CreateNetworkInterface - return interface and security zone assosiated with it
func (c *Client) CreateNetworkInterface(name string) (*NetworkInterface, error) {
	i := Items[NetworkInterface]{}
	//sz := &ReferenceModel{}
	err := doFTDRequest(&i, fmt.Sprintf("/devices/default/interfaces?filter=name:%s", name), "GET", c)
	if err != nil {
		return nil, err
	}

	iface := i.Items[0]
	/*
		if securityZone.ID == "" {
			sz, err = c.GetNetworkInterfaceSecurityZone(iface.ID)
			if err != nil {
				return nil, nil, err
			}
			fmt.Println(securityZone)
		}
	*/
	return &iface, err
}

func (c *Client) UpdateNetworkInterface(n NetworkInterface) (*NetworkInterface, error) {

	err := doFTDRequest(&n, fmt.Sprintf("/devices/default/interfaces/%s", n.ID), "PUT", c)

	return &n, err
}
