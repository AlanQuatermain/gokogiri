package mem
/*
#cgo CFLAGS: -I/usr/include/libxml2
#cgo LDFLAGS: -lxml2

#include <libxml/xmlversion.h>
#include "libxml.h"
*/
import "C"

const LIBXML_VERSION = C.LIBXML_DOTTED_VERSION

func init() {
	C.libxmlGoInit()
}

func AllocSize() int {
	return int(C.libxmlGoAllocSize())
}
