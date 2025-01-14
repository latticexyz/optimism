package altda

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/ethereum-optimism/optimism/op-node/rollup/derive/params"
	"github.com/ethereum/go-ethereum/crypto"
)

// ErrInvalidCommitment is returned when the commitment cannot be parsed into a known commitment type.
var ErrInvalidCommitment = errors.New("invalid commitment")

// ErrCommitmentMismatch is returned when the commitment does not match the given input.
var ErrCommitmentMismatch = errors.New("commitment mismatch")

// CommitmentType is the commitment type prefix.
type CommitmentType byte

// DALayer specifies the DA Layer byte for generic commitments.
type DALayer byte

func CommitmentTypeFromString(s string) (CommitmentType, error) {
	switch s {
	case KeccakCommitmentString:
		return Keccak256CommitmentType, nil
	case GenericCommitmentString:
		return GenericCommitmentType, nil
	default:
		return 0, fmt.Errorf("invalid commitment type: %s", s)
	}
}

// CommitmentType describes the binary format of the commitment.
// KeccakCommitmentType is the default commitment type for the centralized DA storage.
// GenericCommitmentType indicates an opaque bytestring that the op-node never opens.
const (
	Keccak256CommitmentType CommitmentType = 0
	GenericCommitmentType   CommitmentType = 1
	KeccakCommitmentString  string         = "KeccakCommitment"
	GenericCommitmentString string         = "GenericCommitment"

	Keccak256DALayer DALayer = 0xff
)

// CommitmentData is the binary representation of a commitment.
type CommitmentData interface {
	CommitmentType() CommitmentType
	Encode() []byte
	TxData() []byte
	Verify(input []byte) error
	String() string
}

type GenericCommitmentData interface {
	CommitmentData
	DALayer() DALayer
}

type GenericBatchedCommitmentData interface {
	GenericCommitmentData
	GetBatchedCommitments() []CommitmentData
}

// Keccak256Commitment is an implementation of CommitmentData that uses Keccak256 as the commitment function.
type Keccak256Commitment []byte

// GenericCommitment is an implementation of CommitmentData that treats the commitment as an opaque bytestring.
type GenericCommitment []byte

type GenericKeccak256Commitment []byte

// NewCommitmentData creates a new commitment from the given input and desired type.
func NewCommitmentData(t CommitmentType, input []byte) CommitmentData {
	switch t {
	case Keccak256CommitmentType:
		return NewKeccak256Commitment(input)
	case GenericCommitmentType:
		return NewGenericCommitment(input)
	default:
		return nil
	}
}

// DecodeCommitmentData parses the commitment into a known commitment type.
// The input type is determined by the first byte of the raw data.
// The input type is discarded and the commitment is passed to the appropriate constructor.
func DecodeCommitmentData(input []byte) (CommitmentData, error) {
	if len(input) == 0 {
		return nil, ErrInvalidCommitment
	}
	t := CommitmentType(input[0])
	data := input[1:]
	switch t {
	case Keccak256CommitmentType:
		return DecodeKeccak256(data)
	case GenericCommitmentType:
		return DecodeGenericCommitment(data)
	default:
		return nil, ErrInvalidCommitment
	}
}

// NewKeccak256Commitment creates a new commitment from the given input.
func NewKeccak256Commitment(input []byte) Keccak256Commitment {
	return Keccak256Commitment(crypto.Keccak256(input))
}

// DecodeKeccak256 validates and casts the commitment into a Keccak256Commitment.
func DecodeKeccak256(commitment []byte) (Keccak256Commitment, error) {
	// guard against empty commitments
	if len(commitment) == 0 {
		return nil, ErrInvalidCommitment
	}
	// keccak commitments are always 32 bytes
	if len(commitment) != 32 {
		return nil, ErrInvalidCommitment
	}
	return commitment, nil
}

// CommitmentType returns the commitment type of Keccak256.
func (c Keccak256Commitment) CommitmentType() CommitmentType {
	return Keccak256CommitmentType
}

// Encode adds a commitment type prefix that describes the commitment.
func (c Keccak256Commitment) Encode() []byte {
	return append([]byte{byte(Keccak256CommitmentType)}, c...)
}

// TxData adds an extra version byte to signal it's a commitment.
func (c Keccak256Commitment) TxData() []byte {
	return append([]byte{params.DerivationVersion1}, c.Encode()...)
}

// Verify checks if the commitment matches the given input.
func (c Keccak256Commitment) Verify(input []byte) error {
	if !bytes.Equal(c, crypto.Keccak256(input)) {
		return ErrCommitmentMismatch
	}
	return nil
}

