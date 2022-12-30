package ftdClient

import "fmt"

func (c *Client) GetTcpUdpPort(ID string, portType string) (*TcpUdpPort, error) {
	var tcpUdpPort TcpUdpPort
	var err error
	switch portType {
	case "tcpportobject":
		err = doFTDRequest(&tcpUdpPort, fmt.Sprintf("object/tcpports/%s", ID), "GET", c)
	case "udpportobject":
		err = doFTDRequest(&tcpUdpPort, fmt.Sprintf("object/udpports/%s", ID), "GET", c)
	default:
		err = fmt.Errorf("expect tcpportobject or udpportobject, got: %s", portType)
	}

	return &tcpUdpPort, err
}

func (c *Client) CreateTcpUdpPort(tcpUdpPort TcpUdpPort) (*TcpUdpPort, error) {
	var err error
	switch tcpUdpPort.Type {
	case "tcpportobject":
		err = doFTDRequest(&tcpUdpPort, "object/tcpports", "POST", c)
	case "udpportobject":
		err = doFTDRequest(&tcpUdpPort, "object/udpports", "POST", c)
	default:
		err = fmt.Errorf("expect tcpportobject or udpportobject, got: %s", tcpUdpPort.Type)
	}

	return &tcpUdpPort, err
}

func (c *Client) UpdateTcpUdpPort(tcpUdpPort TcpUdpPort) (*TcpUdpPort, error) {
	var err error
	switch tcpUdpPort.Type {
	case "tcpportobject":
		err = doFTDRequest(&tcpUdpPort, fmt.Sprintf("object/tcpports/%s", tcpUdpPort.ID), "PUT", c)
	case "udpportobject":
		err = doFTDRequest(&tcpUdpPort, fmt.Sprintf("object/udpports/%s", tcpUdpPort.ID), "PUT", c)
	default:
		err = fmt.Errorf("expect tcpportobject or udpportobject, got: %s", tcpUdpPort.Type)
	}

	return &tcpUdpPort, err
}

func (c *Client) DeleteTcpUdpPort(tcpUdpPort *TcpUdpPort) error {
	var err error
	switch tcpUdpPort.Type {
	case "tcpportobject":
		err = doFTDRequest(&tcpUdpPort, fmt.Sprintf("object/tcpports/%s", tcpUdpPort.ID), "DELETE", c)
	case "udpportobject":
		err = doFTDRequest(&tcpUdpPort, fmt.Sprintf("object/udpports/%s", tcpUdpPort.ID), "DELETE", c)
	default:
		err = fmt.Errorf("expect tcpportobject or udpportobject, got: %s", tcpUdpPort.Type)
	}

	return err
}
