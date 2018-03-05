package gopathlib

import (
  "os"
  "strings"
)

type Options struct {
  sep string
}

type PathStruct struct{
  segments []string
  *Options
}

func (this *PathStruct) P(segments ...string) (*PathStruct) {
  for _, seg := range segments {
    if seg != "" {
      this.segments = append(this.segments, seg)
    }
  }
  return this
}

func (this *PathStruct) String() string {
  return strings.Join(this.segments, this.sep)
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
