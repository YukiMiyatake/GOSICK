package MyFile

import "os"

type ImyFile interface {
	Read(len int) ([]byte, error)
}

type myFile struct {
	fp       *os.File
	filename string
}

func NewMyFile(filename string, cb func(ImyFile) error) error {
	fp, err := os.Open(filename)
	s := &myFile{filename: filename, fp: fp}
	if err != nil {
		return err
	}
	defer s.fp.Close()

	err = cb(s)

	return err
}

func (s *myFile) Read(len int) ([]byte, error) {
	var err error
	buf := make([]byte, len)
	for {
		n, err := s.fp.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
	}
	return buf, err
}
