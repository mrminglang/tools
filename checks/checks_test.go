package checks_test

import (
	"testing"

	"github.com/mrminglang/tools/checks"
	"github.com/mrminglang/tools/dumps"
	genid "github.com/srlemon/gen-id"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	m.Run()
}

// 身份证正则校验
func TestCheckIDCard(t *testing.T) {
	idCard := genid.NewGeneratorData().IDCard
	dumps.Dump(idCard)
	err := checks.RegexCheckIDCard(idCard)
	assert.Nil(t, err)
	if err != nil {
		dumps.Dump(err.Error())
	}
}

func TestCheckIDCard2(t *testing.T) {
	idCard := "42010219880305003X"
	dumps.Dump(idCard)
	err := checks.RegexCheckIDCard(idCard)
	assert.Nil(t, err)
	if err != nil {
		dumps.Dump(err.Error())
	}
}

// 手机号正则校验
func TestCheckPhone(t *testing.T) {
	//phone := genid.NewGeneratorData().PhoneNum
	phone := genid.NewGeneratorData().IDCard
	dumps.Dump(phone)
	err := checks.RegexCheckPhone(phone)
	assert.Nil(t, err)
	if err != nil {
		dumps.Dump(err.Error())
	}
}

// 联系方式校验
func TestCheckTel(t *testing.T) {
	tel1 := "027-12306"
	err := checks.RegexCheckTel(tel1)
	assert.Nil(t, err)

	tel2 := "027-123067890"
	err = checks.RegexCheckTel(tel2)
	assert.Error(t, err)

	tel3 := "17720503271"
	err = checks.RegexCheckTel(tel3)
	assert.Nil(t, err)

	tel4 := "00852-81700288"
	err = checks.RegexCheckTel(tel4)
	assert.Nil(t, err)

	tel5 := "12306"
	err = checks.RegexCheckTel(tel5)
	assert.Nil(t, err)

	tel6 := "02712306"
	err = checks.RegexCheckTel(tel6)
	assert.Nil(t, err)
}

// 姓名正则校验
func TestRegexCheckName(t *testing.T) {
	name := genid.NewGeneratorData().Name
	dumps.Dump(name)
	err := checks.RegexCheckName(name)
	assert.Nil(t, err)
	if err != nil {
		dumps.Dump(err.Error())
	}
}

// URL正则校验
func TestRegexCheckIsRequestURL(t *testing.T) {
	//url := "https://cx.ukerw.com/uker-psp/download.do?preview=true&fileName=/id_number/1213329723463659522.jpg"
	//url := "ftp://abc.com/id_card/20200910/1250813983174725633.jpg"
	//url := "file://abc.com/id_card/20200910/1250813983174725633.jpg"
	url := "abc.com/id_card/20200910/1250813983174725633.jpg"
	err := checks.RegexCheckIsRequestURL(url)
	assert.Nil(t, err)
	if err != nil {
		dumps.Dump(err.Error())
	}
}

// URL正则校验
func TestRegexCheckURL(t *testing.T) {
	url := "https://cx.ukerw.com/uker-psp/download.do?preview=true&fileName=/id_number/1213329723463659522.jpg"
	//url := "ftp://abc.com/id_card/20200910/1250813983174725633.jpg"
	//url := "file://abc.com/id_card/20200910/1250813983174725633.jpg"
	err := checks.RegexCheckURL(url)
	assert.Nil(t, err)
	if err != nil {
		dumps.Dump(err.Error())
	}
}

// 电子邮箱正则校验
func TestRegexCheckEmail(t *testing.T) {
	email := genid.NewGeneratorData().Email
	dumps.Dump(email)
	err := checks.RegexCheckEmail(email)
	assert.Nil(t, err)
	if err != nil {
		dumps.Dump(err.Error())
	}
}

// 支付宝用户ID正则校验
func TestRegexCheckAliUserID(t *testing.T) {
	userID := "2088123456789012"
	err := checks.RegexCheckAliUserID(userID)
	assert.Nil(t, err)
	if err != nil {
		dumps.Dump(err.Error())
	}
}

// 收款账号为支付宝正则校验
func TestRegexCheckAliFormat(t *testing.T) {
	//identity :=  genid.NewGeneratorData().Email
	identity := genid.NewGeneratorData().PhoneNum
	//identity :=  "2088123456789012"
	dumps.Dump(identity)
	err := checks.RegexCheckAliFormat(identity)
	assert.Nil(t, err)
	if err != nil {
		dumps.Dump(err.Error())
	}
}

// 收款账号为银行卡正则校验
func TestRegexCheckBankFormat(t *testing.T) {
	identity := genid.NewGeneratorData().BankID
	dumps.Dump(identity)
	err := checks.RegexCheckBankFormat(identity)
	assert.Nil(t, err)
	if err != nil {
		dumps.Dump(err.Error())
	}
}

// 收款账号类型
func TestRegexCheckPayeeType(t *testing.T) {
	payeeType := "银行卡"
	identityType := checks.RegexCheckPayeeType(payeeType)
	dumps.Dump(identityType)
}

// 收款账号正则校验：银行卡和支付宝
func TestRegexCheckIdentity(t *testing.T) {
	//payeeType := "银行卡"
	//identityType := checks.RegexCheckPayeeType(payeeType)
	//identity := genid.NewGeneratorData().BankID

	payeeType := "支付宝"
	identityType := checks.RegexCheckPayeeType(payeeType)
	identity := genid.NewGeneratorData().PhoneNum

	dumps.Dump(identity)
	err := checks.RegexCheckIdentity(identityType, identity)
	assert.Nil(t, err)
	if err != nil {
		dumps.Dump(err.Error())
	}
}

// 银行账号名称
func TestCheckBankAccountName(t *testing.T) {
	bankAccountName1 := "明明"
	err := checks.RegexCheckBankAccountName(bankAccountName1)
	assert.Nil(t, err)

	bankAccountName2 := "湖"
	err = checks.RegexCheckBankAccountName(bankAccountName2)
	assert.Error(t, err)
}

// 开户行
func TestCheckBankOfDeposit(t *testing.T) {
	bankOfDeposit1 := "湖北石首农村商业银行"
	err := checks.RegexCheckBankOfDeposit(bankOfDeposit1)
	assert.Nil(t, err)

	bankOfDeposit2 := "AEON信贷财务亚洲有限公司"
	err = checks.RegexCheckBankOfDeposit(bankOfDeposit2)
	assert.Nil(t, err)

	bankOfDeposit3 := "招商"
	err = checks.RegexCheckBankOfDeposit(bankOfDeposit3)
	assert.Error(t, err)
}

// 联系地址校验
func TestCheckAddress(t *testing.T) {
	address1 := "湖北省-武_汉（市）【武】<>[昌]区,，珞珈/山|路16号\"(八一路299号)"
	err := checks.RegexCheckAddress(address1)
	assert.Nil(t, err)

	address2 := "湖北省武汉市江夏区"
	err = checks.RegexCheckAddress(address2)
	assert.Error(t, err)
}
