package gopathlib

import (
  "fmt"
  "os"
  "strings"
)

type Options struct {
  sep string
}

type PathStruct struct{
  segments []string
  addRootSlash bool // default is false
  *Options
}

func (this *PathStruct) P(segments ...string) (*PathStruct) {
  if len(this.segments) == 0 && len(segments) > 0{
    if string(segments[0][0]) == this.sep {
      this.addRootSlash = true
    }
  }
  for _, seg := range segments {
    // Remove spurious slashes and single dot
    for _, seg2 := range strings.Split(seg, this.sep) {
      if seg2 != "" && seg2 != "." {
        this.segments = append(this.segments, seg2)
      }
    }
  }
  return this
}

func (this *PathStruct) String() string {
  out := strings.Join(this.segments, this.sep)
  if this.addRootSlash {
    out = fmt.Sprintf("%s%s", this.sep, out)
  }
  return out
}

func Pwd() (*PathStruct) {
  this := New()
  seg, _ := os.Getwd()
  this.segments = append(this.segments, seg)
  return this
}

func New(segments ...string) *PathStruct{
  return New2(nil, segments...)
}

func New2(options *Options, segments ...string) *PathStruct{
  var opts *Options = options
  if opts == nil { // Use defaults
    opts = &Options{sep: "/"}
  }
  this := &PathStruct{Options: opts}
  this.P(segments...)
  return this
}
