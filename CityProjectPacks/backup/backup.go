package backup

import (
	"cityapp/city"
	"fmt"
	"log"
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

	// listing directory
	dirF, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}
	defer dirF.Close()
	names, err := dirF.Readdirnames(1000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(names)
}

func TestOS() {
	// currentDir := "."
	// fs := os.DirFS(currentDir)
	pid := os.Getpid()
	uid := os.Getuid()
	fmt.Println(pid, uid)
}

func WriteCities(filename string, cities []city.City) {
	// f.Write, f.WriteString, os.WriteFile
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, c := range cities {
		f.WriteString(c.String() + "\n")
	}
} // close en mode defer

func ReadCities(filename string) []city.City {
	// NB : os.ReadFile, f.Read
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)
	fmt.Println(text)
	return []city.City{}
}

func ReadCities2(filename string) []city.City {
	// NB : os.ReadFile, f.Read
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	buffer := make([]byte, 1024)
	for n, err := f.Read(buffer); n > 0; n, err = f.Read(buffer) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d bytes read: %s\n", n, buffer[:n])
	}
	return []city.City{}
}

func TestReadWriteCities() {
	cities := []city.City{
		{Name: "Lyon", Pop: 1500000},
		{Name: "Toulouse", Pop: 1300000},
		{Name: "Pau", Pop: 100000},
	}
	WriteCities("cities.txt", cities)
	cities2 := ReadCities2("cities.txt")
	fmt.Println("Cities read from file:", cities2)
}
