package main

import (
	"fmt"
	// "bufio"
	// "os"
)

func main() {
	//Test comment
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Input ints: ")
	// text, _ := reader.ReadString('\n')
	// perform the parsing
	// so I can just pass in the values with +? wtf
	// input := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	input := []string{"rvefnvyxzbodgpnpkumawhijsc", "rvefqtywzbodglnkkubawhijsc", "rvefqpyxzbozglnpkumawhiqsc", "rvefqtyxzbotgenpkuyawhijsc", "rvefqtyxzbodglnlkumtphijsc", "rwefqtykzbodglnpkumawhijss", "rvynqtyxzbodglnpkumawrijsc", "rvefqtyxlbodgcnpkumawhijec", "rvefqtyxzbodmlnpnumawhijsx", "rvefqtyxzbqdbdnpkumawhijsc", "rvefqtyxzlodblnpkuiawhijsc", "rvefqtyizrodelnpkumawhijsc", "rveffjyxzgodglnpkumawhijsc", "rvefqjyxzbodalnpkumadhijsc", "rvefqtidzbodglnpkumawhdjsc", "hvefqtygzbodglnpkumawhijfc", "rzefqtyxzbodglfhkumawhijsc", "rmefqtyxzbolglnpkumaehijsc", "rnefqqyxzbodglnhkumawhijsc", "rvwfqvyxzbodglnpcumawhijsc", "rvefqtyxzbokgltpkumavhijsc", "rvefciyxzbodglnmkumawhijsc", "rvefptyxzbodglnpkuhashijsc", "rvefqtyxzrodglnpkxmawhiqsc", "rvefqtyxzbotglnpkumawriwsc", "rvufqtyxzbodglnplumawhijvc", "rvefutykzbodglnpkumaahijsc", "rvefqtyxqbodgllprumawhijsc", "rvegqttxzbodgllpkumawhijsc", "dvefqtyxzsodglnpkumawdijsc", "rvefqtyxkbodglnfkumawhijsj", "rvefqtyxzbodnlnpcumawhijnc", "rvefqtyxzbodglfpkuocwhijsc", "rvecqtyxzbbdganpkumawhijsc", "rvefytyxzbodglnpkubgwhijsc", "rvefxtyazbomglnpkumawhijsc", "rvefqgyxzbodglnpkumawyiksc", "avefqtyxzbodglnfkummwhijsc", "fvefqtyxzbbdglnpkumswhijsc", "rvefqtyxzxodglnpkumuuhijsc", "rvezqtyxzbydclnpkumawhijsc", "rvefqtyxzbohglnpkumawdijjc", "rvejqtyxzbodrlnpkumawhijsd", "rvefitzxzbxdglnpkumawhijsc", "rvefutyxzbvdglnikumawhijsc", "rvefqtyazbodgqnbkumawhijsc", "rvefqtyxzbolglnpkwmajhijsc", "rvefqtyxzjodglnpgwmawhijsc", "rvefhtyxzbodglbpaumawhijsc", "mvexqtyxzbodglnpkumawrijsc", "rvefqtyxwbodglnpkumawhbxsc", "rvefqtyxzbodgsnpkudawsijsc", "rvwfqtyxzbonglnwkumawhijsc", "rvefqtyxzjodglnpkfmawhwjsc", "rvefqtyxzbodglntkumughijsc", "rvefctyxzbodglnpkumawhiwsx", "avefqtyvzbodglnpkumawhijsb", "rfefqtyxzlodglnphumawhijsc", "rvefqtyxzfowglnpkumaehijsc", "rvhfvtyxzbodgqnpkumawhijsc", "rfefqtyxzbodglapkumuwhijsc", "rvefqclxzbodglnzkumawhijsc", "qvefqtyxzbodglnckumcwhijsc", "rvefqtyxzkodglnpkymawgijsc", "rvefqtyxzbodgfnpkumafhizsc", "rvefqtyxzbodglnxkumavhijsf", "rvevqtyxzbodgpnpkurawhijsc", "rvefqtyxziodglnpkubawhijss", "rrefqtpxzyodglnpkumawhijsc", "rvefqfyxzbodglcpkxmawhijsc", "rvefdtyxzbodglnpkumvwhijsn", "rverqtyxzbodglnpkwmawhijuc", "rvecjtyxzboxglnpkumawhijsc", "rvefqtyxzbodglnpkqmaxhifsc", "rtnfqtyxzbodglnpkumawhijmc", "lvefqtyxzbodelnpkumawhijsz", "dvefqtyxzbbdgvnpkumawhijsc", "rvefqlyhzbodglnpkumtwhijsc", "roefqtyxlbodglnpkumawhyjsc", "rvefqsydzjodglnpkumawhijsc", "rveybtyxzbodglnpkumawhijsn", "rvefqtyhzbodgvnpmumawhijsc", "rvefqxyazboddlnpkumawhijsc", "vvefqtyxzbohglqpkumawhijsc", "reefhtyxzbodglnpkkmawhijsc", "rvefqtyxzbodglnpkulowhijrc", "rveqqtyxzbodgknpkumawhijsk", "jvefqtqxzbodglnpkumawiijsc", "rvefqtyxzboxglnpvuqawhijsc", "rvefquyxzbodglwwkumawhijsc", "rvefqtyxzbodnlnpkumawhgjbc", "rvdfqthxdbodglnpkumawhijsc", "rvefqtyxzbodllnpkumawhujsb", "evefqtyxzboyglnpkumowhijsc", "rvefktyxzbomglnpzumawhijsc", "rvefqtyxzbodhlnnkrmawhijsc", "rvefqtyxrbodglnpkujaehijsc", "rvefqtyzzbodglnpkumrwhijsb", "evefqtyxzpodglfpkumawhijsc", "rvefqtyxibodglkpyumawhijsc", "rrefqtyxzbodglnpkudawhajsc", "rvifqtyxzbodglxpkumawhijlc", "rxefqtyxzbedglnpkumawhijsp", "rvnfqtyxzbopglnpkuqawhijsc", "rvefqtyxkbodglnpoumawoijsc", "dvefwtyxzbodglnpksmawhijsc", "rvkfqtyxzbodglnpkdmawhijsa", "rcefytyxzzodglnpkumawhijsc", "rvefqtkxzbodglnpktqawhijsc", "nvezqhyxzbodglnpkumawhijsc", "rrefqtyxzbodgunpkumpwhijsc", "rvefqtaxzbodgknpkumawhijic", "pvefqtyxzbodglnpkuxawsijsc", "rvefqtyxzbodglkpvumawhjjsc", "wvefqtyxzkodglnpkumawhhjsc", "rzefqtyxzbotglnpkumawhxjsc", "rvefqtxpzbodglnpkumawzijsc", "bgefqtyxzbodglnpkrmawhijsc", "rvefqlyxzbodglnpkumilhijsc", "cbefqtyxzbodglnpkumawhiesc", "rvefqtyxzbydelnpkumahhijsc", "rvefntyxzbodglnpkumaehijsw", "rverqtyxztodglopkumawhijsc", "rvefqtyxzdodgwrpkumawhijsc", "rvefqtyxibodglnikumawhtjsc", "qvafqtyxzbodglnpkurawhijsc", "rvefqtyxwbodglnpaumawoijsc", "rvefqtyxzoodglndknmawhijsc", "rvdfqtlxzyodglnpkumawhijsc", "rvefqtyxzbodglngfumawhinsc", "rsefqtyxzbodglnpkumawhijek", "rvoestyxzbodglnpkumawhijsc", "svefqtyxzboaglnprumawhijsc", "rvefqtybzbodgwnpkumawwijsc", "rvefqtyxzdwdglnpkvmawhijsc", "rvlfqtyxzbodglnpkrmawhixsc", "rvefqtyxwbodglepkumawhijsd", "rvefqtbxzbodglnqkumawhijmc", "rvefqtzxzbodglnpkumuzhijsc", "rvefqtyxzbodglnpkumawzwnsc", "rvwfqtyxzboiglnpkumawhijsg", "rtehotyxzbodglnpkudawhijsc", "rvegqtyxzbodglnpyumawhijsl", "rvecqtyxzbsdglnpkumawhojsc", "rvefqtyxzbodmlnpkumaghijfc", "rvefqtyxzxodglnpkumanvijsc", "rvefqtyxzbodglnbiugawhijsc", "lvefqtlxzbodglnplumawhijsc", "rvefqtyxvbodglnpkumaldijsc", "rmefqtyxzbodgvnpkuuawhijsc", "rvefqtyxzbodglnpkymeuhijsc", "rvefqtyxzuodganpsumawhijsc", "rxefqtyxzbodglnpkumgwhijwc", "rvefgtyxzbodglnpkudawxijsc", "ahefqtyxzbodglnpkumawhejsc", "rfefqtyxzbzdglnpkusawhijsc", "rvefqtyszqodgljpkumawhijsc", "rvefqtylzboiglnpkumrwhijsc", "rvefqtyxzltdglnpkumawhijsu", "rbefqtyxzbodglnpqumawhijsi", "rvefqtyozpodglnpkumawhijsa", "zvefqtyxzpopglnpkumawhijsc", "rvefqtyxzbodglnfkqmawhijsp", "rvefqtyxzbodgliakumawhijsf", "rvefqtymzrodgfnpkumawhijsc", "ivejqtyxzbodglnpkumawhijuc", "rvefqtyxzbodflnpkxwawhijsc", "dvrfqtyxzbodglnpkumashijsc", "rqefqtyxzbwdglnpkumawvijsc", "tvefqtkxzbodgltpkumawhijsc", "rvefdtyxzbodguxpkumawhijsc", "rveqqtyxvbodglnykumawhijsc", "rvefqtypzcovglnpkumawhijsc", "rvefqnyxzbosglnpkumdwhijsc", "rvefstjxzbodslnpkumawhijsc", "rvefqzyxzpodglnpkummwhijsc", "rvefqkyxzbodglnhgumawhijsc", "rvufqvyxzbodklnpkumawhijsc", "rvefotyxzhodglnpkumawhijsk", "rvefqtyxzbokglnpkumawvcjsc", "lvefqtyxzbolglnpkumawoijsc", "rvefqtywzoodglfpkumawhijsc", "rvehqtqxzbodglnpkumawhcjsc", "rqefqtyxzbodolnpkumjwhijsc", "rvefqtyxzbodglrpkunawgijsc", "rvefqtyxzbodglamkumawdijsc", "rvefvtyzzbodllnpkumawhijsc", "rvefqtyxzbldglnpfcmawhijsc", "rvefppyxzbodglnpkucawhijsc", "rvefquyuzbodglnpkumkwhijsc", "rvefqtyxzbodgqxpkumawhivsc", "rtefotyxzbodglnpkudawhijsc", "rvefqtyxzbodgbnmkuzawhijsc", "ivefqtyxzbodgsnpkumzwhijsc", "rvhfqtyxzbodolnpkumawhijsz", "rvefvtyxzbodwlnpkusawhijsc", "riemqtyxzbodglnpkumawhiasc", "rvtfqtyxzbqdglnpkumawuijsc", "raesqtyxzbodglnpkumawhijsj", "rvefqtyxzbodalmpkumawhihsc", "rvefqtlxzbodgznpkkmawhijsc", "rvefqbyxzbodglgpkubawhijsc", "rvefqtyxnbodgxnpkumswhijsc", "rvefqtyxzkodvlnukumawhijsc", "rvefqtyzzbocglnpkumafhijsc", "rvhfqtyxzbodglmpkumgwhijsc", "rvsfrtyxzbodnlnpkumawhijsc", "rvefqtyxzbxdglnpkujcwhijsc", "rvefqtyvzrodglnphumawhijsc", "reetatyxzbodglnpkumawhijsc", "rvefqtyxzbodglnpzumaoqijsc", "ovefqtyyzbodglnvkumawhijsc", "rvefqbyxzbodnlnpkumawhijsi", "xvefqtyxzbodgrnpkumawrijsc", "rvebqtyxzbodglnpkumazhiasc", "rqeretyxzbodglnpkumawhijsc", "rvefqtyxzyodglapkumvwhijsc", "rvesqxyxzbodglnpvumawhijsc", "rvefqtyxeborglnpkufawhijsc", "rvecqtyxzbodflnpkumawnijsc", "rvefdpyxtbodglnpkumawhijsc", "rvefqtyfzbodclnpkymawhijsc", "rvefqtywzbodglnpxumawhiusc", "rvefqtyxzbodglnpkumawzbjwc", "rvewqtyxdbodglnpxumawhijsc", "rvefqtyxzgocglnpkgmawhijsc", "rvufqtyxzbodggnpkuzawhijsc", "rvefqtynzlodgllpkumawhijsc", "rvedqtyxzbodghnpkumawhujsc", "rvefqtyxlbodgnnpkpmawhijsc", "rvefqtyxzboqglnpkzmawhijec", "rvefqtyxzbodglnpkfmwwyijsc", "rwefqtkxzbodzlnpkumawhijsc", "rvefqtyxvbodglnpkufawhyjsc", "rvefqtyxzbodgltpkumawhqmsc", "rvefctyxzbodglfpkumathijsc", "rvefqtyxzbodgfnpkuuawhijfc", "rvefqttxzbodglnpmumawhijwc", "rvefqtyxzbodglnpkqmawhihsj", "rvefqtyxzbsdglcnkumawhijsc", "rvbiqtyxzbodglnpkumawhijlc", "rnefqtylzvodglnpkumawhijsc", "mvefqtyxzbddglnpkumcwhijsc", "rvefwtyxzbodglnpkgmawhijxc", "rvefqtyxljodglnpkumxwhijsc", "rvefqtyxzbodglnpkuprwhijsd", "rcxfqtyxzbldglnpkumawhijsc", "rvetqtyxzbojglnpkumewhijsc", "rvxfqtyxzbtdglnpkbmawhijsc"}
	var o1 = ""
	var o2 = ""
	counter := 0
	for _, s := range input {
		for _, s2 := range input {
			for i := 0; i < 26; i++ {
				if counter > 1 {
					break
				}
				if s[i] != s2[i] {
					counter++
				}
			}
			if counter == 1 {
				o1 = s
				o2 = s2
				break
			}
			counter = 0
		}
		if counter == 1 {
			break
		}
	}

	fmt.Println("Output:", o1, o2)
}
