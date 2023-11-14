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

// Client communicates with the DA service via JSON RPC.
type Client struct {
	url string
	log log.Logger
}

func New(log log.Logger, url string) *Client {
	return &Client{
		url: url,
		log: log,
	}
}

// GetPreImage returns the input data for the given commitment as well as its
// challenge status. We pass a block number so the DA service knows which l1 block
// the commitment was included in.
func (c *Client) GetPreImage(ctx context.Context, key []byte) ([]byte, error) {
	k := hexutil.Bytes(key)
	resp, err := http.Get(fmt.Sprintf("%s/%s", c.url, k))
	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// SetPreImage sets the input data and returns the keccak256 hash commitment.
func (c *Client) SetPreImage(ctx context.Context, img []byte) ([]byte, error) {
	key := crypto.Keccak256(img)
	k := hexutil.Bytes(key)
	body := bytes.NewReader(img)
	url := fmt.Sprintf("%s/%s", c.url, k)
	resp, err := http.Post(url, "application/octet-stream", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to store preimage: %s", resp.Status)
	}
	return key, nil
}
