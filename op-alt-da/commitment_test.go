package altda

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/rollup/derive/params"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func encodeCommitmentData(commitmentType CommitmentType, data []byte) []byte {
	return append([]byte{byte(commitmentType)}, data...)
}

// TestCommitmentData tests the CommitmentData type and its implementations,
// by encoding and decoding the commitment data and verifying the input data.
func TestCommitmentData(t *testing.T) {

	type tcase struct {
		name        string
		commType    CommitmentType
		commInput   []byte
		commData    []byte
		expectedErr error
	}

	input := []byte{0}
	hash := crypto.Keccak256([]byte{0})

	testCases := []tcase{
		{
			name:        "valid keccak256 commitment",
			commType:    Keccak256CommitmentType,
			commInput:   input,
			commData:    encodeCommitmentData(Keccak256CommitmentType, hash),
			expectedErr: nil,
		},
		{
			name:        "invalid keccak256 commitment",
			commType:    Keccak256CommitmentType,
			commInput:   input,
			commData:    encodeCommitmentData(Keccak256CommitmentType, []byte("ab_baddata_yz012345")),
			expectedErr: ErrInvalidCommitment,
		},
		{
			name:        "valid generic commitment",
			commType:    GenericCommitmentType,
			commInput:   []byte("any input works"),
			commData:    encodeCommitmentData(GenericCommitmentType, []byte("any length of data! wow, that's so generic!")),
			expectedErr: nil, // This should actually be valid now
		},
		{
			name:        "invalid commitment type",
			commType:    9,
			commInput:   []byte("some input"),
			commData:    encodeCommitmentData(CommitmentType(9), []byte("abcdefghijklmnopqrstuvwxyz012345")),
			expectedErr: ErrInvalidCommitment,
		},
		{
			name:        "valid generic keccak256 commitment",
			commType:    GenericCommitmentType,
			commInput:   input,
			commData:    encodeCommitmentData(GenericCommitmentType, append([]byte{Keccak256DALayer}, hash...)), // Single zero hash
			expectedErr: nil,
		},
		{
			name:        "valid generic keccak256 commitment (batched)",
			commType:    GenericCommitmentType,
			commInput:   input,
			commData:    encodeCommitmentData(GenericCommitmentType, append([]byte{Keccak256DALayer}, append(hash, hash...)...)), // Two zero hashes
			expectedErr: nil,
		},
		{
			name:        "invalid generic keccak256 commitment - wrong length",
			commType:    GenericCommitmentType,
			commInput:   input,
			commData:    encodeCommitmentData(GenericCommitmentType, append([]byte{Keccak256DALayer}, hash[1:]...)), // Not multiple of 32
			expectedErr: ErrInvalidCommitment,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.commData)
			comm, err := DecodeCommitmentData(tc.commData)
			require.ErrorIs(t, err, tc.expectedErr)
			if err == nil {
				// Test that the commitment type is correct
				require.Equal(t, tc.commType, comm.CommitmentType())
				// Test that reencoding the commitment returns the same data
				require.Equal(t, tc.commData, comm.Encode())
				// Test that TxData() returns the same data as the original, prepended with a version byte
				require.Equal(t, append([]byte{params.DerivationVersion1}, tc.commData...), comm.TxData())

				// Test that Verify() returns no error for the correct data
				require.NoError(t, comm.Verify(tc.commInput))
				// Test that Verify() returns error for the incorrect data
				// don't do this for GenericCommitmentType, which does not do any verification
				if tc.commType != GenericCommitmentType {
					require.ErrorIs(t, ErrCommitmentMismatch, comm.Verify([]byte("wrong data")))
				}
			}
		})
	}
}


// TestBatchedCommitmentData specifically tests the GenericKeccak256CommitmentType
// with multiple hashes in a batch
func TestBatchedCommitmentData(t *testing.T) {
	// Create some test inputs and their corresponding hashes
	inputs := [][]byte{
		[]byte("first input"),
		[]byte("second input"),
		[]byte("third input"),
	}
	hashes := make([][]byte, len(inputs))
	for i, input := range inputs {
		hashes[i] = crypto.Keccak256(input)
	}

	// Combine all hashes into a single commitment
	batchedData := []byte{Keccak256DALayer}
	for _, hash := range hashes {
		batchedData = append(batchedData, hash...)
	}
	commData := encodeCommitmentData(GenericCommitmentType, batchedData)

	// Decode and verify the commitment
	comm, err := DecodeCommitmentData(commData)
	require.NoError(t, err)
	require.Equal(t, GenericCommitmentType, comm.CommitmentType())

	// Test that we can get the batched commitments
	batchedComm, ok := comm.(BatchedCommitmentData)
	require.True(t, ok)
	batched := batchedComm.BatchedCommitments()
	require.Len(t, batched, len(inputs))

	// Verify each individual commitment in the batch
	for _, comm := range batched {
		// Should be a GenericCommitment with DA Layer byte == 0xff
		require.Equal(t, GenericCommitmentType, comm.CommitmentType())
		require.Equal(t, Keccak256DALayer, comm.Encode()[1]) // index 0 is commitment type, index 1 is DA Layer
	}

	// Test invalid batch sizes
	invalidBatch := append([]byte{Keccak256DALayer}, make([]byte, 31)...) // Not multiple of 32
	invalidComm := encodeCommitmentData(GenericCommitmentType, invalidBatch)
	_, err = DecodeCommitmentData(invalidComm)
	require.ErrorIs(t, err, ErrInvalidCommitment)

	// Test empty batch
	emptyBatch := []byte{Keccak256DALayer}
	emptyComm := encodeCommitmentData(GenericCommitmentType, emptyBatch)
	_, err = DecodeCommitmentData(emptyComm)
	require.ErrorIs(t, err, ErrInvalidCommitment)
}

