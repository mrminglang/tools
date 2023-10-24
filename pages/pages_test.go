package pages_test

import (
	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/pages"
	"testing"
)

func TestGetPages(t *testing.T) {
	sql := `ORDER BY A.ID DESC
$OFFSET_STRING`
	sql = pages.GetPages(sql, 1, 10)
	dumps.Dump(sql)
}
