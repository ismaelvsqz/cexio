// Copyright 2021 Juan Ismael Vasquez <ismael.juan@gmail.com>.
// All rights reserved. Use of this source code is governed by
// a MIT style license that can be found in the LICENSE file.

package cexio

import "testing"

func Test_nonce(t *testing.T) {
	t.Run("nonce", func(t *testing.T) {
		if got := nonce(); len(got) <= 0 {
			t.Errorf("nonce() = %v, want length greater than 0", got)
		}
	})
}

func Test_toHmac256(t *testing.T) {
	type args struct {
		message string
		secret  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"hello", args{"hello", "123456"}, "AC28D602C767424D0C809EDEBF73828BED5CE99CE1556F4DF8E223FAEEC60EDD"},
		{"emptyMessage", args{"", "123456"}, "B946CCC987465AFCDA7E45B1715219711A13518D1F1663B8C53B848CB0143441"},
		{"emptySecret", args{"hello", ""}, "4352B26E33FE0D769A8922A6BA29004109F01688E26ACC9E6CB347E5A5AFC4DA"},
		{"empty", args{"", ""}, "B613679A0814D9EC772F95D778C35FC5FF1697C493715653C6C712144292C5AD"},
		{"undefined", args{}, "B613679A0814D9EC772F95D778C35FC5FF1697C493715653C6C712144292C5AD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toHmac256(tt.args.message, tt.args.secret); got != tt.want {
				t.Errorf("toHmac256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasText(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid", args{"xxx"}, true},
		{"emtpy", args{""}, false},
		{"undefined", args{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasText(tt.args.s); got != tt.want {
				t.Errorf("hasText() = %v, want %v", got, tt.want)
			}
		})
	}
}
