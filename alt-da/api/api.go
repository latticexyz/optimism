package api

import "github.com/ethereum/go-ethereum/common/hexutil"

// DataStatusCode combines challenge status with the local
// availability of the data.
type DataStatusCode int32

const (
	Available               DataStatusCode = 0
	AvailableAfterChallenge DataStatusCode = 1
	MissingChallengeExpired DataStatusCode = 2
	MissingPendingChallenge DataStatusCode = 3
	MissingNotChallenged    DataStatusCode = 4
)

// ChallengeStatusCode maps to the contract enum value.
type ChallengeStatusCode int8

const (
	ChallengeUninitialized ChallengeStatusCode = iota
	ChallengeActive
	ChallengeResolved
	ChallengeExpired
)

// Response associates input data with its challenge contract state.
// It communicates to the op-node derivation pipeline DASource step why data
// is not available. The main distinction is that data should be dropped if
// the challenge is expired whereas the pipeline should halt if missing data
// challenge is waiting to be resolved. Missing data that was never challenged
// should be treated as fatal.
type Response struct {
	Data   *hexutil.Bytes `json:"data"`
	Status DataStatusCode `json:"status"`
}

func NewResponse(status DataStatusCode, data []byte) *Response {
	bytes := hexutil.Bytes(data)
	return &Response{
		Data:   &bytes,
		Status: status,
	}
}

func (r *Response) SetPendingChallenge() {
	r.Status = MissingPendingChallenge
}

func (r *Response) SetExpiredChallenge() {
	r.Status = MissingChallengeExpired
}

func (r *Response) SetFatal() {
	r.Status = MissingNotChallenged
}

func (r *Response) SetResolved(data []byte) {
	r.Status = AvailableAfterChallenge
	bytes := hexutil.Bytes(data)
	r.Data = &bytes
}

// IsExpired returns true if the commitment was challenged but
// the data was not submitted before the challenge window ended.
func (r *Response) IsExpired() bool {
	return r.Status == MissingChallengeExpired
}

// IsAvailable returns true if the data is available regardless of
// whether it was challenged or not.
func (r *Response) IsAvailable() bool {
	return r.Data != nil && (r.Status == Available || r.Status == AvailableAfterChallenge)
}

// IsPending returns true if the data is not available but the challenge
// is still active so it could become available later.
func (r *Response) IsPending() bool {
	return r.Status == MissingPendingChallenge
}
