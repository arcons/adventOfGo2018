package main

import (
	"fmt"
	"strings"
	// "bufio"
	// "os"
)

func day2A() {
	//Test comment
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Input ints: ")
	// text, _ := reader.ReadString('\n')
	// perform the parsing
	// so I can just pass in the values with +? wtf
	// input := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	input := []string{"rvefnvyxzbodgpnpkumawhijsc", "rvefqtyxzsddglnppumawhijsc", "rvefqtywzbodglnkkubawhijsc", "rvefqpyxzbozglnpkumawhiqsc", "rvefqtyxzbotgenpkuyawhijsc", "rvefqtyxzbodglnlkumtphijsc", "rwefqtykzbodglnpkumawhijss", "rvynqtyxzbodglnpkumawrijsc", "rvefqtyxlbodgcnpkumawhijec", "rvefqtyxzbodmlnpnumawhijsx", "rvefqtyxzbqdbdnpkumawhijsc", "rvefqtyxzlodblnpkuiawhijsc", "rvefqtyizrodelnpkumawhijsc", "rveffjyxzgodglnpkumawhijsc", "rvefqjyxzbodalnpkumadhijsc", "rvefqtidzbodglnpkumawhdjsc", "hvefqtygzbodglnpkumawhijfc", "rzefqtyxzbodglfhkumawhijsc", "rmefqtyxzbolglnpkumaehijsc", "rnefqqyxzbodglnhkumawhijsc", "rvwfqvyxzbodglnpcumawhijsc", "rvefqtyxzbokgltpkumavhijsc", "rvefciyxzbodglnmkumawhijsc", "rvefptyxzbodglnpkuhashijsc", "rvefqtyxzrodglnpkxmawhiqsc", "rvefqtyxzbotglnpkumawriwsc", "rvufqtyxzbodglnplumawhijvc", "rvefutykzbodglnpkumaahijsc", "rvefqtyxqbodgllprumawhijsc", "rvegqttxzbodgllpkumawhijsc", "dvefqtyxzsodglnpkumawdijsc", "rvefqtyxkbodglnfkumawhijsj", "rvefqtyxzbodnlnpcumawhijnc", "rvefqtyxzbodglfpkuocwhijsc", "rvecqtyxzbbdganpkumawhijsc", "rvefytyxzbodglnpkubgwhijsc", "rvefxtyazbomglnpkumawhijsc", "rvefqgyxzbodglnpkumawyiksc", "avefqtyxzbodglnfkummwhijsc", "fvefqtyxzbbdglnpkumswhijsc", "rvefqtyxzxodglnpkumuuhijsc", "rvezqtyxzbydclnpkumawhijsc", "rvefqtyxzbohglnpkumawdijjc", "rvejqtyxzbodrlnpkumawhijsd", "rvefitzxzbxdglnpkumawhijsc", "rvefutyxzbvdglnikumawhijsc", "rvefqtyazbodgqnbkumawhijsc", "rvefqtyxzbolglnpkwmajhijsc", "rvefqtyxzjodglnpgwmawhijsc", "rvefhtyxzbodglbpaumawhijsc", "mvexqtyxzbodglnpkumawrijsc", "rvefqtyxwbodglnpkumawhbxsc", "rvefqtyxzbodgsnpkudawsijsc", "rvwfqtyxzbonglnwkumawhijsc", "rvefqtyxzjodglnpkfmawhwjsc", "rvefqtyxzbodglntkumughijsc", "rvefctyxzbodglnpkumawhiwsx", "avefqtyvzbodglnpkumawhijsb", "rfefqtyxzlodglnphumawhijsc", "rvefqtyxzfowglnpkumaehijsc", "rvhfvtyxzbodgqnpkumawhijsc", "rfefqtyxzbodglapkumuwhijsc", "rvefqclxzbodglnzkumawhijsc", "qvefqtyxzbodglnckumcwhijsc", "rvefqtyxzkodglnpkymawgijsc", "rvefqtyxzbodgfnpkumafhizsc", "rvefqtyxzbodglnxkumavhijsf", "rvevqtyxzbodgpnpkurawhijsc", "rvefqtyxziodglnpkubawhijss", "rrefqtpxzyodglnpkumawhijsc", "rvefqfyxzbodglcpkxmawhijsc", "rvefdtyxzbodglnpkumvwhijsn", "rverqtyxzbodglnpkwmawhijuc", "rvecjtyxzboxglnpkumawhijsc", "rvefqtyxzbodglnpkqmaxhifsc", "rtnfqtyxzbodglnpkumawhijmc", "lvefqtyxzbodelnpkumawhijsz", "dvefqtyxzbbdgvnpkumawhijsc", "rvefqlyhzbodglnpkumtwhijsc", "roefqtyxlbodglnpkumawhyjsc", "rvefqsydzjodglnpkumawhijsc", "rveybtyxzbodglnpkumawhijsn", "rvefqtyhzbodgvnpmumawhijsc", "rvefqxyazboddlnpkumawhijsc", "vvefqtyxzbohglqpkumawhijsc", "reefhtyxzbodglnpkkmawhijsc", "rvefqtyxzbodglnpkulowhijrc", "rveqqtyxzbodgknpkumawhijsk", "jvefqtqxzbodglnpkumawiijsc", "rvefqtyxzboxglnpvuqawhijsc", "rvefquyxzbodglwwkumawhijsc", "rvefqtyxzbodnlnpkumawhgjbc", "rvdfqthxdbodglnpkumawhijsc", "rvefqtyxzbodllnpkumawhujsb", "evefqtyxzboyglnpkumowhijsc", "rvefktyxzbomglnpzumawhijsc", "rvefqtyxzbodhlnnkrmawhijsc", "rvefqtyxrbodglnpkujaehijsc", "rvefqtyzzbodglnpkumrwhijsb", "evefqtyxzpodglfpkumawhijsc", "rvefqtyxibodglkpyumawhijsc", "rrefqtyxzbodglnpkudawhajsc", "rvifqtyxzbodglxpkumawhijlc", "rxefqtyxzbedglnpkumawhijsp", "rvnfqtyxzbopglnpkuqawhijsc", "rvefqtyxkbodglnpoumawoijsc", "dvefwtyxzbodglnpksmawhijsc", "rvkfqtyxzbodglnpkdmawhijsa", "rcefytyxzzodglnpkumawhijsc", "rvefqtkxzbodglnpktqawhijsc", "nvezqhyxzbodglnpkumawhijsc", "rrefqtyxzbodgunpkumpwhijsc", "rvefqtaxzbodgknpkumawhijic", "pvefqtyxzbodglnpkuxawsijsc", "rvefqtyxzbodglkpvumawhjjsc", "wvefqtyxzkodglnpkumawhhjsc", "rzefqtyxzbotglnpkumawhxjsc", "rvefqtxpzbodglnpkumawzijsc", "bgefqtyxzbodglnpkrmawhijsc", "rvefqlyxzbodglnpkumilhijsc", "cbefqtyxzbodglnpkumawhiesc", "rvefqtyxzbydelnpkumahhijsc", "rvefntyxzbodglnpkumaehijsw", "rverqtyxztodglopkumawhijsc", "rvefqtyxzdodgwrpkumawhijsc", "rvefqtyxibodglnikumawhtjsc", "qvafqtyxzbodglnpkurawhijsc", "rvefqtyxwbodglnpaumawoijsc", "rvefqtyxzoodglndknmawhijsc", "rvdfqtlxzyodglnpkumawhijsc", "rvefqtyxzbodglngfumawhinsc", "rsefqtyxzbodglnpkumawhijek", "rvoestyxzbodglnpkumawhijsc", "svefqtyxzboaglnprumawhijsc", "rvefqtybzbodgwnpkumawwijsc", "rvefqtyxzdwdglnpkvmawhijsc", "rvlfqtyxzbodglnpkrmawhixsc", "rvefqtyxwbodglepkumawhijsd", "rvefqtbxzbodglnqkumawhijmc", "rvefqtzxzbodglnpkumuzhijsc", "rvefqtyxzbodglnpkumawzwnsc", "rvwfqtyxzboiglnpkumawhijsg", "rtehotyxzbodglnpkudawhijsc", "rvegqtyxzbodglnpyumawhijsl", "rvecqtyxzbsdglnpkumawhojsc", "rvefqtyxzbodmlnpkumaghijfc", "rvefqtyxzxodglnpkumanvijsc", "rvefqtyxzbodglnbiugawhijsc", "lvefqtlxzbodglnplumawhijsc", "rvefqtyxvbodglnpkumaldijsc", "rmefqtyxzbodgvnpkuuawhijsc", "rvefqtyxzbodglnpkymeuhijsc", "rvefqtyxzuodganpsumawhijsc", "rxefqtyxzbodglnpkumgwhijwc", "rvefgtyxzbodglnpkudawxijsc", "ahefqtyxzbodglnpkumawhejsc", "rfefqtyxzbzdglnpkusawhijsc", "rvefqtyszqodgljpkumawhijsc", "rvefqtylzboiglnpkumrwhijsc", "rvefqtyxzltdglnpkumawhijsu", "rbefqtyxzbodglnpqumawhijsi", "rvefqtyozpodglnpkumawhijsa", "zvefqtyxzpopglnpkumawhijsc", "rvefqtyxzbodglnfkqmawhijsp", "rvefqtyxzbodgliakumawhijsf", "rvefqtymzrodgfnpkumawhijsc", "ivejqtyxzbodglnpkumawhijuc", "rvefqtyxzbodflnpkxwawhijsc", "dvrfqtyxzbodglnpkumashijsc", "rqefqtyxzbwdglnpkumawvijsc", "tvefqtkxzbodgltpkumawhijsc", "rvefdtyxzbodguxpkumawhijsc", "rveqqtyxvbodglnykumawhijsc", "rvefqtypzcovglnpkumawhijsc", "rvefqnyxzbosglnpkumdwhijsc", "rvefstjxzbodslnpkumawhijsc", "rvefqzyxzpodglnpkummwhijsc", "rvefqkyxzbodglnhgumawhijsc", "rvufqvyxzbodklnpkumawhijsc", "rvefotyxzhodglnpkumawhijsk", "rvefqtyxzbokglnpkumawvcjsc", "lvefqtyxzbolglnpkumawoijsc", "rvefqtywzoodglfpkumawhijsc", "rvehqtqxzbodglnpkumawhcjsc", "rqefqtyxzbodolnpkumjwhijsc", "rvefqtyxzbodglrpkunawgijsc", "rvefqtyxzbodglamkumawdijsc", "rvefvtyzzbodllnpkumawhijsc", "rvefqtyxzbldglnpfcmawhijsc", "rvefppyxzbodglnpkucawhijsc", "rvefquyuzbodglnpkumkwhijsc", "rvefqtyxzbodgqxpkumawhivsc", "rtefotyxzbodglnpkudawhijsc", "rvefqtyxzbodgbnmkuzawhijsc", "ivefqtyxzbodgsnpkumzwhijsc", "rvhfqtyxzbodolnpkumawhijsz", "rvefvtyxzbodwlnpkusawhijsc", "riemqtyxzbodglnpkumawhiasc", "rvtfqtyxzbqdglnpkumawuijsc", "raesqtyxzbodglnpkumawhijsj", "rvefqtyxzbodalmpkumawhihsc", "rvefqtlxzbodgznpkkmawhijsc", "rvefqbyxzbodglgpkubawhijsc", "rvefqtyxnbodgxnpkumswhijsc", "rvefqtyxzkodvlnukumawhijsc", "rvefqtyzzbocglnpkumafhijsc", "rvhfqtyxzbodglmpkumgwhijsc", "rvsfrtyxzbodnlnpkumawhijsc", "rvefqtyxzbxdglnpkujcwhijsc", "rvefqtyvzrodglnphumawhijsc", "reetatyxzbodglnpkumawhijsc", "rvefqtyxzbodglnpzumaoqijsc", "ovefqtyyzbodglnvkumawhijsc", "rvefqbyxzbodnlnpkumawhijsi", "xvefqtyxzbodgrnpkumawrijsc", "rvebqtyxzbodglnpkumazhiasc", "rqeretyxzbodglnpkumawhijsc", "rvefqtyxzyodglapkumvwhijsc", "rvesqxyxzbodglnpvumawhijsc", "rvefqtyxeborglnpkufawhijsc", "rvecqtyxzbodflnpkumawnijsc", "rvefdpyxtbodglnpkumawhijsc", "rvefqtyfzbodclnpkymawhijsc", "rvefqtywzbodglnpxumawhiusc", "rvefqtyxzbodglnpkumawzbjwc", "rvewqtyxdbodglnpxumawhijsc", "rvefqtyxzgocglnpkgmawhijsc", "rvufqtyxzbodggnpkuzawhijsc", "rvefqtynzlodgllpkumawhijsc", "rvedqtyxzbodghnpkumawhujsc", "rvefqtyxlbodgnnpkpmawhijsc", "rvefqtyxzboqglnpkzmawhijec", "rvefqtyxzbodglnpkfmwwyijsc", "rwefqtkxzbodzlnpkumawhijsc", "rvefqtyxvbodglnpkufawhyjsc", "rvefqtyxzbodgltpkumawhqmsc", "rvefctyxzbodglfpkumathijsc", "rvefqtyxzbodgfnpkuuawhijfc", "rvefqttxzbodglnpmumawhijwc", "rvefqtyxzbodglnpkqmawhihsj", "rvefqtyxzbsdglcnkumawhijsc", "rvbiqtyxzbodglnpkumawhijlc", "rnefqtylzvodglnpkumawhijsc", "mvefqtyxzbddglnpkumcwhijsc", "rvefwtyxzbodglnpkgmawhijxc", "rvefqtyxljodglnpkumxwhijsc", "rvefqtyxzbodglnpkuprwhijsd", "rcxfqtyxzbldglnpkumawhijsc", "rvetqtyxzbojglnpkumewhijsc", "rvxfqtyxzbtdglnpkbmawhijsc"}

	twoSum := 0
	threeSum := 0
	var splitS = []string{}
	// loop through all of the strings
	// var usedCharsTwo = ""
	// var usedCharsThree = ""
	var numberOfChar = 0
	hasTwo := false
	hasThree := false
	for _, s := range input {
		// split the given string into characters
		splitS = strings.Split(s, "")
		fmt.Print("Performing parsing on string ", s)
		fmt.Print("\n")
		for _, c := range splitS {
			// get the number of times a char occurs
			numberOfChar = strings.Count(s, c)
			if numberOfChar > 1 {
				if numberOfChar == 2 && !hasTwo {
					fmt.Print("char ", c)
					fmt.Print(" has occured two times for the first time\n")
					hasTwo = true
					twoSum++
					// usedCharsTwo += c
				}
				//just add it to the map and check if it contains the variable
				// if numberOfChar == 3 && !strings.Contains(usedCharsThree, c) {
				if numberOfChar == 3 && !hasThree {
					fmt.Print("char ", c)
					fmt.Print(" has occured three times for the first time\n")
					hasThree = true
					threeSum++
					// usedCharsThree += c
				}
			}
			numberOfChar = 0
		}
		// reset on next word
		hasThree = false
		hasTwo = false
	}
	fmt.Println("two, three", twoSum, threeSum)
	sum := twoSum * threeSum
	fmt.Println("sum:", sum)

}
