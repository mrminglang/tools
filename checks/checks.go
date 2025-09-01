package checks

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/bluesky335/IDCheck/IdNumber"
	"github.com/mrminglang/tools/checks/regex"
)

// RegexCheckIDCard 身份证号正则校验
func RegexCheckIDCard(idCard string) (err error) {
	if strings.Contains(idCard, "x") == true {
		err = errors.New(regex.RegularIDCardXMsg)
		return
	}
	id := IDNumber.New(idCard)
	if !id.IsValid() {
		err = errors.New(regex.RegularIDCardMsg)
	}
	return
}

// RegexCheckPhone 手机号正则校验
func RegexCheckPhone(phone string) (err error) {
	reg := regexp.MustCompile(regex.RegularPhone)
	if ok := reg.MatchString(phone); !ok {
		err = errors.New(regex.RegularPhoneMsg)
	}
	return
}

// RegexCheckTel 联系方式正则校验（座机号码或者手机号）
func RegexCheckTel(tel string) (err error) {
	reg := regexp.MustCompile(regex.RegularTel)
	if ok := reg.MatchString(tel); !ok {
		err = errors.New(regex.RegularTelMsg)
	}
	return
}

// RegexCheckPayeeType 收款类型校验EN
func RegexCheckPayeeType(payeeType string) (identityType string) {
	switch payeeType {
	case regex.PayeeTypeBankCN:
		identityType = regex.PayeeTypeBankEN
	case regex.PayeeTypeAliCN:
		identityType = regex.PayeeTypeAliEN
	default:
		identityType = ""
	}
	return identityType
}

// RegexCheckIdentityTypeCN 收款类型校验CN
func RegexCheckIdentityTypeCN(payeeType string) (identityType string) {
	switch payeeType {
	case regex.PayeeTypeBankEN:
		identityType = regex.PayeeTypeBankCN
	case regex.PayeeTypeAliEN:
		identityType = regex.PayeeTypeAliCN
	default:
		identityType = ""
	}
	return identityType
}

// RegexCheckIdentity 收款账号正则校验
func RegexCheckIdentity(identityType string, identity string) (err error) {
	switch identityType {
	case regex.PayeeTypeBankEN:
		err = RegexCheckBankFormat(identity)
	case regex.PayeeTypeAliEN:
		err = RegexCheckAliFormat(identity)
	default:
		err = errors.New(regex.RegularIdentityMsg)
	}

	return
}

// RegexCheckBankFormat 收款账号为银行卡
func RegexCheckBankFormat(identity string) (err error) {
	if len(identity) == 11 { // 手机号
		reg := regexp.MustCompile(regex.RegularPhone)
		if ok := reg.MatchString(identity); ok {
			err = errors.New(regex.RegularBankCardMsg)
			return
		}
	}

	reg := regexp.MustCompile(regex.RegularBankCard)
	if ok := reg.MatchString(identity); !ok {
		err = errors.New(regex.RegularBankCardMsg)
	}
	return
}

// RegexCheckAliFormat 收款账号为支付宝
func RegexCheckAliFormat(identity string) (err error) {
	if strings.Contains(identity, "@") { // 电子邮箱
		if err = RegexCheckEmail(identity); err != nil {
			err = errors.New(fmt.Sprintf("%s，%s", regex.RegularIdentityMsg, err.Error()))
			return
		}
	} else if len(identity) == 16 { // 支付宝用户ID
		if err = RegexCheckAliUserID(identity); err != nil {
			return
		}
	} else if len(identity) == 11 { // 手机号
		if err = RegexCheckPhone(identity); err != nil {
			err = errors.New(fmt.Sprintf("%s，%s", regex.RegularIdentityMsg, err.Error()))
			return
		}
	} else {
		err = errors.New(regex.RegularIdentityMsg)
		return
	}

	return
}

// RegexCheckAliUserID 支付宝用户ID正则校验
func RegexCheckAliUserID(userID string) (err error) {
	reg := regexp.MustCompile(regex.RegularAlipayUserID)
	if ok := reg.MatchString(userID); !ok {
		err = errors.New(regex.RegularAlipayUserIDMsg)
	}
	return
}

// RegexCheckEmail 邮箱正则校验
func RegexCheckEmail(email string) (err error) {
	if !govalidator.IsEmail(email) {
		err = errors.New(regex.RegularEmailMsg)
	}
	return
}

// RegexCheckName 姓名正则校验
func RegexCheckName(name string) (err error) {
	reg := regexp.MustCompile(regex.RegularName)
	if ok := reg.MatchString(name); !ok {
		err = errors.New(regex.RegularNameMsg)
	}
	return
}

// RegexCheckIsRequestURL URL正则校验
func RegexCheckIsRequestURL(url string) (err error) {
	if !govalidator.IsRequestURL(url) {
		err = errors.New(regex.RegularUrlMsg)
	}
	return
}

// RegexCheckURL URL正则校验
func RegexCheckURL(url string) (err error) {
	regularUrl := fmt.Sprintf("^%s$", regex.RegularUrl)
	reg := regexp.MustCompile(regularUrl)
	if ok := reg.MatchString(url); !ok {
		err = errors.New(regex.RegularUrlMsg)
	}
	return
}

// RegexCheckAddress 联系地址
func RegexCheckAddress(address string) (err error) {
	reg := regexp.MustCompile(regex.RegularAddress)
	if ok := reg.MatchString(address); !ok {
		err = errors.New(regex.RegularAddressMsg)
	}
	return
}

// RegexCheckBankAccountName 银行账户名称
func RegexCheckBankAccountName(bankAccountName string) (err error) {
	reg := regexp.MustCompile(regex.RegularBankAccountName)
	if ok := reg.MatchString(bankAccountName); !ok {
		err = errors.New(regex.RegularBankAccountNameMsg)
	}
	return
}

// RegexCheckBankOfDeposit 开户行
func RegexCheckBankOfDeposit(bankOfDeposit string) (err error) {
	reg := regexp.MustCompile(regex.RegularBankOfDeposit)
	if ok := reg.MatchString(bankOfDeposit); !ok {
		err = errors.New(regex.RegularBankOfDepositMsg)
	}
	return
}
