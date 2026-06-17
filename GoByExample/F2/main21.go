package main

import (
	"errors"
	"fmt"
)

type ConnectionState int

const (
	DISCONNECTED ConnectionState = iota
	CONNECTING
	CONNECTED
	FAILED
)

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
	case FAILED:
		return "Failed"
	}
	return "Unknown"
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

func (c *Connection) String() string {
	return "Connection is " + c.state.String()
}

func main() {
	c := &Connection{state: DISCONNECTED}

	fmt.Println(c.Transition(CONNECTING))   // nil
	fmt.Println(c.Transition(CONNECTED))    // nil
	fmt.Println(c.Transition(DISCONNECTED)) // error
}
