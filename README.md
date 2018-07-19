# go-mediainfo

Go bindings for [MediaInfo](https://github.com/MediaArea/MediaInfoLib)

## Dependencies

### ubuntu 18.04

```bash
sudo apt-get install libmediainfo0v5 libmediainfo-dev
```

## Usage

### Open file

```go
file, err := mediainfo.Open("/file/path/filename.mp4")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
```

### Bitrate

```go
videoBitrate := file.Parameter(mediainfo.StreamVideo, 0, "BitRate")
audioBitrateTrack1 := file.Parameter(mediainfo.StreamAudio, 0, "BitRate")
audioBitrateTrack2 := file.Parameter(mediainfo.StreamAudio, 1, "BitRate")
```

### Stream count

```go
audioTracks := file.Parameter(mediainfo.StreamAudio, 0, "StreamCount")
```

### Available parameters

```go
parameters := file.Option("info_parameters", "")
```

## Author

[Aleksandr Zelenin](https://github.com/zelenin/), e-mail: [aleksandr@zelenin.me](mailto:aleksandr@zelenin.me)
