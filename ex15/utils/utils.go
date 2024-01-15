package utils

import (
	"container/list"
	"fmt"
)

type Entry struct {
	Label     string
	Operation string
	FocalLength       int
	BoxHash   int
}

func (e Entry) String() string {
	return fmt.Sprintf("LABEL <%s> FocalLength<%d>", e.Label, e.FocalLength)
}


//k
type HashMap struct {
	Fields [256]Box
}


func (h *HashMap) Print(idx int) {
    if idx < 0 || idx >= len(h.Fields) {
        fmt.Println("Invalid index")
        return
    }

    fmt.Printf("BOX %d: %v\n", idx, &h.Fields[idx])
}


func GetNewHashMap() HashMap {
	newHashMap := HashMap{}
	for idx := range newHashMap.Fields {
		newHashMap.Fields[idx] = getNewBox()
	}
	return newHashMap
}

func (h *HashMap) Add(e Entry) {
	
	currentBox := h.Fields[e.BoxHash]
	foundElement := currentBox.Find(e)
	if foundElement != nil {

		tempEntry := foundElement.Value.(Entry)
		tempEntry.FocalLength = e.FocalLength
		foundElement.Value = tempEntry
	} else {
		currentBox.Add(e)
	}
	
}

func (h *HashMap) Remove(e Entry) {
	currentBox := h.Fields[e.BoxHash]
	foundElement := currentBox.Find(e)
	if foundElement != nil {
		currentBox.Remove(foundElement)
	}
}


type Box struct {
	List *list.List
}

func getNewBox() Box {
	return Box{List: list.New()}
}

func (l *Box) String() string {
	result := "BoxData -> ["

	for e := l.List.Front(); e != nil; e = e.Next() {
		entry := e.Value.(Entry)
		result += fmt.Sprintf("{ %v }, ", entry)
	}

	result += "]\n"
	return result
}

func (l *Box) Add(e Entry) {
	l.List.PushBack(e)
}

func (l *Box) Remove(e *list.Element) {
	l.List.Remove(e)
}

func (l *Box) Find(entry Entry) *list.Element{
	for e := l.List.Front(); e != nil; e = e.Next() {
		if e.Value.(Entry).Label == entry.Label {
			return e
		}
	}
	return nil
}

