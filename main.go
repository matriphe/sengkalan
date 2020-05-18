package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RadhiFadlillah/go-hijri"
)

const versi = "0.1"

var angkaWatak = map[int][][]string{
	0: {
		{"byoma", "musna", "nis", "mlethik", "langit"},
		{"sirna", "ilang", "kombul", "awang-awang"},
		{"mesat", "muluk", "gegana", "nglès"},
		{"tumenga", "nenga", "luhur"},
		{"suwung", "sonya", "muksa", "doh", "tebih"},
		{"swarga", "tanpa", "barakan"},
		{"tan", "rusak", "brastha", "swuh"},
		{"walang", "kos", "pejah", "akasa"},
		{"tawang", "wiyat", "oncat", "windu", "widik-widik"},
		{"nir", "wuk", "sat", "surud", "sempal"},
	},
	1: {
		{"tunggal", "gusti", "sujanma", "semedi"},
		{"badan", "nabi", "rupa", "maha", "buddha"},
		{"niyata", "luwih", "pamasé"},
		{"wong", "buweng", "rat", "lèk", "iku"},
		{"surya", "candra", "kartika", "bumi"},
		{"wiji", "urip", "ron", "éka"},
		{"prabu", "kenya", "nekung"},
		{"raja", "putra", "sasa", "dhara"},
		{"peksi", "dara", "tyas", "wungkul", "sudira", "budi"},
		{"wani", "hyang", "jagad", "nata"},
	},
	2: {
		{"asta", "kalih", "ro", "nembah", "ngabekti"},
		{"nétra", "kembar", "myat", "mandeng", "nayana"},
		{"swiwi", "lar", "sikara", "gandhèng"},
		{"paksa", "apasang", "sungu"},
		{"athi-athi", "talingan", "dresthi"},
		{"carana", "tangan", "karna"},
		{"bau", "suku", "caksuh"},
		{"mata", "paningal", "locana"},
		{"ama", "nebah", "karnan", "ngrengga", "pengantèn", "dwi"},
		{"kanthi", "buja", "bujana"},
	},
	3: {
		{"bahni", "tiga", "ujwala", "kaèksi"},
		{"katon", "murub", "dahana", "payudah"},
		{"katingalan", "kaya", "bentèr"},
		{"nala", "uninga", "kawruh"},
		{"lir", "wrin", "wéda", "naut", "nauti"},
		{"teken", "siking", "pawaka"},
		{"kukus", "api", "apyu"},
		{"brama", "rana", "rananggana"},
		{"utawaka", "uta", "ujel", "kobar", "agni"},
		{"wignya", "guna", "tri", "jatha"},
	},
	4: {
		{"catur", "warna", "wahana", "pat", "warih"},
		{"waudadi", "dadya", "keblat", "papat"},
		{"toya", "suci", "udaka", "we"},
		{"woh", "nadi", "jladri", "sindu"},
		{"yoga", "gawé", "tlaga", "hèr", "wening"},
		{"udan", "bun", "tirta", "marta"},
		{"karya", "sumber", "sumur"},
		{"masuh", "marna", "karti", "karta"},
		{"jalaniddhi", "samodra", "udaya", "tasik"},
		{"tawa", "segara", "wédang"},
	},
	5: {
		{"pandhawa", "lima", "wisikan", "gati"},
		{"indri", "indriya", "warastra", "wrayang"},
		{"astra", "lungid", "sara", "saré"},
		{"guling", "raseksa", "diyu"},
		{"buta", "galak", "wil", "yaksa", "yaksi"},
		{"saya", "wisaya", "bana"},
		{"jemparing", "cakra", "hru"},
		{"tata", "hanata", "bayu", "bajra"},
		{"samirana", "pawaka", "maruta", "angin"},
		{"panca", "marga", "margana"},
	},
	6: {
		{"rasa", "nenem", "rinaras", "artati"},
		{"lona", "tikta", "madura", "sarkara"},
		{"amla", "kayasa", "karaséng"},
		{"oyag", "obah", "nem", "kayu"},
		{"wreksa", "glinggang", "prabatang", "oyig"},
		{"sad", "anggas", "anggang-anggang"},
		{"mangsa", "naya", "retu"},
		{"wayang", "winayang", "anggana"},
		{"ilat", "kilat", "lidhah", "lindhu", "carem", "manis"},
		{"tahen", "osik", "karengya"},
	},
	7: {
		{"sapta", "prawata", "acala", "giri"},
		{"ardi", "gora", "prabata", "himawan"},
		{"pandhita", "pitu", "kaswarèng"},
		{"resi", "sogata", "wiku"},
		{"yogi", "swara", "dwija", "suyati"},
		{"wulang", "weling", "wasita"},
		{"tunggang", "turangga", "gung"},
		{"swa", "aswa", "titihan", "kuda"},
		{"ajar", "arga", "sabda", "nabda", "angsa", "muni"},
		{"suka", "biksu", "biskuka"},
	},
	8: {
		{"astha", "basu", "anggusthi", "basuki"},
		{"slira", "murti", "bujangga", "manggala"},
		{"taksaka", "menyawak", "tekèk"},
		{"dwipa", "dwipangga", "bajul"},
		{"gajah", "liman", "dwirada", "èsthi"},
		{"éstha", "matengga", "brahma"},
		{"brahmana", "wewolu"},
		{"baya", "bebaya", "kunjara"},
		{"tanu", "sarpa", "samaja", "madya", "mangèsthi"},
		{"panagan", "ula", "naga"},
	},
	9: {
		{"bolong", "nawa", "dwara", "pintu", "kori"},
		{"bedhah", "lawang", "wiwara", "gapura"},
		{"rong", "song", "wilasita", "anglèng"},
		{"trustha", "trusthi", "trus", "butul"},
		{"déwa", "sanga", "jawata", "manjing"},
		{"arum", "ganda", "kusuma"},
		{"muka", "rudra", "masuk"},
		{"rago", "angrong", "guwa", "menga"},
		{"babahan", "lèng", "ambuka", "gatra", "anggangsir"},
		{"nanda", "wangi", "wadana"},
	},
}

