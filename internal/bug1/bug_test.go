package bug_test

import (
	"testing"

	bug "github.com/kujira-company/kujira-sc/internal/bug1"
)

func TestRun(t *testing.T) {
	bug.Run(1, 1)
	bug.Run(1, 2)
	bug.Run(1, 3)
	bug.Run(1, 4)
	bug.Run(2, 1)
	bug.Run(2, 2)
	bug.Run(2, 3)
	bug.Run(2, 4)
}
