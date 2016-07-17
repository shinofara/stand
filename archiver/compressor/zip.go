package compressor

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"time"
	"os"
	"fmt"
	"sync"

	"github.com/shinofara/stand/config"

	"golang.org/x/net/context"
)

type (
	ZipCompressor struct {
		ctx context.Context
		cfg *config.Config
	}

	zipCompression struct {
		p  chan int
		wg *sync.WaitGroup
		zw *zip.Writer
	}
)

func (c *zipCompression) perform(f string) {
	c.p <- 1
	defer c.wg.Done()

	fw, err := os.Open(f)
	if err != nil {
		<- c.p
		return
	}
	defer fw.Close()
	
	info, err := fw.Stat()
	if err != nil {
		<- c.p		
		return
	}
	
	if info.IsDir() {
		<- c.p		
		return 
	}	

	hdr, err := setFileHeader(info)
	if err != nil {
		<- c.p		
		return		
	}

	wf, err := c.zw.CreateHeader(hdr)
	if err != nil {
		<- c.p
		return		
	}
	contents, _ := ioutil.ReadAll(fw)
	wf.Write(contents)
	<- c.p	
}

func NewZipCompressor(ctx context.Context, cfg *config.Config) *ZipCompressor {
	return &ZipCompressor{
		ctx: ctx,
		cfg: cfg,
	}
}

func (c *ZipCompressor) Compress(compressedFile io.Writer, files []string) error {
	zw := zip.NewWriter(compressedFile)


	zc := &zipCompression{
		p:  make(chan int, 14),
		wg: new(sync.WaitGroup),
		zw: zw,
	}
	
	for _, f := range files {
		zc.wg.Add(1)
		go zc.perform(fmt.Sprintf("%s/%s", c.cfg.Path, f))
	}

	zc.wg.Wait()
	zw.Close()
	return nil
}

func setFileHeader(info os.FileInfo) (*zip.FileHeader, error) {
	hdr, err := zip.FileInfoHeader(info)
	if err != nil {
		return nil, err
	}

	local := time.Now().Local()

	//現時刻のオフセットを取得
	_, offset := local.Zone()

	//差分を追加
	hdr.SetModTime(hdr.ModTime().Add(time.Duration(offset) * time.Second))

	return hdr, nil
}
