package adguard

import (
	"github.com/sagernet/sing-box/common/ruls"
	"github.com/sagernet/sing-box/option"
	"io"
)

func Convert(reader io.Reader) ([]option.HeadlessRule, error) {
	return ruls.Convert(reader)
}
