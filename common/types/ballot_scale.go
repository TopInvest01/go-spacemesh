// Code generated by github.com/spacemeshos/go-scale/scalegen. DO NOT EDIT.

package types

import (
	"github.com/spacemeshos/go-scale"
)

func (t *Ballot) EncodeScale(enc *scale.Encoder) (total int, err error) {
	if n, err := t.InnerBallot.EncodeScale(enc); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeByteSlice(enc, t.Signature); err != nil {
		return total, err
	} else {
		total += n
	}
	return total, nil
}

func (t *Ballot) DecodeScale(dec *scale.Decoder) (total int, err error) {
	if n, err := t.InnerBallot.DecodeScale(dec); err != nil {
		return total, err
	} else {
		total += n
	}
	if field, n, err := scale.DecodeByteSlice(dec); err != nil {
		return total, err
	} else {
		total += n
		t.Signature = field
	}
	return total, nil
}

func (t *InnerBallot) EncodeScale(enc *scale.Encoder) (total int, err error) {
	if n, err := scale.EncodeByteArray(enc, t.AtxID[:]); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeStructSlice(enc, t.EligibilityProofs); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := t.Votes.EncodeScale(enc); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeByteArray(enc, t.RefBallot[:]); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeOption(enc, t.EpochData); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := t.LayerIndex.EncodeScale(enc); err != nil {
		return total, err
	} else {
		total += n
	}
	return total, nil
}

func (t *InnerBallot) DecodeScale(dec *scale.Decoder) (total int, err error) {
	if n, err := scale.DecodeByteArray(dec, t.AtxID[:]); err != nil {
		return total, err
	} else {
		total += n
	}
	if field, n, err := scale.DecodeStructSlice[VotingEligibilityProof](dec); err != nil {
		return total, err
	} else {
		total += n
		t.EligibilityProofs = field
	}
	if n, err := t.Votes.DecodeScale(dec); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.DecodeByteArray(dec, t.RefBallot[:]); err != nil {
		return total, err
	} else {
		total += n
	}
	if field, n, err := scale.DecodeOption[EpochData](dec); err != nil {
		return total, err
	} else {
		total += n
		t.EpochData = field
	}
	if n, err := t.LayerIndex.DecodeScale(dec); err != nil {
		return total, err
	} else {
		total += n
	}
	return total, nil
}

func (t *Votes) EncodeScale(enc *scale.Encoder) (total int, err error) {
	if n, err := scale.EncodeByteArray(enc, t.Base[:]); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeStructSlice(enc, t.Support); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeStructSlice(enc, t.Against); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeStructSlice(enc, t.Abstain); err != nil {
		return total, err
	} else {
		total += n
	}
	return total, nil
}

func (t *Votes) DecodeScale(dec *scale.Decoder) (total int, err error) {
	if n, err := scale.DecodeByteArray(dec, t.Base[:]); err != nil {
		return total, err
	} else {
		total += n
	}
	if field, n, err := scale.DecodeStructSlice[BlockID](dec); err != nil {
		return total, err
	} else {
		total += n
		t.Support = field
	}
	if field, n, err := scale.DecodeStructSlice[BlockID](dec); err != nil {
		return total, err
	} else {
		total += n
		t.Against = field
	}
	if field, n, err := scale.DecodeStructSlice[LayerID](dec); err != nil {
		return total, err
	} else {
		total += n
		t.Abstain = field
	}
	return total, nil
}

func (t *EpochData) EncodeScale(enc *scale.Encoder) (total int, err error) {
	if n, err := scale.EncodeStructSlice(enc, t.ActiveSet); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeByteArray(enc, t.Beacon[:]); err != nil {
		return total, err
	} else {
		total += n
	}
	return total, nil
}

func (t *EpochData) DecodeScale(dec *scale.Decoder) (total int, err error) {
	if field, n, err := scale.DecodeStructSlice[ATXID](dec); err != nil {
		return total, err
	} else {
		total += n
		t.ActiveSet = field
	}
	if n, err := scale.DecodeByteArray(dec, t.Beacon[:]); err != nil {
		return total, err
	} else {
		total += n
	}
	return total, nil
}

func (t *VotingEligibilityProof) EncodeScale(enc *scale.Encoder) (total int, err error) {
	if n, err := scale.EncodeCompact32(enc, uint32(t.J)); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeByteSlice(enc, t.Sig); err != nil {
		return total, err
	} else {
		total += n
	}
	return total, nil
}

func (t *VotingEligibilityProof) DecodeScale(dec *scale.Decoder) (total int, err error) {
	if field, n, err := scale.DecodeCompact32(dec); err != nil {
		return total, err
	} else {
		total += n
		t.J = uint32(field)
	}
	if field, n, err := scale.DecodeByteSlice(dec); err != nil {
		return total, err
	} else {
		total += n
		t.Sig = field
	}
	return total, nil
}

func (t *DBBallot) EncodeScale(enc *scale.Encoder) (total int, err error) {
	if n, err := t.InnerBallot.EncodeScale(enc); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeByteArray(enc, t.ID[:]); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeByteSlice(enc, t.Signature); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.EncodeByteSlice(enc, t.SmesherID); err != nil {
		return total, err
	} else {
		total += n
	}
	return total, nil
}

func (t *DBBallot) DecodeScale(dec *scale.Decoder) (total int, err error) {
	if n, err := t.InnerBallot.DecodeScale(dec); err != nil {
		return total, err
	} else {
		total += n
	}
	if n, err := scale.DecodeByteArray(dec, t.ID[:]); err != nil {
		return total, err
	} else {
		total += n
	}
	if field, n, err := scale.DecodeByteSlice(dec); err != nil {
		return total, err
	} else {
		total += n
		t.Signature = field
	}
	if field, n, err := scale.DecodeByteSlice(dec); err != nil {
		return total, err
	} else {
		total += n
		t.SmesherID = field
	}
	return total, nil
}
