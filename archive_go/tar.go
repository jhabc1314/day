package archive_go

import (
	"archive/tar"
	"errors"
	"io"
	"os"
	"time"
)

//tar 包是对tar打包文件的相关功能实现 只打包，不压缩
//使用流程 
//创建一个 .tar 的文件
//定义一个写对象实例
//读取一个被打包文件的头信息
//将头信息写入
//将文件内容写入
//关闭文件，关闭被打包的文件


//常量
const (
	TypeReg  = tar.TypeReg  //字节'0' 表示常规文件
	TypeRegA = tar.TypeRegA // 已经不推荐使用，使用TypeReg 代替 十六进制的00
	// Type '1' to '6' are header-only flags and may not have a data body. 类型“ 1”到“ 6”是仅标头标志，并且可能没有数据主体。
	TypeLink    = tar.TypeLink    //字节'1' 硬链接
	TypeSymLink = tar.TypeSymlink // 字节 '2' 50 符号链接
	TypeChar    = tar.TypeChar    //角色设备节点
	TypeBlock   = tar.TypeBlock   //块设备节点
	TypeDir     = tar.TypeDir     //目录
	TypeFifo    = tar.TypeFifo    //fifo节点
	TypeCont    = tar.TypeCont    //保留字段

	TypeXHeader       = tar.TypeXHeader       //PAX格式使用“ x”类型存储仅与下一个文件相关的键值记录。该软件包透明地处理这些类型。
	TypeXGlobalHeader = tar.TypeXGlobalHeader //PAX格式使用“ g”类型存储与所有后续文件相关的键值记录。该程序包仅支持解析和组合此类标头，但当前不支持在文件之间持久保存全局状态。
	TypeGNUSparse     = tar.TypeGNUSparse     //类型“ S”表示GNU格式的稀疏文件。
	TypeGNULongName   = tar.TypeGNULongName
	TypeGNULongLink   = tar.TypeGNULongLink //GNU格式将“ L”和“ K”类型用于图元文件，该图元文件用于存储下一个文件的路径或链接名称。该软件包透明地处理这些类型。
)

//变量
var (
	ErrHeader          = errors.New("archive/tar:invalid tar header")
	ErrWriteTooLong    = errors.New("archive/tar:write too long")
	ErrorFieldTooLong  = errors.New("archive/tar:header field too long")
	ErrWriteAfterClose = errors.New("archive/tar:write after close")
)

//原始的tar格式是在Unix V7中引入的。 从那时起，已经有多种竞争格式试图标准化或扩展V7格式以克服其局限性。 最常见的格式是USTAR，PAX和GNU格式，每种格式都有其自身的优点和局限性。

//Writer当前不支持稀疏文件。
type Format int

//标识各种tar格式的常量。
const (
	//FormatUnknown表示格式未知。
	FormatUnknown = iota

	FormatUSTAR

	FormatPAX

	FormatGNU
)

type Header struct {
	TypeFlag byte   // 类型 零值将自动升为TypeReg或TypeDir，具体取决于Name中是否存在斜杠。
	Name     string //文件实体名称
	LinkName string //链接的目标文件名称 链接类型时有效
	Size     int64  //逻辑文件大小（字节单位）
	Mode     int64  //权限和模式位
	Uid      int    //拥有者id
	Gid      int    //拥有组id
	Uname    string //用户名
	Gname    string //组名

	// 如果未指定Format，则Writer.WriteHeader将ModTime舍入到最接近的秒数，并忽略AccessTime和ChangeTime字段。
	// 要使用AccessTime或ChangeTime，请将格式指定为PAX或GNU。
	// 要使用亚秒级分辨率，请将格式指定为PAX
	ModTime    time.Time //修改时间
	AccessTime time.Time //最近访问时间
	ChangeTime time.Time //最近修改时间
	Devmajor   int64     //主设备号（对TypeChar或TypeBlock有效）
	Devminor   int64     //次设备号（对TypeChar或TypeBlock有效）

	Xattrs map[string]string //不推荐使用 使用 PAXRecords替代

	//PAXRecords是PAX扩展头记录的映射。
	//用户定义的记录应具有以下格式的键：
	// VENDOR.keyword
	// VENDOR是所有大写形式的命名空间，而关键字可以
	//不包含'='字符（例如“ GOLANG.pkg.version”）。
	//键和值应为非空的UTF-8字符串。

	//调用Writer.WriteHeader时，从
	//标头中的其他字段优先于PAXRecords。
    PAXRecords map[string]string
    
    // Format specifies the format of the tar header.
    //
    // This is set by Reader.Next as a best-effort guess at the format.
    // Since the Reader liberally reads some non-compliant files,
    // it is possible for this to be FormatUnknown.
    //
    // If the format is unspecified when Writer.WriteHeader is called,
    // then it uses the first format (in the order of USTAR, PAX, GNU)
    // capable of encoding this Header (see Format).
    Format Format //GO1.10
}

