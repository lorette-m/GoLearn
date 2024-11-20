package main

func main() {
	//FcodesFunc() // Function witch name starts with capital letter - kind of "global" (in fact in package), — Lesson 06

	// UnixCatFunc()

	// 2 ways to run it in PowerShell
	// go run . a.txt b.txt
	// go run . (Get-ChildItem -Filter *.txt | ForEach-Object { $_.FullName })

	//CalcFileSize1()

	/* Output:
	The file has 678 bytes
	The file has 302 bytes
	*/

	CalcFileSize2()

	/* Output:
	11      50     656 C:\Users\raise\GolandProjects\GoLearn\L07\a.txt
	 3      19     296 C:\Users\raise\GolandProjects\GoLearn\L07\b.txt
	*/

	// Exercise: line "Total"
	// done
}
