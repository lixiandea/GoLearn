package musicLib

import (
	"errors"
)

type MusicManager struct {
	musics [] MusicEntry
}

func NewMusicManager() * MusicManager {
	return &MusicManager{make([] MusicEntry, 0)}
}

func (m *MusicManager)Len() int {
	return len(m.musics)
}

func (m * MusicManager) Get(index int) (music * MusicEntry,err error) {
	if index > m.Len(){
		music = nil
		err = errors.New("out of range")
	}else {
		music = &m.musics[index]
		err = nil
	}
	return music, err
}

func (m *MusicManager)Find(name *string) (music *MusicEntry, index int ) {
	if m.Len() == 0 {
		music, index = nil, -1
	}
	for i, m := range m.musics{
		if m.Name == *name{
			music, index = &m, i
		}
	}
	return music, index
}

func (m * MusicManager) Add(music * MusicEntry)  {
	m.musics = append(m.musics, * music)
}

func (m * MusicManager) Remove(index int)  (music * MusicEntry){
	if m.Len()< index || index < 0{
		music = nil
	}
	music = &m.musics[index]
	m.musics = append(m.musics[:index], m.musics[index+1:]...)
	return music
}

func (m * MusicManager) RemoveByName(name string)  (music * MusicEntry, err error){
	music, index := m.Find(&name)
	if index < 0 {
		err = errors.New("Can't find music want to remove")
	}else {
		m.Remove(index)
	}
	return music, err
}



