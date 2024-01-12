package mediainfo

// #include <mediainfo.h>
// #cgo CFLAGS: -DUNICODE
// #cgo LDFLAGS: -ldl -lmediainfo
import "C"
import (
    "unsafe"
    "errors"
    "runtime"
)

type StreamKind int

const (
    StreamGeneral StreamKind = iota
    StreamVideo
    StreamAudio
    StreamText
    StreamOther
    StreamImage
    StreamMenu
    streamMax
)

type InfoKind int

const (
    // Unique name of parameter
    InfoName InfoKind = iota
    // Value of parameter
    InfoText
    // Unique name of measure unit of parameter
    InfoMeasure
    // See infooptions_t
    InfoOptions
    // Translated name of parameter
    InfoNameText
    // Translated name of measure unit
    InfoMeasureText
    // More information about the parameter
    InfoInfo
    // Information : how data is found
    InfoHowTo
    infoMax
)

func init() {
    C.setlocale(C.LC_CTYPE, C.CString(""))
    C.MediaInfoDLL_Load()
}

type File struct {
    handle unsafe.Pointer
}

func (file *File) Close() {
    C.CloseFile(file.handle)
}

func (file *File) Inform() string {
    return C.GoString(C.Inform(file.handle, 0))
}

func (file *File) Get(streamKind StreamKind, streamNumber int, parameter string, infoKind InfoKind) string {
    cParameter := C.CString(parameter)
    defer C.free(unsafe.Pointer(cParameter))

    return C.GoString(C.Get(file.handle, C.MediaInfo_stream_C(streamKind), C.size_t(streamNumber), cParameter, C.MediaInfo_info_C(infoKind), C.MediaInfo_info_C(0)))
}

func (file *File) GetI(streamKind StreamKind, streamNumber int, parameter int, infoKind InfoKind) string {
    return C.GoString(C.GetI(file.handle, C.MediaInfo_stream_C(streamKind), C.size_t(streamNumber), C.size_t(parameter), C.MediaInfo_info_C(infoKind)))
}

func (file *File) Parameter(streamKind StreamKind, streamNumber int, parameter string) string {
    return file.Get(streamKind, streamNumber, parameter, InfoText)
}

func (file *File) Option(parameter string, value string) string {
    cParameter := C.CString(parameter)
    defer C.free(unsafe.Pointer(cParameter))

    cValue := C.CString(value)
    defer C.free(unsafe.Pointer(cValue))

    return C.GoString(C.Option(file.handle, cParameter, cValue))
}

func Open(filePath string) (*File, error) {
    cFilePath := C.CString(filePath)
    defer C.free(unsafe.Pointer(cFilePath))

    handle := C.OpenFile(cFilePath)
    if handle == nil {
        return nil, errors.New("Cannot open file.")
    }

    file := &File{
        handle: handle,
    }

    runtime.SetFinalizer(file, func(file *File) {
        file.Close()
    })

    return file, nil
}
