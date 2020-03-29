package compress

//包lzw实现了Lempel-Ziv-Welch压缩数据格式，在T.A. Welch，“高性能数据压缩技术”，计算机，17（6）（1984年6月），第8-19页中进行了介绍。

//特别是，它实现了GIF和PDF文件格式所使用的LZW，这意味着最多12位的可变宽度代码，而前两个非文字代码是明码和EOF代码。

//TIFF文件格式使用LZW算法的相似但不兼容的版本。 有关实现，请参阅golang.org/x/image/tiff/lzw软件包。
