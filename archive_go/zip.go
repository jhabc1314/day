package archive_go

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"time"
)

//提供zip包的压缩和解压

//该软件包不支持磁盘扩展。

//压缩方法
const (
	Store uint16 = 0 //不压缩
	Deflate uint16 = 8 //使用deflate算法压缩
)

var (
	ErrFormat = errors.New("zip:not a valid zip file")
	ErrAlgorithm = errors.New("zip:unsupported compression algorithm")
	ErrCheckSum = errors.New("zip:check sum error")
)

//RegisterCompressor为指定的方法ID注册自定义压缩器。内置了常用的Store和Deflate方法。
//func RegisterCompressor(method uint16, comp Compressor)

//RegisterDecompressor允许为指定的方法ID自定义解压缩器。内置了常用的Store和Deflate方法。
//func RegisterDecompressor(method uint16, dcomp Decompressor)

//Compressor 返回一个新的压缩编写器，写入w。 必须使用WriteCloser的Close方法将待处理数据刷新到w。 Compressor本身必须能够安全地同时从多个goroutine调用，但是每个返回的writer一次只能由一个goroutine使用。
type Compressor func(w io.Writer)(io.WriteCloser, error)

//Decompressor 解压缩器返回一个新的解压缩读取器，从r读取。 必须使用ReadCloser的Close方法来释放关联的资源。 同时从多个goroutine调用Decompressor本身必须安全，但是每个返回的读取器一次只能由一个goroutine使用。
type Decompressor func(r io.Reader) io.ReadCloser

type File struct {
	FileHeader
}
//DataOffset返回文件的可能压缩数据相对于zip文件开头的偏移量。 相反，大多数调用者应该使用Open，它可以透明地解压缩数据并验证校验和。
//func (f *File) DataOffset() (offset int64, err error)

//Open返回一个ReadCloser，它提供对文件内容的访问。可以同时读取多个文件。
//func (f *File) Open() (io.ReadCloser, error)

//FileHeader 描述zip文件中的文件。有关详细信息，请参见zip规范。
type FileHeader struct {
	Name string //文件名 必须是相对路径 必须使用正斜杠不是反斜杠
	Comment string //文件注释，用户定义
	NonUTF8 bool //表示名称和注释不适用utf8编码 Go1.10

	CreatorVersion uint16
	ReaderVersion uint16
	Flags uint16

	Method uint16 //压缩方法 如果是0，默认使用Store

	Modified time.Time //修改时间 Go1.10
	ModifiedTime uint16 //修改时间 不推荐使用，使用上面的代替
	ModifiedDate uint16 //修改时间 不推荐使用，使用上面的代替

	CRC32 uint32 //循环冗余校验散列函数

	CompressedSize uint32 //不推荐
	UncompressedSize uint32 //不推荐
	CompressedSize64 uint64 //GO1.1 推荐
	UncompressedSize64 uint64 //Go1.1 推荐
	Extra []byte //
	ExternalAttrs uint32 //含义取决于CreatorVersion 
}

//FileInfoHeader从os.FileInfo创建一个部分填充的FileHeader。 由于os.FileInfo的Name方法仅返回其描述的文件的基本名称，因此可能有必要修改返回的标头的Name字段以提供文件的完整路径名。 如果需要压缩，则调用者应设置FileHeader.Method字段； 默认情况下未设置。
//func FileInfoHeader(fi os.FileInfo) (*FileHeader, error)

//FileInfo returns an os.FileInfo for the FileHeader.
//func (h *FileHeader) FileInfo() os.FileInfo

//不推荐使用
//func (h *FileHeader) ModTime() time.Time

//Mode返回FileHeader的权限和模式位。
//func (h *FileHeader) Mode() (mode os.FileMode)

//SetModTime将Modified，ModifiedTime和ModifiedDate字段设置为UTC中的给定时间。
//func (h *FileHeader) SetModTime(t time.Time)

//SetMode更改FileHeader的权限和模式位。
//func (h *FileHeader) SetMode(mode os.FileMode)

// type ReadCloser struct {
// 	Reader
// }

//OpenReader将打开按名称指定的Zip文件，并返回ReadCloser。
//func OpenReader(name string) (*ReadCloser, error)

//Close 关闭Zip文件，使其无法用于I / O。.
//func (rc *ReadCloser) Close() error



