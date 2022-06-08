fn main() {
   arr := ["aaabbcdddd", 
   		 "BBBBBBBBBBBBNBBBBBBBBBBBBNNNBBBBBBBBBBBBBBBBBBBBBBBBNBBBBBBBBBBBBBB",
		"abcdefghijklmnopqrstuvwxyz"]

	for str in arr {
		println(str)
		println(run_length_encoding(str))
		println("----------------------------------")
	}

}

fn run_length_encoding(input string ) string {
	mut count := 0
	mut	output := ""

	for i := 0; i < input.len; i++ {

		// first time	
		if i == 0 {
			count = 1
			continue
		}
		
		previous := input[i-1]
		current := input[i]

		// the same character was found
		if current == previous {
			count++
			continue
		}

		// a different char was found
		if count >  1 {
			output += count.str()
		}
		output += previous.ascii_str()
		count = 1
	}

	// Last char
	if count > 1 {
		output += count.str()
	}
	output += input[input.len-1].ascii_str()
	return output
}
