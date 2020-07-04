package main

import (
	"fmt"
	"strings"
	"os/exec"
	"os"
)

type Student struct {
	id      int
	name    string
	chinese float64
	math    float64
	english float64
}

func main() {
	var local_db []Student = []Student{}
	home_page(&local_db)
}

func home_page(local_db *[]Student) {

	for true {
		print_menu()
		switch input_int() {
		case 1:
			show_transcript(*local_db)
			fmt.Print("按ENTER返回主菜单...")
			fmt.Scanln()
		case 2:
			input_transcript(local_db)
		case 3:
			modify_student(local_db)
		case 4:
			result := remove_student(local_db)
			local_db = &result
		case 5:
			filter_student(local_db)
		case 6:
			os.Exit(0)
		}
	}
}

func print_menu() {
	clear_screen()
	fmt.Println("1. 查看学生成绩单")
	fmt.Println("2. 输入学生成绩")
	fmt.Println("3. 修改学生成绩")
	fmt.Println("4. 删除学生成绩")
	fmt.Println("5. 查询学生成绩")
	fmt.Println("6. 退出系统")
	fmt.Print("请输入你的选择: ")
}

func filter_student(local_db *[]Student) {
	student_id := input_int()
	for _, student := range *local_db {
		if student.id == student_id {
			fmt.Println(student)
		}
	}
}

func remove_student(local_db *[]Student) []Student {
	fmt.Print("请输入学号:")
	student_id := input_int()
	var deleted_index int

	for i, _ := range *local_db {
		if (*local_db)[i].id == student_id {
			deleted_index = i
			break
		}
	}

	fmt.Print("是否继续删除(Y/N):")
	if strings.ToUpper(input_string()) == "Y" {
		result := remove_student_by_index(deleted_index, *local_db)
		return remove_student(&result)
	}

	return remove_student_by_index(deleted_index, *local_db)
}

func modify_student(local_db *[]Student) {
	fmt.Print("请输入学号:")
	student_id := input_int()
	for i, _ := range *local_db {
		if student_id == (*local_db)[i].id {
			fmt.Print("请重新输入学生成绩(语文 数学 英语):")
			(*local_db)[i].chinese = input_float()
			(*local_db)[i].math = input_float()
			(*local_db)[i].english = input_float()

			fmt.Print("是否继续(Y/N):")
			if strings.ToUpper(input_string()) == "Y" {
				modify_student(local_db)
			}
		}

	}
	fmt.Println("没找到学号为%d的学生信息", student_id)
}

func clear_screen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func show_transcript(local_db []Student) {
	fmt.Println("学号 姓名 语文 数学 英语")
	for _, student := range local_db {
		fmt.Println(student)
	}
}

func input_transcript(local_db *[]Student) []Student {
	for true {
		student := Student{input_int(), input_string(), input_float(), input_float(), input_float()}
		*local_db = append(*local_db, student)
		fmt.Print("是否继续输入(输入Y继续):")
		if strings.ToUpper(input_string()) != "Y" {
			break
		}
	}
	return *local_db
}

func remove_student_by_index(index int, students []Student) []Student {
	return append(students[:index], students[index+1:]...)
}

func input_int() int {
	var inputValue int
	fmt.Scan(&inputValue)
	return inputValue
}

func input_string() string {
	var inputValue string
	fmt.Scan(&inputValue)
	return inputValue
}

func input_float() float64 {
	var inputValue float64
	fmt.Scan(&inputValue)
	return inputValue
}
