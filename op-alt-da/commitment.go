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

	Keccak256DALayer byte = 0xff
)

// CommitmentData is the binary representation of a commitment.
type CommitmentData interface {
	CommitmentType() CommitmentType
	Encode() []byte
	TxData() []byte
	Verify(input []byte) error
	String() string
}

type BatchedCommitmentData interface {
	BatchedCommitments() []CommitmentData
}

// Keccak256Commitment is an implementation of CommitmentData that uses Keccak256 as the commitment function.
type Keccak256Commitment []byte

// GenericCommitment is an implementation of CommitmentData that treats the commitment as an opaque bytestring.
type GenericCommitment []byte

type GenericKeccak256Commitment struct {
	GenericCommitment
}

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
func NewGenericCommitment(input []byte) GenericCommitment {
	return GenericCommitment(input)
}

// DecodeGenericCommitment validates and casts the commitment into a GenericCommitment.
func DecodeGenericCommitment(commitment []byte) (CommitmentData, error) {
	if len(commitment) == 0 {
		return nil, ErrInvalidCommitment
	}

	daByte := commitment[0]
	switch daByte {
		case Keccak256DALayer:
			return DecodeGenericKeccak256Commitment(commitment)
		default:
			return GenericCommitment(commitment[:]), nil
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

// DecodeGenericKeccak256Commitment validates and creates a GenericKeccak256Commitment
func DecodeGenericKeccak256Commitment(commitment []byte) (GenericKeccak256Commitment, error) {
	if len(commitment) == 0 {
		return GenericKeccak256Commitment{}, ErrInvalidCommitment
	}
	// Strip the DALayer byte
	if (len(commitment)-1)%32 != 0 {
		return GenericKeccak256Commitment{}, fmt.Errorf("invalid length for Generic Keccak256 commitment: %d", len(commitment))
	}
	return GenericKeccak256Commitment{ GenericCommitment: commitment }, nil
}

// BatchedCommitments implements the BatchedCommitmentData interface
func (c GenericKeccak256Commitment) BatchedCommitments() []CommitmentData {
	// Skip the DA layer byte and calculate how many 32-byte Keccak hashes we have
	commitment := c.GenericCommitment[1:]
	numCommitments := len(commitment) / 32

	// Each 32-byte chunk represents a Keccak256 commitment
	comms := make([]CommitmentData, numCommitments)
	for i := 0; i < numCommitments; i++ {
		start := i * 32
		end := start + 32
		comms[i] = GenericKeccak256Commitment{GenericCommitment:append([]byte{Keccak256DALayer}, commitment[start:end]...)}
	}
	return comms
}

