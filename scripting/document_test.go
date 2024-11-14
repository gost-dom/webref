package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("V8 Document", func() {
	ctx := InitializeContextWithEmptyHtml()

	itShouldBeADocument := func() {
		It("Should be an instance of Document", func() {
			Expect(ctx.RunTestScript("actual instanceof Document")).To(BeTrue())
		})
		It("Should be an instance of Node", func() {
			Expect(ctx.RunTestScript("actual instanceof Node")).To(BeTrue())
		})
		It("Should be an instance of EventTarget", func() {
			Expect(ctx.RunTestScript("actual instanceof EventTarget")).To(BeTrue())
		})
		It("Should be an instance of Object", func() {
			Expect(ctx.RunTestScript("actual instanceof Object")).To(BeTrue())
		})
		It("Should have a class hierarchy of 4 classes", func() {
			Expect(ctx.RunTestScript(`
        let baseClassCount = 0
        let current = actual
        while(current = Object.getPrototypeOf(current))
          baseClassCount++
        baseClassCount;
      `)).To(BeEquivalentTo(4))
		})
	}

	Describe("Class Hierarchy of new Document()", func() {
		BeforeEach(func() {
			ctx.RunTestScript("const actual = new Document()")
		})
		itShouldBeADocument()
	})

	Describe("Class Hierarchy of `window.document`", func() {
		BeforeEach(func() {
			ctx.RunTestScript("const actual = window.document")
		})
		itShouldBeADocument()
	})

	Describe("Constructor", func() {
		It("Should be instance of Document", func() {
			Expect(ctx.RunTestScript(`
        const doc = new Document();
        doc instanceof Document && doc != document;
      `)).To(BeTrue())
		})

		It("Should have `createElement` as a function", func() {
			Expect(
				ctx.RunTestScript(`typeof (new Document().createElement)`),
			).To(Equal("function"))
		})

		It("Should support Document functions", func() {
			Skip("createElement and HTMLElement are missing")
			Expect(
				ctx.RunTestScript(`document.createElement("div") instanceof HTMLElement`),
			).Error().
				ToNot(HaveOccurred())
		})
	})
})
