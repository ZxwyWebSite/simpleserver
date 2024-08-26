package app

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"unsafe"

	"github.com/gin-gonic/gin"
)

var (
	wdir, edir string
	host, port string
	sdir, fdir string
)

func Init() {
	wdir, _ = os.Getwd()
	edir, _ = os.Executable()
	edir = filepath.Dir(edir)
	flag.StringVar(&host, `h`, `0.0.0.0`, `Host`)
	flag.StringVar(&port, `p`, `8100`, `Port`)
	flag.StringVar(&sdir, `s`, `/`, `Sdir`)
	flag.StringVar(&fdir, `f`, `.`, `Fdir`)
	flag.Parse()
	os.Stdout.WriteString("  ____  _                 _         _   _   \n / ___|(_)_ __ ___  _ __ | | ___   / \\ | |  \n \\___ \\| | '_ ` _ \\| '_ \\| |/ _ \\ /_ _\\| |  \n  ___) | | | | | | | |_) | |  __/  | |_| |_ \n |____/|_|_| |_| |_| .__/|_|\\___|  | |\\   / \n                   |_|             |_| \\_/  \n============================================\n")
	os.Stdout.WriteString(" Welcome To Simple Server v1.0.0 ↑↓\n\n")
	fmt.Fprintf(os.Stdout, "Wdir: %v\nEdir: %v\nHost: %v\nPort: %v\nSdir: %v\nFdir: %v\n\n", wdir, edir, host, port, sdir, fdir)
	// gin.SetMode(gin.ReleaseMode)
}

func Main() {
	g := gin.Default()
	var rdir string
	if filepath.IsAbs(fdir) {
		rdir = fdir
	} else {
		rdir = filepath.Join(wdir, fdir)
	}
	g.StaticFS(sdir, gin.Dir(rdir, true))
	// g.Static(sdir, rdir)
	b := make([]byte, len(host)+1+len(port))
	p := copy(b, host)
	b[p] = ':'
	copy(b[p+1:], port)
	addr := unsafe.String(unsafe.SliceData(b), len(b))
	fmt.Fprintf(os.Stdout, "Rdir: %v\nAddr: %v\n\n", rdir, addr)
	e := g.Run(addr)
	if e != nil {
		fmt.Fprintln(os.Stdout, e)
	}
}
