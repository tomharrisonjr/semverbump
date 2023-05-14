package semverbump

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BumpSemVer(t *testing.T) {
	type args struct {
		currentVersion  string
		majorMinorPatch string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "no version major passed create v1.0.0",
			args:    args{"", "major"},
			want:    "v1.0.0",
			wantErr: assert.NoError,
		},
		{
			name:    "no version minor passed create v0.1.0",
			args:    args{"", "minor"},
			want:    "v0.1.0",
			wantErr: assert.NoError,
		},
		{
			name:    "no release found patch passed create v0.0.1",
			args:    args{"", "patch"},
			want:    "v0.0.1",
			wantErr: assert.NoError,
		},
		{
			name:    "invalid semVer tag xyz",
			args:    args{"xyz", "major"},
			want:    "",
			wantErr: assert.Error,
		},
		{
			name:    "invalid semVer tag leading 0",
			args:    args{"05", "major"},
			want:    "",
			wantErr: assert.Error,
		},
		{
			name:    "major version only major bump",
			args:    args{"v1", "major"},
			want:    "v2.0.0",
			wantErr: assert.NoError,
		},
		{
			name:    "major minor version only major bump",
			args:    args{"v1.0", "major"},
			want:    "v2.0.0",
			wantErr: assert.NoError,
		},
		{
			name:    "no leading v major bump",
			args:    args{"1.1", "major"},
			want:    "v2.0.0",
			wantErr: assert.NoError,
		},
		{
			name:    "minor bump",
			args:    args{"1.0.1", "minor"},
			want:    "v1.1.0",
			wantErr: assert.NoError,
		},
		{
			name:    "patch bump",
			args:    args{"1.0.0", "patch"},
			want:    "v1.0.1",
			wantErr: assert.NoError,
		},
		{
			name:    "invalid bump type",
			args:    args{"1.0.0", "foo"},
			want:    "",
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Bump(tt.args.currentVersion, tt.args.majorMinorPatch)
			if !tt.wantErr(t, err, fmt.Sprintf("Bump(%v, %v)", tt.args.currentVersion, tt.args.majorMinorPatch)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Bump(%v, %v)", tt.args.currentVersion, tt.args.majorMinorPatch)
		})
	}
}
