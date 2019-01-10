package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"testing"
	"time"

	"github.com/gogo/protobuf/proto"
	pb "github.com/hlts2/protobuf_unmarshal_v2_json_unmarshal/proto"
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

func genJSONData(dirPath string, filesLimit int) ([]byte, error) {
	filesMap, err := getFiles(dirPath, filesLimit)
	if err != nil {
		return nil, errors.Wrap(err, "faild to get files")
	}

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

func genProtobuf(dirPath string, filesLimit int) ([]byte, error) {
	filesMap, err := getFiles(dirPath, filesLimit)
	if err != nil {
		return nil, errors.Wrap(err, "faild to get files")
	}

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

func jsonPerformance(data []byte, imageCnt int) {
	req := ImageRequest{}
	json.Unmarshal(data, &req)
}

func protoPerformance(data []byte, imageCnt int) {
	req := &pb.ImageRequest{}
	proto.Unmarshal(data, req)
}

func BenchmarkJsonPerformanceTest(b *testing.B) {
	b.Run("Json performance: 4", func(b *testing.B) {
		count := 4

		for i := 0; i < b.N; i++ {

			b.StopTimer()

			data, err := genJSONData("./images", count)
			if err != nil {
				b.Fatal(errors.Wrap(err, "faild to generate json data"))
			}

			b.StartTimer()

			jsonPerformance(data, count)
		}
	})

	b.Run("Json performance: 10", func(b *testing.B) {
		count := 10

		for i := 0; i < b.N; i++ {

			b.StopTimer()

			data, err := genJSONData("./images", count)
			if err != nil {
				b.Fatal(errors.Wrap(err, "faild to generate json data"))
			}

			b.StartTimer()

			jsonPerformance(data, count)
		}
	})

	b.Run("Json performance: 20", func(b *testing.B) {
		count := 20

		for i := 0; i < b.N; i++ {

			b.StopTimer()

			data, err := genJSONData("./images", count)
			if err != nil {
				b.Fatal(errors.Wrap(err, "faild to generate json data"))
			}

			b.StartTimer()

			jsonPerformance(data, count)
		}
	})

	b.Run("Json performance: 50", func(b *testing.B) {
		count := 50

		for i := 0; i < b.N; i++ {

			b.StopTimer()

			data, err := genJSONData("./images", count)
			if err != nil {
				b.Fatal(errors.Wrap(err, "faild to generate json data"))
			}

			b.StartTimer()

			jsonPerformance(data, count)
		}
	})

	b.Run("Json performance: 100", func(b *testing.B) {
		count := 100

		for i := 0; i < b.N; i++ {

			b.StopTimer()

			data, err := genJSONData("./images", count)
			if err != nil {
				b.Fatal(errors.Wrap(err, "faild to generate json data"))
			}

			b.StartTimer()

			jsonPerformance(data, count)
		}
	})
}

func BenchmarkProtoPerformanceTest(b *testing.B) {
	b.Run("protobuf performance: 4", func(b *testing.B) {
		count := 4

		for i := 0; i < b.N; i++ {

			b.StopTimer()

			data, err := genProtobuf("./images", count)
			if err != nil {
				b.Fatal(errors.Wrap(err, "faild to generate protobuf"))
			}

			b.StartTimer()

			jsonPerformance(data, count)
		}
	})

	b.Run("protobuf performance: 10", func(b *testing.B) {
		count := 10

		for i := 0; i < b.N; i++ {

			b.StopTimer()

			data, err := genProtobuf("./images", count)
			if err != nil {
				b.Fatal(errors.Wrap(err, "faild to generate protobuf"))
			}

			b.StartTimer()

			jsonPerformance(data, count)
		}
	})

	b.Run("protobuf performance: 20", func(b *testing.B) {
		count := 20

		for i := 0; i < b.N; i++ {

			b.StopTimer()

			data, err := genProtobuf("./images", count)
			if err != nil {
				b.Fatal(errors.Wrap(err, "faild to generate protobuf"))
			}

			b.StartTimer()

			jsonPerformance(data, count)
		}
	})

	b.Run("protobuf performance: 50", func(b *testing.B) {
		count := 50

		for i := 0; i < b.N; i++ {

			b.StopTimer()

			data, err := genProtobuf("./images", count)
			if err != nil {
				b.Fatal(errors.Wrap(err, "faild to generate protobuf"))
			}

			b.StartTimer()

			jsonPerformance(data, count)
		}
	})

	b.Run("protobuf performance: 100", func(b *testing.B) {
		count := 100

		for i := 0; i < b.N; i++ {

			b.StopTimer()

			data, err := genProtobuf("./images", count)
			if err != nil {
				b.Fatal(errors.Wrap(err, "faild to generate protobuf"))
			}

			b.StartTimer()

			jsonPerformance(data, count)
		}
	})
}
