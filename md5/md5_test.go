package md5

import (
	"testing"
	"fmt"
	"giftone/ico-audit/lib"
)

func TestMd5(t *testing.T) {
	md5:= lib.Md5("000000")
	fmt.Println(len(md5))
	fmt.Println(md5)


}

func TestToken(t *testing.T) {
	md5:= lib.Md5("123456") // 96e79218965eb72c92a549dd5a330112
	fmt.Println(len(md5))
	fmt.Println(md5)
}