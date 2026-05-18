package thumbnail_test

import (
	"fmt"
	"golandTry/gopl/ch8/thumbnail"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// makeThumbanils makes thumbnails of the specified files.
func makeThumbnails(filenames []string) {
	startTime := time.Now()
	fmt.Println("started")
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
	endTime := time.Now()
	fmt.Println("end:", endTime.Sub(startTime))
}

func makeThumbnails2(filenames []string) {
	startTime := time.Now()
	fmt.Println("started")
	for _, f := range filenames {
		go thumbnail.ImageFile(f) // NOTE: ignoring errors
	}
	endTime := time.Now()
	fmt.Println("end:", endTime.Sub(startTime))
}

func makeTumbnailsByChannel(filenames []string) {
	startTime := time.Now()
	ch := make(chan struct{})
	for _, filename := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(filename)
	}
	for range ch {
		// 等待所有goroutine完成
	}
	end := time.Now()
	fmt.Printf(" total time = %d\n", end.Sub(startTime))
}

func makeThumbnailsByWaitGroup(filenames []string) {
	var wg sync.WaitGroup
	for _, f := range filenames {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			thumbnail.ImageFile(filename)
		}(f)
	}
	wg.Wait()
}

// makeThumbnails6 makes thumbnails for each file received from the channel.
// It returns the number of bytes occupied by the files it creates.
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	// 此处for range会循环直到channel关闭
	for size := range sizes {
		total += size
	}
	return total
}

func scaneFolderAndBuildStrings(folderPath string) []string {
	// 支持的图片文件扩展名
	imageExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
		".tiff": true,
		".webp": true,
	}

	var imageFiles []string

	// 读取目录内容
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		log.Printf("读取目录失败: %v", err)
		return nil
	}

	// 遍历文件，筛选图片文件
	for _, file := range files {
		if !file.IsDir() {
			ext := filepath.Ext(file.Name())
			if imageExtensions[strings.ToLower(ext)] {
				imageFiles = append(imageFiles, filepath.Join(folderPath, file.Name()))
			}
		}
	}

	return imageFiles
}

func main() {
	// 测试用的图片文件路径列表（请替换为实际存在的图片文件）
	testFiles := scaneFolderAndBuildStrings("/Users/leokchen/GolandProjects/golandTry/gopl/ch8/thumbnail/resource")

	fmt.Println("=== 测试顺序处理 ===")
	makeThumbnails(testFiles)

	fmt.Println("\n=== 测试并行处理（不等待完成）===")
	makeThumbnails2(testFiles)

	fmt.Println("\n=== 测试Channel同步 ===")
	makeTumbnailsByChannel(testFiles)

	fmt.Println("\n=== 测试WaitGroup同步 ===")
	makeThumbnailsByWaitGroup(testFiles)
}
