package protocol

type Null struct{}

func (this *Null) Decode(bs []byte) []byte { return bs }
