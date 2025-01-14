// Code generated by github.com/spacemeshos/go-scale/scalegen. DO NOT EDIT.

package types

import (
	"github.com/spacemeshos/go-scale"
)

func (t *TxHeader) EncodeScale(enc *scale.Encoder) (total int, err error) {
	if n, err := scale.EncodeByteArray(enc, t.Principal[:]); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeByteArray(enc, t.Template[:]); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeCompact8(enc, uint8(t.Method)); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := t.Nonce.EncodeScale(enc); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := t.LayerLimits.EncodeScale(enc); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeCompact64(enc, uint64(t.MaxGas)); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeCompact64(enc, uint64(t.GasPrice)); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeCompact64(enc, uint64(t.MaxSpend)); err != nil {
		return total, err
	} else {
		total += n
	}
	return total, nil
}

func (t *TxHeader) DecodeScale(dec *scale.Decoder) (total int, err error) {
	if n, err := scale.DecodeByteArray(dec, t.Principal[:]); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.DecodeByteArray(dec, t.Template[:]); err != nil {
		return total, err
	} else {
		total += n
	}
	if field, n, err := scale.DecodeCompact8(dec); err != nil {
		return total, err
	} else {
		total += n
		t.Method = uint8(field)
	}
	if n, err := t.Nonce.DecodeScale(dec); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := t.LayerLimits.DecodeScale(dec); err != nil {
		return total, err
	} else {
		total += n
	}
	if field, n, err := scale.DecodeCompact64(dec); err != nil {
		return total, err
	} else {
		total += n
		t.MaxGas = uint64(field)
	}
	if field, n, err := scale.DecodeCompact64(dec); err != nil {
		return total, err
	} else {
		total += n
		t.GasPrice = uint64(field)
	}
	if field, n, err := scale.DecodeCompact64(dec); err != nil {
		return total, err
	} else {
		total += n
		t.MaxSpend = uint64(field)
	}
	return total, nil
}

func (t *LayerLimits) EncodeScale(enc *scale.Encoder) (total int, err error) {
	if n, err := scale.EncodeCompact32(enc, uint32(t.Min)); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeCompact32(enc, uint32(t.Max)); err != nil {
		return total, err
	} else {
		total += n
	}
	return total, nil
}

func (t *LayerLimits) DecodeScale(dec *scale.Decoder) (total int, err error) {
	if field, n, err := scale.DecodeCompact32(dec); err != nil {
		return total, err
	} else {
		total += n
		t.Min = uint32(field)
	}
	if field, n, err := scale.DecodeCompact32(dec); err != nil {
		return total, err
	} else {
		total += n
		t.Max = uint32(field)
	}
	return total, nil
}

func (t *Nonce) EncodeScale(enc *scale.Encoder) (total int, err error) {
	if n, err := scale.EncodeCompact64(enc, uint64(t.Counter)); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeCompact8(enc, uint8(t.Bitfield)); err != nil {
		return total, err
	} else {
		total += n
	}
	return total, nil
}

func (t *Nonce) DecodeScale(dec *scale.Decoder) (total int, err error) {
	if field, n, err := scale.DecodeCompact64(dec); err != nil {
		return total, err
	} else {
		total += n
		t.Counter = uint64(field)
	}
	if field, n, err := scale.DecodeCompact8(dec); err != nil {
		return total, err
	} else {
		total += n
		t.Bitfield = uint8(field)
	}
	return total, nil
}
