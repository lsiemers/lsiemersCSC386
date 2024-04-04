package main

import (
	"fmt"
	"time"
)

// Disk represents the virtual disk
var Disk [6 * 1024 * 1024]byte 	// 6MB virtual disk Array

type Inode struct {
	IsValid      bool		//(True if this inode is an allocated file)
	IsDirectory  bool		//(True if this inode represents a directory rather than a regular file)
	DataBlocks   [4]int 	//(three direct block and one indirect block)
	CreatedTime  time.Time	//Created Time stamp
	LastModified time.Time	//Last modified Time stamp
}

// DirectoryEntry represents an entry in a directory
type DirectoryEntry struct {
	Name  string
	Inode Inode
}

// Directory represents a directory
type Directory struct {
	Entries []DirectoryEntry
}

// AddFile adds a file entry to the directory
func (d *Directory) AddFile(name string, inode Inode) {
	entry := DirectoryEntry{Name: name, Inode: inode}
	d.Entries = append(d.Entries, entry)
}

// RemoveFile removes a file entry from the directory
func (d *Directory) RemoveFile(name string) {
	for i, entry := range d.Entries {
		if entry.Name == name {
			d.Entries = append(d.Entries[:i], d.Entries[i+1:]...)
			return
		}
	}
}

// ListFiles lists all files in the directory
func (d *Directory) ListFiles() []string {
	files := make([]string, len(d.Entries))
	for i, entry := range d.Entries {
		files[i] = entry.Name
	}
	return files
}
// ReadFromDisk reads data from the virtual disk
func ReadFromDisk(offset int64, data []byte) {
	copy(data, Disk[offset:])
}

// WriteToDisk writes data to the virtual disk
func WriteToDisk(offset int64, data []byte) {
	copy(Disk[offset:], data)
}
// Bitmap represents an allocation bitmap
type Bitmap struct {
	Data []byte
}

// AllocateBlock allocates a block in the bitmap
func (b *Bitmap) AllocateBlock() int {
	for i, val := range b.Data {
		if val == 0 {
			b.Data[i] = 1
			return i
		}
	}
	return -1 // No free blocks available
}

// FreeBlock frees a block in the bitmap
func (b *Bitmap) FreeBlock(index int) {
	b.Data[index] = 0
}

// OpenFile opens a file
func OpenFile(filename string, parentDir *Directory) (*Inode, error) {
	// Implementation goes here
	return nil, nil
}

// UnlinkFile unlinks a file
func UnlinkFile(filename string, parentDir *Directory) error {
	// Implementation goes here
	return nil
}

// ReadFile reads from a file
func ReadFile(filename string, offset int64, length int) ([]byte, error) {
	// Implementation goes here
	return nil, nil
}

// WriteFile writes to a file
func WriteFile(filename string, data []byte) error {
	// Implementation goes here
	return nil
}

func main() {
	fmt.Println("Virtual File System")
	// You can add more code here to interact with the virtual file system
}
