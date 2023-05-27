package internal

import "zakirullin/dumpbot/internal/fs"

func shouldSplitChecklist(checklist string) bool {
	for _, unsplittableChecklist := range []string{fs.DirRead, fs.DirWatch} {
		if checklist == unsplittableChecklist {
			return false
		}
	}
	return true
}
