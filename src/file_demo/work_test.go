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
	GoRoutineNum = 6            // 限制读取文件的携程数量，防止短时间加载大量文件内容导致内存溢出
)

var (
	keySlice = make([]string, 0)
	wg       sync.WaitGroup
	mu       sync.Mutex
	fileMu   sync.Mutex
)

func TestFile05(t *testing.T) {
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
	defer close(fileCh)

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileCh <- filepath.Join(DataDir, file.Name())
	}

	for i := 0; i < GoRoutineNum; i++ {
		wg.Add(1)
		go rangFile(fileCh)
	}

	wg.Wait()

	resultFile, err := os.Create(MergeFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resultFile.Close()

	keys := unique(keySlice)
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

// rangFile 遍历文件
func rangFile(fileCh chan string) {
	defer wg.Done()
	for {
		select {
		case filePath := <-fileCh:
			writeTempFile(filePath)
		default:
			return
		}
	}

}

// writeTempFile写入临时文件
func writeTempFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("打开文件错误:%v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fileWriters := make(map[string]*os.File)
	tempKeys := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		key := line[:3]
		tempKeys = append(tempKeys, key)

		tempFile := filepath.Join(TempDir, key+".txt")

		f, ok := fileWriters[tempFile]

		fileMu.Lock()
		if !ok {
			f, err = os.OpenFile(tempFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				continue
			}

			fileWriters[tempFile] = f
		}

		if _, err = f.WriteString(line + "\n"); err != nil {
			log.Printf("写入临时文件错误:%v", err)
		}
		fileMu.Unlock()
	}

	for _, file := range fileWriters {
		file.Close()
	}

	mu.Lock()
	for _, key := range tempKeys {
		keySlice = append(keySlice, key)
	}
	mu.Unlock()
}

// unique 去重
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
