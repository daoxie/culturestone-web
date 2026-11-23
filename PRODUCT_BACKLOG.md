# Product Backlog — Culturstone

Dokumen ini memetakan fitur Culturstone ke dalam **Epic → User Story → Fungsi → Use Case → Backlog**. Isinya disusun berdasarkan struktur repository dan panduan yang tersedia: situs publik katalog batu alam dan panel admin; frontend React/Vite; backend Go dengan SQLite.

> **Cara membaca status:** `Tersedia` berarti kemampuan tersebut disebutkan atau terlihat pada implementasi saat ini. `Perlu verifikasi` berarti dokumentasi menyebutnya, tetapi perilaku lengkapnya belum dipastikan dari inventaris ini. `Usulan` adalah pekerjaan backlog, bukan klaim bahwa fitur sudah ada.
>
> Prioritas: **P0** = penting untuk operasi/keamanan; **P1** = peningkatan utama; **P2** = penyempurnaan. Estimasi dan urutan sprint perlu dikonfirmasi oleh pemilik produk.

## Ringkasan produk

Culturstone adalah situs profil perusahaan dan katalog produk batu alam. Pengunjung dapat mengenal perusahaan, menelusuri produk dan portofolio, membaca testimoni, serta menghubungi perusahaan. Admin mengelola konten melalui panel admin. Model data yang teridentifikasi mencakup produk, kategori, portofolio, testimoni, pesan kontak, visitor, dan admin.

## Peta Epic

| Epic | Outcome | Fungsi yang ditempatkan di dalam Epic |
|---|---|---|
| E-01 — Informasi perusahaan & navigasi | Pengunjung memahami Culturstone dan mudah menemukan informasi | Beranda, profil perusahaan, navigasi, informasi kontak |
| E-02 — Katalog produk & kategori | Pengunjung menemukan produk yang sesuai kebutuhannya | Daftar produk, pencarian nama produk, detail spesifikasi, kategori |
| E-03 — Portofolio proyek | Pengunjung dapat menilai hasil pekerjaan | Daftar portofolio, detail proyek dan gambar |
| E-04 — Testimoni pelanggan | Pengunjung dapat membaca pengalaman klien | Daftar testimoni, foto/video, relasi dengan portofolio |
| E-05 — Kontak & prospek | Pengunjung dapat mengirim pertanyaan dan admin menindaklanjutinya | Form kontak, penyimpanan pesan, daftar pesan admin |
| E-06 — Administrasi & analitik | Admin dapat mengelola konten dan memantau aktivitas situs | Login admin, dashboard statistik, visitor dashboard, pengelolaan konten |
| E-07 — Keamanan, kualitas & operasional | Sistem dan data dikelola dengan aman dan andal | Validasi, otorisasi, upload media, error handling, testing, dokumentasi |

---

## E-01 — Informasi perusahaan & navigasi

**Fungsi:** beranda, profil perusahaan, navigasi halaman, informasi perusahaan.

### User stories
- **US-01.1 (P1):** Sebagai pengunjung, saya ingin melihat ringkasan perusahaan di beranda agar memahami produk dan layanan Culturstone.
- **US-01.2 (P1):** Sebagai pengunjung, saya ingin membaca profil, sejarah, visi, dan misi perusahaan agar dapat menilai kredibilitas perusahaan.
- **US-01.3 (P1):** Sebagai pengunjung, saya ingin berpindah ke halaman produk, portofolio, testimoni, dan kontak melalui navigasi yang konsisten.

### Use cases
| ID | Aktor | Use case | Hasil yang diharapkan |
|---|---|---|---|
| UC-01 | Pengunjung | Membuka beranda | Ringkasan perusahaan dan tautan ke konten utama ditampilkan |
| UC-02 | Pengunjung | Membaca profil perusahaan | Informasi perusahaan disajikan dengan jelas |
| UC-03 | Pengunjung | Berpindah halaman | Halaman yang dipilih terbuka dan navigasi tetap dapat digunakan |

