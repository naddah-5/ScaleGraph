package scalegraph


import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

type contact struct {
	nodeIP  [4]byte
	udpPort int
	nodeID  [5]uint32
}

func (c *contact) IP() [4]byte {
	return c.nodeIP
}

func (c *contact) Port() int {
	return c.udpPort
}

func (c *contact) ID() [5]uint32 {
	return c.nodeID
}

func BuildContact(ip [4]byte, port int, id [5]uint32) contact {
	var newContact contact
	newContact = contact{
		nodeIP:  ip,
		udpPort: port,
		nodeID:  id,
	}
	return newContact
}

// Provides a empty contact instance.
func EmptyContact() contact {
	var empty contact = contact{}
	return empty
}

// Generates and returns a new, validated, contact with pseudo-random values.
func NewRandomContact() contact {
	var port int = 80
	var ip [4]byte
	var id [5]uint32 = [5]uint32{rand.Uint32(), rand.Uint32(), rand.Uint32(), rand.Uint32(), rand.Uint32()}

	for i := 0; i < 4; i++ {
		seg, _ := randU32(0, 256)
		ip[i] = byte(seg)
	}
	newContact := BuildContact(ip, port, id)
	return newContact
}



func validateContactInfo(ip [4]byte, port int, id [5]uint32) error {
	var errMsg []error
	err := validateIPStructure(ip)
	if err != nil {
		errMsg = append(errMsg, err)
	}
	err = validateUDPPort(port)
	if err != nil {
		errMsg = append(errMsg, err)
	}
	err = validateNodeID(id)
	if err != nil {
		errMsg = append(errMsg, err)
	}
	if len(errMsg) == 0 {
		return nil
	} else {
		var errString string = ""
		for i := 0; i < len(errMsg); i++ {
			errString += errString + "\n" + errMsg[i].Error()
		}
		err := errors.New(errString)
		return err
	}
}

func validateIPStructure(ip [4]byte) error {
	var errMsg string = ""
	for i := 0; i < len(ip); i++ {
		if ip[i] < 0 || ip[i] > 255 {
			var err string = fmt.Sprintf("ip segment out of bounds, valid for 0 <= segment <= 255, received: %v in address %v\n", ip[i], ip)
			errMsg = errMsg + err
		}
	}
	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}

func validateUDPPort(port int) error {
	if port < 0 || port > 1023 {
		return errors.New("forbidden UDP port submitted, use ports in the range; 0 <= port <= 1023\t received port " + strconv.Itoa(port))
	}
	return nil
}

func validateNodeID(id [5]uint32) error {
	// This will validate id once it is implemented
	// i.e. the signing of it or something like that
	return nil
}
