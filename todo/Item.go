package todo

import "time"

type Item struct{
    name string
    isDone bool
    dateOfCreation time.Time
}

func (i Item)NewItem(name string) Item{
    return Item{
        name: name,
        isDone: false,
        dateOfCreation: time.Now(),
    }
}

func (i* Item)ItemDone(){
    i.isDone = true
}