// type Reader struct {
// 	File []*File
// 	Comment string
// 	//等等一堆未导出的
// }

//NewReader从r返回一个新的Reader读数，假定具有给定的字节大小。
//func NewReader(r io.ReaderAt, size int64) (*Reader, error)

//RegisterDecompressor为特定的方法ID注册或覆盖自定义解压缩器。如果找不到给定方法的解压缩器，则Reader将默认在包级别查找该解压缩器。
//func (z *Reader) RegisterDecompressor(method uint16, dcomp Decompressor) GO1.6

//Writer实现一个zip文件writer。
// type Writer struct {

// }
//NewWriter返回一个新的Writer，将一个zip文件写入w。
//func NewWriter(w io.Writer) *Writer

//Close通过写入中央目录来完成zip文件的写入。它不会关闭基础编写器。
//func (w *Writer) Close() error

//Create使用提供的名称将文件添加到zip文件中。 它返回应该将文件内容写入的Writer。 文件内容将使用Deflate方法压缩。 该名称必须是相对路径：不得以驱动器号（例如C :）或前斜杠开头，并且只能使用正斜杠。 要创建目录而不是文件，请在名称后添加斜杠。 在下一次调用Create，CreateHeader或Close之前，必须将文件的内容写入io.Writer。
//func (w *Writer) Create(name string) (io.Writer, error)

//CreateHeader使用提供的FileHeader作为文件元数据将文件添加到zip存档中。 Writer 拥有fh的所有权，并可以对其字段进行更改。 调用CreateHeader后，调用者不得修改fh。
//这将返回一个Writer，应该将文件内容写入其中。 在下一次调用Create，CreateHeader或Close之前，必须将文件的内容写入io.Writer。
//func (w *Writer) CreateHeader(fh *FileHeader) (io.Writer, error)


//RegisterCompressor注册或覆盖用于特定方法ID的自定义压缩器。 如果找不到用于给定方法的压缩器，Writer将默认在包级别查找该压缩器。
//func (w *Writer) RegisterCompressor(method uint16, comp Compressor) GO1.6

//刷新将所有缓冲的数据刷新到基础写入器。通常不需要调用Flush。调用Close就足够了。
//func (w *Writer) Flush() error GO1.4

//
//func (w *Writer) SetComment(comment string) error GO1.10

//SetOffset设置基础编写器中zip数据开头的偏移量。将zip数据附加到现有文件（例如二进制可执行文件）时，应使用该文件。在写入任何数据之前必须先调用它。
//func (w *Writer) SetOffset(n int64) GO1.5


//Zip demo
func Zip() {
	//创建一个zip文件
	z,err := os.Create("./output.zip")
	if err != nil {
		panic(err)
	}
	//基于z创建一个新的写对象实例
	w := zip.NewWriter(z)
	defer w.Close()
	fileList := [2]string{"./archive_go/tar_file1.txt", "./archive_go/tar_file2.log"}

	for _,path := range fileList {
		//打开文件
		fo,err := os.Open(path)
		if err != nil {
			panic(err)
		}
		//获取文件信息
		fileInfo,err := fo.Stat()
		if err != nil {
			panic(err)
		}
		//基于文件头创建写入的文件头
		fheader,err := zip.FileInfoHeader(fileInfo)
		//基于文件头创建写入的文件
		zw,err := w.CreateHeader(fheader) //返回一个存储写的io.writer对象，
		if err != nil {
			panic(err)
		}
		//写入文件
		_, err = io.Copy(zw, fo)
		if err != nil {
			panic(err)
		}
	}

}

//Unzip 解压
func Unzip(z string) {
	//打开压缩包
	fo,err := zip.OpenReader(z)
	defer fo.Close()
	if err != nil {
		panic(err)
	}
	//循环解压
	for _,file := range fo.File {
		//判断是否为文件夹
		if file.FileInfo().IsDir() {
			if err = os.MkdirAll(file.Name, file.Mode()); err != nil {
				panic(err)
			}
			continue
		}
		//打开压缩包文件
		fw,err := file.Open()
		defer fw.Close()
		//创建文件
		fr,err := os.OpenFile(file.Name, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			panic(err)
		}
		defer fr.Close()
		//复制文件
		_,err = io.Copy(fr, fw)
	}
}
