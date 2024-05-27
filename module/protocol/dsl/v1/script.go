package dsl

import (
	"github.com/injoyai/conv"
	"github.com/injoyai/gateway/module/protocol/internal/common"
	"github.com/injoyai/goutil/oss"
	"github.com/injoyai/goutil/script/dsl"
	"io/fs"
	"os"
	"path/filepath"
)

func New(dir string, options ...func(*dsl.Decode)) (map[string]common.Decoder, error) {
	m := map[string]common.Decoder{}
	oss.RangeFileInfo(dir, func(info fs.FileInfo) (bool, error) {
		if !info.IsDir() {
			bs, err := os.ReadFile(filepath.Join(dir, info.Name()))
			if err != nil {
				return false, err
			}
			d, err := dsl.NewDecode(bs, options...)
			if err != nil {
				return false, err
			}
			m[d.Name] = &decode{d}
		}
		return true, nil
	})
	return m, nil
}

type decode struct {
	d *dsl.Decode
}

func (this *decode) Decode(bs []byte) ([]byte, error) {
	_, result, err := this.d.Do(bs)
	return conv.Bytes(result), err
}
