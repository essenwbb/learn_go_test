package Escape

import (
	"encoding/json"
	"testing"
)

func TestEscape(t *testing.T) {
	type args struct {
		req []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "No Escapes",
			args: args{
				req: []byte(`{"Name":"Bob"}`),
			},
			want:    `{"Name":"Bob"}`,
			wantErr: false,
		},
		{
			name: "Escapes",
			args: args{
				req: []byte("{\"Name\":\"Bob\"}"),
			},
			want:    `{"Name":"Bob"}`,
			wantErr: false,
		},
		{
			name: "int value",
			args: args{
				req: []byte(`{"1":1}`),
			},
			want:    `{"1":1}`,
			wantErr: false,
		},
		{
			name: "int value escape",
			args: args{
				req: []byte("{\"1\":1}"),
			},
			want:    `{"1":1}`,
			wantErr: false,
		},
		{
			name: "test not json",
			args: args{
				req: []byte("hello \"五月\"！！"),
			},
			want:    `hello "五月"！！`,
			wantErr: false,
		},
		{
			name: "not json test",
			args: args{
				req: []byte("hello 五月！\"\""),
			},
			want:    `hello 五月！""`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := escape(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("escape() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("escape() got = %v, want %v", got, tt.want)
			}
		})
	}

}

type Product struct {
	Name       string  `json:"name"`
	Weight     float64 `json:"weight,string"`
	WeightUnit string  `json:"weight_unit"`
	OnSale     bool    `json:"on_sale"`
}

type Payload struct {
	Objects     []interface{} `json:"objects,omitempty"`
	Description string        `json:"description"`
}

type ResData struct {
	Status  string  `json:"status"`
	Payload Payload `json:"payload"`
}

func TestMarshalMap(t *testing.T) {
	t.Helper()
	s1 := map[string]ResData{"return": {
		Status: "ok",
		Payload: Payload{
			Objects: []interface{}{
				Product{
					Name:       "Horse",
					Weight:     100.21,
					WeightUnit: "KG",
					OnSale:     false,
				},
			},
			Description: "ok",
		},
	}}
	res, err := json.Marshal(s1)
	if err != nil {
		t.Logf("err: %v", err)
	}
	t.Log(string(res))
}

func TestMarshalStruct(t *testing.T) {
	t.Helper()
	s1 := ResData{
		Status: "ok",
		Payload: Payload{
			Objects: []interface{}{
				Product{
					Name:       "Horse",
					Weight:     100.21,
					WeightUnit: "KG",
					OnSale:     false,
				},
			},
			Description: "ok",
		},
	}
	res, err := json.Marshal(s1)
	if err != nil {
		t.Logf("err: %v", err)
	}
	t.Log(string(res))
	var s2 ResData
	err = json.Unmarshal(res, &s2)
	if err != nil {
		t.Logf("err: %v", err)
	}
	t.Logf("s2: %v", s2)

	res, err = json.Marshal(s2)
	if err != nil {
		t.Logf("err: %v", err)
	}
	t.Log(string(res))
}
