// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package whisperv6

import (
	"encoding/json"

	"github.com/one-currency/one-coin/common/hexutil"
)

var _ = (*criteriaOverride)(nil)

func (c Criteria) MarshalJSON() ([]byte, error) {
	type Criteria struct {
		SymKeyID     string        `json:"symKeyID"`
		PrivateKeyID string        `json:"privateKeyID"`
		Sig          hexutil.Bytes `json:"sig"`
		MinPow       float64       `json:"minPow"`
		Topics       []TopicType   `json:"topics"`
		AllowP2P     bool          `json:"allowP2P"`
	}
	var enc Criteria
	enc.SymKeyID = c.SymKeyID
	enc.PrivateKeyID = c.PrivateKeyID
	enc.Sig = c.Sig
	enc.MinPow = c.MinPow
	enc.Topics = c.Topics
	enc.AllowP2P = c.AllowP2P
	return json.Marshal(&enc)
}

func (c *Criteria) UnmarshalJSON(input []byte) error {
	type Criteria struct {
		SymKeyID     *string       `json:"symKeyID"`
		PrivateKeyID *string       `json:"privateKeyID"`
		Sig          hexutil.Bytes `json:"sig"`
		MinPow       *float64      `json:"minPow"`
		Topics       []TopicType   `json:"topics"`
		AllowP2P     *bool         `json:"allowP2P"`
	}
	var dec Criteria
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.SymKeyID != nil {
		c.SymKeyID = *dec.SymKeyID
	}
	if dec.PrivateKeyID != nil {
		c.PrivateKeyID = *dec.PrivateKeyID
	}
	if dec.Sig != nil {
		c.Sig = dec.Sig
	}
	if dec.MinPow != nil {
		c.MinPow = *dec.MinPow
	}
	if dec.Topics != nil {
		c.Topics = dec.Topics
	}
	if dec.AllowP2P != nil {
		c.AllowP2P = *dec.AllowP2P
	}
	return nil
}
