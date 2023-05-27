//go:build linux

package fs

var Ctime = func(fi os.FileInfo) int64 {
	stat := fi.Sys().(*syscall.Stat_t)

	return stat.Ctim.Sec
}
