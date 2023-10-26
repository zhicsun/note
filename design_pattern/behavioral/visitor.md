# 访问者模式

访问者模式可以给一系列对象透明的添加功能，并且把相关代码封装到一个类中，对象只要预留访问者接口Accept则后期为对象添加功能的时候就不需要改动对象。

````go
package main

import (
	"fmt"
	"path"
)

func main() {
	file, err := NewResourceFile("/root/t.pdf")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = file.Accept(&Compressor{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

type Element interface {
	Accept(Observer) error
}

type Observer interface {
	Visit(Element) error
}

func NewResourceFile(filepath string) (Element, error) {
	switch path.Ext(filepath) {
	case ".ppt":
		return &PPTFile{path: filepath}, nil
	case ".pdf":
		return &PdfFile{path: filepath}, nil
	default:
		return nil, fmt.Errorf("not found file type: %s", filepath)
	}
}

type PdfFile struct {
	path string
}

func (f *PdfFile) Accept(visitor Observer) error {
	return visitor.Visit(f)
}

type PPTFile struct {
	path string
}

func (f *PPTFile) Accept(visitor Observer) error {
	return visitor.Visit(f)
}

type Compressor struct{}

func (c *Compressor) Visit(r Element) error {
	switch f := r.(type) {
	case *PPTFile:
		return c.VisitPPTFile(f)
	case *PdfFile:
		return c.VisitPDFFile(f)
	default:
		return fmt.Errorf("not found resource typr: %#v", r)
	}
}

func (c *Compressor) VisitPPTFile(f *PPTFile) error {
	fmt.Println("this is ppt file")
	return nil
}

func (c *Compressor) VisitPDFFile(f *PdfFile) error {
	fmt.Println("this is pdf file")
	return nil
}

````