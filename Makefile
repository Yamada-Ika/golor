test:
	cat testdata/test.txt | go run . -target hoge
	cat testdata/test.txt | go run . -color blue -target hoge
	cat testdata/test.txt | go run . -t hoge
	cat testdata/test.txt | go run . -c blue -t hoge
	cat testdata/multibyte.txt | go run . -c blue -t 世界
