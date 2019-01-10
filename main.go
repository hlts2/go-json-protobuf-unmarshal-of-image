package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
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

func jsonPerformance(filesMap map[string][]byte, imageCnt int) error {
	data, err := genJSONData(filesMap)
	if err != nil {
		return errors.Wrap(err, "faild to execute json performance test")
	}

	fmt.Printf("json unmarshal check. image count: %d\n", imageCnt)
	start := time.Now()

	req := ImageRequest{}
	json.Unmarshal(data, &req)

	fmt.Println(time.Since(start))

	return nil
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

func protoPerformance(filesMap map[string][]byte, imageCnt int) error {
	data, err := genProtoData(filesMap)
	if err != nil {
		return errors.Wrap(err, "faild to execute proto protoPerformance check")
	}

	fmt.Printf("proto unmarshal check. image count: %d\n", imageCnt)

	start := time.Now()

	req := &pb.ImageRequest{}
	proto.Unmarshal(data, req)

	fmt.Println(time.Since(start))
	return nil
}

func genProtoData(filesMap map[string][]byte) ([]byte, error) {
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

func main() {
	counts := []int{4, 10, 20, 50, 100}
	for _, count := range counts {
		m, err := getFiles("./images", count)
		if err != nil {
			log.Fatal(err)
		}

		jsonPerformance(m, count)
		protoPerformance(m, count)
	}
}
