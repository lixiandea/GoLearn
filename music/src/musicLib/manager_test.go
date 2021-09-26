package musicLib

import "testing"

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("New MusicManger failed")
	}
	if mm.Len() != 0 {
		t.Error("New MusicManager fail, not empty")
	}

	m0 := &MusicEntry{"1", "My heart will go on", "celion dion", "music.qq.com", "Pop"}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("Add MusicEntry failed")
	}

	m := mm.Find(&m0.Name)
	if m == nil {
		t.Error("find MusicEntry failed")
	}

	if m.Id != m0.Id || m.Artist != m.Artist || m.Type != m.Type || m.Name != m.Name {
		t.Error("find MusicEntry failed, wrong music")
	}
	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManager get failed.", err)
	}
	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove failed.", err)
	}

}
