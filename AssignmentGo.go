package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"

)

type Pair struct {
	values [2] string
}
func MakePair(k, v string) Pair {
 return Pair{values:[2]string{k, v}}
}
func (p Pair) Get(i int) string {
 return p.values[i]
}




func Reverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result
  }
  return 
}

var cgMap = make(map[string] string)

func paramdiff(dx string, cg map[int]Pair, indices int) string {
	flag := 0

	for i := 0; i<indices; i++ {
		var p Pair = cg[i]
		str := p.Get(1)

		lenStr := len(str)

		var ret string = ""

		for j:= 0; j < lenStr; j++ {
			if dx[0] == str[j] {
				flag = 1

				if j == 0 {
					if str[j+1] == '*' {
						j = j + 2

						for j < lenStr {
							ret = ret + string(str[j])
							j = j + 1
						}
					} else if str[j+1] == '^' {
						if str[j+2] > '1' {
							ret = ret + string(str[j+2])
							ret = ret + "*"
							ret = ret + string(str[j]) 
							ret = ret + "^"
							numx := str[j+2] - '0'
							numxInt  := numx - 1
							ret = ret + strconv.Itoa(int(numxInt))
						}
					} else if str[j+1] == '+' {
						ret = "i"
					} else if str[j+1] == '-' {
						ret = "1"
					}
				}
				return ret 

			}
		}
	}

	if flag == 0 {
		return "didn't find"
	}

	return "problem"

}






func differentiator(y int, CG map[int] Pair, idx int)  {
	//fmt.Println()
	//.Println("*********************************************")
//	fmt.Println("len of cgMap : ", len(cgMap))
	//fmt.Println("i", y, " : ", CG[y].Get(0) , " | ", CG[y].Get(1))
	 
	var dy string = "i" + strconv.Itoa(y)
	var p Pair = CG[y]
	term := p.Get(1)
	lenTerm := len(term)
	var newStr string = ""

//	fmt.Println("term : ", term, "len : ", lenTerm)

	for i := 0; i<lenTerm; i++ {



		if term[i] == 'i' {


//	fmt.Println("term[i] : ", string(term[i]) )

			var key string = dy + "/"
			appOfI := i
			var dx string = "i"
			i = i + 1

			for term[i] >= '0' && term[i] <= '9'  {
				dx = dx +  string(term[i])
				i = i + 1

				if i >= lenTerm {
					break
				}
			}
			endOfI := i
			i = i - 1


			if appOfI == 0 && endOfI == lenTerm {
				newStr = newStr + "1"
			}

			if appOfI > 0 {

				if term[appOfI-1] == '*' {
					pos := appOfI - 2

					for pos >= 0 {
						newStr = newStr +  string(term[pos])
						pos = pos - 1
						newStr = Reverse(newStr)
					} 
				} else if term[appOfI-1] == '+' {
						//"do nothing"
					//	fmt.Println(".")

				} else if term[appOfI-1] == '-' {
						newStr = newStr + "-"

				}
			}


			if endOfI != lenTerm && term[endOfI] == '^' {
				newStr = newStr +  "2*" + dx
			} 

			key = key + dx

			if newStr == "" || newStr[0] == '-' && len(newStr) == 1 {
				newStr = newStr + "1"
			}

			

			cgMap[key] = newStr

		//	fmt.Println("in if -> key : ", key, "newStr : ", newStr, "cgMap push 1 == ", cgMap[key], " ", len(cgMap[key]))
			fmt.Println("cgMap[", key , "] = ", cgMap[key])

		} else if term[i] == 'a' || term[i] == 'b' ||term[i] == 'c' || term[i] == 'd' || term[i] == 'w' {
			var dx string = ""
			dx = dx + string(term[i])
			key := dy + "/"

			if term[i+1] == '*' {
				i = i + 2

				for i < lenTerm {
					newStr = newStr + string(term[i])
					i = i + 1
				}
			} else if term[i+1] == '+' || term[i+1] == '-' {
				newStr = newStr + "1"
			}

			key = key + dx
			cgMap[key] = newStr

			fmt.Println("cgMap[", key , "] = ", cgMap[key])

			//fmt.Println("in if -> key : ", key, "newStr : ", newStr, "cgMap push 2 == ", cgMap[key], " ", len(cgMap[key]))
		} 
	}
}



