package file_demo

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"testing"
)

const (
	DataDir      = "./data"     // 源文件夹
	TempDir      = "./temp"     // 临时文件夹
	MergeFile    = "result.txt" // 合并后文件
	GoRoutineNum = 6            //  限制读取文件的携程数量，防止短时间加载大量文件内容导致内存溢出
)

var (
	keySetCh = make(chan []string, GoRoutineNum)
	wg       sync.WaitGroup
)

func TestFileMerge(t *testing.T) {
	files, err := os.ReadDir(DataDir)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir(TempDir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(TempDir)

	fileCh := make(chan string, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileCh <- filepath.Join(DataDir, file.Name())
	}
	for i := 0; i < GoRoutineNum; i++ {
		wg.Add(1)
		go readFile(fileCh)
	}

	wg.Wait()

	resultFile, err := os.Create(MergeFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resultFile.Close()

	strings := <-keySetCh
	keys := unique(strings)
	sort.Strings(keys)

	for _, key := range keys {
		tempFile := filepath.Join(TempDir, key+".txt")
		file, err := os.Open(tempFile)
		if err != nil {
			continue
		}
		defer file.Close()
		if _, err := io.Copy(resultFile, file); err != nil {
			continue
		}
	}
}

func readFile(fileCh chan string) {
	defer wg.Done()
	tempkeys := make([]string, 0)
	for filePath := range fileCh {
		file, err := os.Open(filePath)
		if err != nil {
			log.Printf("打开文件错误:%v", err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		fileWriters := make(map[string]*bufio.Writer)
		files := make([]*os.File, 0)
		for scanner.Scan() {
			line := scanner.Text()
			key := line[:3]
			tempkeys = append(tempkeys, key)
			tempFile := filepath.Join(TempDir, key+".txt")
			if _, ok := fileWriters[tempFile]; !ok {
				f, err := os.Create(tempFile)
				if err != nil {
					continue
				}
				files = append(files, f)
				writer := bufio.NewWriter(f)
				fileWriters[tempFile] = writer
			}
			w := fileWriters[tempFile]
			if _, err = w.WriteString(line + "\n"); err != nil {
				log.Printf("写入临时文件错误:%v", err)
			}
			w.Flush()
		}

		for _, file := range files {
			file.Close()
		}
	}
	keySetCh <- tempkeys
}

func unique(strs []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, str := range strs {
		if !seen[str] {
			seen[str] = true
			result = append(result, str)
		}
	}

	return result
}
