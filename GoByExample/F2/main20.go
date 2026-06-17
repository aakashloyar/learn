package main 
import (
	"fmt"
	"errors"
)

type ConnectionState int


const (
	DISCONNECTED ConnectionState = iota
	CONNECTING
	CONNECTED
	FAILED

)

type Conn interface {
	Transition(to ConnectionState) error
}


type Connection struct {
    state ConnectionState
}

var transitions = map[ConnectionState]map[ConnectionState]bool{
    DISCONNECTED: {
        CONNECTING: true,
    },
    CONNECTING: {
        CONNECTED: true,
        FAILED:    true,
    },
    FAILED: {
        CONNECTING: true,
    },
}

func (s ConnectionState) String() string {
    switch s {
    case DISCONNECTED:
        return "Disconnected"
    case CONNECTING:
        return "Connecting"
    case CONNECTED:
        return "Connected"
    default:
        return "Failed"
    }
}

func (c *Connection) Transition(to ConnectionState) error {
	allowed, exists := transitions[c.state]
	if !exists {
		return errors.New("no transitions allowed from " + c.state.String())
	}

	if !allowed[to] {
		return errors.New(
			"invalid transition from " +
				c.state.String() + " to " + to.String(),
		)
	}
	c.state = to
	return nil
}
func (c *Connection) String() string{
	return "Connection is "+c.state.String();
}

func main() {
	var c Connection = Connection{state: DISCONNECTED}
	fmt.Println(c.Transition(CONNECTING))
	fmt.Println(c.Transition(CONNECTED))
	fmt.Println(c.Transition(DISCONNECTED)) // must error
}