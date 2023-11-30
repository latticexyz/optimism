package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

var ErrNotFound = errors.New("not found")

type AltDARPC interface {
	GetPreImage(ctx context.Context, key []byte) ([]byte, error)
	SetPreImage(ctx context.Context, data []byte) ([]byte, error)
}

var _ AltDARPC = &Client{}

type MetricsRecorder interface {
	RecordStorageError()
}

// Client communicates with the DA service via JSON RPC.
type Client struct {
	url     string
	log     log.Logger
	metrics MetricsRecorder
}

func New(log log.Logger, metrics MetricsRecorder, url string) *Client {
	return &Client{
		url:     url,
		log:     log,
		metrics: metrics,
	}
}

// GetPreImage returns the input data for the given commitment as well as its
// challenge status. We pass a block number so the DA service knows which l1 block
// the commitment was included in.
func (c *Client) GetPreImage(ctx context.Context, key []byte) ([]byte, error) {
	k := hexutil.Bytes(key)
	resp, err := http.Get(fmt.Sprintf("%s/get/%s", c.url, k))
	if resp.StatusCode == http.StatusNotFound {
		c.metrics.RecordStorageError()
		return nil, ErrNotFound
	}
	if err != nil {
		c.metrics.RecordStorageError()
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.metrics.RecordStorageError()
		return nil, err
	}
	return bytes, nil
}

// SetPreImage sets the input data and returns the keccak256 hash commitment.
func (c *Client) SetPreImage(ctx context.Context, img []byte) ([]byte, error) {
	key := crypto.Keccak256(img)
	k := hexutil.Bytes(key)
	body := bytes.NewReader(img)
	url := fmt.Sprintf("%s/put/%s", c.url, k)
	resp, err := http.Post(url, "application/octet-stream", body)
	if err != nil {
		c.metrics.RecordStorageError()
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		c.metrics.RecordStorageError()
		return nil, fmt.Errorf("failed to store preimage: %s", resp.Status)
	}
	return key, nil
}
