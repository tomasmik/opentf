package replay

import (
	"os"
	"path/filepath"

	"github.com/placeholderplaceholderplaceholder/opentf/internal/builtin/provisioners/file"
)

const applyFile = "replay.log"

type StreamCopier struct {
	file *os.File
}

func NewCopier() (*StreamCopier, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	fp := filepath.Join(wd, ".terraform", applyFile)
	file, err := os.Create(fp)
	if err != nil {
		return nil, err
	}

	return &StreamCopier{
		file: file,
	}, nil
}

func (c *StreamCopier) Write(p []byte) (n int, err error) {
	return c.file.Write(p)
}

func (c *StreamCopier) Close() {
	file.New().Close()
}
