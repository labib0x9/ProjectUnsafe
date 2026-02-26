package main

import (
	"html/template"
	"log"
	"log/slog"
	"net/http"
)

type Lab struct {
	Id          int
	Title       string
	Badge       string
	Description string
	Hints       []string
}

// var LinuxLabs []Lab

type TopicPage struct {
	Label    string
	Name     string
	Category string
	Labs     []Lab
}

var LINUX TopicPage

// var XssLabs []Lab
// var SqliLabs []Lab
// var SegfaultLabs []Lab

type Topic struct {
	TopicUrl         string
	BoxTopicReveal   string
	BoxTag           string
	Topic            string
	ShortDescription string
	LabCount         int
}

var AllTopics []Topic

func init() {
	xss := Topic{
		TopicUrl:         "/xss",
		BoxTopicReveal:   "box box-xss reveal",
		BoxTag:           "Web Security",
		Topic:            "XSS",
		LabCount:         0,
		ShortDescription: "UPCOMING",
	}

	sqli := Topic{
		TopicUrl:         "/sqli",
		BoxTopicReveal:   "box box-sql reveal",
		BoxTag:           "Database Attack",
		Topic:            "SQL Injection",
		LabCount:         0,
		ShortDescription: "UPCOMING",
	}

	segfault := Topic{
		TopicUrl:         "/segfault",
		BoxTopicReveal:   "box box-seg reveal",
		BoxTag:           "Memory Error",
		Topic:            "Segmentation Fault",
		LabCount:         0,
		ShortDescription: "UPCOMING",
	}

	linux := Topic{
		TopicUrl:         "/linux",
		BoxTopicReveal:   "box box-linux reveal",
		BoxTag:           "Operating System",
		Topic:            "Linux",
		LabCount:         3,
		ShortDescription: "Get comfortable — it's your primary weapon.",
	}

	AllTopics = append(AllTopics, []Topic{xss, sqli, segfault, linux}...)

	linuxLab1 := Lab{
		Id:    1,
		Title: "Filesystem Navigation &amp; File Permissions",
		Badge: "EASY",
		Description: `You are dropped into a shell on a Linux machine. A flag is hidden somewhere in
          the filesystem inside a file called <code>flag.txt</code>. Using only shell commands,
          locate the file, understand its permissions, and read its contents. Along the way
          you will encounter files owned by other users and directories you cannot enter —
          learn to read permission strings and work around restrictions.`,
		Hints: []string{
			`Use <code>find / -name "flag.txt" 2>/dev/null</code> to search the whole filesystem while suppressing permission-denied errors. Once found, use <code>ls -la</code> to check permissions and <code>cat</code> or <code>less</code> to read it.`,
		},
	}

	linuxLab2 := Lab{
		Id:    1,
		Title: "Filesystem Navigation &amp; File Permissions",
		Badge: "EASY",
		Description: `You are dropped into a shell on a Linux machine. A flag is hidden somewhere in
          the filesystem inside a file called <code>flag.txt</code>. Using only shell commands,
          locate the file, understand its permissions, and read its contents. Along the way
          you will encounter files owned by other users and directories you cannot enter —
          learn to read permission strings and work around restrictions.`,
		Hints: []string{
			`Use <code>find / -name "flag.txt" 2>/dev/null</code> to search the whole filesystem while suppressing permission-denied errors. Once found, use <code>ls -la</code> to check permissions and <code>cat</code> or <code>less</code> to read it.`,
		},
	}

	linuxLab3 := Lab{
		Id:    1,
		Title: "Filesystem Navigation &amp; File Permissions",
		Badge: "EASY",
		Description: `You are dropped into a shell on a Linux machine. A flag is hidden somewhere in
          the filesystem inside a file called <code>flag.txt</code>. Using only shell commands,
          locate the file, understand its permissions, and read its contents. Along the way
          you will encounter files owned by other users and directories you cannot enter —
          learn to read permission strings and work around restrictions.`,
		Hints: []string{
			`Use <code>find / -name "flag.txt" 2>/dev/null</code> to search the whole filesystem while suppressing permission-denied errors. Once found, use <code>ls -la</code> to check permissions and <code>cat</code> or <code>less</code> to read it.`,
		},
	}

	LINUX = TopicPage{
		Label:    "OS & Shell · Linux",
		Name:     "Linux",
		Category: "Fundamentals",
		Labs:     []Lab{linuxLab1, linuxLab2, linuxLab3},
	}
}

var funcMap = template.FuncMap{
	"add": func(a, b int) int { return a + b },
}

var homeTmpl = template.Must(template.New("index.html").Funcs(funcMap).ParseFiles("templates/index.html"))
var linuxTmpl = template.Must(template.New("linux.html").Funcs(funcMap).ParseFiles("templates/linux.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {

	slog.Info("/	Path=" + r.URL.Path + "   " + r.Method)

	if r.Method != http.MethodGet {
		return
	}

	if r.URL.Path != "/" {
		return
	}

	homeTmpl.Execute(
		w,
		map[string]any{
			"TopicItems": AllTopics,
		},
	)
}

func xssHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/upcoming", 302)
}

func sqliHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/upcoming", 302)
}

func segfaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/upcoming", 302)
}

func linuxHandler(w http.ResponseWriter, r *http.Request) {
	linuxTmpl.Execute(
		w,
		map[string]any{
			"Label":    LINUX.Label,
			"Name":     LINUX.Name,
			"Category": LINUX.Category,
			"Labs":     LINUX.Labs,
		},
	)
}

func upcomingHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "UPCOMING....")
	http.ServeFile(w, r, "./www/upcoming.html")
}

func main() {

	// fs := http.FileServer(http.Dir("./www"))
	sfs := http.FileServer(http.Dir("./www/static"))
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/xss", xssHandler)
	mux.HandleFunc("/sqli", sqliHandler)
	mux.HandleFunc("/segfault", segfaultHandler)
	mux.HandleFunc("/linux", linuxHandler)
	mux.HandleFunc("/upcoming", upcomingHandler)

	mux.Handle("/static/", http.StripPrefix("/static", sfs))
	// mux.Handle("/favicon.ico", http.StripPrefix("/", fs))

	slog.Info("Starting Server at http://127.0.0.1:8080/")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
