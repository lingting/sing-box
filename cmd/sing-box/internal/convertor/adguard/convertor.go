package adguard

import (
	"github.com/sagernet/sing-box/common/rules"
	"github.com/sagernet/sing-box/option"
	"io"
)

func Convert(reader io.Reader) ([]option.HeadlessRule, error) {
	return rules.Convert(reader)
}
