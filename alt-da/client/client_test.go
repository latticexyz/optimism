package client

import (
	"context"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/alt-da/api"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum-optimism/optimism/op-service/testutils"
)

type mockRPCAPI struct {
	store    ethdb.KeyValueStore
	challNum uint64
	safeHead eth.L2BlockRef
}

func NewMockRPCAPI() *mockRPCAPI {
	return &mockRPCAPI{
		store:    memorydb.New(),
		challNum: 0,
	}
}

func (a *mockRPCAPI) GetPreImage(ctx context.Context, key hexutil.Bytes, bn hexutil.Uint64) (*api.Response, error) {
	val, err := a.store.Get(key)
	res := api.NewResponse(api.Available, val)
	if err != nil {
		res.SetFatal()
		return res, nil
	}
	return res, nil
}

func (a *mockRPCAPI) ExpiredChallengeNumber(ctx context.Context) (uint64, error) {
	return a.challNum, nil
}

func (a *mockRPCAPI) SetPreImage(ctx context.Context, img hexutil.Bytes) (hexutil.Bytes, error) {

	key := crypto.Keccak256(img)

	a.store.Put(key, img)

	return key, nil
}

func (a *mockRPCAPI) SetSubjectiveSafeHead(ctx context.Context, ref eth.L2BlockRef) error {
	a.safeHead = ref
	return nil
}

func (a *mockRPCAPI) setChallNum(cn uint64) {
	a.challNum = cn
}

// This mainly tests rpc client serialization/deserialization
func TestDAProvider(t *testing.T) {
	m := NewMockRPCAPI()
	srv := rpc.NewServer()
	srv.RegisterName("da", m)

	nodeHandler := node.NewHTTPHandlerStack(srv, []string{"*"}, []string{"*"}, nil)

	mux := http.NewServeMux()
	mux.Handle("/", nodeHandler)

	tsrv := httptest.NewServer(mux)

	client, err := New(context.Background(), testlog.Logger(t, log.LvlDebug), tsrv.URL)
	require.NoError(t, err)

	data := []byte("hello world")
	key, err := client.SetPreImage(context.Background(), data)
	require.NoError(t, err)

	require.Equal(t, key, crypto.Keccak256(data))

	resp, err := client.GetPreImage(context.Background(), key, 1)
	require.NoError(t, err)

	require.Equal(t, api.Available, resp.Status)
	require.Equal(t, data, []byte(*resp.Data))

	m.setChallNum(1111)
	cs, err := client.ExpiredChallengeNumber(context.Background())
	require.NoError(t, err)
	require.Equal(t, uint64(1111), cs)

	rng := rand.New(rand.NewSource(1234))
	l2block := testutils.RandomL2BlockRef(rng)

	require.NoError(t, client.SetSubjectiveSafeHead(context.Background(), l2block))
	require.Equal(t, l2block, m.safeHead)
}

// Test against a local DA service to ensure data compatibility with the Rust RPC.
func TestCompat(t *testing.T) {
	t.Skip()
	cl, err := New(context.Background(), testlog.Logger(t, log.LvlDebug), "http://localhost:8064")
	require.NoError(t, err)

	data := []byte("hello world")
	key, err := cl.SetPreImage(context.TODO(), data)
	require.NoError(t, err)

	require.Equal(t, key, crypto.Keccak256(data))

	resp, err := cl.GetPreImage(context.TODO(), key, 12)
	require.NoError(t, err)

	require.Equal(t, resp.Status, api.Available)
	require.Equal(t, data, []byte(*resp.Data))

	cn, err := cl.ExpiredChallengeNumber(context.Background())
	require.NoError(t, err)
	require.Equal(t, uint64(0), cn)
}
