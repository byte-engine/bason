package files

import (
	"os"
	"path/filepath"
)

func RealFileInfo(filePath string) (fileInfo os.FileInfo, err error) {
	var (
		logFileInfoIsLink = true
		nextFilePath      string
	)
	for logFileInfoIsLink {
		fileInfo, err = os.Lstat(filePath)
		if err != nil {
			panic(err)
		}
		logFileInfoIsLink = (fileInfo.Mode().Type() & os.ModeSymlink) == os.ModeSymlink
		if logFileInfoIsLink {
			nextFilePath, err = os.Readlink(filePath)
			if err != nil {
				panic(err)
			}
			nextFilePath = filepath.Join(filepath.Dir(filePath), nextFilePath)
			filePath = nextFilePath
		}
	}
	return
}
