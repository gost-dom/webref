package idl_test

import (
	"testing"

	. "github.com/gost-dom/webref/idl"

	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

type IdlInterfacesTestSuite struct {
	suite.Suite
	gomega.Gomega
	html Spec
	url  Spec
}

func (s *IdlInterfacesTestSuite) SetupSuite() {
	var err error
	s.html, err = Load("html")
	s.Assert().NoError(err)
	s.url, err = Load("url")
	s.Assert().NoError(err)
}

func (s *IdlInterfacesTestSuite) SetupTest() {
	s.Gomega = gomega.NewWithT(s.T())
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

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(IdlInterfacesTestSuite))
}
