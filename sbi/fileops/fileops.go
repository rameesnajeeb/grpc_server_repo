package fileops

import (
	fileModel "grpc_server_repo/models/files"
)

type File struct {
	fileModel.File
}

func (F File) WriteJson() {

	// data, err := json.Marshal(p)
	// if err != nil {
	// 	// handle error
	// }
	// fmt.Println(string(data))

}

func (F File) ReadJson() {

}