func (c Keccak256Commitment) String() string {
	return hex.EncodeToString(c.Encode())
}

// NewGenericCommitment creates a new commitment from the given input.
func NewGenericCommitment(input []byte) CommitmentData {
	if len(input) == 0 {
		return nil
	}

	c := GenericCommitment(input)
	switch c.DALayer() {
	case Keccak256DALayer:
		return NewGenericKeccak256Commitment(input[1:])
	default:
		return c
	}
}

// DecodeGenericCommitment validates and casts the commitment into a GenericCommitment.
func DecodeGenericCommitment(commitment []byte) (CommitmentData, error) {
	if len(commitment) == 0 {
		return nil, ErrInvalidCommitment
	}

	c := GenericCommitment(commitment)
	switch c.DALayer() {
	case Keccak256DALayer:
		return DecodeGenericKeccak256Commitment(commitment)
	default:
		return c, nil
	}
}

// CommitmentType returns the commitment type of Generic Commitment.
func (c GenericCommitment) CommitmentType() CommitmentType {
	return GenericCommitmentType
}

// Encode adds a commitment type prefix self describing the commitment.
func (c GenericCommitment) Encode() []byte {
	return append([]byte{byte(GenericCommitmentType)}, c...)
}

// TxData adds an extra version byte to signal it's a commitment.
func (c GenericCommitment) TxData() []byte {
	return append([]byte{params.DerivationVersion1}, c.Encode()...)
}

// Verify always returns true for GenericCommitment because the DA Server must validate the data before returning it to the op-node.
func (c GenericCommitment) Verify(input []byte) error {
	return nil
}

func (c GenericCommitment) String() string {
	return hex.EncodeToString(c.Encode())
}

func (c GenericCommitment) DALayer() DALayer {
	return DALayer(c[0])
}

func NewGenericKeccak256Commitment(input []byte) GenericKeccak256Commitment {
	hash := crypto.Keccak256(input)
	commitment := make([]byte, len(hash)+1)
	commitment[0] = byte(Keccak256DALayer)
	copy(commitment[1:], hash)
	return GenericKeccak256Commitment(commitment)
}

// DecodeGenericKeccak256Commitment validates and creates a GenericKeccak256Commitment
func DecodeGenericKeccak256Commitment(commitment []byte) (GenericKeccak256Commitment, error) {
	if len(commitment) == 0 {
		return nil, ErrInvalidCommitment
	}
	if DALayer(commitment[0]) != Keccak256DALayer {
		return nil, fmt.Errorf("invalid DA layer for Keccak256 commitment: %d", commitment[0])
	}
	if (len(commitment)-1)%32 != 0 {
		return nil, fmt.Errorf("invalid length for Keccak256 commitment: %d", len(commitment))
	}
	return GenericKeccak256Commitment(commitment), nil
}


func (c GenericKeccak256Commitment) CommitmentType() CommitmentType {
	return GenericCommitment(c).CommitmentType()
}

func (c GenericKeccak256Commitment) Encode() []byte {
	return GenericCommitment(c).Encode()
}

// TxData adds an extra version byte to signal it's a commitment.
func (c GenericKeccak256Commitment) TxData() []byte {
	return GenericCommitment(c).TxData()
}

// Verify checks if the commitment matches the given input
func (c GenericKeccak256Commitment) Verify(input []byte) error {
	if !bytes.Equal(c[1:], crypto.Keccak256(input)) {
		return ErrCommitmentMismatch
	}
	// TODO: support batched commitments
	return fmt.Errorf("cannot verify batched commitment against single input")
}

// String provides a custom string representation
func (c GenericKeccak256Commitment) String() string {
	return "keccak256:" + hex.EncodeToString(c)
}

// GetBatchedCommitments returns all individual commitments contained in the commitment
func (c GenericKeccak256Commitment) GetBatchedCommitments() []CommitmentData {
	numCommitments := (len(c) - 1) / 32
	comms := make([]CommitmentData, numCommitments)

	for i := 0; i < numCommitments; i++ {
		start := 1 + (i * 32)
		end := start + 32
		// Create a new commitment with the Keccak256DALayer prefix
		comm:= make([]byte, 33) // 1 byte prefix + 32 bytes hash
		comm[0] = byte(Keccak256DALayer)
		copy(comm[1:], c[start:end])
		comms[i] = GenericKeccak256Commitment(comm)
	}

	return comms
}

func (c GenericKeccak256Commitment) DALayer() DALayer {
	return GenericCommitment(c).DALayer()
}
