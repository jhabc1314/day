package main

import "github.com/jhabc1314/day/archive_go"

func main() {
    archive_go.Tar()
    archive_go.Untar("./output.tar")

    archive_go.Zip()
    archive_go.Unzip("./output.zip")
}
