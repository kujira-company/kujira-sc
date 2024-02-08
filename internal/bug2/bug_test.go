package bug_test

import (
	"testing"

	bug "github.com/kujira-company/kujira-sc/internal/bug2"
)

func TestBug(t *testing.T) {
	bug.Run(true)
	bug.Run(false)
}
