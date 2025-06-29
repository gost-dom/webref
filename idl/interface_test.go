package idl_test

import (
	"testing"

	"github.com/gost-dom/webref/idl"
	. "github.com/gost-dom/webref/idl"

	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"github.com/stretchr/testify/suite"
)

type IdlInterfacesTestSuite struct {
	suite.Suite
	gomega.Gomega
	html    Spec
	url     Spec
	dom     Spec
	xhr     Spec
	fetch   Spec
	streams Spec
}

func (s *IdlInterfacesTestSuite) SetupSuite() {
	var err error
	s.dom, err = Load("dom")
	s.Assert().NoError(err)
	s.html, err = Load("html")
	s.Assert().NoError(err)
	s.url, err = Load("url")
	s.Assert().NoError(err)
	s.xhr, err = Load("xhr")
	s.Assert().NoError(err)
	s.fetch, err = Load("fetch")
	s.Assert().NoError(err)
	s.streams, err = Load("streams")
	s.Assert().NoError(err)
}

func (s *IdlInterfacesTestSuite) SetupTest() {
	s.Gomega = gomega.NewWithT(s.T())
}

func (s *IdlInterfacesTestSuite) TestConstructor() {
	element := s.dom.Interfaces["Element"]
	s.Expect(element.Constructors).To(BeEmpty())

	text := s.dom.Interfaces["Text"]
	s.Expect(text.Constructors).To(HaveLen(1))
	cons := text.Constructors[0]
	s.Expect(cons.Arguments).To(HaveLen(1))
	arg := cons.Arguments[0]
	s.Expect(arg.Name).To(Equal("data"))
	s.Expect(arg.Variadic).To(BeFalse())
	s.Expect(arg.Optional).To(BeTrue())
	s.Expect(arg.Type).To(Equal(Type{
		Kind:      KindSimple,
		Name:      "DOMString",
		Nullable:  false,
		TypeParam: nil,
	}))
}

func (s *IdlInterfacesTestSuite) TestAnchorIncludeHyperlinkUtils() {
	actual := s.html.Interfaces["HTMLAnchorElement"].Includes
	s.Expect(actual).To(ContainElement(HaveField("Name", "HTMLHyperlinkElementUtils")))
}

func (s *IdlInterfacesTestSuite) TestUrlHasToJSONMethod() {
	ops := s.url.Interfaces["URL"].Operations
	s.Expect(ops).To(ContainElement(HaveField("Name", "toJSON")))
}

func (s *IdlInterfacesTestSuite) TestUrlParseIsStatic() {
	ops := s.url.Interfaces["URL"].Operations
	s.Expect(ops).To(
		ContainElement(SatisfyAll(
			HaveField("Name", "parse"),
			HaveField("Static", true)),
		))
}

func (s *IdlInterfacesTestSuite) TestNodeInheritsFromEventTarget() {
	node, ok := s.dom.Interfaces["Node"]
	s.Assert().True(ok)
	s.Assert().Equal("EventTarget", node.Inheritance)
}

func (s *IdlInterfacesTestSuite) TestHTMLHyperlinkElementUtilsIsAMixingButAnchorIsNot() {
	anchorIsAMixin := s.html.Interfaces["HTMLAnchorElement"].Mixin
	s.Assert().False(anchorIsAMixin, "Anchor is a mixin")

	hyperLinkUtilsIsAMixin := s.html.Interfaces["HTMLHyperlinkElementUtils"].Mixin
	s.Assert().True(hyperLinkUtilsIsAMixin, "HyperlinkElementUtils is a mixin")
}

func (s *IdlInterfacesTestSuite) TestAttributeTypeOnHTMLCollection() {
	intf := s.dom.Interfaces["HTMLCollection"]
	op, found := intf.GetOperation("item")
	s.True(found)
	s.Expect(op.Arguments[0].Type).To(Equal(idl.Type{
		Name:     "unsigned long",
		Nullable: false,
	}))
	s.Expect(op).ToNot(BeNil())
}

func (s *IdlInterfacesTestSuite) TestParetNodeVeriadicArguments() {
	intf := s.dom.Interfaces["ParentNode"]
	op, found := intf.GetOperation("append")
	s.True(found)
	s.Expect(op.Arguments[0].Variadic).To(BeTrue())

	op, found = intf.GetOperation("querySelector")
	s.True(found)
	s.Expect(op.Arguments[0].Variadic).To(BeFalse())
}

