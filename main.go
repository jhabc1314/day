package main

import (
	"github.com/jhabc1314/day/archive_go"
	"github.com/jhabc1314/day/bufio"
)

func main() {
    archive_go.Tar()
    archive_go.Untar("./cache/output.tar")

    archive_go.Zip()
    archive_go.Unzip("./cache/output.zip")

    bufio.Read("./tar_file1.txt")
    bufio.Write()
}
