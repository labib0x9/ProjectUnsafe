package main

// import (
// 	"encoding/json"
// 	"html/template"
// 	"log"
// 	"log/slog"
// 	"net/http"
// )

// type IO struct {
// 	Input  string
// 	Output string
// }

// type CProblem struct {
// 	Id          int
// 	Title       string
// 	Description string
// 	Examples    []IO
// 	Hints       []string
// 	Solution    string
// 	StarterC    string
// }

// type PageData struct {
// 	Problem  CProblem
// 	Problems []CProblem // sidebar list
// 	WSUrl    string     // e.g. "ws://localhost:8080/ws"
// 	solved   map[int]bool
// }

// // status = 0 (inactive), 1 (active)
// type Lab struct {
// 	Id          int
// 	Title       string
// 	Badge       string
// 	Description string
// 	Hints       []string
// 	Status      bool
// }

// // flag = 1 (start), 2 (reset), 3 (terminate)
// type LabHandler struct {
// 	Id   int `json:"lab_id"`
// 	Flag int `json:"flag"`
// }

// // var LinuxLabs []Lab

// type TopicPage struct {
// 	Label    string
// 	Name     string
// 	Category string
// 	Labs     []Lab
// }

// var LINUX, CCODE TopicPage
// var CPLAYGORUND_PAGE PageData

// // var XssLabs []Lab
// // var SqliLabs []Lab
// // var SegfaultLabs []Lab

// type Topic struct {
// 	TopicUrl         string
// 	BoxTopicReveal   string
// 	BoxTag           string
// 	Topic            string
// 	ShortDescription string
// 	LabCount         int
// }

// var AllTopics []Topic

// func init() {
// 	xss := Topic{
// 		TopicUrl:         "/xss",
// 		BoxTopicReveal:   "box box-xss reveal",
// 		BoxTag:           "Web Security",
// 		Topic:            "XSS",
// 		LabCount:         0,
// 		ShortDescription: "UPCOMING",
// 	}

// 	sqli := Topic{
// 		TopicUrl:         "/sqli",
// 		BoxTopicReveal:   "box box-sql reveal",
// 		BoxTag:           "Database Attack",
// 		Topic:            "SQL Injection",
// 		LabCount:         0,
// 		ShortDescription: "UPCOMING",
// 	}

// 	segfault := Topic{
// 		TopicUrl:         "/segfault",
// 		BoxTopicReveal:   "box box-seg reveal",
// 		BoxTag:           "Memory Error",
// 		Topic:            "Segmentation Fault",
// 		LabCount:         0,
// 		ShortDescription: "UPCOMING",
// 	}

// 	linux := Topic{
// 		TopicUrl:         "/linux",
// 		BoxTopicReveal:   "box box-linux reveal",
// 		BoxTag:           "Operating System",
// 		Topic:            "Linux",
// 		LabCount:         3,
// 		ShortDescription: "Get comfortable — it's your primary weapon.",
// 	}

// 	ccode := Topic{
// 		TopicUrl:         "/ccode",
// 		BoxTopicReveal:   "box box-ccode reveal",
// 		BoxTag:           "C Code",
// 		Topic:            "C Programming",
// 		LabCount:         1,
// 		ShortDescription: "Learn how program works - Coding...",
// 	}

// 	// Dont comment / remove
// 	_ = xss
// 	_ = sqli
// 	_ = segfault
// 	AllTopics = append(AllTopics, []Topic{linux, ccode}...)

// 	linuxLab1 := Lab{
// 		Id:          1,
// 		Title:       "Connect using ssh",
// 		Badge:       "EASY",
// 		Description: `You just need to connect using ssh`,
// 		Hints: []string{
// 			`Use ssh to connect.. ssh runs on port 22`,
// 		},
// 	}

// 	linuxLab2 := Lab{
// 		Id:    2,
// 		Title: "Filesystem Navigation and File Permissions",
// 		Badge: "EASY",
// 		Description: `You are dropped into a shell on a Linux machine. A flag is hidden somewhere in
//           the filesystem inside a file called <code>flag.txt</code>. Using only shell commands,
//           locate the file, understand its permissions, and read its contents. Along the way
//           you will encounter files owned by other users and directories you cannot enter —
//           learn to read permission strings and work around restrictions.`,
// 		Hints: []string{
// 			`Use <code>find / -name "flag.txt" 2>/dev/null</code> to search the whole filesystem while suppressing permission-denied errors. Once found, use <code>ls -la</code> to check permissions and <code>cat</code> or <code>less</code> to read it.`,
// 		},
// 	}

