// Package charnet implements spec.Network to provide functionality to analyse,
// connect and gather data with respect to the given input.
package charnet

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	ObjectTypeCharNet spec.ObjectType = "char-net"
)

type Config struct {
	FactoryClient spec.Factory
	Log           spec.Log
	Storage       spec.Storage

	EvalNet  spec.Network
	ExecNet  spec.Network
	PatNet   spec.Network
	PredNet  spec.Network
	StratNet spec.Network
}

func DefaultConfig() Config {
	newConfig := Config{
		FactoryClient: factoryclient.NewFactory(factoryclient.DefaultConfig()),
		Log:           log.NewLog(log.DefaultConfig()),
		Storage:       memorystorage.NewMemoryStorage(memorystorage.DefaultConfig()),

		EvalNet:  nil,
		ExecNet:  nil,
		PatNet:   nil,
		PredNet:  nil,
		StratNet: nil,
	}

	return newConfig
}

// NewCharNet returns a new configured character network.
func NewCharNet(config Config) (spec.Network, error) {
	newNet := &charNet{
		Booted: false,
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeCharNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type charNet struct {
	Config

	Booted bool
	ID     spec.ObjectID
	Mutex  sync.Mutex
	Type   spec.ObjectType
}

func (cn *charNet) Boot() {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()

	if cn.Booted {
		return
	}
	cn.Booted = true

	cn.Log.WithTags(spec.Tags{L: "D", O: cn, T: nil, V: 13}, "call Boot")
}

func (cn *charNet) Shutdown() {
	cn.Log.WithTags(spec.Tags{L: "D", O: cn, T: nil, V: 13}, "call Shutdown")
}

func (cn *charNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	cn.Log.WithTags(spec.Tags{L: "D", O: cn, T: nil, V: 13}, "call Trigger")

	// Dynamically walk impulse through the other networks.
	var err error
	for {
		imp, err = cn.StratNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = cn.PredNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = cn.ExecNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = cn.EvalNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}

		break
	}

	// Note that the impulse returned here is not actually the same as received
	// at the beginning of the call, but was manipulated during its walk through
	// the networks.
	return imp, nil
}
