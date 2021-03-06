[![Build Status](https://travis-ci.com/olaure/chardet.svg?branch=master)](https://travis-ci.com/olaure/chardet)

Chardet: The Universal Character Encoding Detector
==================================================


Detects :
 - ASCII, UTF-8, UTF-16 (2 variants), UTF-32 (4 variants)
 - Big5, GB2312, EUC-TW, HZ-GB-2312, ISO-2022-CN (Traditional and Simplified Chinese)
 - EUC-JP, SHIFT_JIS, CP932 (aka MS932), ISO-2022-JP (Japanese)
 - EUC-KR, ISO-2022-KR (Korean)
 - KOI8-R, x-mac-cyrillic (prev MacCyrillic), IBM855, IBM866, ISO-8859-5, windows-1251 (Cyrillic)
 - ISO-8859-5, windows-1251 (Bulgarian)
 - ISO-8859-1, windows-1252 (Western European languages)
 - ISO-8859-7, windows-1253 (Greek)
 - ISO-8859-8, windows-1255 (Visual and Logical Hebrew)
 - TIS-620 (Thai)

**note** :  
   The ISO-8859-2 and windows-1250 (Hungarian) probers have been temporarily
   disabled until the models can be retrained.

About
=====

This is a port to go of the excellent python `chardet library<https://github.com/chardet/chardet>`.
It is based on the mozilla statistical encoding detector.
v0.0.7 is based on the chardet version [0.0.4 (Dec 20)](https://github.com/chardet/chardet/releases/tag/4.0.0)


Usage
=====

The simplest way to use chardet is simply the package-level exported Detect method:

```go
package main

import (
	"fmt"
	"github.com/olaure/chardet"
)

func main() {
	data := []byte("नमस्कार")
	detected := chardet.Detect(data)
	fmt.Printf(
		"Detectected character set : %v with confidence %v\n",
		detected.Encoding, detected.Confidence,
	)
}
```

Another way uses the method DetectShortestUTF8 that will look for the decoded string with the lowest count of unicode categories C (control), S (symbol), P (punctuation):

```go
package main

import (
	"fmt"
	"github.com/olaure/chardet"
)

func main() {
	data := []byte("नमस्कार")
	detected := chardet.DetectShortestUTF8(data)
	fmt.Printf(
		"Detectected character set : %v with confidence %v\n",
		detected.Encoding, detected.Confidence,
	)
}
```

This function thus will not necessarily yield the highest probability decoder, unless the probability is maximum.
