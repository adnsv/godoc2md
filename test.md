

# sourcerer


* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## Installation
	go get bitbucket.org/tessonics/sourcerer

## <a name="pkg-overview">Overview</a>
Package sourcerer is a small library which helps writing automated source code generators
for different languages.

The key features are:
- Support for flexible tabs (via text/tabwriter).
- Automatic management of indentation levels.
- Advanced handling of EOLs and empty lines between blocks of code
- Support for multiple buffers (e.g. writing .hpp/.cpp file pairs)


### Usage Example
```GO
buf1 := os.Stdout
buf2 := &strings.Builder{}

w := Writer{}
w.AddBuffer("buf1", buf1, 8)
w.AddBuffer("buf2", buf2, 8)

w.P("#pragma once").NL(2)
w.B("namespace test {", func() {
    w.P("indented")
}, "}")

// blah
fmt.Println("--- content of second buffer: ---")
fmt.Print(buf2.String())
// Output: Hello
```


## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [type Writer](#Writer)
  * [func (w *Writer) AddBuffer(name string, output io.Writer, tabwidth int) error](#Writer.AddBuffer)
  * [func (w *Writer) B(before string, f func(), after string) *Writer](#Writer.B)
  * [func (w *Writer) F(format string, a ...interface{}) *Writer](#Writer.F)
  * [func (w *Writer) FlushAll() error](#Writer.FlushAll)
  * [func (w *Writer) FlushCurrent() error](#Writer.FlushCurrent)
  * [func (w *Writer) I(f func()) *Writer](#Writer.I)
  * [func (w *Writer) NL(numEols int) *Writer](#Writer.NL)
  * [func (w *Writer) P(a ...interface{}) *Writer](#Writer.P)
  * [func (w *Writer) Print(a ...interface{}) (n int, err error)](#Writer.Print)
  * [func (w *Writer) Printf(format string, a ...interface{}) (n int, err error)](#Writer.Printf)
  * [func (w *Writer) SelectBuffer(name string) error](#Writer.SelectBuffer)
  * [func (w *Writer) Write(buf []byte) (n int, err error)](#Writer.Write)

#### <a name="pkg-examples">Examples</a>
* [Package](#example_)
* [Writer](#example_Writer)

#### <a name="pkg-files">Package files</a>
[doc.go](/src/bitbucket.org/tessonics/sourcerer/doc.go) [writer.go](/src/bitbucket.org/tessonics/sourcerer/writer.go) 



## <a name="pkg-variables">Variables</a>
``` go
var ErrMissingBuffer = errors.New("missing output buffer")
```
ErrMissingBuffer is returned when trying to write to a writer that
does not have a buffer selected for output.




## <a name="Writer">type</a> [Writer](/src/target/writer.go?s=340:410#L15)
``` go
type Writer struct {
    // contains filtered or unexported fields
}

```
Writer supports writing formatted indented text into one or more buffers.



### Example
```GO
w := Writer{}
w.FlushAll()
fmt.Println("Hello")
```







### <a name="Writer.AddBuffer">func</a> (\*Writer) [AddBuffer](/src/target/writer.go?s=558:635#L27)
``` go
func (w *Writer) AddBuffer(name string, output io.Writer, tabwidth int) error
```
AddBuffer adds a new buffer, to which the writer can redirect its output.




### <a name="Writer.B">func</a> (\*Writer) [B](/src/target/writer.go?s=3299:3364#L128)
``` go
func (w *Writer) B(before string, f func(), after string) *Writer
```
B writes a block of code in which lines produced within the callback
scope are automatically indented:


	BeforeContent + [EOL]
	  Line #1 produced from within the f() call
	  Line #2 produced from within the f() call
	AfterContent + [EOL]




### <a name="Writer.F">func</a> (\*Writer) [F](/src/target/writer.go?s=2687:2746#L104)
``` go
func (w *Writer) F(format string, a ...interface{}) *Writer
```
F provides formatted printing with support for chaining.




### <a name="Writer.FlushAll">func</a> (\*Writer) [FlushAll](/src/target/writer.go?s=1524:1557#L63)
``` go
func (w *Writer) FlushAll() error
```
FlushAll flushes all the contained buffers.




### <a name="Writer.FlushCurrent">func</a> (\*Writer) [FlushCurrent](/src/target/writer.go?s=1333:1370#L55)
``` go
func (w *Writer) FlushCurrent() error
```
FlushCurrent flushes the currently active buffer.




### <a name="Writer.I">func</a> (\*Writer) [I](/src/target/writer.go?s=3575:3611#L142)
``` go
func (w *Writer) I(f func()) *Writer
```
I provides an indented scope for lines written within the callback function.




### <a name="Writer.NL">func</a> (\*Writer) [NL](/src/target/writer.go?s=2869:2909#L110)
``` go
func (w *Writer) NL(numEols int) *Writer
```
NL sets the number of EOLs to be emitted before the next non-eol content.




### <a name="Writer.P">func</a> (\*Writer) [P](/src/target/writer.go?s=2546:2590#L98)
``` go
func (w *Writer) P(a ...interface{}) *Writer
```
P provides non-formatted printing with support for chaining.




### <a name="Writer.Print">func</a> (\*Writer) [Print](/src/target/writer.go?s=2008:2067#L83)
``` go
func (w *Writer) Print(a ...interface{}) (n int, err error)
```
Print outputs non-formatted content, similar to standard fmt.Println.
Notice that the terminating EOL can be cancelled with NL(0).




### <a name="Writer.Printf">func</a> (\*Writer) [Printf](/src/target/writer.go?s=2353:2428#L93)
``` go
func (w *Writer) Printf(format string, a ...interface{}) (n int, err error)
```
Printf outputs formatted content, similar to standard fmt.Printf.
Unlike the standard Printf, however, it automaticall adds terminating EOL,
which can be cancelled with NL(0).




### <a name="Writer.SelectBuffer">func</a> (\*Writer) [SelectBuffer](/src/target/writer.go?s=1069:1117#L43)
``` go
func (w *Writer) SelectBuffer(name string) error
```
SelectBuffer chooses which of the previously added buffers to use for output.




### <a name="Writer.Write">func</a> (\*Writer) [Write](/src/target/writer.go?s=1716:1769#L74)
``` go
func (w *Writer) Write(buf []byte) (n int, err error)
```
Write implements io.Writer interface.









- - -
Generated by [godoc2md](http://godoc.org/github.com/adnsv/godoc2md)
