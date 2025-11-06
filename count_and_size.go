// Copyright 2025 The LevelDB-Go and Pebble Authors. All rights reserved. Use
// of this source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package deletepacer

// CountAndSize tracks the count and total size of a set of items.
type CountAndSize struct {
	// Count is the number of files.
	Count uint64

	// Bytes is the total size of all files.
	Bytes uint64
}

// Inc increases the count and size for a single item.
func (cs *CountAndSize) Inc(fileSize uint64) {
	cs.Count++
	cs.Bytes += fileSize
}

// TableCountsAndSizes contains counts and sizes for tables, broken down by
// locality.
type TableCountsAndSizes struct {
	// All contains counts for all tables (local and remote).
	All CountAndSize
	// Local contains counts for local tables only.
	Local CountAndSize
}

// Inc increases the count and size for a single table.
func (cs *TableCountsAndSizes) Inc(tableSize uint64, isLocal bool) {
	cs.All.Inc(tableSize)
	if isLocal {
		cs.Local.Inc(tableSize)
	}
}

// BlobFileCountsAndSizes contains counts and sizes for blob files, broken down
// by locality.
type BlobFileCountsAndSizes struct {
	// All contains counts for all blob files (local and remote).
	All CountAndSize
	// Local contains counts for local blob files only.
	Local CountAndSize
}

// Inc increases the count and size for a single blob file.
func (cs *BlobFileCountsAndSizes) Inc(fileSize uint64, isLocal bool) {
	cs.All.Inc(fileSize)
	if isLocal {
		cs.Local.Inc(fileSize)
	}
}

// FileCountsAndSizes contains counts and sizes for all file types.
type FileCountsAndSizes struct {
	// Tables contains counts and sizes for tables.
	Tables TableCountsAndSizes

	// BlobFiles contains counts and sizes for blob files.
	BlobFiles BlobFileCountsAndSizes

	// Other contains counts and sizes for other file types (log, manifest, etc).
	// These are not separated by locality.
	Other CountAndSize
}

// Inc increases the relevant count and size for a single file.
func (cs *FileCountsAndSizes) Inc(fileType FileType, fileSize uint64, isLocal bool) {
	switch fileType {
	case FileTypeTable:
		cs.Tables.Inc(fileSize, isLocal)
	case FileTypeBlob:
		cs.BlobFiles.Inc(fileSize, isLocal)
	default:
		cs.Other.Inc(fileSize)
	}
}

type FileType int

// The FileType enumeration.
const (
	FileTypeLog FileType = iota
	FileTypeLock
	FileTypeTable
	FileTypeManifest
	FileTypeOptions
	FileTypeOldTemp
	FileTypeTemp
	FileTypeBlob
)