func main() {
	fmt.Println("Sengkalan versi ", versi)
	fmt.Println("")

	args := os.Args[1:]
	if len(args) == 0 {
		tampilkanBantuan("Tidak disertai tahun")
		return
	}

	kamusID, err := muatKamus("./kamus/id.json")
	if err != nil {
		tampilkanKesalahan(fmt.Sprintf("Tidak dapat memuat kamus ID, karena: %s", err.Error()))
		return
	}

	kamusJV, err := muatKamus("./kamus/jv.json")
	if err != nil {
		tampilkanKesalahan(fmt.Sprintf("Tidak dapat memuat kamus JV, karena: %s", err.Error()))
		return
	}

	tahunMasehi := args[0]
	tahun, err := time.Parse("2006", tahunMasehi)
	if err != nil {
		tampilkanBantuan("Tahun tidak valid")
		return
	}

	tahunHijriyah, _, _ := hijri.ToHijri(tahun)
	tahunJawa := strconv.Itoa(tahunHijriyah + 512)

	sengkalanMasehi := buatSengkalan(tahunMasehi)
	artiMasehi := cariArtiSengkalan(sengkalanMasehi, kamusID, kamusJV)

	sengkalanJawa := buatSengkalan(tahunJawa)
	artiJawa := cariArtiSengkalan(sengkalanJawa, kamusID, kamusJV)

	fmt.Println("📅 Tahun Masehi:", tahunMasehi)
	fmt.Println("☀️ Surya Sengkala:", strings.Join(sengkalanMasehi, " "))
	fmt.Println("📜 Makna Surya Sengkala:")
	tampilkanArti(artiMasehi)

	fmt.Println("")

	fmt.Println("📅 Tahun Jawa:", tahunJawa)
	fmt.Println("🌙 Candra Sengkala:", strings.Join(sengkalanJawa, " "))
	fmt.Println("📜 Makna Candra Sengkala:")
	tampilkanArti(artiJawa)

	fmt.Println("")
}

func tampilkanKesalahan(pesan string) {
	msg := fmt.Sprintf("❌  %s", pesan)
	numMsg := len(msg)

	fmt.Println(strings.Repeat("=", numMsg))
	fmt.Println(msg)
	fmt.Println(strings.Repeat("-", numMsg))
}

func tampilkanBantuan(pesan string) {
	tampilkanKesalahan(pesan)
	fmt.Println("ℹ️ Penggunaan: sengkalan [tahun]")
	fmt.Println("🤖 Contoh: sengkalan", time.Now().Year())
}

func tampilkanArti(arti map[string]map[string]string) {
	for k, ar := range arti {
		fmt.Printf("   > %s:\n", k)
		for i, a := range ar {
			if a != "" {
				fmt.Printf("     >> (%s) %s\n", i, a)
			}
		}
	}
}

func buatSengkalan(tahun string) []string {
	var watak = make([]string, len(tahun))
	for i, t := range reverse(tahun) {
		it, _ := strconv.Atoi(fmt.Sprintf("%c", t))

		w := angkaWatak[it]

		rand.Seed(time.Now().UnixNano())
		c := rand.Intn(len(w) - 1)
		ch := w[c]

		rand.Seed(time.Now().UnixNano())
		wch := rand.Intn(len(ch))

		watak[i] = strings.Title(ch[wch])
	}

	return watak
}

func muatKamus(filename string) (map[string]string, error) {
	j, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	kamus := make(map[string]string)
	err = json.Unmarshal(j, &kamus)
	if err != nil {
		return nil, err
	}

	return kamus, nil
}

func cariArti(kata string, kamus map[string]string) string {
	kata = strings.ToLower(kata)

	eWords := []string{"\\u00e9", "\\u00e8", "\\u00ea", "é", "è", "ê"}
	for _, e := range eWords {
		kata = strings.Replace(kata, e, "e", len(kata))
	}

	if arti, ok := kamus[kata]; ok {
		return arti
	}

	return ""
}

func cariArtiSengkalan(sengkalan []string, kamusID, kamusJV map[string]string) map[string]map[string]string {
	arti := make(map[string]map[string]string, len(sengkalan))
	for _, k := range sengkalan {
		arti[k] = make(map[string]string, 0)
		arti[k]["ID"] = cariArti(k, kamusID)
		arti[k]["JV"] = cariArti(k, kamusJV)
	}

	return arti
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
