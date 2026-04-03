package cmd

// func init() {
// 	lab1 := model.Lab{
// 		Id:          "lab-ssh-basics",
// 		Title:       "SSH Basics",
// 		Difficulty:  "easy",
// 		Category:    "networking",
// 		Description: "Learn to connect to remote machines, manage keys, and tunnel ports using SSH.",
// 		LongDescription: `In this lab you will practice the fundamentals of SSH. You will:
// - Connect to a remote host using password and key-based auth
// - Generate and manage SSH key pairs
// - Use SSH tunneling to forward local ports
// - Explore the ~/.ssh/config file

// The container runs a minimal Debian image with sshd configured and ready.`,
// 		Hints: []string{
// 			"Use ssh-keygen -t ed25519 to generate a modern key pair",
// 			"Copy your public key with ssh-copy-id user@host",
// 			"Check /etc/ssh/sshd_config for server settings",
// 		},
// 		Completions:   0,
// 		EstimatedTime: "45 min",
// 		Tags:          []string{"ssh", "networking", "linux"},
// 	}

// 	lab2 := model.Lab{
// 		Id:          "lab-fd-hunt-01",
// 		Title:       "FD Hunt",
// 		Difficulty:  "easy",
// 		Category:    "linux",
// 		Description: "Learn to find opened fd by a process.",
// 		LongDescription: `In this lab you will practice the fundamentals of /proc folder. You will:
// - Find opened file by monitoring /proc folder
// - No need to use lsof tool
// - Submit how many files are opened by the process.

// The container runs a minimal Debian image with sshd configured and ready.`,
// 		Hints: []string{
// 			"You will need pid of that process",
// 			"monitor /proc/<pid> directory",
// 		},
// 		Completions:   0,
// 		EstimatedTime: "45 min",
// 		Tags:          []string{"ssh", "process", "linux"},
// 	}

// 	lab3 := model.Lab{
// 		Id:          "lab-process-01",
// 		Title:       "Process Status",
// 		Difficulty:  "easy",
// 		Category:    "linux",
// 		Description: "Learn to Monitor process status.",
// 		LongDescription: `In this lab you will practice theof /proc folder. You will:
// - Find opened file by monitoring /proc folder
// - No need to use lsof tool
// - You will submit how much ram, cpu usage by the process

// The container runs a minimal Debian image with sshd configured and ready.`,
// 		Hints: []string{
// 			"You will need pid of that process",
// 			"monitor /proc/<pid> directory",
// 		},
// 		Completions:   0,
// 		EstimatedTime: "45 min",
// 		Tags:          []string{"ssh", "process", "linux"},
// 	}

// 	model.LabList = append(model.LabList, lab1, lab2, lab3)
// }