### Backlog
- **BL-01 (P1, Usulan):** Tinjau kelengkapan, akurasi, dan konsistensi konten beranda/profil.
- **BL-02 (P1, Usulan):** Uji navigasi seluruh halaman pada desktop dan perangkat mobile.
- **BL-03 (P2, Usulan):** Tambahkan pengelolaan konten profil dari panel admin bila perubahan konten saat ini masih memerlukan perubahan kode.

## E-02 — Katalog produk & kategori

**Fungsi:** menampilkan katalog, mencari produk berdasarkan nama, menyajikan kategori dan atribut produk (deskripsi, kualitas, ukuran, finishing, gambar).

### User stories
- **US-02.1 (P1):** Sebagai pengunjung, saya ingin melihat daftar produk agar mengetahui pilihan batu alam yang tersedia.
- **US-02.2 (P1):** Sebagai pengunjung, saya ingin mencari produk berdasarkan nama agar cepat menemukan produk tertentu.
- **US-02.3 (P1):** Sebagai pengunjung, saya ingin melihat spesifikasi produk agar dapat membandingkan pilihan.
- **US-02.4 (P0):** Sebagai admin, saya ingin menambah dan memperbarui produk beserta gambar dan spesifikasinya agar katalog tetap akurat.
- **US-02.5 (P1):** Sebagai admin, saya ingin mengelola kategori produk agar katalog tersusun.

### Use cases
| ID | Aktor | Use case | Hasil yang diharapkan |
|---|---|---|---|
| UC-04 | Pengunjung | Melihat katalog produk | Daftar produk ditampilkan dari API publik |
| UC-05 | Pengunjung | Mencari produk | Daftar dibatasi berdasarkan teks pencarian nama |
| UC-06 | Pengunjung | Membuka informasi produk | Deskripsi dan atribut produk tersedia untuk dibaca |
| UC-07 | Admin | Membuat atau mengubah produk | Data produk tersimpan; gambar dapat diunggah |
| UC-08 | Admin | Mengelola kategori | Data kategori dapat dikelola dan digunakan pada katalog |

### Backlog
- **BL-04 (P0, Tersedia; verifikasi):** Verifikasi alur admin membuat/mengubah produk dan mengunggah gambar.
- **BL-05 (P1, Tersedia; verifikasi):** Uji endpoint publik produk dan pencarian nama, termasuk hasil kosong.
- **BL-06 (P1, Perlu verifikasi):** Pastikan penghapusan produk/kategori, validasi input, dan pesan kesalahan berfungsi serta terotorisasi.
- **BL-07 (P2, Usulan):** Tambahkan filter kategori, pengurutan, dan pagination jika katalog membutuhkan skala lebih besar.
- **BL-08 (P2, Usulan):** Tambahkan optimasi gambar dan state loading/kosong/error pada katalog.

## E-03 — Portofolio proyek

**Fungsi:** menampilkan dan mengelola proyek, deskripsi serta gambar; menghubungkan proyek dengan testimoni.

### User stories
- **US-03.1 (P1):** Sebagai pengunjung, saya ingin menelusuri portofolio agar dapat melihat proyek yang telah diselesaikan.
- **US-03.2 (P1):** Sebagai pengunjung, saya ingin membaca deskripsi dan melihat gambar proyek agar memahami hasil pekerjaan.
- **US-03.3 (P1):** Sebagai admin, saya ingin menambah atau memperbarui entri portofolio agar contoh pekerjaan tetap relevan.

### Use cases
| ID | Aktor | Use case | Hasil yang diharapkan |
|---|---|---|---|
| UC-09 | Pengunjung | Melihat daftar portofolio | Daftar proyek ditampilkan |
| UC-10 | Pengunjung | Membaca proyek | Judul, deskripsi, dan gambar proyek dapat dilihat |
| UC-11 | Admin | Membuat/mengubah portofolio | Data proyek tersimpan dengan gambar opsional |

