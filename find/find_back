package find

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

type FileInfos []os.FileInfo
type ByName struct{ FileInfos }

func (fi ByName) Len() int {
	return len(fi.FileInfos)
}
func (fi ByName) Swap(i, j int) {
	fi.FileInfos[i], fi.FileInfos[j] = fi.FileInfos[j], fi.FileInfos[i]
}
func (fi ByName) Less(i, j int) bool {
	return fi.FileInfos[j].ModTime().Unix() < fi.FileInfos[i].ModTime().Unix()
}

// 指定されたファイル名がディレクトリかどうか調べる
func IsDirectory(name string) (isDir bool, err error) {
	fInfo, err := os.Stat(name) // FileInfo型が返る。
	if err != nil {
		return false, err // もしエラーならエラー情報を返す
	}
	// ディレクトリかどうかチェック
	return fInfo.IsDir(), nil
}

func All(path string) ([]string, error) {
	// ディレクトリ内のファイル情報の読み込み[] *os.FileInfoが返る。
	fileInfos, err := ioutil.ReadDir(path)

	// ディレクトリの読み込みに失敗したらエラーで終了
	if err != nil {
		return nil, fmt.Errorf("Directory cannot read %s\n", err)
	}

	// ファイル情報を一つずつ表示する
	var files []string
	sort.Sort(ByName{fileInfos})
	for _, fileInfo := range fileInfos {
		// *FileInfo型
		var name = (fileInfo).Name()
		files = append(files, name)
	}

	return files, nil
}
