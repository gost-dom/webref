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
	html Spec
	url  Spec
	dom  Spec
}

func (s *IdlInterfacesTestSuite) SetupSuite() {
	var err error
	s.dom, err = Load("dom")
	s.Assert().NoError(err)
	s.html, err = Load("html")
	s.Assert().NoError(err)
	s.url, err = Load("url")
	s.Assert().NoError(err)
}

func (s *IdlInterfacesTestSuite) SetupTest() {
	s.Gomega = gomega.NewWithT(s.T())
}

func (s *IdlInterfacesTestSuite) TestConstructor() {
	element := s.dom.Interfaces["Element"]
	s.Expect(element.Constructors).To(BeEmpty())

	text := s.dom.Interfaces["Text"]
	s.Expect(text.Constructors).To(HaveExactElements(
		Constructor{Arguments: []Argument{{
			Name: "data",
			Type: Type{
				Kind:      KindSimple,
				Name:      "DOMString",
				Nullable:  false,
				TypeParam: nil,
			},
			Variadic: false,
			Optional: true,
		}}},
	))
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

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(IdlInterfacesTestSuite))
}

func BeAStringifier() types.GomegaMatcher {
	return WithTransform(
		func(a Attribute) bool { return a.Stringifier },
		BeTrue())
}
