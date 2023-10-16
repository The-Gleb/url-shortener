package storage

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_storage_AddURL(t *testing.T) {
	var s storage
	s.m = &sync.Map{}
	s.m.Store("id1", "url1")
	s.m.Store("id2", "url2")
	type args struct {
		id  string
		url string
	}
	tests := []struct {
		name string
		s    *storage
		args args
		err  error
	}{
		{
			name: "pos test #1",
			s:    &s,
			args: args{"id1", "url1"},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.AddURL(tt.args.id, tt.args.url)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}
			val, ok := s.m.Load(tt.args.id)
			require.Equal(t, true, ok)
			assert.Equal(t, tt.args.url, val)
		})
	}
}

func Test_storage_GetURL(t *testing.T) {
	var s storage
	s.m = &sync.Map{}
	s.m.Store("id1", "url1")
	s.m.Store("id2", "url2")
	type args struct {
		id string
	}
	tests := []struct {
		name string
		s    *storage
		args args
		want string
		err  error
	}{
		{
			name: "pos test #1",
			s:    &s,
			args: args{"id1"},
			want: "url1",
			err:  nil,
		},
		{
			name: "neg test #1",
			s:    &s,
			args: args{"id0"},
			want: "",
			err:  errors.New("url not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetURL(tt.args.id)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
