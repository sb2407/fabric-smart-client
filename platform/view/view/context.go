/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package view

import "context"

// RunViewOptions models the options to run a view
type RunViewOptions struct {
	Session     Session
	AsInitiator bool
}

// CompileRunViewOptions compiles a set of RunViewOption to a RunViewOptions
func CompileRunViewOptions(opts ...RunViewOption) (*RunViewOptions, error) {
	txOptions := &RunViewOptions{}
	for _, opt := range opts {
		if err := opt(txOptions); err != nil {
			return nil, err
		}
	}
	return txOptions, nil
}

// RunViewOption models a function that set options to run a view
type RunViewOption func(*RunViewOptions) error

// AsResponder sets the context's session to the passed session
func AsResponder(session Session) RunViewOption {
	return func(o *RunViewOptions) error {
		o.Session = session
		return nil
	}
}

// AsInitiator tells the context to initialize the initiator to the executing view
func AsInitiator() RunViewOption {
	return func(o *RunViewOptions) error {
		o.AsInitiator = true
		return nil
	}
}

// Context gives a view information about the environment in which it is in execution
type Context interface {
	// GetService returns an instance of the given type
	GetService(v interface{}) (interface{}, error)

	// ID returns the identifier of this context
	ID() string

	// RunView runs the passed view on input this context
	RunView(view View, opts ...RunViewOption) (interface{}, error)

	// Me returns the identity bound to this context
	Me() Identity

	// IsMe returns true if the passed identity is an alias
	// of the identity bound to this context, false otherwise
	IsMe(id Identity) bool

	// Initiator returns the View that initiate a call
	Initiator() View

	// GetSession returns a session to the passed remote
	// party for the given view caller.
	// Cashing may be be used.
	GetSession(caller View, party Identity) (Session, error)

	// GetSessionByID returns a session to the passed remote party and id.
	// Cashing may be used.
	GetSessionByID(id string, party Identity) (Session, error)

	// Session returns the session created to respond to a
	// remote party, nil if the context was created
	// not to respond to a remote call
	Session() Session

	// Context return the associated context.Context
	Context() context.Context

	// OnError appends to passed callback function to the list of functions called when
	// the current execution return an error or panic.
	// This is useful to release resources.
	OnError(callback func())
}
