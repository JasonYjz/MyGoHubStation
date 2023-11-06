package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ModelCfg struct {
	Id                  float64 `json:"id"`
	Version             string  `json:"version"`
	Type                string  `json:"type"`
	Name                string  `json:"name"`
	Path                string  `json:"path"`
	Size                []int   `json:"size"`
	OutputShape         []int   `json:"output_shape"`
	InputBlobName       string  `json:"input_blob_name"`
	OutputBlobNames     string  `json:"output_blob_names"`
	ClassifierThreshold float64 `json:"classifier_threshold"`
	DeviceType          int     `json:"device_type"`
	Data                string  `json:"data"`
	LabelFile           string  `json:"label_file"`
	ShowResult          bool    `json:"show_result"`
	OperateOnAieId      int     `json:"operate_on_aie_id"`
	OperateOnClassIds   int     `json:"operate_on_class_ids"`
}

func main() {
	bytes, err := ioutil.ReadFile("model.config")
	if err != nil {
		panic(err)
	}

	var models []ModelCfg
	if err := json.Unmarshal(bytes, &models); err != nil {
		panic(err)
	}

	for _, model := range models {
		fmt.Printf("model:%+v\n", model)
	}
}