//FileInfoHeader从fi创建部分填充的Header。 如果fi描述了符号链接，则FileInfoHeader会将链接记录为链接目标。 如果fi描述目录，则在名称后添加斜杠。

//由于os.FileInfo的Name方法仅返回其描述的文件的基本名称，因此可能有必要修改Header. Name以提供文件的完整路径名。
//func FileInfoHeader(fi os.FileInfo, link string) (*Header, error)

//FileInfo返回头信息的文件信息os.FileInfo。
//func (h *Header) FileInfo() os.FileInfo

//Reader 提供了对tar存档内容的顺序访问。 Reader. Next前进到存档中的下一个文件（包括第一个文件），然后Reader可以被视为io.Reader来访问文件的数据。
type Reader struct {
	tar.Reader //含未导出的字段
}
//基于r 创建一个新的Reader实例
//func NewReader(r io.Reader) *Reader

//Next方法前进到tar存档中的下一个条目。 Header.Size确定可以为下一个文件读取多少个字节。当前文件中的所有剩余数据将被自动丢弃。 io.EOF在输入的末尾返回。
//func (tr *Reader) Next() (*Header, error)

/*
Read方法 从tar归档文件中的当前文件读取。 当到达该文件的末尾时，它将返回（0，io.EOF），直到调用Next前进到下一个文件为止。

如果当前文件稀疏，则标记为孔的区域将以NUL字节的形式读回。

不管Header.Size声明什么，在TypeLink，TypeSymlink，TypeChar，TypeBlock，TypeDir和TypeFifo等特殊类型上调用Read都会返回（0，io.EOF）。
*/
//func (tr *Reader) Read(b []byte) (int, error)

//Writer 提供tar存档的顺序写入。 Write.WriteHeader使用提供的Header开始新文件，然后Writer可以被视为io.Writer来提供该文件的数据。
type Writer struct {
	tar.Writer
}

//基于w创建一个新写入实例
//func NewWriter(w io.Writer) *Writer

//Close通过刷新填充并写入页脚来关闭tar存档。如果当前文件（从对WriteHeader的先前调用中）未完全写入，则将返回错误。
//func (tw *Writer) Close() error

//刷新完成当前文件的块填充的写入。必须先完全写入当前文件，然后才能调用Flush。
//这是不必要的，因为下次调用WriteHeader或Close会隐式清除文件的填充。
//func (tw *Writer) Flush() error

//Write 将b 写入tar存档中。如果在WriteHeader之后写入了超过Header.Size字节，Write返回错误ErrWriteTooLong。
//不管Header.Size声明什么，都对特殊类型（例如TypeLink，TypeSymlink，TypeChar，TypeBlock，TypeDir和TypeFifo）调用Write返回（0，ErrWriteTooLong）。
//func (tw *Writer) Write(b []byte) (int, error)

//WriteHeader写入hdr头信息 并准备接受文件的内容。 Header.Size确定可以为下一个文件写入多少个字节。 如果当前文件未完全写入，则返回错误。 这将在写入标题之前隐式刷新所有必要的填充。
//func (tw *Writer) WriteHeader(hdr *Header) error

//Tar 打包几个文件
func Tar() {
	tarFile,err := os.Create("./output.tar") //tarFile 为os.File 类型 实现了io.Writer 接口
	if err != nil {
		panic(err)
	}
	defer tarFile.Close()

	w := tar.NewWriter(tarFile) //基于tarFile创建写对象实例
	defer w.Close()
	files := [2]string{"./archive_go/tar_file1.txt", "./archive_go/tar_file2.log"}
	for _,fileName := range files {
		//获取文件1的信息
		file1Info,err := os.Stat(fileName)
		if err != nil {
			panic(err)
		}
		hdr,err := tar.FileInfoHeader(file1Info, "") //从file1Info中读取文件头信息
		if err != nil {
			panic(err)
		}
		//写入头信息
		err = w.WriteHeader(hdr)
		if err != nil {
			panic(err)
		}
		//写入文件内容
		file1,err := os.Open(fileName)
		defer file1.Close()
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(w, file1) //会调用w.Write 方法复制文件内容
		if err != nil {
			panic(err)
		}
	}
	
}

//Untar 解包
func Untar(f string) {
	//打开文件
	file,err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := tar.NewReader(file) //基于文件创建一个读取实例对象

	//循环读取文件到当前文件夹
	for hdr,err := r.Next(); err != io.EOF; hdr,err = r.Next() { //i := 0; i<2;i++ 一样的模式
		if err != nil {
			panic(err)
		}
		fileInfo := hdr.FileInfo() //文件信息
		//创建新文件
		newF,err := os.Create("new_" + fileInfo.Name())
		if err != nil {
			panic(err)
		}
		defer newF.Close()
		//写入
		_, err = io.Copy(newF, r) //调用r.Read 方法来读取内容
		if err != nil {
			panic(err)
		}
	}

}
