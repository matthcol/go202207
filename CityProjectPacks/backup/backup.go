package backup

import (
	"fmt"
	"os"
	"path"
)

func TestPaths() {
	path1 := "sql/cities.sql"
	dir := path.Dir(path1)
	filename := path.Base(path1)
	ext := path.Ext(path1)
	fmt.Println(dir, filename, ext, path.IsAbs(path1))
	path2 := path.Join(dir, "data", "data-cities.sql")
	fmt.Println(path2)
}

func TestOS() {
	// currentDir := "."
	// fs := os.DirFS(currentDir)
	pid := os.Getpid()
	uid := os.Getuid()
	fmt.Println(pid, uid)
}
