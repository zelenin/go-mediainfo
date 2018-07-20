#include <stdlib.h>
#include <string.h>
#include <wchar.h>
#include <MediaInfoDLL/MediaInfoDLL.h>

const wchar_t *toWchar(const char *c)
{
    const size_t cSize = strlen(c) + 1;
    wchar_t *wc = malloc(cSize * sizeof(wchar_t));
    mbstowcs(wc, c, cSize);

    return wc;
}

const char *toChar(const wchar_t *c)
{
    const size_t cSize = wcslen(c) + 1;
    char *wc = malloc(cSize * sizeof(char));
    wcstombs(wc, c, cSize);

    return wc;
}

void *OpenFile(char *filePath)
{
    void *handle = MediaInfo_New();
    if (!handle)
    {
        return NULL;
    }

    size_t res = MediaInfo_Open(handle, toWchar(filePath));
    if (!res)
    {
        return NULL;
    }

    return handle;
}

void CloseFile(void *handle)
{
    MediaInfo_Close(handle);
    MediaInfo_Delete(handle);
}

const char *Get(void *handle, MediaInfo_stream_C streamKind, size_t StreamNumber, char *Parameter, MediaInfo_info_C KindOfInfo, MediaInfo_info_C KindOfSearch)
{
    return toChar(MediaInfo_Get(handle, streamKind, StreamNumber, toWchar(Parameter), KindOfInfo, KindOfSearch));
}

const char *Option(void *handle, char *Parameter, char *Value)
{
    return toChar(MediaInfo_Option(handle, toWchar(Parameter), toWchar(Value)));
}

const char *Inform(void *handle, size_t Reserved)
{
    return toChar(MediaInfo_Inform(handle, Reserved));
}
