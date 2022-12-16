package ftdClient

import "fmt"

func (c *Client) GetTcpUdpPort(ID string, portType string) (*TcpUdpPort, error) {
	var tcpUdpPort TcpUdpPort
	var err error
	switch portType {
	case "TCP":
		err = doFTDRequest(&tcpUdpPort, fmt.Sprintf("object/tcpports/%s", ID), "GET", c)
	case "UDP":
		err = doFTDRequest(&tcpUdpPort, fmt.Sprintf("object/udpports/%s", ID), "GET", c)
	default:
		err = fmt.Errorf("expect TCP or UDP, got: %s", portType)
	}

	return &tcpUdpPort, err
}

func (c *Client) CreateTcpUdpPort(tcpUdpPort TcpUdpPort, portType string) (*TcpUdpPort, error) {
	var err error
	switch portType {
	case "TCP":
		err = doFTDRequest(&tcpUdpPort, "object/tcpports", "POST", c)
	case "UDP":
		err = doFTDRequest(&tcpUdpPort, "object/udpports", "POST", c)
	default:
		err = fmt.Errorf("expect TCP or UDP, got: %s", portType)
	}

	return &tcpUdpPort, err
}

func (c *Client) UpdateTcpUdpPort(tcpUpdPort TcpUdpPort, portType string) (*TcpUdpPort, error) {
	var tcpUdpPort TcpUdpPort
	var err error
	switch portType {
	case "TCP":
		err = doFTDRequest(&tcpUdpPort, fmt.Sprintf("object/tcpports/%s", tcpUdpPort.ID), "PUT", c)
	case "UDP":
		err = doFTDRequest(&tcpUdpPort, fmt.Sprintf("object/udpports/%s", tcpUdpPort.ID), "PUT", c)
	default:
		err = fmt.Errorf("expect TCP or UDP, got: %s", portType)
	}

	return &tcpUdpPort, err
}

func (c *Client) DeleteTcpUdpPort(tcpUdpPort *TcpUdpPort, portType string) error {
	var err error
	switch portType {
	case "TCP":
		err = doFTDRequest(&tcpUdpPort, fmt.Sprintf("object/tcpports/%s", tcpUdpPort.ID), "DELETE", c)
	case "UDP":
		err = doFTDRequest(&tcpUdpPort, fmt.Sprintf("object/udpports/%s", tcpUdpPort.ID), "DELETE", c)
	default:
		err = fmt.Errorf("expect TCP or UDP, got: %s", portType)
	}

	return err
}