### Backlog
- **BL-09 (P1, Tersedia; verifikasi):** Uji daftar portofolio publik serta pembuatan/perubahan portofolio oleh admin.
- **BL-10 (P1, Perlu verifikasi):** Pastikan penghapusan portofolio dan penanganan kegagalan upload tersedia.
- **BL-11 (P2, Usulan):** Tambahkan urutan/tanggal proyek dan tampilan detail proyek bila diperlukan oleh kebutuhan bisnis.

## E-04 — Testimoni pelanggan

**Fungsi:** menampilkan testimoni dan media, mengelola data testimoni, mengaitkan testimoni dengan portofolio.

### User stories
- **US-04.1 (P1):** Sebagai pengunjung, saya ingin membaca testimoni agar memahami pengalaman klien.
- **US-04.2 (P1):** Sebagai pengunjung, saya ingin melihat foto atau video testimoni bila tersedia agar konteks testimoni lebih meyakinkan.
- **US-04.3 (P1):** Sebagai admin, saya ingin menambah, mengubah, dan menghapus testimoni agar konten tetap relevan.
- **US-04.4 (P2):** Sebagai admin, saya ingin menghubungkan testimoni dengan proyek portofolio terkait.

### Use cases
| ID | Aktor | Use case | Hasil yang diharapkan |
|---|---|---|---|
| UC-12 | Pengunjung | Membaca testimoni | Testimoni ditampilkan; respons kosong tetap berupa daftar kosong |
| UC-13 | Admin | Membuat/mengubah testimoni | Nama klien, isi, dan media tersimpan |
| UC-14 | Admin | Mengaitkan testimoni dengan proyek | Relasi portofolio valid atau kosong |
| UC-15 | Admin | Menghapus testimoni | Data yang dipilih dihapus dengan hasil yang jelas |

### Backlog
- **BL-12 (P1, Tersedia; verifikasi):** Uji CRUD testimoni, termasuk kondisi data portofolio tidak ditemukan.
- **BL-13 (P1, Tersedia; verifikasi):** Uji upload media foto/video dan kompatibilitas penampilannya di halaman publik.
- **BL-14 (P2, Usulan):** Tambahkan konfirmasi sebelum penghapusan dan validasi ukuran/jenis file.

## E-05 — Kontak & prospek

**Fungsi:** form kontak, penyimpanan pesan, daftar pesan untuk admin, penanda pesan dibaca.

### User stories
- **US-05.1 (P0):** Sebagai pengunjung, saya ingin mengirim nama, email/telepon, dan pesan agar dapat menghubungi perusahaan.
- **US-05.2 (P0):** Sebagai admin, saya ingin melihat pesan masuk agar dapat menindaklanjuti calon pelanggan.
- **US-05.3 (P1):** Sebagai admin, saya ingin membedakan pesan yang sudah dibaca agar tidak melewatkan permintaan.

### Use cases
| ID | Aktor | Use case | Hasil yang diharapkan |
|---|---|---|---|
| UC-16 | Pengunjung | Mengirim formulir kontak | Pesan tersimpan dan pengunjung memperoleh konfirmasi yang benar |
| UC-17 | Admin | Melihat pesan masuk | Pesan ditampilkan dengan urutan terbaru |
| UC-18 | Admin | Menandai/mengelola pesan | Status tindak lanjut dapat diketahui |

### Backlog
- **BL-15 (P0, Tersedia; verifikasi):** Uji submit JSON dan form, validasi field, serta respons ketika penyimpanan gagal.
- **BL-16 (P0, Perlu verifikasi):** Pastikan endpoint admin untuk pesan hanya dapat diakses setelah autentikasi/otorisasi.
- **BL-17 (P1, Perlu verifikasi):** Pastikan status baca pesan dapat diperbarui; model memiliki atribut `IsRead`.
- **BL-18 (P2, Usulan):** Tambahkan perlindungan spam/rate limiting dan kebijakan retensi pesan.

