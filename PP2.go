package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"time"
)

const (
	blockSize    = 1024
	diskSize     = 6 * 1024 * 1024
	numInodes    = 80
	inodeSize    = 64  // Size of inode in bytes
	directoryMax = 256 // Maximum number of directory entries
)

var Disk [diskSize]byte // 6MB virtual disk

// Inode represents an inode
type Inode struct {
	IsValid      bool
	IsDirectory  bool
	DataBlocks   [4]int // Indexes of data blocks
	CreatedTime  int64
	LastModified int64
	// Add more properties as needed
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

// Bitmap represents an allocation bitmap
type Bitmap struct {
	Data []byte
}






















// ReadFromDisk reads data from the virtual disk
func ReadFromDisk(offset int64, data []byte) {
	copy(data, Disk[offset:])
}

// WriteToDisk writes data to the virtual disk
func WriteToDisk(offset int64, data []byte) {
	// Resize Disk array if necessary
	if int(offset)+len(data) > len(Disk) {
		newSize := int(offset) + len(data)
		if newSize > len(Disk) {
			fmt.Printf("Resizing Disk to %d bytes\n", newSize)
			var newDisk [diskSize]byte
			copy(newDisk[:], Disk[:]) // Convert newDisk to a slice
			Disk = newDisk
		}
	}

	copy(Disk[offset:], data)
}

// InodeToBytes converts Inode to byte array
func InodeToBytes(inode Inode) []byte {
	b := make([]byte, inodeSize)
	var isValid, isDirectory uint64
	if inode.IsValid {
		isValid = 1
	}
	if inode.IsDirectory {
		isDirectory = 1
	}
	binary.LittleEndian.PutUint64(b[:8], isValid) // Update slice bounds
	binary.LittleEndian.PutUint64(b[8:], isDirectory)
	for i := 0; i < 4; i++ {
		binary.LittleEndian.PutUint64(b[16+(i*8):], uint64(inode.DataBlocks[i]))
	}
	binary.LittleEndian.PutUint64(b[48:], uint64(inode.CreatedTime))
	binary.LittleEndian.PutUint64(b[56:], uint64(inode.LastModified))
	return b
}

// BytesToInode converts byte array to Inode
func BytesToInode(b []byte) Inode {
	inode := Inode{}
	inode.IsValid = binary.LittleEndian.Uint64(b) != 0
	inode.IsDirectory = binary.LittleEndian.Uint64(b[8:]) != 0
	for i := 0; i < 4; i++ {
		inode.DataBlocks[i] = int(binary.LittleEndian.Uint64(b[16+(i*8):]))
	}
	inode.CreatedTime = int64(binary.LittleEndian.Uint64(b[48:]))
	inode.LastModified = int64(binary.LittleEndian.Uint64(b[56:]))
	return inode
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
func OpenFile(filename string, parentDir *Directory, bitmap *Bitmap) (*Inode, error) {
	for _, entry := range parentDir.Entries {
		if entry.Name == filename {
			return &entry.Inode, nil
		}
	}
	// Allocate a new inode
	inodeIndex := bitmap.AllocateBlock()
	if inodeIndex == -1 {
		return nil, errors.New("no free inode available")
	}
	inode := Inode{IsValid: true, CreatedTime: time.Now().Unix()}
	WriteToDisk(int64(inodeIndex)*blockSize, InodeToBytes(inode))
	// Update directory
	parentDir.Entries = append(parentDir.Entries, DirectoryEntry{Name: filename, Inode: inode})
	return &parentDir.Entries[len(parentDir.Entries)-1].Inode, nil
}

// UnlinkFile unlinks a file
func UnlinkFile(filename string, parentDir *Directory, bitmap *Bitmap) error {
	for i, entry := range parentDir.Entries {
		if entry.Name == filename {
			// Mark all data blocks associated with the inode as available
			for _, blockIndex := range entry.Inode.DataBlocks {
				if blockIndex != -1 {
					bitmap.FreeBlock(blockIndex)
				}
			}
			// Remove entry from directory
			parentDir.Entries = append(parentDir.Entries[:i], parentDir.Entries[i+1:]...)
			return nil
		}
	}
	return errors.New("file not found")
}

// ReadFile reads from a file
func ReadFile(inode *Inode, offset int64, length int) ([]byte, error) {
	if !inode.IsValid {
		return nil, errors.New("invalid inode")
	}
	data := make([]byte, length)
	for i, blockIndex := range inode.DataBlocks {
		if blockIndex != -1 {
			blockStart := int64(blockIndex * blockSize)
			ReadFromDisk(blockStart+offset, data[i*blockSize:])
		}
	}
	return data, nil
}

// WriteFile writes to a file
func WriteFile(inode *Inode, offset int64, data []byte) error {
	if !inode.IsValid {
		return errors.New("invalid inode")
	}
	for i, blockIndex := range inode.DataBlocks {
		if blockIndex != -1 {
			blockStart := int64(blockIndex * blockSize)
			WriteToDisk(blockStart+offset, data[i*blockSize:])
		}
	}
	return nil
}

func main() {
	fmt.Println("Virtual File System")

	// Initialize bitmap
	bitmap := &Bitmap{Data: make([]byte, diskSize/blockSize)}

	// Initialize inodes
	inodeOffset := diskSize / blockSize
	for i := 0; i < numInodes; i++ {
		inode := Inode{}
		WriteToDisk(int64(inodeOffset+i)*blockSize, InodeToBytes(inode))
	}

	// Test functions
	rootDir := &Directory{}
	inode, err := OpenFile("file1", rootDir, bitmap) // Pass bitmap to OpenFile
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Opened file1 with inode:", inode)

	err = UnlinkFile("file1", rootDir, bitmap) // Pass bitmap to UnlinkFile
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Unlinked file1")

	// Read and Write tests can be performed here
}
