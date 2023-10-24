package times

const (
	TimeZoneUTC   = "UTC"
	TimeZoneLocal = "Local"
	TimeZoneSH    = "Asia/Shanghai"  // 上海时区
	TimeZoneCQ    = "Asia/Chongqing" // 重庆时区

	YYYY = "2006" // 年
	YY   = "06"   // 年
	MM   = "01"   // 月
	M    = "1"    // 月
	DD   = "02"   // 日
	D    = "2"    // 日
	hh   = "15"   // 时
	mm   = "04"   // 分
	ss   = "05"   // 秒

	YYYYMM    = "200601"  // 年月
	YMFormat  = "2006-01" // 年-月
	YMFormat2 = "2006/01" // 年/月
	YMFormat3 = "2006年1月" // 年/月

	YYYYMMDD   = "20060102"   // 年月日
	YMDFormat  = "2006-01-02" // 年-月-日
	YMDFormat2 = "2006/01/02" // 年/月/日
	YMDFormat3 = "2006年1月2日"  // 年月日

	YYMMDD       = "060102"   // 年月日
	YYMMDFormat  = "06-01-02" // 年-月-日
	YYMMDFormat2 = "06/01/02" // 年/月/日
	YYMMDFormat3 = "06年1月2日"  // 年月日

	YYYYMMDDhh  = "2006010215"    // 年月日时
	YMDhFormat  = "2006-01-02 15" // 年-月-日 时
	YMDhFormat2 = "2006/01/02 15" // 年/月/日 时
	YMDhFormat3 = "2006年1月2日 15时" // 年月日 时

	YYYYMMDDhhmm = "200601021504"     // 年月日时分
	YMDhmFormat  = "2006-01-02 15:04" // 年-月-日 时:分
	YMDhmFormat2 = "2006/01/02 15:04" // 年/月/日 时:分
	YMDhmFormat3 = "2006年1月2日 15时4分"  // 年月日 时分

	YYYYMMDDhhmmss = "20060102150405"      // 年月日时分秒
	TimeFormat     = "2006-01-02 15:04:05" // 年-月-日 时:分:秒
	TimeFormat2    = "2006/01/02 15:04:05" // 年/月/日 时:分:秒
	TimeFormat3    = "2006年1月2日 15时4分5秒"   // 年月日 时分秒

	YYMMDDhhmmss  = "060102150405"      // 年月日时分秒
	YMDhmsFormat  = "06-01-02 15:04:05" // 年-月-日 时:分:秒
	YMDhmsFormat2 = "06/01/02 15:04:05" // 年/月/日 时:分:秒
	YMDhmsFormat3 = "06年1月2日 15时4分5秒"   // 年月日 时分秒

	HHmmss     = "150405"   // 时分秒
	HmsFormat  = "15:04:05" // 时:分:秒
	HmsFormat2 = "15时4分5秒"  // 时分秒
)
