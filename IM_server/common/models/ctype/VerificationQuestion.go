package ctype

type VerificationQuestion struct {
	Problem1 *string `json:"problem1"` // 问题1
	Problem2 *string `json:"problem2"` // 问题2
	Problem3 *string `json:"problem3"` // 问题3
	Answer1  *string `json:"answer1"`  // 答案1
	Answer2  *string `json:"answer2"`  // 答案2
	Answer3  *string `json:"answer3"`  // 答案3
}