// 	linuxLab3 := Lab{
// 		// Id:    1,
// 		// Title: "Filesystem Navigation &amp; File Permissions",
// 		// Badge: "EASY",
// 		// Description: `You are dropped into a shell on a Linux machine. A flag is hidden somewhere in
// 		//   the filesystem inside a file called <code>flag.txt</code>. Using only shell commands,
// 		//   locate the file, understand its permissions, and read its contents. Along the way
// 		//   you will encounter files owned by other users and directories you cannot enter —
// 		//   learn to read permission strings and work around restrictions.`,
// 		// Hints: []string{
// 		// 	`Use <code>find / -name "flag.txt" 2>/dev/null</code> to search the whole filesystem while suppressing permission-denied errors. Once found, use <code>ls -la</code> to check permissions and <code>cat</code> or <code>less</code> to read it.`,
// 		// },
// 	}

// 	cLab1 := Lab{
// 		Id:          1,
// 		Title:       "Playground",
// 		Badge:       "...",
// 		Description: `A container sandbox to run c code`,
// 		Hints: []string{
// 			`No Hints..`,
// 		},
// 	}

// 	LINUX = TopicPage{
// 		Label:    "OS & Shell · Linux",
// 		Name:     "Linux",
// 		Category: "Fundamentals",
// 		Labs:     []Lab{linuxLab1, linuxLab2, linuxLab3},
// 	}

// 	CCODE = TopicPage{
// 		Label:    "Programming",
// 		Name:     "C Programming",
// 		Category: "",
// 		Labs:     []Lab{cLab1},
// 	}

// 	cp_prob1 := CProblem{
// 		Id:          1,
// 		Title:       "Two Sum",
// 		Description: "Given an array of integers <strong>nums</strong> and an integer <strong>target</strong>, return the <strong>indices</strong> of the two numbers such that they add up to target. Exactly one solution exists.",
// 		Examples: []IO{
// 			{Input: "nums=[2,7,11,15], target=9", Output: "[0,1]"},
// 			{Input: "nums=[3,2,4], target=6", Output: "[1,2]"},
// 		},
// 		Hints: []string{
// 			"brute force O(n²): check every pair. can you do better?",
// 			"for each x, you need (target-x). store seen values in a hash map.",
// 			"one pass: check if complement exists in map, else store current index.",
// 		},
// 		Solution: "hash map one-pass. for each num check if complement exists. O(n) time, O(n) space.",
// 		StarterC: `#include <stdio.h>\n#include <stdlib.h>\n\nint main() {\n   return 0;\n}`,
// 	}

// 	_ = cp_prob1
// 	CPLAYGORUND_PAGE = PageData{
// 		Problem:  cp_prob1,
// 		Problems: []CProblem{cp_prob1},
// 		WSUrl:    "ws://localhost:8080/ws",
// 		solved:   map[int]bool{},
// 	}

// }

// var funcMap = template.FuncMap{
// 	// {{add $i 1}}
// 	"add": func(a, b int) int { return a + b },

// 	// {{.Problem | toJSON}}
// 	"toJSON": func(v any) (template.JS, error) {
// 		b, err := json.Marshal(v)
// 		if err != nil {
// 			return "", err
// 		}
// 		return template.JS(b), nil
// 	},
// }

// var homeTmpl = template.Must(template.New("index.html").Funcs(funcMap).ParseFiles("templates/index.html"))
// var linuxTmpl = template.Must(template.New("linux.html").Funcs(funcMap).ParseFiles("templates/linux.html"))
// var cTmpl = template.Must(template.New("ccode.html").Funcs(funcMap).ParseFiles("templates/ccode.html"))
// var editorTmpl = template.Must(template.New("editor.html").Funcs(funcMap).ParseFiles("templates/editor.html"))

// func homeHandler(w http.ResponseWriter, r *http.Request) {

// 	slog.Info("/	Path=" + r.URL.Path + "   " + r.Method)

// 	if r.Method != http.MethodGet {
// 		return
// 	}

