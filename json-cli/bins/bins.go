package bins

import "time"

type Bin struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

// BinList представляет собой список файлов (Bin)
type BinList struct {
	Bins []Bin `json:"bins"` // Список всех JSON файлов
}

// NewBin создает новый Bin с заданными параметрами
func NewBin(id string, private bool, name string) Bin {
	return Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

// NewBinList создает новый пустой список BinList
func NewBinList() BinList {
	return BinList{
		Bins: []Bin{},
	}
}
