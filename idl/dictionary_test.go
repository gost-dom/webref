package idl_test

import (
	"testing"

	"github.com/gost-dom/webref/idl"
	"github.com/stretchr/testify/suite"
)

type DictionaryTestSuite struct {
	suite.Suite
	fetch    idl.Spec
	uievents idl.Spec
}

func (s *DictionaryTestSuite) SetupTest() {
	var err error
	s.fetch, err = idl.Load("fetch")
	s.Assert().NoError(err)
	s.uievents, err = idl.Load("uievents")
	s.Assert().NoError(err)
}

func (s *DictionaryTestSuite) TestRequestInitDictionary() {
	requestInit := s.fetch.Dictionaries["RequestInit"]
	method, ok := requestInit.Get("method")
	s.Assert().True(ok, "Dictionary has entry: method")
	s.Assert().Equal("method", method.Key)
	s.Assert().Equal("ByteString", method.Value.Name)
}

func (s *DictionaryTestSuite) TestDictionaryInheritance() {
	keybaordEventInit := s.uievents.Dictionaries["KeyboardEventInit"]
	s.Assert().Equal("EventModifierInit", keybaordEventInit.Inheritance)
}

func TestDictionary(t *testing.T) {
	suite.Run(t, new(DictionaryTestSuite))
}
