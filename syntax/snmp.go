package trapd

/*
 * Start a trap receiver locally at port 162, using default community: public
 */

import (
	"fmt"
	"log"
	"net"
	"strings"

	g "github.com/soniah/gosnmp"
)

var MsgChan = make(chan *Alarm, 32)
var site = "0.0.0.0:162"

func UnMarshal(pdu []SnmpPDU, data interface{}) {

}
func myTrapHandler(packet *g.SnmpPacket, addr *net.UDPAddr) {

	var alarm Alarm
	alarm.Source = fmt.Sprintf("%s", addr.IP)

	for _, v := range packet.Variables {
		if strings.Contains(v.Name, "6876.4.3.304") {
			alarm.OldStatus = string(v.Value.([]byte))
		}
		if strings.Contains(v.Name, "6876.4.3.305") {
			alarm.NewStatus = string(v.Value.([]byte))
		}
		if strings.Contains(v.Name, "6876.4.3.306") {
			alarm.Detail = string(v.Value.([]byte))
		}
		if strings.Contains(v.Name, "6876.4.3.307") {
			alarm.Object = string(v.Value.([]byte))
		}
	}

	//fmt.Printf("alarm: %v\n", alarm)
	MsgChan <- &alarm
}

func StartTrapd(logger Logger) {
	tl := g.NewTrapListener()
	tl.OnNewTrap = myTrapHandler
	tl.Params = g.Default
	tl.Params.Logger = logger

	err := tl.Listen(site)
	if err != nil {
		log.Panicf("error in listen: %s", err)
	}
}
