
type pemain struct {
	nama  string
	poin  int
}

const NMAX int = 1024

type dataPemain [NMAX]pemain

func isiArray(himpunan *dataPemain, n int) {
	for i := 0; i < n; i++ {
		var nama string
		var poin int
		fmt.Scanln(&nama, &poin)
		himpunan[i].nama = nama
		himpunan[i].poin = poin
	}
}

func selectionSort(himpunan *dataPemain, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if himpunan[j].poin < himpunan[minIdx].poin {
				minIdx = j
			}
		}
		himpunan[i], himpunan[minIdx] = himpunan[minIdx], himpunan[i]
	}
}

func showArray(himpunan dataPemain, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(himpunan[i].nama, himpunan[i].poin)
	}
}
