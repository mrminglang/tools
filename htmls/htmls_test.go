package htmls_test

import (
	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/htmls"
	"testing"
)

func TestConvertToText(t *testing.T) {
	htmlContent := `<div class="page" page_index=0>
<p class="text title">2024&nbsp;~~-·txllai~ft*:X~B&lt;J&nbsp;</p>
<p class="text title">¥!1$~£~&nbsp;</p>
<p class="text block"><span class="text">~$&nbsp;</span></p>
<p class="text paragraph" > ~•mm~~~~~M~.&nbsp;~~&nbsp;A~~•m~T=&nbsp;</p>
<p class="text paragraph" > 中文港股内房股震荡走低，融创中国跌超14%，富力地产等跟跌。</p>
</div>`
	text := htmls.ConvertToText(htmlContent)

	dumps.Dump(text)
}
