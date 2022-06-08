package defaultmap

import "testing"

func TestNewMap(t *testing.T) {
	var factory DefaultFactory[struct{}] = func() struct{} {
		return struct{}{}
	}

	m := NewMap[string](factory)

	if m.m == nil {
		t.Fatalf(`Internal map is nil`)
	}
	if m.defaultF() != factory() {
		t.Fatalf(`DefaultFactory return values do not match`)
	}
}

func TestInsert(t *testing.T) {
	m := NewMap[string](func() string { return "" })

	m.Insert("key", "value")

	v, ok := m.m["key"]
	if !ok {
		t.Fatalf(`m["key"] does not exist`)
	}
	if v != "value" {
		t.Fatalf(`m["key"] = %q, want "value"`, v)
	}
}

func TestGet(t *testing.T) {
	m := NewMap[string](func() string { return "new string" })

	m.m["exists"] = "yes"

	tests := []struct {
		key  string
		want string
	}{
		{"exists", "yes"},
		{"doesn't exist", "new string"},
	}

	for _, tt := range tests {
		if v := m.Get(tt.key); v != tt.want {
			t.Errorf(`m.Get(%q) = %q, want %q`, tt.key, v, tt.want)
		}
	}

}

func TestGetOr(t *testing.T) {
	m := NewMap[string](func() string { return "new string" })

	m.m["exists"] = "yes"

	tests := []struct {
		key         string
		defaultV    string
		want        string
		existsAfter bool
	}{
		{"exists", "no", "yes", true},
		{"doesn't exist", "new string", "new string", false},
	}

	for _, tt := range tests {
		if v := m.GetOr(tt.key, tt.defaultV); v != tt.want {
			t.Errorf(`m.Get(%q) = %q, want %q`, tt.key, v, tt.want)
		}
		if _, ok := m.m[tt.key]; ok != tt.existsAfter {
			t.Fatalf(`m.m[%q]: ok = %v, want %v`, tt.key, ok, tt.existsAfter)
		}
	}
}

func TestDelete(t *testing.T) {
	m := NewMap[string](func() string { return "new string" })

	m.m["exists"] = "yes"
	m.m["deleted"] = "soon"

	m.Delete("deleted")

	tests := []struct {
		key string
		ok  bool
	}{
		{"exists", true},
		{"deleted", false},
	}

	for _, tt := range tests {
		if _, ok := m.m[tt.key]; ok != tt.ok {
			t.Errorf(`m.m[%q]: ok = %v, want %v`, tt.key, ok, tt.ok)
		}
	}
}

func TestContains(t *testing.T) {
	m := NewMap[string](func() string { return "new string" })

	m.m["exists"] = "yes"

	m.Delete("deleted")

	tests := []struct {
		key      string
		contains bool
		ok       bool
	}{
		{"exists", true, true},
		{"doesn't", false, false},
	}

	for _, tt := range tests {
		if ok := m.Contains(tt.key); ok != tt.contains {
			t.Errorf(`m.Contains(%q) = %v, want %v`, tt.key, ok, tt.contains)
		}
		if _, ok := m.m[tt.key]; ok != tt.ok {
			t.Fatalf(`m.m[%q]: ok = %v, want %v`, tt.key, ok, tt.ok)
		}
	}
}
