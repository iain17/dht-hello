package repo

import (
	"errors"

	filestore "gx/ipfs/QmTxUjSZnG7WmebrX2U7furEPNSy33pLgA53PtpJYJSZSn/go-ipfs/filestore"
	keystore "gx/ipfs/QmTxUjSZnG7WmebrX2U7furEPNSy33pLgA53PtpJYJSZSn/go-ipfs/keystore"
	"gx/ipfs/QmTxUjSZnG7WmebrX2U7furEPNSy33pLgA53PtpJYJSZSn/go-ipfs/repo/config"

	ma "gx/ipfs/QmW8s4zTsUoX1Q6CeYxVKPyqSKbF7H1YDUyTostBtZ8DaG/go-multiaddr"
)

var errTODO = errors.New("TODO: mock repo")

// Mock is not thread-safe
type Mock struct {
	C config.Config
	D Datastore
	K keystore.Keystore
}

func (m *Mock) Config() (*config.Config, error) {
	return &m.C, nil // FIXME threadsafety
}

func (m *Mock) SetConfig(updated *config.Config) error {
	m.C = *updated // FIXME threadsafety
	return nil
}

func (m *Mock) SetConfigKey(key string, value interface{}) error {
	return errTODO
}

func (m *Mock) GetConfigKey(key string) (interface{}, error) {
	return nil, errTODO
}

func (m *Mock) Datastore() Datastore { return m.D }

func (m *Mock) GetStorageUsage() (uint64, error) { return 0, nil }

func (m *Mock) Close() error { return errTODO }

func (m *Mock) SetAPIAddr(addr ma.Multiaddr) error { return errTODO }

func (m *Mock) Keystore() keystore.Keystore { return nil }

func (m *Mock) SwarmKey() ([]byte, error) {
	return nil, nil
}

func (m *Mock) FileManager() *filestore.FileManager { return nil }