## E-06 — Administrasi & analitik

**Fungsi:** login admin, dashboard ringkasan, statistik visitor, navigasi panel, pengelolaan konten.

### User stories
- **US-06.1 (P0):** Sebagai admin, saya ingin login agar dapat mengakses panel pengelolaan.
- **US-06.2 (P0):** Sebagai admin, saya ingin dashboard ringkasan agar dapat melihat jumlah produk, pesan, portofolio, testimoni, dan kategori.
- **US-06.3 (P1):** Sebagai admin, saya ingin melihat statistik pengunjung harian agar dapat memahami penggunaan situs.
- **US-06.4 (P0):** Sebagai admin, saya ingin hanya admin terautentikasi yang dapat mengubah konten.

### Use cases
| ID | Aktor | Use case | Hasil yang diharapkan |
|---|---|---|---|
| UC-19 | Admin | Login | Kredensial valid menghasilkan token; kredensial tidak valid ditolak |
| UC-20 | Admin | Membuka dashboard | Ringkasan jumlah konten dan visitor ditampilkan |
| UC-21 | Sistem | Mencatat kunjungan | Jumlah kunjungan harian diperbarui |
| UC-22 | Admin | Mengelola konten | Operasi admin tersedia sesuai hak akses |

### Backlog
- **BL-19 (P0, Tersedia; verifikasi):** Uji login, hash password, masa berlaku/validasi token, dan alur logout di frontend.
- **BL-20 (P0, Perlu verifikasi):** Audit seluruh route admin: route terlindungi, tidak ada kredensial/token di log atau respons yang tidak perlu.
- **BL-21 (P1, Tersedia; verifikasi):** Cocokkan angka dashboard dengan data aktual, termasuk data kosong dan error database.
- **BL-22 (P1, Perlu verifikasi):** Tinjau akurasi visitor counter, pengecualian request non-pengunjung, dan pencatatan berulang.
- **BL-23 (P2, Usulan):** Tambahkan audit trail perubahan konten dan filter pesan/visitor bila kebutuhan admin mengharuskan.

## E-07 — Keamanan, kualitas & operasional

**Fungsi:** otorisasi, validasi request, upload aman, penanganan error, stabilitas UI/API, dokumentasi dan testing.

### User stories
- **US-07.1 (P0):** Sebagai pemilik sistem, saya ingin route administrasi terlindungi agar konten tidak dapat diubah oleh pihak tanpa hak.
- **US-07.2 (P0):** Sebagai pengguna, saya ingin mendapat respons yang jelas ketika form atau API gagal.
- **US-07.3 (P1):** Sebagai pengelola situs, saya ingin upload file dibatasi dan divalidasi agar penyimpanan tetap aman.
- **US-07.4 (P1):** Sebagai tim pengembang, saya ingin memiliki pengujian dan dokumentasi agar perubahan dapat dirilis dengan lebih andal.

### Use cases
| ID | Aktor | Use case | Hasil yang diharapkan |
|---|---|---|---|
| UC-23 | Sistem | Memvalidasi request | Input tidak valid ditolak tanpa merusak data |
| UC-24 | Sistem | Memeriksa akses route admin | Request tanpa akses yang sah ditolak |
| UC-25 | Admin | Mengunggah media | File yang valid tersimpan dan URL dapat digunakan |
| UC-26 | Pengunjung | Menghadapi kegagalan API | Pesan error dapat dipahami dan halaman tetap stabil |