func (s *IdlInterfacesTestSuite) TestUrlHasStringifierOnHref() {
	assert := s.Assert()
	intf := s.url.Interfaces["URL"]

	attr, found := intf.GetAttribute("href")
	assert.True(found)
	assert.True(attr.Stringifier, "URL.href is a stringifier")

	attr, found = intf.GetAttribute("host")
	assert.True(found)
	assert.False(attr.Stringifier, "URL.host is a stringifier")
}

func (s *IdlInterfacesTestSuite) TestUrlSearchParamsHasStringifierOnEmptyFunction() {
	assert := s.Assert()

	intf, found := s.url.Interfaces["URLSearchParams"]
	assert.True(found)

	attr, found := intf.GetOperation("")
	assert.True(found)
	assert.True(attr.Stringifier, "URLSearchParams has a toString()")

	attr, found = intf.GetOperation("append")
	assert.True(found)
	assert.False(attr.Stringifier, "URL.appand is a stringifier")
}

func (s *IdlInterfacesTestSuite) TestSequenceReturnValue() {
	assert := s.Assert()

	intf, found := s.url.Interfaces["URLSearchParams"]
	assert.True(found)

	get, found := intf.GetOperation("get")
	assert.Equal(KindSimple, get.ReturnType.Kind)
	assert.Equal("USVString", get.ReturnType.Name)

	getAll, found := intf.GetOperation("getAll")
	assert.Equal(KindSequence, getAll.ReturnType.Kind, "getAll is a sequence")
	assert.Equal("USVString", getAll.ReturnType.TypeParam.Name)
}

func (s *IdlInterfacesTestSuite) TestOptionalArgs() {
	assert := s.Assert()

	intf, found := s.url.Interfaces["URLSearchParams"]
	assert.True(found)

	op, found := intf.GetOperation("has")
	s.Assert().True(found)
	s.Assert().False(op.Arguments[0].Optional, "URLSearchParams.has - first argument optional")
	s.Assert().True(op.Arguments[1].Optional, "URLSearchParams.has - second argument optional")
}

func (s *IdlInterfacesTestSuite) TestNodeInsertBefore() {
	assert := s.Assert()

	intf, found := s.dom.Interfaces["Node"]
	assert.True(found)

	op, found := intf.GetOperation("insertBefore")
	s.Assert().False(op.Arguments[0].Type.Nullable, "Node.insertBefore - first argument nullable")
	s.Assert().True(op.Arguments[1].Type.Nullable, "Node.insertBefore - second argument nullable")
}

func (s *IdlInterfacesTestSuite) TestIterable() {
	assert := s.Assert()

	url, foundURL := s.url.Interfaces["URL"]
	assert.True(foundURL)
	assert.Len(url.IterableTypes, 0)

	usp, foundUSP := s.url.Interfaces["URLSearchParams"]
	assert.True(foundUSP)
	assert.Len(usp.IterableTypes, 2)
	assert.Equal("USVString", usp.IterableTypes[0].Name)
	assert.Equal("USVString", usp.IterableTypes[1].Name)
}

func (s *IdlInterfacesTestSuite) TestUnionType() {
	assert := s.Assert()

	intf, found := s.xhr.Interfaces["XMLHttpRequest"]
	assert.True(found)

	send, found := intf.GetOperation("send")
	assert.True(found, "XHR has a send method")
	assert.Equal(KindUnion, send.Arguments[0].Type.Kind)
	assert.True(send.Arguments[0].Type.Nullable)
	assert.Equal("Document", send.Arguments[0].Type.Types[0].Name)
	assert.Equal("XMLHttpRequestBodyInit", send.Arguments[0].Type.Types[1].Name)
}

func (s *IdlInterfacesTestSuite) TestFetchResponse() {
	assert := s.Assert()

	response, found := s.fetch.Interfaces["Response"]
	assert.True(found)

	args := response.Constructors[0].Arguments
	assert.NotNil(args[0].Default)
	assert.Equal("null", string(args[0].Default.Type))
}

func (s *IdlInterfacesTestSuite) TestReadReturnsAPromise() {
	assert := s.Assert()
	reader, ok := s.streams.Interfaces["ReadableStreamDefaultReader"]
	assert.True(ok)

	read, ok := reader.GetOperation("read")
	assert.Equal(idl.KindPromise, read.ReturnType.Kind, "Reader.read returns a promise")
	typeParam := read.ReturnType.TypeParam
	assert.Equal(idl.KindSimple, typeParam.Kind)
	assert.Equal("ReadableStreamReadResult", typeParam.Name)
}
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(IdlInterfacesTestSuite))
}

func BeAStringifier() types.GomegaMatcher {
	return WithTransform(
		func(a Attribute) bool { return a.Stringifier },
		BeTrue())
}
