nomer 1

Deskripsi Fungsi:
Fungsi faktorial(n int) int:

Fungsi ini menghitung faktorial dari bilangan n. Faktorial didefinisikan sebagai hasil kali dari semua bilangan bulat positif yang lebih kecil atau sama dengan n. Jika n adalah 0, maka hasilnya adalah 1.
Fungsi permutasi(n, r int) int:

Fungsi ini menghitung nilai permutasi dari n elemen yang diambil r elemen. Permutasi menghitung berapa banyak cara untuk menyusun r elemen dari n elemen. Rumusnya adalah P(n, r) = n! / (n-r)!.
Fungsi kombinasi(n, r int) int:

Fungsi ini menghitung nilai kombinasi dari n elemen yang diambil r elemen. Kombinasi menghitung berapa banyak cara untuk memilih r elemen dari n elemen tanpa memperhatikan urutan. Rumusnya adalah C(n, r) = n! / (r!(n-r)!).
Proses pada main:
Input:
Program meminta pengguna untuk memasukkan empat nilai integer: a, b, c, dan d. Nilai ini diinput dalam satu baris dengan spasi sebagai pemisah.
Perhitungan:
Program menghitung permutasi dan kombinasi untuk dua set angka yang berbeda:
Set pertama menggunakan nilai a dan c, dengan hasil berupa permutasi P_a_c dan kombinasi C_a_c.
Set kedua menggunakan nilai b dan d, dengan hasil berupa permutasi P_b_d dan kombinasi C_b_d.
Output:
Hasil perhitungan permutasi dan kombinasi untuk set pertama (a, c) dan set kedua (b, d) ditampilkan dalam format dua baris:
Baris pertama menampilkan hasil permutasi dan kombinasi untuk set pertama.
Baris kedua menampilkan hasil permutasi dan kombinasi untuk set kedua.

nomer 2

Deskripsi Fungsionalitas:
Konstanta MAX_TIME:

MAX_TIME diset ke nilai 301, yang merepresentasikan batas waktu maksimum yang dianggap sebagai soal yang tidak terselesaikan. Jika waktu pengerjaan soal lebih dari 300 menit, soal dianggap tidak selesai.
Prosedur hitungSkor(waktu []int, soal *int, skor *int):

Prosedur ini menghitung dua hal:
Jumlah soal yang diselesaikan oleh seorang peserta. Jika waktu pengerjaan sebuah soal kurang dari MAX_TIME, soal tersebut dianggap terselesaikan.
Total skor berupa akumulasi dari waktu yang dibutuhkan untuk menyelesaikan soal-soal yang berhasil diselesaikan (di bawah 301 menit).
Nilai jumlah soal yang diselesaikan dan total skor dikembalikan melalui parameter pointer soal dan skor.
Fungsi main:

Program dimulai dengan meminta input dari pengguna, yaitu:
Jumlah peserta dalam kompetisi.
Nama peserta dan waktu pengerjaan untuk setiap soal (8 soal per peserta).
Proses per peserta:
Untuk setiap peserta, program meminta waktu pengerjaan untuk masing-masing dari 8 soal.
Waktu tersebut digunakan oleh prosedur hitungSkor untuk menghitung jumlah soal yang diselesaikan dan total waktu yang diperlukan (skor).
Program kemudian membandingkan hasil dengan peserta pemenang saat ini:
Jika peserta tersebut menyelesaikan lebih banyak soal dibanding pemenang saat ini, maka peserta tersebut menjadi pemenang sementara.
Jika jumlah soal yang diselesaikan sama, maka peserta dengan total waktu yang lebih sedikit menjadi pemenang.
Penentuan Pemenang:

Pemenang akhir ditentukan berdasarkan siapa yang menyelesaikan lebih banyak soal. Jika dua peserta menyelesaikan jumlah soal yang sama, pemenang adalah yang memiliki total waktu (skor) paling rendah.
Output:

Setelah memproses semua peserta, program menampilkan nama pemenang, jumlah soal yang diselesaikan oleh pemenang, dan total waktu yang dibutuhkan pemenang.

nomer 3


Deskripsi Fungsionalitas:
Fungsi cetakDeret(n int):

Fungsi ini digunakan untuk mencetak setiap elemen deret Collatz yang dimulai dari bilangan n yang diberikan pengguna. Aturannya adalah:
Jika bilangan n genap, maka n dibagi dua (n = n / 2).
Jika bilangan n ganjil, maka n dikalikan tiga dan ditambah satu (n = 3*n + 1).
Proses ini diulangi hingga n mencapai nilai 1. Setiap nilai n dicetak di layar selama proses berjalan.
Fungsi main:

Program pertama-tama meminta pengguna untuk memasukkan sebuah bilangan bulat positif n yang kurang dari 1 juta.
Setelah nilai n diinput, program akan memanggil fungsi cetakDeret untuk menghitung dan mencetak deret Collatz berdasarkan aturan yang telah dijelaskan.
Cetak Deret:

Program akan mencetak setiap elemen dari deret Collatz pada baris yang sama dengan spasi sebagai pemisah hingga deret tersebut mencapai angka 1, yang kemudian diakhiri dengan baris baru (\n).
 
 
