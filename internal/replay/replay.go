package replay

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/placeholderplaceholderplaceholder/opentf/internal/terminal"
)

func ReplayApply(s *terminal.Streams) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	fp := filepath.Join(wd, ".terraform", applyFile)

	content, err := os.ReadFile(fp)
	if err != nil {
		return fmt.Errorf("Error opening file: %v\n", err)
	}

	s.Println(string(content))

	return nil
}
