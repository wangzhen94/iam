package reflect

import (
	"reflect"
	"testing"
)

func TestGetObjFieldsMap(t *testing.T) {
	type Obj struct {
		A int
		B string
		C bool
	}
	org := &Obj{
		A: 1,
		B: "kk",
		C: true,
	}
	type args struct {
		obj    interface{}
		fields []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "normal obj",
			args: args{
				obj:    org,
				fields: nil,
			},
			want:    map[string]interface{}{"A": 1, "B": "kk", "C": true},
			wantErr: false,
		},
		/*		{
				name: "nest obj",
				args: args{
					obj: args{
						obj: org,
					},
					fields: []string{},
				},
				want:    map[string]interface{}{"obj": map[string]interface{}{"A": 1, "B": "kk", "C": true}},
				wantErr: false,
			},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetObjFieldsMap(tt.args.obj, tt.args.fields)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetObjFieldsMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopyObj(t *testing.T) {
	type Obj struct {
		A int
		B int
		C int
	}

	org := &Obj{
		A: 1,
		B: 2,
		C: 3,
	}

	des := &Obj{
		A: 4,
		B: 5,
		C: 6,
	}

	type args struct {
		from   interface{}
		to     interface{}
		fields []string
	}
	tests := []struct {
		name        string
		args        args
		wantChanged bool
		wantErr     bool
	}{
		{
			name: "change",
			args: args{
				from: org,
				to:   des,
				//fields: []string{"A"},
			},
			wantChanged: true,
			wantErr:     false,
		},
		{
			name: "no change",
			args: args{
				from: org,
				to:   org,
				//fields: []string{"A"},
			},
			wantChanged: false,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChanged, err := CopyObj(tt.args.from, tt.args.to, tt.args.fields)
			if (err != nil) != tt.wantErr {
				t.Errorf("CopyObj() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotChanged != tt.wantChanged {
				t.Errorf("CopyObj() gotChanged = %v, want %v", gotChanged, tt.wantChanged)
			}
		})
	}
}
