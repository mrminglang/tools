package regex

const (
	// 收款类型-银行卡/支付宝
	PayeeTypeBankCN string = "银行卡"
	PayeeTypeBankEN string = "bank_card"
	PayeeTypeAliCN  string = "支付宝"
	PayeeTypeAliEN  string = "alipay"

	// 加入了可以验证港澳台手机号的规则，我们的客户中可能会有这样的收款人
	RegularPhone    string = "^((1[3-9][0-9]{9})|(([6|9])[0-9]{7})|([0][9][0-9]{8})|([6]([8|6])[0-9]{5}))$"
	RegularPhoneMsg string = "不是一个有效的中国内地手机号码"

	//全中文，2～10
	RegularName    string = "^[·\u4e00-\u9fa5]{2,20}$"
	RegularNameMsg string = "姓名错误，只能是中文，点号分隔符·，且须2～20个字符。"

	// URL
	RegularUrl    string = "^(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]$"
	RegularUrlMsg string = "不是一个有效的URL地址"

	// 电子邮箱
	RegularEmailMsg = "电子邮箱格式错误"

	// 验证支付宝userid
	RegularAlipayUserID    = "^2088[0-9]{12}$"
	RegularAlipayUserIDMsg = "收款账号格式错误，须是以2088开头的16位纯数字组成"

	// 收款账号
	RegularIdentityMsg string = "收款账号格式错误"

	// 身份证号
	RegularIDCardMsg  string = "不是一个有效的身份证号"
	RegularIDCardXMsg string = "身份证号最后一位必须是大写的X"

	// 银行卡
	RegularBankCard    string = `^([0-9]{1}[0-9]{6,22})$`
	RegularBankCardMsg string = "收款账号格式错误，不是一个有效的银行卡号"

	// 银行账号名称
	RegularBankAccountName    string = "^[·\u4e00-\u9fa5]{2,20}"
	RegularBankAccountNameMsg string = "账号名称错误，只能是中文，点号分隔符·，且须2~20个字符"

	// 开户行
	RegularBankOfDeposit    string = "^[A-Z\u4e00-\u9fa5]{4,20}$"
	RegularBankOfDepositMsg string = "开户行错误，只能是中文，大写字母，且须4~20个字符"

	// 联系方式(座机号码或者手机号【目前紧支持大陆，港澳台地区】正则)
	RegularTel    string = "^((([0][0-9]{2,4}-*)*[0-9]{5,8}(-[0-9]{1,4})*)|(1[3-9][0-9]{9})|(([6|9])[0-9]{7})|([0][9][0-9]{8})|([6]([8|6])[0-9]{5}))$"
	RegularTelMsg string = "不是一个有效的电话号码"

	// 联系地址
	RegularAddress    string = "^[a-zA-Z0-9\\_\\-()（）\\[\\]【】<>/|:'\".·,，\u4e00-\u9fa5]{10,100}$"
	RegularAddressMsg string = "联系地址错误，只能是中文，英文字母，数字，或符号_-()（）[]【】<>/|:'\".·,，且须10~100个字符"
)
