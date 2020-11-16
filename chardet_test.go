package chardet

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

var missingEncoding = []string{
	"iso-8859-2",
	"iso-8859-6",
	"windows-1250",
	"windows-1254",
	"windows-1256",
}

var expectedFailures = []string{
	"tests/iso-8859-7-greek/disabled.gr.xml",
	"tests/iso-8859-9-turkish/divxplanet.com.xml",
	"tests/iso-8859-9-turkish/subtitle.srt",
	"tests/iso-8859-9-turkish/wikitop_tr_ISO-8859-9.txt",
}

var listedExts = []string{".html", ".txt", ".xml", ".srt"}

func dirLister(root string) []string {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if path == root {
				return nil
			}
			name := strings.ToLower(path)
			for _, miss := range missingEncoding {
				if strings.Contains(name, miss) {
					return nil
				}
			}
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return []string{}
	}
	return files
}

func fileLister(root string) []string {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			for _, miss := range expectedFailures {
				if strings.HasSuffix(path, miss) {
					return nil
				}
			}
			ret := true
			for _, allowed := range listedExts {
				if filepath.Ext(path) == allowed {
					ret = false
					break
				}
			}
			if ret {
				fmt.Printf("Ignoring file %v because of ext not in whtelist %v\n", path, filepath.Ext(path))
				return nil
			}
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return []string{}
	}
	return files
}

type testCase struct {
	path     string
	encoding string
}

func getTests() []testCase {
	tests := []testCase{}
	wd, _ := os.Getwd()
	encodings := dirLister(path.Join(wd, "tests/"))
	fmt.Printf("Found %v encodings\n", len(encodings))
	for _, encoding := range encodings {
		enc := path.Base(encoding)
		for _, postfix := range []string{"-arabic", "-bulgarian", "-cyrillic", "-greek", "-hebrew", "-hungarian", "-turkish"} {
			if strings.HasSuffix(enc, postfix) && enc != "x-mac-cyrillic" {
				enc = strings.TrimSuffix(enc, postfix)
				break
			}
		}
		files := fileLister(encoding)
		if len(files) == 0 {
			fmt.Printf("No file found for encoding %v\n", encoding)
			continue
		}
		folderTests := make([]testCase, len(files))
		for idx, file := range files {
			folderTests[idx] = testCase{file, enc}
		}
		tests = append(tests, folderTests...)
	}
	return tests
}

func getFileContent(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content := []byte{}
	for {
		buf := make([]byte, 1024*1024)
		charRead, err := file.Read(buf)
		if err == io.EOF {
			content = append(content, buf[:charRead]...)
			break
		}
		if err != nil {
			return nil, err
		}
		content = append(content, buf[:charRead]...)
	}
	return content, nil
}

func toUTF8(content []byte, encoding string) string {
	var decoded string
	byteReader := bytes.NewReader(content)
	reader, err := charset.NewReaderLabel(encoding, byteReader)
	if err != nil {
		decoded = string(content)
	} else {
		data, _ := ioutil.ReadAll(reader)
		if err != nil {
			decoded = string(content)
		} else {
			decoded = string(data)
		}
	}
	return decoded
}

func compareDecodings(content []byte, found, expected string) bool {
	foundString := toUTF8(content, found)
	expectedString := toUTF8(content, expected)
	return foundString == expectedString
}

func TestEncodingDetection(t *testing.T) {
	tests := getTests()
	success := 0
	//fmt.Printf("Entering test. %v loaded.\n", len(tests))
	for _, test := range tests {
		content, err := getFileContent(test.path)
		if err != nil {
			t.Errorf("Error encountered when reading file %v: %v", test.path, err)
			continue
		}
		log.Debugf("Test %v. Expecting %v, %v bytes read.", test.path, test.encoding, len(content))
		result := Detect(content)
		if result == nil {
			t.Errorf("Nothing detected for test %v. Expected %v.", test.path, test.encoding)
			continue
		}
		if !strings.EqualFold(result.Encoding, test.encoding) {
			sameDecoding := compareDecodings(content, result.Encoding, test.encoding)
			if !sameDecoding {
				t.Errorf("Mismatch for test %v. Expecting %v - got %v.", test.path, test.encoding, result.Encoding)
			} else {
				log.Infof("Found differing encoding (got %v vs %v expected) but same decoded string for test %v", result.Encoding, test.encoding, test.path)
				success++
			}
		} else {
			log.Debugf("FOUND test %v!! Got %v\n", test.path, test.encoding)
			success++
		}
	}
	log.Infof("Success: %v / %v", success, len(tests))
}
