package v1

import (
	"github.com/injoyai/goutil/oss"
	"github.com/injoyai/goutil/script/dsl"
	"os"
)

var Infos = []*Info(nil)

type Info struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

func Loading(dir string) error {
	infos := []*Info(nil)
	if err := oss.RangeFile(dir, func(info os.FileInfo, f *os.File) (bool, error) {

		d, err := dsl.NewDecode(f)
		if err != nil {
			return false, err
		}
		infos = append(infos, &Info{
			Name: d.Name,
			Key:  info.Name(),
		})
		return true, nil
	}); err != nil {
		return err
	}
	Infos = infos
	return nil
}
