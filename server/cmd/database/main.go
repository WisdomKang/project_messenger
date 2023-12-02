package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// database.InitDB()

	exePath, _ := os.Getwd()

	// 실행 파일이 있는 디렉토리 경로를 찾음
	exeDir := filepath.Dir(exePath)

	// 파일 기준으로 상대 경로를 사용하여 파일 찾기 예제
	filePath := filepath.Join(exeDir, "relative_path.txt")

	// 이제 filePath를 사용하여 파일을 열거나 다른 작업을 수행할 수 있음
	fmt.Printf("File path: %s\n", filePath)
}