### Backlog
- **BL-24 (P0, Usulan):** Audit pendaftaran admin; pastikan endpoint register tidak membuka pembuatan akun admin tanpa kontrol yang disengaja.
- **BL-25 (P0, Usulan):** Terapkan validasi wajib untuk kredensial, pesan kontak, serta payload create/update.
- **BL-26 (P0, Usulan):** Batasi jenis dan ukuran unggahan; pastikan nama file dan penyimpanan tidak membuka akses berbahaya.
- **BL-27 (P1, Usulan):** Tambahkan test backend untuk autentikasi, CRUD penting, pencarian produk, upload, dan form kontak.
- **BL-28 (P1, Usulan):** Tambahkan test frontend untuk navigasi, form kontak, login admin, serta state loading/error.
- **BL-29 (P1, Usulan):** Dokumentasikan kontrak endpoint dan konfigurasi deployment/secret tanpa menaruh secret di repository.
- **BL-30 (P2, Usulan):** Tinjau aksesibilitas, responsivitas, performa gambar, dan metadata SEO halaman publik.

---

## Urutan backlog lintas Epic

| Prioritas | ID | Item | Epic | Status awal |
|---|---|---|---|---|
| P0 | BL-19 | Verifikasi login, token, dan logout | E-06 | Verifikasi |
| P0 | BL-20 | Audit perlindungan seluruh route admin | E-06 | Perlu verifikasi |
| P0 | BL-24 | Audit kontrol pendaftaran admin | E-07 | Usulan |
| P0 | BL-25 | Validasi input endpoint utama | E-07 | Usulan |
| P0 | BL-15 | Verifikasi form kontak dan penanganan kegagalan | E-05 | Verifikasi |
| P0 | BL-16 | Verifikasi akses admin untuk pesan masuk | E-05 | Perlu verifikasi |
| P0 | BL-04 | Verifikasi CRUD produk dan upload gambar | E-02 | Verifikasi |
| P1 | BL-05 | Verifikasi katalog dan pencarian produk | E-02 | Verifikasi |
| P1 | BL-09 | Verifikasi daftar dan pengelolaan portofolio | E-03 | Verifikasi |
| P1 | BL-12 | Verifikasi CRUD serta relasi testimoni | E-04 | Verifikasi |
| P1 | BL-21 | Uji dashboard dengan data nyata/kosong/error | E-06 | Verifikasi |
| P1 | BL-26 | Validasi keamanan file upload | E-07 | Usulan |
| P1 | BL-27 | Tambahkan test backend untuk alur kritis | E-07 | Usulan |
| P1 | BL-28 | Tambahkan test frontend untuk alur kritis | E-07 | Usulan |
| P2 | BL-07 | Filter kategori, pengurutan, pagination | E-02 | Usulan |
| P2 | BL-18 | Perlindungan spam dan retensi pesan | E-05 | Usulan |
| P2 | BL-23 | Audit trail dan filter admin | E-06 | Usulan |
| P2 | BL-30 | Aksesibilitas, SEO, dan performa | E-07 | Usulan |

## Definition of Done

Sebuah backlog item dapat dinyatakan selesai apabila:
1. Kriteria penerimaan/user story terpenuhi dan perubahan sesuai Epic yang dituju.
2. Validasi akses diterapkan untuk operasi admin; tidak ada secret yang dikomit.
3. Kondisi sukses, data kosong, input tidak valid, dan kegagalan dependency ditangani dengan wajar.
4. Pengujian yang relevan lulus atau keterbatasan pengujian dicatat.
5. Dokumentasi pengguna/developer diperbarui bila perilaku atau konfigurasi berubah.
6. Perubahan frontend diuji pada ukuran layar utama dan perubahan API tetap kompatibel dengan klien.

## Catatan perencanaan

Backlog ini adalah baseline untuk refinement, bukan klaim bahwa semua usulan telah dikerjakan. Tim perlu mengonfirmasi prioritas bisnis, kriteria penerimaan, estimasi, pemilik, serta apakah item verifikasi sudah tertangani sebelum memindahkannya ke sprint. Story point tidak ditentukan karena belum ada kapasitas atau velocity tim yang diberikan.
