package structural

import "fmt"

/*
# Composite Pattern

## Concept
Consists of creating objects into a tree-like structure, while the client only
has to work as if it was a singular object.
*/

// Component interface
type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

// func (f *File) getName() string {
// 	return f.name
// }

// Composite
type Folder struct {
	components []Component
	name       string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Searching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}

// Leaf
type Component interface {
	search(string)
}

// Client Code
func RunCompositeExample() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{name: "Folder1"}

	folder1.add(file1)

	folder2 := &Folder{name: "Folder2"}

	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}
