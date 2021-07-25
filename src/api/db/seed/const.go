package seed

// seed package内のみのスコープとしたい場合は、小文字から始める。
// 大文字で始めるとパッケージ外からアクセスできてしまう。

// pref, city, ward mstで使用する定数
const (
	code_Index       int = 0
	prefecture_Index int = 1
	city_Index       int = 2
)

const (
	ward_code_Index int = 0
	ward_Index      int = 1
)
