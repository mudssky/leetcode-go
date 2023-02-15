package util

import "testing"

var sortedList = []int{287,
	337,
	461,
	632,
	658,
	971,
	1001,
	1090,
	1094,
	1104,
	1117,
	1201,
	1258,
	1423,
	1437,
	1460,
	1698,
	2084,
	2096,
	2217,
	2417,
	2465,
	2555,
	2707,
	2859,
	2918,
	2945,
	2967,
	2968,
	2979,
	3197,
	3556,
	3562,
	3610,
	3623,
	3649,
	3846,
	3895,
	3904,
	4065,
	4139,
	4140,
	4165,
	4191,
	4302,
	4359,
	4500,
	4529,
	4555,
	4565,
	4577,
	4657,
	4952,
	5051,
	5140,
	5153,
	5278,
	5468,
	5689,
	5773,
	5855,
	5885,
	5901,
	5943,
	5971,
	5992,
	6226,
	6335,
	6431,
	6565,
	6622,
	6737,
	6792,
	6844,
	6929,
	7014,
	7115,
	7153,
	7319,
	7337,
	7553,
	7817,
	7955,
	7971,
	8126,
	8203,
	8254,
	8287,
	8419,
	8452,
	8473,
	8691,
	9002,
	9006,
	9041,
	9266,
	9276,
	9537,
	9668,
	9917}
var testCase = []struct {
	Input1 []int
	Input2 int
	Expect int
}{
	{sortedList, 9276, RandomInt(len(sortedList))},
}

func BenchmarkLinearSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LinearSearch(testCase[0].Input1, testCase[0].Input2)
	}
}
func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearch(testCase[0].Input1, testCase[0].Input2)
	}
}

func BenchmarkBinarySearchCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binarySearchCompare(testCase[0].Input1, testCase[0].Input2)
	}
}
