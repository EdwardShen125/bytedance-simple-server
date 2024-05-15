package oss

import "testing"

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
