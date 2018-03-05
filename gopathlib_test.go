package gopathlib

import (
	"testing"
  "fmt"
	"github.com/dk1027/gotesttools"
  "os"
)

func Test1(t *testing.T) {
	path := New("root")
  fmt.Println(path)
  gotesttools.AssertEqual(t, "root", path.String())
}

func Test2(t *testing.T) {
	path := New("root").P("user")
  fmt.Println(path)
  gotesttools.AssertEqual(t, "root/user", path.String())
}

func Test3(t *testing.T) {
	path := New("root", "user")
  fmt.Println(path)
  gotesttools.AssertEqual(t, "root/user", path.String())
}

func Test4(t *testing.T) {
	path := New("root", "user").P("dk1027")
  fmt.Println(path)
  gotesttools.AssertEqual(t, "root/user/dk1027", path.String())
}

func Test5(t *testing.T) {
	path := Pwd().P("foo")
  fmt.Println(path)
  expected, _ := os.Getwd()
  gotesttools.AssertEqual(t, fmt.Sprintf("%s/foo", expected), path.String())
}

func Test6(t *testing.T) {
  path := New2(&Options{sep: `\`}, "root", "user").P("dk1027")
  fmt.Println(path)
  gotesttools.AssertEqual(t, `root\user\dk1027`, path.String())
}