func dp(str string, str2 string, computationalGrpah map[int]Pair, valuemap map[string]float64, indices int) float64 {
	x := 1
	num := 0
	fmt.Println()
	fmt.Println()
	fmt.Print("str : ", str, " ")

	for x < len(str) {
		num *= 10
		num = num + int(str[x] - '0')
		x = x + 1 
	}

	var p Pair = computationalGrpah[num]
	term := p.Get(1)
	l := len(term)
    //base case:

	if str == "i0" {
		return 0
	}


	for i := 0; i<l; i++ {


		if term[i] == str2[0] {

			var ans float64 = 0
			var  dummy string = str + "/" + str2
			diffterm := cgMap[dummy]

			if diffterm[0] == '1' {
				ans = 1
			} else if diffterm[0]  == '-'{
				ans = -1
			} else if diffterm[0] == 'i' && len(diffterm) < 3 {
				ans = valuemap[diffterm]
			} else {
				if diffterm[0] >= '1' && diffterm[0] <= '9' && diffterm[1] == '*' {
					var dum string = ""
					pos2 := 2

					for pos2 < len(diffterm) {
						dum = dum + string(diffterm[pos2])
						pos2 = pos2 + 1
					}

					ans = valuemap[dum]
					ans = ans * float64(int(diffterm[0] - '0'))
					fmt.Println("dum : ", dum, " ans : ", ans , " valuemap : ", valuemap[dum])

				}

			}

			return ans
		}
	}


	for i := 0; i<l; i++ {
		var dx string = ""
		var ans float64 = 0

		if term[i] == 'i' {
			pos := i+1
			num = 0
			dx = dx + "i"

			for term[pos] >= '0' && term[pos] <= '9' {
				dx = dx + string(term[pos])
				pos = pos + 1

				if pos >= l {
					break
				}
			}
			fmt.Print("dx : ", dx, " ")
			str3 := str + "/" + dx
			if cgMap[str3] != "" {
				diffterm := cgMap[str3]
				fmt.Println("str3 ; ", str3, " diffterm ; ", diffterm)


				//diffterm = 1/ -1 / i0/ i7 / 2 * i6
				if diffterm[0] == '1' {
					ans = 1
				} else if diffterm[0] == '-' {
					ans = -1
				} else if diffterm[0] == 'i' && len(diffterm) < 3 {
					ans = valuemap[diffterm]
				} else {
					if diffterm[0] >= '1' && diffterm[0] <= '9' && diffterm[1] == '*' {
						dum := ""
						pos2 := 2

						for pos2 < len(diffterm) {
							dum = dum + string(diffterm[pos2])
							pos2 = pos2 + 1
						}

						ans = valuemap[dum]
						ans = ans * float64(int(diffterm[0] - '0'))

					fmt.Println("dum : ", dum, " ans : ", ans , " valuemap : ", valuemap[dum])
					}
				}
			} 

			fmt.Println("ans : ", ans)

			if ans == 0 {
				continue
			} else {
				var ans2 float64 = dp( dx , str2, computationalGrpah, valuemap, indices)

				if ans2 == 0 {
					continue
				} else {
					return ans * ans2 
				}
			}
		} // if term[i] == 'i' ends here
		
	}
	fmt.Println("$$$dp -> num : ", num, " term ; ", term)

	return 0
}





