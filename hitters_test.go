package hitters

import (
	"reflect"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want *Hitters
	}{
		{
			name: "Empty Hitters",
			args: args{k: 0},
			want: &Hitters{},
		},
		{
			name: "Correct constructor call",
			args: args{k: 1},
			want: &Hitters{capacity: 1, items: make(map[string]int)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHitters() = %v, want %v", got, tt.want)
			}
			if tt.args.k == 0 && err == nil { // If k == 0 then err should not be nil
				t.Errorf("NewHitters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHitters_addOne(t *testing.T) {
	type fields struct {
		items          map[string]int
		capacity       int
		processedCount int
		RWMutex        sync.RWMutex
	}
	type args struct {
		item string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Add a single element",
			fields: fields{
				items:    make(map[string]int),
				capacity: 3,
			},
			args: args{"anitem"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := &Hitters{
				items:          tt.fields.items,
				capacity:       tt.fields.capacity,
				processedCount: tt.fields.processedCount,
				RWMutex:        tt.fields.RWMutex,
			}
			tk.addOne(tt.args.item)
			if len(tk.items) != 1 || tk.items["anitem"] != 1 {
				t.Errorf("Adding item to Hitters failed")
			}
		})
	}
}

func TestHitters_Add(t *testing.T) {
	type fields struct {
		items          map[string]int
		capacity       int
		processedCount int
		RWMutex        sync.RWMutex
	}
	type args struct {
		items []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Add multiple elements",
			fields: fields{
				items:    make(map[string]int),
				capacity: 3,
			},
			args: args{[]string{"anitem", "anotheritem", "5", "anitem"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := &Hitters{
				items:          tt.fields.items,
				capacity:       tt.fields.capacity,
				processedCount: tt.fields.processedCount,
				RWMutex:        tt.fields.RWMutex,
			}
			tk.Add(tt.args.items...)
			if tk.items["anitem"] != 1 {
				t.Errorf("multi-item add error %#v\n", tk.items)
			}
		})
	}
}

func TestHitters_Get(t *testing.T) {
	type fields struct {
		items          map[string]int
		capacity       int
		processedCount int
		RWMutex        sync.RWMutex
	}
	type args struct {
		k string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Add a single element",
			fields: fields{
				items:    make(map[string]int),
				capacity: 3,
			},
			args: args{"anitem"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := &Hitters{
				items:          tt.fields.items,
				capacity:       tt.fields.capacity,
				processedCount: tt.fields.processedCount,
				RWMutex:        tt.fields.RWMutex,
			}
			tk.Add("anitem")
			if got := tk.Get(tt.args.k); got != tt.want {
				t.Errorf("Hitters.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHitters_Items(t *testing.T) {
	type fields struct {
		items          map[string]int
		capacity       int
		processedCount int
		RWMutex        sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]int
	}{
		{
			name: "items getter",
			fields: fields{
				items: map[string]int{
					"abc": 1,
					"b":   2,
					"5":     3,
				},
				capacity: 10,
			},
			want: map[string]int{
				"abc": 1,
				"b":   2,
				"5":     3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := &Hitters{
				items:          tt.fields.items,
				capacity:       tt.fields.capacity,
				processedCount: tt.fields.processedCount,
				RWMutex:        tt.fields.RWMutex,
			}
			if got := tk.Items(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hitters.Items() = %v, want %v", got, tt.want)
			}
		})
	}
}
