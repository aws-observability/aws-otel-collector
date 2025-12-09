package encoders

type Int64Encoder struct {
	Uint64Encoder
}

func (e *Int64Encoder) IsEqual(val int64) bool {
	return e.Uint64Encoder.IsEqual(uint64(val))
}

func (e *Int64Encoder) Encode(val int64) {
	e.Uint64Encoder.Encode(uint64(val))
}

type Int64Decoder struct {
	Uint64Decoder
}

func (d *Int64Decoder) Decode(dst *int64) error {
	tsDeltaOfDelta, err := d.buf.ReadVarint()
	if err != nil {
		return err
	}

	tsDelta := d.lastDelta + uint64(tsDeltaOfDelta)
	d.lastDelta = tsDelta

	d.lastVal += tsDelta

	*dst = int64(d.lastVal)
	return nil
}
