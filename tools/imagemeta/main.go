package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const (
	markerTEM   = 0x01
	markerRST0  = 0xD0
	markerRST7  = 0xD7
	markerEOI   = 0xD9
	markerSOS   = 0xDA
	markerAPP1  = 0xE1
	markerAPP2  = 0xE2
	markerAPP13 = 0xED
	markerCOM   = 0xFE

	tagOrientation = 0x0112
)

var (
	jpegSignature = []byte{0xFF, 0xD8, 0xFF}
	pngSignature  = []byte("\x89PNG\r\n\x1a\n")
)

func main() {
	if len(os.Args) < 3 || (os.Args[1] != "check" && os.Args[1] != "strip") {
		fmt.Fprintln(os.Stderr, "usage: go run ./tools/imagemeta check|strip DIR...")
		os.Exit(2)
	}
	strip := os.Args[1] == "strip"

	found := false
	for _, root := range os.Args[2:] {
		if err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.Type().IsRegular() {
				return nil
			}

			var stripped []byte
			{
				data, err := os.ReadFile(path)
				if err != nil {
					return fmt.Errorf("failed to read file: %w", err)
				}
				if stripped, err = stripMetadata(data); err != nil {
					return fmt.Errorf("failed to strip metadata from %s: %w", path, err)
				}
				if bytes.Equal(stripped, data) {
					return nil
				}
			}

			found = true
			if !strip {
				fmt.Fprintf(os.Stderr, "%s: contains metadata\n", path)
				return nil
			}
			if err := os.WriteFile(path, stripped, 0o644); err != nil {
				return fmt.Errorf("failed to write file: %w", err)
			}
			fmt.Printf("%s: stripped metadata\n", path)
			return nil
		}); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	if found && !strip {
		fmt.Fprintln(os.Stderr, `Run "task fmt" to strip the metadata.`)
		os.Exit(1)
	}
}

func stripMetadata(data []byte) ([]byte, error) {
	switch {
	case bytes.HasPrefix(data, jpegSignature):
		return stripJPEG(data)
	case bytes.HasPrefix(data, pngSignature):
		return stripPNG(data)
	default:
		return data, nil
	}
}

func stripJPEG(data []byte) ([]byte, error) {
	out := append([]byte(nil), data[:2]...)
	pos := 2
	for {
		for pos+1 < len(data) && data[pos] == 0xFF && data[pos+1] == 0xFF {
			pos++
		}
		if pos+2 > len(data) {
			return nil, errors.New("invalid jpeg: must end with EOI")
		}
		if data[pos] != 0xFF {
			return nil, errors.New("invalid jpeg: must have a marker before each segment")
		}

		marker := data[pos+1]
		if marker == markerEOI {
			return append(out, data[pos:pos+2]...), nil
		}
		if marker == markerTEM || markerRST0 <= marker && marker <= markerRST7 {
			out = append(out, data[pos:pos+2]...)
			pos += 2
			continue
		}

		if pos+4 > len(data) {
			return nil, errors.New("invalid jpeg: must end with EOI")
		}
		{
			end := pos + 2 + int(binary.BigEndian.Uint16(data[pos+2:]))
			if end < pos+4 {
				return nil, errors.New("invalid jpeg segment: must have a length of at least 2")
			}
			if end > len(data) {
				return nil, errors.New("invalid jpeg: must end with EOI")
			}

			payload := data[pos+4 : end]
			switch {
			case marker == markerAPP1 && bytes.HasPrefix(payload, []byte("Exif\x00\x00")):
				o, err := orientation(payload[6:])
				if err != nil {
					return nil, err
				}
				if o != 1 {
					return nil, fmt.Errorf("unsupported exif orientation: %d", o)
				}
			case marker == markerAPP1, marker == markerAPP13, marker == markerCOM:
			case marker == markerAPP2 && bytes.HasPrefix(payload, []byte("MPF\x00")):
			default:
				out = append(out, data[pos:end]...)
			}
			pos = end
		}

		if marker == markerSOS {
			start := pos
			for pos+1 < len(data) && (data[pos] != 0xFF || data[pos+1] == 0x00 || markerRST0 <= data[pos+1] && data[pos+1] <= markerRST7) {
				pos++
			}
			out = append(out, data[start:pos]...)
		}
	}
}

func stripPNG(data []byte) ([]byte, error) {
	out := append([]byte(nil), pngSignature...)
	pos := len(pngSignature)
	for {
		if pos+12 > len(data) {
			return nil, errors.New("invalid png: must end with IEND")
		}
		end := pos + 12 + int(binary.BigEndian.Uint32(data[pos:]))
		if end > len(data) {
			return nil, errors.New("invalid png: must end with IEND")
		}

		switch string(data[pos+4 : pos+8]) {
		case "eXIf":
			o, err := orientation(data[pos+8 : end-4])
			if err != nil {
				return nil, err
			}
			if o != 1 {
				return nil, fmt.Errorf("unsupported exif orientation: %d", o)
			}
		case "tEXt", "zTXt", "iTXt", "tIME":
		case "IEND":
			return append(out, data[pos:end]...), nil
		default:
			out = append(out, data[pos:end]...)
		}
		pos = end
	}
}

func orientation(tiff []byte) (uint16, error) {
	if len(tiff) < 8 {
		return 0, errors.New("invalid exif: must not be truncated")
	}

	var order binary.ByteOrder
	switch string(tiff[:2]) {
	case "II":
		order = binary.LittleEndian
	case "MM":
		order = binary.BigEndian
	default:
		return 0, fmt.Errorf("unsupported exif byte order: %q", tiff[:2])
	}

	ifd := int(order.Uint32(tiff[4:]))
	if ifd+2 > len(tiff) {
		return 0, errors.New("invalid exif: must not be truncated")
	}
	n := int(order.Uint16(tiff[ifd:]))
	if ifd+2+12*n > len(tiff) {
		return 0, errors.New("invalid exif: must not be truncated")
	}

	for i := range n {
		entry := tiff[ifd+2+12*i:]
		if order.Uint16(entry) == tagOrientation {
			return order.Uint16(entry[8:]), nil
		}
	}
	return 1, nil
}