func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	fmt.Println(text)
	str1 := strings.Fields(text)
	
	for testcase := 0; testcase < len(str1); testcase++ {
		fmt.Println(str1[testcase])
		expr := str1[testcase]
		fmt.Println(expr)

		x := 2.0
		w := 1.0
		a := 2.0
		b := 3.0
		c := -1.0

		y := 4.0

		fmt.Println(x, " ", y)
		fmt.Println(w, " ", b)
		fmt.Println(a, " ", c)

		var mymap map[rune] int

		mymap = make(map[rune]int)

		mymap['^'] = 5
		mymap['/'] = 4
		mymap['*'] = 3
		mymap['+'] = 2
		mymap['-'] = 1
	

		var stk []rune
		var outputQueue []string

		lenOFexpr := len(expr)

		// Shunting Yard Algorithm



		for i := 0; i < lenOFexpr; i++ {


			if expr[i] >= 'a' && expr[i] <= 'z' {
				// fmt.Printf("%+q ", expr[i])
				// r := rune(expr[i])
				// fmt.Println(mymap[r])

				
				outputQueue = append(outputQueue, string(byte(expr[i])))


			}

			if expr[i] >= '0' && expr[i] <= '9'	{


				if expr[i+1] >= '0' && expr[i+1] <= '9' {
					num := expr[i] - '0'
					num *= 10
					num += expr[i] - '0'
					//fmt.Println(num)
					
					outputQueue = append(outputQueue, string([]byte {expr[i], expr[i+1]}))

				} else {

					oq := string(byte(expr[i]))
					outputQueue = append(outputQueue, oq)					
				}
				
			}

			if expr[i] == '+' || expr[i] == '-' || expr[i] == '*' || expr[i] == '/' || expr[i] == '^' {

			    
				//fmt.Println( " ******* ", mymap[rune(expr[i])])
			//	fmt.Println("i : ", i , " expr[i] : ", string(expr[i])  )

			
				if len(stk) >= 1 {

					topEl := len(stk) - 1
					stkTop := stk[topEl]

					for topEl >= 0 && mymap[rune(expr[i])] < mymap[stkTop] {
						//fmt.Println("mark 1")
						outputQueue = append(outputQueue, string(stkTop) )
						stk = stk[:topEl] // popping stack
						topEl = len(stk) - 1  // getting new size of stack 
						if topEl >= 0 {
						 stkTop = stk[topEl] // getting the next top element
						}

					}
				}
				stk = append(stk, rune(expr[i]))

			}





			if expr[i] == '(' {
				stk = append(stk, rune(expr[i]))
			}
			if expr[i] == ')' {
				topEl := len(stk) - 1

				for stk[topEl] != '(' {
					stkEle := stk[topEl]
					outputQueue = append(outputQueue, string(stkEle) )
					stk = stk[:topEl]
					topEl = len(stk) - 1
				}

				
			}

			//fristElementOfQueue := len(outputQueue) - 1
			//fmt.Println("i : ", i, " ", outputQueue[fristElementOfQueue])

		}

		for len(stk) >= 1 {
			lenOfStk := len(stk) - 1
			outputQueue = append(outputQueue, string(stk[lenOfStk]))

			stk = stk[:lenOfStk] //pop

		}

		fmt.Println("Length of outputQueue : ", len(outputQueue))


		testQueue := outputQueue

		testLenforque := len(testQueue) - 1
fmt.Println("outputQueue : ")
		for testLenforque >= 0 {
			fmt.Print(testQueue[testLenforque], " ")
			testQueue = testQueue[:testLenforque]
			testLenforque = len(testQueue) - 1
		}
		fmt.Println()

















		//computational graph

		var secondStack []string
        var indices int = 0
		computationalGrpah := make(map[int] Pair)
		checkMap := make(map[string] int) 
		computationalGrpah[indices] = MakePair("x", "i0")
		indices += 1

		for len(outputQueue) - 1 >= 0 {
			str := outputQueue[0] 
			outputQueue = outputQueue[1:] // queue popped

			st := rune(str[0])
			//fmt.Println("str : ", str)

			if mymap[st] > 0 { // str[0] == +/*/^

				//fmt.Println("if statement 1")
				var fVar, sVar, first string

				if len(secondStack) - 2 >= 0 {

				//fmt.Println("if statement 1 then 1")
					first = secondStack[len(secondStack) - 1] //getting stack top
					secondStack = secondStack[:len(secondStack) - 1] // pop
					fVar = first
					first = secondStack[len(secondStack) - 1] //getting stack second top 
					sVar = first
					secondStack = secondStack[:len(secondStack) - 1] // pop
				}

				//fmt.Println(" 1 fvar : ",fVar, " sVar : ", sVar," first :", first)
				var flag bool = false
				//var cont int = 0

				for idx := 0; idx < indices; idx++ {

				//fmt.Println("loop statement 1")

					p := computationalGrpah[idx]

					if p.Get(0) == fVar {


				//fmt.Println("if statement 2 CGfirst", p.Get(0), " = fvar ", fVar)
						var dummy string = ""
						if len(fVar) == 1 && fVar[0] >= '0' && fVar[0] <= '9' {
							Swap := fVar
							fVar = sVar
							sVar = Swap

				//fmt.Println("if statement 3")
						}
						first = ""
						first = first + sVar
						first = first + string(st)
						first = first + fVar
						flag = true 
					//	fmt.Println(flag ,"|", first)



						if checkMap[sVar] != 0 {
					//		fmt.Println( "if statement 3 then 1")
							dummy = "i"
							dummy = dummy + strconv.Itoa(checkMap[sVar])

			//	fmt.Println("if statement 4")
						} else {
							//fmt.Println( "else statement 0 ")

							dummy = dummy + sVar
						}
							dummy = dummy + string(st)
							dummy = dummy + "i"
							dummy = dummy + strconv.Itoa(idx)
							secondStack = append(secondStack, first)
							checkMap[first] = indices

							//fmt.Println("graph push 1 : first -> ", first, " dummy -> : ", dummy)
							computationalGrpah[indices] = MakePair(first, dummy)
							checkMap[first] = indices
							
							indices += 1

							break
						 
					} else if p.Get(0) == sVar {
//fmt.Println("else statement 1 ", p.Get(0), " = ", sVar)
				

						var dummy string = ""
						first = ""
						flag2 := false

						if len(fVar) == 1 && fVar[0] >= '0' && fVar[0] <= '9' {
							Swap := fVar
						    fVar = sVar
						    sVar = Swap
							flag2 = true
						} 
						first = first + fVar
						first = first + string(st)
						first = first + sVar

						if flag2 == true {
						    Swap := fVar
						    fVar = sVar
						    sVar = Swap
							flag = true
							dummy = dummy + "i"
							dummy = dummy + strconv.Itoa(idx)
							dummy = dummy + string(st)

				//fmt.Println("if statement 5")

							if checkMap[fVar] != 0 {
							//	fmt.Println(" if statement 5 then 1")
								dummy = dummy + "i"
								dummy = dummy + strconv.Itoa(checkMap[fVar])
							} else {

							//	fmt.Println(" ef statement 1 then 1")
								dummy = dummy + fVar
							} 
						} else {

			//	fmt.Println("else statement 2")
							flag = true
							if checkMap[fVar] != 0 {

							//	fmt.Println(" if statement 5 then 2")
								dummy = dummy + "i";
								dummy = dummy + strconv.Itoa(checkMap[fVar])
							} else {

						///		fmt.Println(" else statement 2 then 1")
								dummy = dummy + fVar
							}
							dummy = dummy + string(st)
							dummy = dummy + "i"
							dummy = dummy + strconv.Itoa(idx)


						}


						checkMap[first] = indices
						secondStack = append(secondStack, first)
					//	fmt.Println("graph push 2 : first -> ", first, " dummy -> : ", dummy)

						computationalGrpah[indices] = MakePair(first, dummy)
						checkMap[first] = indices
						indices += 1

						break
					}
				}

				if flag == false {


				//fmt.Println("if statement 6")
				//	fmt.Println("graph push 3 : first -> ", first, " second -> : /", )
					checkMap[first] = indices

					computationalGrpah[indices] = MakePair(first, "/")
					indices += 1
				}
			} else { // str[0] != +/*/^

			//	fmt.Println("else statement 3")

			//	fmt.Println("else : str ->", str)

				if str[0] >= '0' && str[0] <= '9' && len(str) == 1 {

			//	fmt.Println("if statement 7")
					test := outputQueue[0]
					tt := rune(test[0])

					if mymap[tt] > 0 || len(secondStack) == 0 {
						if test[0] == '+' ||test[0] == '-' ||test[0] == '*' ||test[0] == '/' {
							var dummy string = ""
							dummy = dummy + secondStack[len(secondStack) - 1]
							secondStack = secondStack[:(len(secondStack) - 1)]
							dummy = dummy + str
							secondStack = append(secondStack, dummy)
						} else {
							secondStack = append(secondStack, str)
						}
					} else {
						var dummy string = ""
						dummy = dummy + secondStack[len(secondStack) - 1]
						secondStack = secondStack[:(len(secondStack) - 1)]
						dummy = dummy + str
						secondStack = append(secondStack, dummy)
					}
				} else {

			//	fmt.Println("else statement 4 secondStack push")
					secondStack = append(secondStack, str)
				}
			}

		//	fmt.Println("iteration complete*****************************************")
			//fmt.Println()
			//fmt.Println()
		} 




/*

	/// if method == "ReLU" then it will use ReLU as activation
	CHANGE HERE FOR THE PHASE 2 IMPLEMENTATION 


*/
		method := "MSE"


		if method == "MSE" {


			fmt.Println("Method : ", method)

			fmt.Println("Computational Grpah : ")
			var str1 string = "i" + strconv.Itoa(indices-1)
			computationalGrpah[indices] = MakePair("y^", str1)
			indices = indices + 1
			str1 = "y-i" + strconv.Itoa(indices-1)
			computationalGrpah[indices] = MakePair("y-y^", str1)
			indices = indices + 1
			str1 = "i" + strconv.Itoa(indices-1) + "^2"
			computationalGrpah[indices] = MakePair("(y-y^)^2", str1)
			indices = indices + 1

			for i := 0; i<indices; i++ {
				fmt.Println("i", i, " : ", computationalGrpah[i].Get(0) , " | ", computationalGrpah[i].Get(1) )
			}

			fmt.Println("DIfferciation : ")


			for i := indices-1; i >= 0; i-- {
				differentiator(i, computationalGrpah, indices)
			}
			

/*
			for i := 1; i<indices; i++ {
				str2 := "i" + strconv.Itoa(i) + "/" + "i" + strconv.Itoa(i-1)

				fmt.Println(str2, " cgmap : ",cgMap[str2] , " len: " , len(cgMap) )

			}*/

			valuemap := make(map[string]float64)
			valuemap["x"] = float64(x)
			valuemap["i0"] = valuemap["x"]

			//fmt.Println(valuemap["x"]," ", valuemap["i0"])

			//var cnt int = 0

			fmt.Println()



			for ij  := 0; ij<indices; ij++ {
				var pairStr Pair
				pairStr = computationalGrpah[ij]

				var pairStrStr string = pairStr.Get(1)

				len2 := len(pairStrStr) 
				//fmt.Println("pairStrStr ; ", pairStrStr)

				var num1, num2 float64 = 0, 0
				var opp rune = '/'
				var iflag int = 0
				var temp float64 = 0

				for i := 0; i<len2; i++ {

					if(pairStrStr[i] == 'i') {
						if iflag == 1 {
							iflag = 2
						} else {
							iflag = 1
						}

						temp = num1

						var ss string = "i"

						i = i + 1
                   
						for pairStrStr[i] >= '0' && pairStrStr[i] <= '9'{
							ss = ss + string(pairStrStr[i])
							i = i + 1
							if i >= len2 {
								break
							}
						}

						if valuemap[ss] != 0.0 {
							num1 = valuemap[ss]
						}

						i = i - 1
					
					}

					if pairStrStr[i] == 'w' {
						num2 = w
					} else if pairStrStr[i] == 'a' {
						num2 = a
					} else if pairStrStr[i] == 'b' {
						num2 = b 
					} else if pairStrStr[i] == 'c' {
						num2 = c
					} else if pairStrStr[i] == 'y' {
						num2 = y
					} 

					if pairStrStr[i] == '+' || pairStrStr[i] == '*' || pairStrStr[i] == '^' || pairStrStr[i] == '-'  {
						if pairStrStr[i] == '^' {
							opp = rune(pairStrStr[i])
							i = i + 1
							
						} else {
							opp = rune(pairStrStr[i])
						}
					}
				}

					var ans float64 = 0

					if iflag == 2 {
						num2 = temp
					}

					if opp == '^' {
						ans = num1 * num1
					} else {
						if opp == '+' {
							ans = num1 + num2

						} else if opp == '*' {
							ans = num1 * num2

						} else if opp == '-' {
							ans = num2 - num1 /// y(true) - y(predict)

						} else if opp == '/' {
							if num1 != 0 {
								ans = num1
							}

						}
					}

					var inp string = "i" + strconv.Itoa(ij)

					valuemap[inp] = ans

				}



				fmt.Println()
				fmt.Println("Differentiated values : ")

				for ij := 0; ij<indices; ij++ {
					var inp string = "i" + strconv.Itoa(ij)
					fmt.Println("valuemap[",inp,"] : ", valuemap[inp])
				}

				fmt.Println()
				fmt.Println("DP starts : ")



				str4 := "i" + strconv.Itoa(indices-1)
			
				if(paramdiff("w", computationalGrpah, indices) != "didn't find") {
					deltaParam_w := dp(str4, "w", computationalGrpah, valuemap, indices)
					fmt.Println()
					fmt.Println("*****************Delta of w = ", deltaParam_w )
				}

				if(paramdiff("a", computationalGrpah, indices) != "didn't find") {
				
					deltaParam_a := dp(str4, "a", computationalGrpah, valuemap, indices)
					fmt.Println()
					fmt.Println("****************Delta of a = ", deltaParam_a )
				}

				if(paramdiff("b", computationalGrpah, indices) != "didn't find") {
					deltaParam_b := dp(str4, "b", computationalGrpah, valuemap, indices)
					fmt.Println()
					fmt.Println("***************Delta of b = ", deltaParam_b )
				}

				if(paramdiff("c", computationalGrpah, indices) != "didn't find") {
					deltaParam_c := dp(str4, "c", computationalGrpah, valuemap, indices)
					fmt.Println()
					fmt.Println("****************Delta of c = ", deltaParam_c )
				}

				cgMap = make(map[string]string)

		} else if method == "relu" {





			fmt.Println("Method : ", method)

			fmt.Println("Computational Grpah : ")
			var str1 string = "i" + strconv.Itoa(indices-1)
			computationalGrpah[indices] = MakePair("ReLU(SUM.y^)", str1)
			indices = indices + 1
			str1 = "y-i" + strconv.Itoa(indices-1)
			computationalGrpah[indices] = MakePair("y-y^", str1)
			indices = indices + 1
			str1 = "i" + strconv.Itoa(indices-1) + "^2"
			computationalGrpah[indices] = MakePair("(y-y^)^2", str1)
			indices = indices + 1


			for i := 0; i<indices; i++ {
				fmt.Println("i", i, " : ", computationalGrpah[i].Get(0) , " | ", computationalGrpah[i].Get(1) )
			}

			fmt.Println("DIfferciation : ")


			for i := indices-1; i >= 0; i-- {
				differentiator(i, computationalGrpah, indices)
			}



			valuemap := make(map[string]float64)
			valuemap["x"] = float64(x)
			valuemap["i0"] = valuemap["x"]




			fmt.Println()

			xArray := [2, -1, 0, 1, 3, 4, 5]
			xLen := len(xArray)



			for kk := 0; kk < xLen; kk++ {


			for ij  := 0; ij<indices; ij++ {
				var pairStr Pair
				pairStr = computationalGrpah[ij]

				var pairStrStr string = pairStr.Get(1)

				len2 := len(pairStrStr) 
				//fmt.Println("pairStrStr ; ", pairStrStr)

				var num1, num2 float64 = 0, 0
				var opp rune = '/'
				var iflag int = 0
				var temp float64 = 0

				for i := 0; i<len2; i++ {

					if(pairStrStr[i] == 'i') {
						if iflag == 1 {
							iflag = 2
						} else {
							iflag = 1
						}

						temp = num1

						var ss string = "i"

						i = i + 1
                   
						for pairStrStr[i] >= '0' && pairStrStr[i] <= '9'{
							ss = ss + string(pairStrStr[i])
							i = i + 1
							if i >= len2 {
								break
							}
						}

						if valuemap[ss] != 0.0 {
							num1 = valuemap[ss]
						}

						i = i - 1
					
					}

					if pairStrStr[i] == 'w' {
						num2 = w
					} else if pairStrStr[i] == 'a' {
						num2 = a
					} else if pairStrStr[i] == 'b' {
						num2 = b 
					} else if pairStrStr[i] == 'c' {
						num2 = c
					} else if pairStrStr[i] == 'y' {
						num2 = y
					} 

					if pairStrStr[i] == '+' || pairStrStr[i] == '*' || pairStrStr[i] == '^' || pairStrStr[i] == '-'  {
						if pairStrStr[i] == '^' {
							opp = rune(pairStrStr[i])
							i = i + 1
							
						} else {
							opp = rune(pairStrStr[i])
						}
					}
				}

					var ans float64 = 0

					if iflag == 2 {
						num2 = temp
					}

					if opp == '^' {
						ans = num1 * num1
					} else {
						if opp == '+' {
							ans = num1 + num2

						} else if opp == '*' {
							ans = num1 * num2

						} else if opp == '-' {
							ans = num2 - num1 /// y(true) - y(predict)

						} else if opp == '/' {
							if num1 != 0 {
								ans = num1
							}

						}
					}

					var inp string = "i" + strconv.Itoa(ij)

					valuemap[inp] = ans

				}

			}



				fmt.Println()
				fmt.Println("Differentiated values : ")

				for ij := 0; ij<indices; ij++ {
					var inp string = "i" + strconv.Itoa(ij)
					fmt.Println("valuemap[",inp,"] : ", valuemap[inp])
				}

				fmt.Println()
				fmt.Println("DP starts : ")



			



			

		}


				fmt.Println()
				fmt.Println()

	}

}