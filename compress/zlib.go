package compress

import "compress/flate"

//包zlib实现RFC 1950中指定的读写zlib格式的压缩数据。
//该实现提供了在读取过程中解压缩并在写入过程中压缩的过滤器。例如，将压缩数据写入缓冲区：
// var b bytes.Buffer
// w := zlib.NewWriter(&b)
// w.Write([]byte("hello, world\n"))
// w.Close()

//and to read that data back:
// r, err := zlib.NewReader(&b)
// io.Copy(os.Stdout, r)
// r.Close()

const (
	NoCompression      = flate.NoCompression
	BestSpeed          = flate.BestSpeed
	BestCompression    = flate.BestCompression
	DefaultCompression = flate.DefaultCompression
	HuffmanOnly        = flate.HuffmanOnly
)

var (
	// ErrChecksum is returned when reading ZLIB data that has an invalid checksum.
	ErrChecksum = errors.New("zlib: invalid checksum")
	// ErrDictionary is returned when reading ZLIB data that has an invalid dictionary.
	ErrDictionary = errors.New("zlib: invalid dictionary")
	// ErrHeader is returned when reading ZLIB data that has an invalid header.
	ErrHeader = errors.New("zlib: invalid header")
)

//NewReader创建一个新的ReadCloser。 从返回的ReadCloser中读取数据，然后从r中读取并解压缩数据。 如果r未实现io.ByteReader，则解压缩器可能会从r读取比所需更多的数据。 完成后，调用方有责任在ReadCloser上调用Close。
//NewReader返回的ReadCloser也实现了Resetter。
//func NewReader(r io.Reader) (io.ReadCloser, error)

/*
buff := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207,
		47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
	b := bytes.NewReader(buff)

	r, err := zlib.NewReader(b)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, r)

	r.Close()
*/

//NewReaderDict类似于NewReader，但使用预设字典。 如果压缩数据未引用字典，则NewReaderDict将忽略该字典。 如果压缩数据引用其他字典，则NewReaderDict返回ErrDictionary。

//NewReaderDict返回的ReadCloser也实现了Resetter。
//func NewReaderDict(r io.Reader, dict []byte) (io.ReadCloser, error)

//重置程序重置由NewReader或NewReaderDict返回的ReadCloser，以切换到新的基础Reader。这允许重新使用ReadCloser而不是分配新的ReadCloser。 GO1.4
type Resetter interface {
    // Reset discards any buffered data and resets the Resetter as if it was
    // newly initialized with the given reader.
    Reset(r io.Reader, dict []byte) error
}

//Writer接收写入其中的数据，并将该数据的压缩形式写入基础写入器（请参见NewWriter）。
type Writer struct {
    // contains filtered or unexported fields
}
//NewWriter创建一个新的Writer。对返回的Writer的写入将被压缩并写入w。 完成后，调用方有责任在Writer上调用Close。写入可能会被缓冲，直到关闭才刷新。
//func NewWriter(w io.Writer) *Writer

/** eg
var b bytes.Buffer

	w := zlib.NewWriter(&b)
	w.Write([]byte("hello, world\n"))
	w.Close()
	fmt.Println(b.Bytes())
*/

//func NewWriterLevel(w io.Writer, level int) (*Writer, error)

//func NewWriterLevelDict(w io.Writer, level int, dict []byte) (*Writer, error)

//func (z *Writer) Close() error
//func (z *Writer) Flush() error
//func (z *Writer) Reset(w io.Writer)

//Write将p的压缩形式写入基础io.Writer。在Writer关闭或显式刷新之前，不必刷新压缩字节。
//func (z *Writer) Write(p []byte) (n int, err error)
