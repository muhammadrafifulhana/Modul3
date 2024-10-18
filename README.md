# Latihan Modul3

## Soal Latihan 3.5 (No. 1)
### Program Menghitung Permutasi & Kombinasi
Program difungksikan untuk menghitung dan menampilkan jumlah permutasi dan kombinasi dari dua pasang bilangan yang diberikan pengguna.

### 1. Fungsi `main`
```go
func main() {
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)
    fmt.Println(permutation(a, c), combination(a, c))
    fmt.Println(permutation(b, d), combination(b, d))
}
```
- Di dalam fungsi `main`, empat variabel integer (`a`, `b`, `c`, `d`) dideklarasikan.
- `fmt.Scan(&a, &b, &c, &d)` membaca empat angka dari input pengguna.
- Kemudian, program mencetak hasil dari fungsi `permutation` dan `combination` untuk dua set angka yang berbeda: `a` dan `c`, serta `b` dan `d`.

---

### 2. Fungsi `factorial`
```go
func factorial(n int) int {
    var factorial int = 1
    for i := n; i >= 1; i-- {
        factorial *= i
    }
    return factorial
}
```
Fungsi `factorial` menghitung faktorial dari angka `n`. 
Dengan menggunakan loop, faktorial dihitung dengan mengalikan `n` dengan semua angka di bawahnya hingga 1.

---

### 3. Fungsi `permutation`
```go
func permutation(n, r int) int {
    var permutate int
    nfactorial := factorial(n)
    nrfactorial := factorial(n - r)
    permutate = nfactorial / nrfactorial
    return permutate
}
```
Fungsi tersebut menghitung permutasi dari `n` mengambil `r` (`P(n, r) = n! / (n - r)!`). 
Ini menggunakan fungsi `factorial` untuk mendapatkan faktorial dari `n` dan `n - r`, kemudian menghitung hasilnya.

---

### 4. Fungsi `combination`
```go
func combination(n, r int) int {
    var combinate int
    nfactorial := factorial(n)
    rnrfactorial := factorial(r) * factorial(n - r)
    combinate = nfactorial / rnrfactorial
    return combinate
}
```
Fungsi tersebut menghitung kombinasi dari `n` mengambil `r` (`C(n, r) = n! / (r! * (n - r)!)`). 
Seperti sebelumnya, fungsi ini menggunakan `factorial` untuk menghitung faktorial yang diperlukan dan mengembalikan hasil kombinasi.

---

## Soal Latihan 3.5 (No. 2)
### Program Menghitung Fungsi Tertentu
Program difungsikan untuk menghitung dan menampilkan hasil dari beberapa fungsi matematika berdasarkan input yang diberikan pengguna.

### 1. Fungsi `main`
```go
func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)
    fmt.Println(fogoh(a))
    fmt.Println(gohof(b))
    fmt.Println(hofog(c))
}
```
- Di dalam fungsi `main`, tiga variabel integer (`a`, `b`, `c`) dideklarasikan.
- `fmt.Scan(&a, &b, &c)` membaca tiga angka dari input pengguna.
- Kemudian, program mencetak hasil dari fungsi `fogoh`, `gohof`, dan `hofog` untuk nilai `a`, `b`, dan `c`.

---

### 2. Fungsi `f`
```go
func f(x int) int {
    return x * x
}
```
Fungsi ini mengembalikan kuadrat dari input `x`.

---

### 3. Fungsi `g`
```go
func g(x int) int {
    return x - 2
}
```
Fungsi ini mengurangi input `x` dengan 2.

---

### 4. Fungsi `h`
```go
func h(x int) int {
    return x + 1
}
```
Fungsi ini menambahkan 1 pada input `x`.

---

### 5. Fungsi `fogoh`
```go
func fogoh(x int) int {
    return f(g(h(x)))
}
```
Fungsi ini menghitung hasil dari komposisi fungsi `f`, `g`, dan `h` secara berurutan.

---

### 6. Fungsi `gohof`
```go
func gohof(x int) int {
    return g(h(f(x)))
}
```
Fungsi ini menghitung hasil dari komposisi fungsi `g`, `h`, dan `f` secara berurutan.

---

### 7. Fungsi `hofog`
```go
func hofog(x int) int {
    return h(f(g(x)))
}
```
Fungsi ini menghitung hasil dari komposisi fungsi `h`, `f`, dan `g` secara berurutan.

---


## Soal Latihan 3.5 (No. 3)
### Program Menghitung Posisi Titik Terhadap Lingkaran
Program tersebut dirancang untuk menentukan apakah suatu titik berada di dalam satu atau kedua lingkaran berdasarkan input yang diberikan pengguna.

### 1. Fungsi `jarak`
```go
func jarak(x1, y1, x2, y2 int) float64 {
    return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}
```
Fungsi tersebut menghitung jarak antara dua titik `(x1, y1)` dan `(x2, y2)` menggunakan rumus jarak Euclidean.

---

### 2. Fungsi `dalamLingkaran`
```go
func dalamLingkaran(cx, cy, r, x, y int) bool {
    return jarak(cx, cy, x, y) < float64(r)
}
```
Fungsi tersebut menentukan apakah titik `(x, y)` berada di dalam lingkaran yang memiliki pusat `(cx, cy)` dan jari-jari `r`. 
Fungsi mengembalikan nilai `true` jika titik berada di dalam lingkaran, dan `false` jika tidak.

---

### 3. Fungsi `main`
```go
func main() {
    var cx1, cy1, r1, cx2, cy2, r2, x, y int

    fmt.Scan(&cx1, &cy1, &r1)
    fmt.Scan(&cx2, &cy2, &r2)
    fmt.Scan(&x, &y)

    lingkaran1 := dalamLingkaran(cx1, cy1, r1, x, y)
    lingkaran2 := dalamLingkaran(cx2, cy2, r2, x, y)

    if lingkaran1 && lingkaran2 {
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    } else if lingkaran1 {
        fmt.Println("Titik di dalam lingkaran 1")
    } else if lingkaran2 {
        fmt.Println("Titik di dalam lingkaran 2")
    } else {
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}
```
- Di dalam fungsi `main`, enam variabel integer (`cx1`, `cy1`, `r1`, `cx2`, `cy2`, `r2`) dideklarasikan untuk menyimpan data tentang dua lingkaran, dan dua variabel (`x`, `y`) untuk titik yang akan diperiksa.
- `fmt.Scan` digunakan untuk membaca input pengguna mengenai pusat dan jari-jari kedua lingkaran, serta koordinat titik.
- Program kemudian memeriksa apakah titik berada di dalam lingkaran 1, lingkaran 2, atau keduanya, dan mencetak hasil yang sesuai.

---


