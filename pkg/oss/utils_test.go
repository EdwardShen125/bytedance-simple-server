package oss

import (
	"bytes"
	"testing"
)

func TestPublishVideoToPublic(t *testing.T) {
	type args struct {
		video    []byte
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PublishVideoToPublic(tt.args.video, tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("PublishVideoToPublic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPublishVideoToOss(t *testing.T) {
	type args struct {
		objectKey string
		filePath  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PublishVideoToOss(tt.args.objectKey, tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("PublishVideoToOss() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueryOssVideoURL(t *testing.T) {
	type args struct {
		objectKey string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryOssVideoURL(tt.args.objectKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryOssVideoURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("QueryOssVideoURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPublishCoverToOss(t *testing.T) {
	type args struct {
		objectKey   string
		coverReader *bytes.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PublishCoverToOss(tt.args.objectKey, tt.args.coverReader); (err != nil) != tt.wantErr {
				t.Errorf("PublishCoverToOss() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
