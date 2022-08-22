package types

type RandomnessChainInfoRow struct {
	PublicKey   string `db:"public_key"`
	Period      uint64 `db:"period"`
	GenesisTime uint64 `db:"genesis_time"`
	Hash        string `db:"hash"`
}

type UnprovenRandomness struct {
	Round uint64 `db:"round"`
}

type ProvenRandomness struct {
	Round             uint64 `db:"round"`
	Randomness        string `db:"randomness"`
	Signature         string `db:"signature"`
	PreviousSignature string `db:"previous_signature"`
}
