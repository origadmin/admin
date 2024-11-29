/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

//func TestSaveConf(t *testing.T) {
//	type args struct {
//		path string
//		conf *Bootstrap
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		{
//			name: "test",
//			args: args{
//				path: "test.toml",
//				conf: DefaultBootstrap(),
//			},
//		},
//		{
//			name: "test",
//			args: args{
//				path: "test.yml",
//				conf: DefaultBootstrap(),
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := SaveConf(tt.args.path, tt.args.conf); (err != nil) != tt.wantErr {
//				t.Errorf("SaveConf() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//		opt := protojson.MarshalOptions{
//			EmitUnpopulated: true,
//			Indent:          " ",
//		}
//		bs, _ := opt.Marshal(DefaultBootstrap())
//		_ = os.WriteFile("test.json", bs, os.ModePerm)
//	}
//}

//func TestLoadConf(t *testing.T) {
//	type args struct {
//		path string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    *configs.Bootstrap
//		wantErr bool
//	}{
//		{
//			name: "test",
//			args: args{
//				path: "test.toml",
//			},
//			want:    configs.DefaultBootstrap(),
//			wantErr: false,
//		},
//		{
//			name: "test",
//			args: args{
//				path: "test.json",
//			},
//			want:    configs.DefaultBootstrap(),
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := LoadConf(tt.args.path)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("LoadConf() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("LoadConf() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