// 	if r.URL.Path != "/" {
// 		return
// 	}

// 	homeTmpl.Execute(
// 		w,
// 		map[string]any{
// 			"TopicItems": AllTopics,
// 		},
// 	)
// }

// func xssHandler(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "/upcoming", 302)
// }

// func sqliHandler(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "/upcoming", 302)
// }

// func segfaultHandler(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "/upcoming", 302)
// }

// func linuxHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		linuxTmpl.Execute(
// 			w,
// 			map[string]any{
// 				"Label":    LINUX.Label,
// 				"Name":     LINUX.Name,
// 				"Category": LINUX.Category,
// 				"Labs":     LINUX.Labs,
// 			},
// 		)
// 	case http.MethodPost:
// 		// linuxTmpl.Execute(
// 		// 	w,
// 		// 	map[string]any{
// 		// 		"Label":    LINUX.Label,
// 		// 		"Name":     LINUX.Name,
// 		// 		"Category": LINUX.Category,
// 		// 		"Labs":     LINUX.Labs,
// 		// 	},
// 		// )

// 		var labHndlr LabHandler
// 		decoder := json.NewDecoder(r.Body)
// 		if err := decoder.Decode(&labHndlr); err != nil {
// 			slog.Error("JSON Decoding failed")
// 			http.Error(w, "give proper json", 400)
// 			return
// 		}
// 		slog.Info("Lab Id", "lab_id", labHndlr.Id)

// 		switch labHndlr.Flag {
// 		case 1: // Start
// 			encoder := json.NewEncoder(w)
// 			encoder.Encode(
// 				map[string]string{
// 					"lab_url": "http://127.0.0.1:8383",
// 				},
// 			)
// 		case 2: // Reset
// 		case 3: // Terminate
// 		default:
// 			http.Error(w, "Flag not allowed", 400)
// 		}
// 	default:
// 		http.Error(w, "method not allowed", 400)
// 	}
// }

// func codeHandler(w http.ResponseWriter, r *http.Request) {
// 	// http.Redirect(w, r, "/upcoming", 302)
// 	switch r.Method {
// 	case http.MethodGet:
// 		cTmpl.Execute(
// 			w,
// 			map[string]any{
// 				"Label":    CCODE.Label,
// 				"Name":     CCODE.Name,
// 				"Category": CCODE.Category,
// 				"Labs":     CCODE.Labs,
// 			},
// 		)
// 	case http.MethodPost:
// 		var labHndlr LabHandler
// 		decoder := json.NewDecoder(r.Body)
// 		if err := decoder.Decode(&labHndlr); err != nil {
// 			slog.Error("JSON Decoding failed")
// 			http.Error(w, "give proper json", 400)
// 			return
// 		}
// 		slog.Info("Lab Id", "lab_id", labHndlr.Id)

// 		switch labHndlr.Flag {
// 		case 1: // Start
// 			encoder := json.NewEncoder(w)
// 			encoder.Encode(
// 				map[string]string{
// 					"lab_url": "/playground",
// 				},
// 			)
// 		case 2: // Reset
// 		case 3: // Terminate
// 		default:
// 			http.Error(w, "Flag not allowed", 400)
// 		}
// 	default:
// 		http.Error(w, "Method not allowed", 400)
// 	}
// }

// func codePlayground(w http.ResponseWriter, r *http.Request) {
// 	// http.ServeFile(w, r, "./www/tmp/ccode.html")
// 	switch r.Method {
// 	case http.MethodGet:
// 		editorTmpl.Execute(
// 			w,
// 			CPLAYGORUND_PAGE,
// 		)
// 	case http.MethodPost:
// 	default:
// 		http.Error(w, "Method not allowed", 400)
// 	}
// }

// func upcomingHandler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintln(w, "UPCOMING....")
// 	http.ServeFile(w, r, "./www/upcoming.html")
// }

// func main() {

// 	// fs := http.FileServer(http.Dir("./www"))
// 	sfs := http.FileServer(http.Dir("./www/static"))
// 	mux := http.NewServeMux()

