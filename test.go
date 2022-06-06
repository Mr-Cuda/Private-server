package test

import (
	"reflect"
	"sync"
	"testing"
)

func TestID_Decimal(t *testing.T) {
	tests := []struct {
		name string
		i    ID
		want string
	}{
		{
			name: "Basic",
			i:    ID(1234),
			want: "1234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Decimal(); got != tt.want {
				t.Errorf("ID.Decimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestID_String(t *testing.T) {
	tests := []struct {
		name string
		i    ID
		want string
	}{
		{
			name: "Basic",
			i:    ID(0x12345555AAAA4321),
			want: "048D15556AAA-24-321",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("ID.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	type args struct {
		worker uint64
		seq    uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "basic",
			args: args{
				worker: 25,
				seq:    12,
			},
			wantErr: false,
		},
		{
			name: "worker zero error",
			args: args{
				worker: 0,
				seq:    12,
			},
			wantErr: true,
		},
		{
			name: "worker error",
			args: args{
				worker: 100,
				seq:    12,
			},
			wantErr: true,
		},
		{
			name: "sequence zero error",
			args: args{
				seq:    0,
				worker: 12,
			},
			wantErr: true,
		},
		{
			name: "sequence error",
			args: args{
				seq:    5000,
				worker: 12,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Generate(tt.args.worker, tt.args.seq)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_generate(t *testing.T) {
	type args struct {
		epoch  uint64
		worker uint64
		seq    uint64
	}
	tests := []struct {
		name string
		args args
		want ID
	}{
		{
			name: "basic",
			args: args{
				epoch:  uint64(0x00FF00),
				worker: uint64(0x0C),
				seq:    uint64(0x0F0),
			},
			want: ID(0x3FC00C0F0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generate(tt.args.epoch, tt.args.worker, tt.args.seq); got != tt.want {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		worker uint64
	}
	tests := []struct {
		name    string
		args    args
		want    *Generator
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				worker: 0,
			},
			wantErr: true,
		},
		{
			name: "error",
			args: args{
				worker: 64,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.worker)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_Generate(t *testing.T) {
	type fields struct {
		epoch  uint64
		seq    uint64
		worker uint64
	}
	tests := []struct {
		name    string
		fields  fields
		want    ID
		wantErr bool
	}{
		{
			name: "basic",
			fields: fields{
				worker: 0xF,
			},
			want: ID(0xF001),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				epoch:  tt.fields.epoch,
				seq:    tt.fields.seq,
				worker: tt.fields.worker,
				mutex:  sync.Mutex{},
			}
			got, err := g.Generate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Generator.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got &= 0xFFFF
			if got != tt.want {
				t.Errorf("Generator.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
