// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package genesis

import (
	"encoding/json"
	"math/big"

	"github.com/celo-org/celo-blockchain/common/decimal/bigintstr"
)

var _ = (*StableTokenExchangeLimitsMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (s StableTokenExchangeLimit) MarshalJSON() ([]byte, error) {
	type StableTokenExchangeLimit struct {
		StableToken       string               `json:"stableToken"`
		MinExchangeAmount *bigintstr.BigIntStr `json:"minExchangeAmount"`
		MaxExchangeAmount *bigintstr.BigIntStr `json:"maxExchangeAmount"`
	}
	var enc StableTokenExchangeLimit
	enc.StableToken = s.StableToken
	enc.MinExchangeAmount = (*bigintstr.BigIntStr)(s.MinExchangeAmount)
	enc.MaxExchangeAmount = (*bigintstr.BigIntStr)(s.MaxExchangeAmount)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (s *StableTokenExchangeLimit) UnmarshalJSON(input []byte) error {
	type StableTokenExchangeLimit struct {
		StableToken       *string              `json:"stableToken"`
		MinExchangeAmount *bigintstr.BigIntStr `json:"minExchangeAmount"`
		MaxExchangeAmount *bigintstr.BigIntStr `json:"maxExchangeAmount"`
	}
	var dec StableTokenExchangeLimit
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.StableToken != nil {
		s.StableToken = *dec.StableToken
	}
	if dec.MinExchangeAmount != nil {
		s.MinExchangeAmount = (*big.Int)(dec.MinExchangeAmount)
	}
	if dec.MaxExchangeAmount != nil {
		s.MaxExchangeAmount = (*big.Int)(dec.MaxExchangeAmount)
	}
	return nil
}