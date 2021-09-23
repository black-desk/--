# Position test

```bash
git clone https://github.com/black-desk/lsp-position-test
cd lsp-position-test
go build
nvim ./README.md
```

a𐐀b

> A position inside a document (see Position definition below) is expressed as a zero-based line and character offset. The offsets are based on a UTF-16 string representation. So a string of the form a𐐀b the character offset of the character a is 0, the character offset of 𐐀 is 1 and the character offset of b is 3 since 𐐀 is represented using two code units in UTF-16.

https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocuments
