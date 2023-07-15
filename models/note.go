package models

type UserNote struct {
	ID    uint64 `gorm:"primaryKey"`
	Notes string `gorm:"size:255"`
}

func NotesAll(user *UserS) *[]UserNote {
	var notes2 []UserNote
	DB.Table(NotesTable).Select("id", "notes").Where("deleted_at is NULL and user_id = ?", user.ID).Find(&notes2)
	return &notes2
}

func NoteCreate(user *UserS, notes string) *Note {
	entry := Note{Notes: notes, UserID: user.ID}
	DB.Create(&entry)
	return &entry
}

func NotesFind(user *UserS, id uint64) *Note {
	var note Note
	DB.Where("id = ? and user_id = ?", id, user.ID).First(&note)
	return &note
}

func NotesMarkDelete(user *UserS, id uint64) {
	DB.Where("id = ? and user_id = ?", id, user.ID).Delete(&Note{})
}

func LastNoteID(user *UserS) *Note {
	var note Note
	DB.Last(&note)
	return &note
}