// 	// mux.HandleFunc("/", homeHandler)
// 	// mux.HandleFunc("/xss", xssHandler)
// 	// mux.HandleFunc("/sqli", sqliHandler)
// 	// mux.HandleFunc("/segfault", segfaultHandler)
// 	// mux.HandleFunc("/linux", linuxHandler)
// 	// mux.HandleFunc("/upcoming", upcomingHandler)
// 	// mux.HandleFunc("/ccode", codeHandler)
// 	// mux.HandleFunc("/ccode/playground", codePlayground)

// 	// mux.HandleFunc("/login", userLogin)
// 	// mux.HandleFunc("/login-anon", anonLogin)
// 	// mux.HandleFunc("/login-admin", adminLogin)
// 	// mux.HandleFunc("/profile", profileHandle)
// 	// mux.HandleFunc("/logout", userLogout)

// 	mux.Handle("/static/", http.StripPrefix("/static", sfs))
// 	// mux.Handle("/favicon.ico", http.StripPrefix("/", fs))

// 	slog.Info("Starting Server at http://127.0.0.1:8080/")
// 	log.Fatal(http.ListenAndServe(":8080", mux))
// }

// mux.Handle("POST /auth/login', { method: 'POST', body: JSON.stringify({ username, password }) })
// mux.Handle("POST /auth/signup', { method: 'POST', body: JSON.stringify({ username, email, password }) })
// mux.Handle("POST /auth/anonymous', { method: 'POST' })
// mux.Handle("POST /auth/reset-password', { method: 'POST', body: JSON.stringify({ email }) })
// mux.Handle("GET /labs')
// mux.Handle("GET /lab/${id}`)
// mux.Handle("POST /lab/start', { method: 'POST', body: JSON.stringify({ labId }) })
// mux.Handle("POST /lab/reset', { method: 'POST', body: JSON.stringify({ containerId }) })
// mux.Handle("POST /lab/terminate', { method: 'POST', body: JSON.stringify({ containerId }) })
// mux.Handle("GET /lab/hints?labId=${labId}`)
// mux.Handle("GET /problems')
// mux.Handle("GET /problem/${id}`)
// mux.Handle("POST /code/run-custom', { method: 'POST', body: JSON.stringify({ problemId, code, customCommand }) })
// mux.Handle("POST /code/run', { method: 'POST', body: JSON.stringify({ problemId, code, customCommand }) })
// mux.Handle("GET /admin/containers')
// mux.Handle("GET /admin/users')
// mux.Handle("POST /admin/terminate', { method: 'POST', body: JSON.stringify({ containerId }) })

// Lab APIs
// mux.Handle("GET /labs", http.HandlerFunc(handlers.GetAllLabs))
// mux.Handle("GET /lab/${id}", http.HandlerFunc())
// mux.Handle("POST /lab/start", http.HandlerFunc())
// mux.Handle("POST /lab/reset'", http.HandlerFunc())
// mux.Handle("POST /lab/terminate'", http.HandlerFunc())
// mux.Handle("GET /lab/hints?labId=${labId}", http.HandlerFunc())

// Auth APIs
// mux.Handle("POST /auth/login", http.HandlerFunc())
// mux.Handle("POST /auth/signup", http.HandlerFunc())
// mux.Handle("POST /auth/anonymous", http.HandlerFunc(handlers.AnonLogin))
// mux.Handle("POST /auth/reset-password", http.HandlerFunc())

// Playground APIs
// mux.Handle("GET /problems", http.HandlerFunc())
// mux.Handle("GET /problem/${id}", http.HandlerFunc())
// mux.Handle("POST /code/run-custom", http.HandlerFunc())
// mux.Handle("POST /code/run", http.HandlerFunc())

// Admin APIs
// mux.Handle("GET /admin/containers", http.HandlerFunc())
// mux.Handle("GET /admin/users", http.HandlerFunc())
// mux.Handle("POST /admin/terminate", http.HandlerFunc())

// // Lab APIs
// GET /labs
// GET /lab/${id}
// POST /lab/start
// POST /lab/reset
// POST /lab/terminate
// GET /lab/hints?labId=${labId}

// // Auth APIs
// POST /auth/login
// POST /auth/signup
// POST /auth/anonymous
// POST /auth/reset-password

// // Playground APIs
// GET /problems
// GET /problem/${id}
// POST /code/run-custom
// POST /code/run

// // Admin APIs
// GET /admin/containers
// GET /admin/users
// POST /admin/terminate
