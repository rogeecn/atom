package fs

import "testing"

func TestFilePathInfo(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name     string
		args     args
		wantPath string
		wantName string
		wantExt  string
	}{
		{
			"1.",
			args{
				file: "/a/bc.ext",
			},
			"/a",
			"bc",
			".ext",
		},
		{
			"1.",
			args{
				file: "/a/c/c/c/bc.ext",
			},
			"/a/c/c/c",
			"bc",
			".ext",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPath, gotName, gotExt := FilePathInfo(tt.args.file)
			if gotPath != tt.wantPath {
				t.Errorf("FilePathInfo() gotPath = %v, want %v", gotPath, tt.wantPath)
			}
			if gotName != tt.wantName {
				t.Errorf("FilePathInfo() gotName = %v, want %v", gotName, tt.wantName)
			}
			if gotExt != tt.wantExt {
				t.Errorf("FilePathInfo() gotExt = %v, want %v", gotExt, tt.wantExt)
			}
		})
	}
}
