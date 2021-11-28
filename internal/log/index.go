package log

import (
	"io"
	"os"

	"github.com/tysonmote/gommap"
)

// Defines the number of bytes that make up each entry
var (
	offWidth uint64 = 4
	posWidth uint64 = 8

	// used to jump straight to the position of an entry given its
	// offset since the positino in the file is offset * entWidth
	entWidth = offWidth + posWidth
)

type index struct {
	file *os.File
	mmap gommap.MMap
	size uint64
}

// creates an index for the given file, We create the index
// and then save the current size of the file so we can track
// the amount of data in the index file as we add entries. We
// grow the file to the max index size before memory-mapping
// the file and then return the created index to the caller
func newIndex(f *os.File, c Config) (*index, error) {
	idx := &index{
		file: f,
	}
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}
	idx.size = uint64(fi.Size())
	if err = os.Truncate(
		f.Name(), int64(c.Segment.MaxIndexBytes),
	); err != nil {
		return nil, err
	}
	if idx.mmap, err = gommap.Map(
		idx.file.Fd(),
		gommap.PROT_READ|gommap.PROT_WRITE,
		gommap.MAP_SHARED,
	); err != nil {
		return nil, err
	}
	return idx, nil
}

// Makes sure the memory mapped file has synced its data
// to the persisted file and that the persisted file
// has flushed its contents to a stable storage. It then
// truncates the persisted file to the amount of data thats
// actually in it and closes out the file
func (i *index) Close() error {
	if err := i.mmap.Sync(gommap.MS_SYNC); err != nil {
		return err
	}
	if err := i.file.Sync(); err != nil {
		return err
	}
	if err := i.file.Truncate(int64(i.size)); err != nil {
		return err
	}
	return i.file.Close()
}

// Read takes in an offset and returns the associated records position in
// in the store. The given offset is relative to the segments base offset;
// 0 is always the offset of the indexes first entry, 1,2,3 etc...
// We use relative offsets since it reduces the the size of indexes to uint32.
// If we used absolute offsets as uint64 it would require + (4bytes * num of records)
func (i *index) Read(in int64) (out uint32, pos uint64, err error) {
	if i.size == 0 {
		return 0, 0, io.EOF
	}
	if in == -1 {
		out = uint32((i.size / entWidth) - 1)
	} else {
		out = uint32(in)
	}
	pos = uint64(out) * entWidth
	if i.size < pos+entWidth {
		return 0, 0, io.EOF
	}
	out = enc.Uint32(i.mmap[pos : pos+offWidth])
	pos = enc.Uint64(i.mmap[pos+offWidth : pos+entWidth])
	return out, pos, nil
}

// Appends the give offset and position to the index. First
// we validate that we have space to add the entry. If we do,
// then we increment the position where the next write lives
func (i *index) Write(off uint32, pos uint64) error {
	if uint64(len(i.mmap)) < i.size+entWidth {
		return io.EOF
	}
	enc.PutUint32(i.mmap[i.size:i.size+offWidth], off)
	enc.PutUint64(i.mmap[i.size+offWidth:i.size+entWidth], pos)
	i.size += uint64(entWidth)
	return nil
}

func (i *index) Name() string {
	return i.file.Name()
}
