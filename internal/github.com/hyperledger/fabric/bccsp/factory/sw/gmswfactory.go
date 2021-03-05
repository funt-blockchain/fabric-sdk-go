package sw

import (
	"github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/bccsp/sw"
	"github.com/pkg/errors"
)

const (
	// GuomiBasedFactoryName is the name of the factory of the software-based BCCSP implementation
	GMSoftwareBasedFactoryName = "GMSW"
)

// GMFactory is the factory of the guomi-based BCCSP.
type GMSWFactory struct{}

// Name returns the name of this factory
func (f *GMSWFactory) Name() string {
	return GMSoftwareBasedFactoryName
}

// Get returns an instance of BCCSP using Opts.
func (f *GMSWFactory) Get(swOpts *SwOpts) (bccsp.BCCSP, error) {
	// Validate arguments
	if swOpts == nil {
		return nil, errors.New("Invalid config. It must not be nil.")
	}

	var ks bccsp.KeyStore
	switch {
	case swOpts.Ephemeral:
		ks = sw.NewDummyKeyStore()
	case swOpts.FileKeystore != nil:
		fks, err := sw.NewFileBasedKeyStore(nil, swOpts.FileKeystore.KeyStorePath, false)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to initialize software key store")
		}
		ks = fks
	case swOpts.InmemKeystore != nil:
		ks = sw.NewInMemoryKeyStore()
	default:
		// Default to ephemeral key store
		ks = sw.NewDummyKeyStore()
	}

	return sw.NewWithParams(swOpts.SecLevel, swOpts.HashFamily, ks)
}

