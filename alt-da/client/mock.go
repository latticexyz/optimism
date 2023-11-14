package client

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/log"
)

type MockClient struct {
	store ethdb.KeyValueStore
	log   log.Logger
}

func NewMockClient(log log.Logger) *MockClient {
	return &MockClient{
		store: memorydb.New(),
		log:   log,
	}
}

func (c *MockClient) GetPreImage(ctx context.Context, key []byte) ([]byte, error) {
	bytes, err := c.store.Get(key)
	if err != nil {
		return nil, ErrNotFound
	}
	return bytes, nil
}

func (c *MockClient) SetPreImage(ctx context.Context, data []byte) ([]byte, error) {
	key := crypto.Keccak256(data)
	c.store.Put(key, data)

	return key, nil
}

func (c *MockClient) DeleteData(key []byte) {
	c.store.Delete(key)
}

type DAErrFaker struct {
	Client *MockClient

	getPreImageErr error
	setPreImageErr error
}

func (f *DAErrFaker) GetPreImage(ctx context.Context, key []byte) ([]byte, error) {
	if err := f.getPreImageErr; err != nil {
		f.getPreImageErr = nil
		return nil, err
	}
	return f.Client.GetPreImage(ctx, key)
}

func (f *DAErrFaker) SetPreImage(ctx context.Context, data []byte) ([]byte, error) {
	if err := f.setPreImageErr; err != nil {
		f.setPreImageErr = nil
		return nil, err
	}
	return f.Client.SetPreImage(ctx, data)
}

func (f *DAErrFaker) ActGetPreImageFail() {
	f.getPreImageErr = errors.New("get preimage failed")
}

func (f *DAErrFaker) ActSetPreImageFail() {
	f.setPreImageErr = errors.New("set preimage failed")
}
