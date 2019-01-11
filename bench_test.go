package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"testing"
	"time"

	"github.com/gogo/protobuf/proto"
	pb "github.com/hlts2/go-json-protobuf-unmarshal-of-image/proto"
	"github.com/pkg/errors"
)

// ImageRequest --
type ImageRequest struct {
	Images []*Image
}

// Image --
type Image struct {
	UUID    string
	Encoded string
}

func getFiles(dirPath string, filesLimit int) (map[string][]byte, error) {
	rand.Seed(time.Now().UnixNano())

	infos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, errors.Wrap(err, "faild to read directory")
	}

	files := make(map[string][]byte, filesLimit)

	cnt := 0
	for i := 0; i < filesLimit; i++ {
		if cnt == filesLimit {
			break
		}

		path := filepath.Join(dirPath, infos[rand.Intn(len(infos))].Name())

		d, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, errors.Wrap(err, "faild to read file")
		}
		files[path] = d

		cnt++
	}

	return files, nil
}

func genJSONData(filesMap map[string][]byte) ([]byte, error) {
	request := &ImageRequest{
		Images: make([]*Image, 0, len(filesMap)),
	}

	for uuid, d := range filesMap {
		request.Images = append(request.Images, &Image{
			UUID:    uuid,
			Encoded: base64.StdEncoding.EncodeToString(d),
		})
	}

	return json.Marshal(request)
}

func genProtobuf(filesMap map[string][]byte) ([]byte, error) {
	request := &pb.ImageRequest{
		Images: make([]*pb.ImageRequest_Image, 0, len(filesMap)),
	}

	for uuid, d := range filesMap {
		request.Images = append(request.Images, &pb.ImageRequest_Image{
			Uuid: uuid,
			Data: d,
		})
	}

	return proto.Marshal(request)
}

func jsonUnmarshal(data []byte) {
	req := ImageRequest{}
	json.Unmarshal(data, &req)
}

func protoUnmarshal(data []byte) {
	req := pb.ImageRequest{}
	proto.Unmarshal(data, &req)
}

func BenchmarkUnmarshal(b *testing.B) {
	cases := []struct {
		imageCnt int
	}{
		{imageCnt: 4},
		{imageCnt: 10},
		{imageCnt: 20},
		{imageCnt: 50},
		{imageCnt: 100},
		{imageCnt: 150},
	}

	for _, bc := range cases {
		filesMap, err := getFiles("./images", bc.imageCnt)
		if err != nil {
			b.Fatal(errors.Wrap(err, "faild to get files from directory"))
		}

		pbData, err := genProtobuf(filesMap)
		if err != nil {
			b.Fatal(errors.Wrap(err, "faild to generate protobuf"))
		}

		jsonData, err := genJSONData(filesMap)
		if err != nil {
			b.Fatal(errors.Wrap(err, "faild to generate json data"))
		}

		b.ResetTimer()

		b.Run(fmt.Sprintf("JsonUnmarshalOf%dfiles", bc.imageCnt), func(b *testing.B) {
			b.Log(fmt.Sprintf("protobuf size: %d", len(jsonData)))
			for i := 0; i < b.N; i++ {
				jsonUnmarshal(jsonData)
			}
		})

		b.Run(fmt.Sprintf("ProtoUnmarshalOf%dfiles", bc.imageCnt), func(b *testing.B) {
			b.Log(fmt.Sprintf("protobuf size: %d", len(pbData)))
			for i := 0; i < b.N; i++ {
				protoUnmarshal(pbData)
			}
		})
	}
}
