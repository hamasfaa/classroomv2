# ClassroomV2

ClassroomV2 adalah aplikasi sederhana yang memfasilitasi pengelolaan kelas, murid, dan tugas secara online. Proyek ini memanfaatkan Vue.js sebagai frontend dan Go sebagai backend untuk kinerja yang cepat dan fungsionalitas yang fleksibel.

## Fitur Utama

1. Manajemen Kelas: Membuat, mengedit, dan menghapus kelas.
2. Pengelolaan Pengguna: Menambahkan serta mengelola data murid dan guru.
3. Penugasan dan Evaluasi: Mengunggah tugas, mengumpulkan jawaban, dan memberikan nilai.

## Teknologi yang Digunakan

- **Frontend (Vue.js)**
  - Mengelola tampilan antarmuka pengguna.
  - Struktur SPA (Single Page Application) untuk pengalaman yang responsif.
- **Backend (Go)**
  - API RESTful untuk pemrosesan data kelas, murid, dan tugas.
  - Efisien serta mudah dikembangkan berkat sifat Go yang sederhana dan cepat.

## Cara Menjalankan

1. **Kloning Repository**  
   git clone https://github.com/hamasfaa/classroomv2.git

2. **Frontend (Vue.js)**

   - Masuk ke folder frontend, misalnya cd frontend
   - Install dependencies:  
     npm install
   - Jalankan aplikasi:  
     npm run serve

3. **Backend (Go)**

   - Masuk ke folder backend, misalnya cd backend
   - Pastikan Go telah terpasang di sistem.
   - Jalankan perintah:  
     go run main.go

4. **Akses Aplikasi**
   - Buka http://localhost:8080 (atau port yang telah ditentukan di server Go) di browser untuk API.
   - Buka http://localhost:5173 (atau port yang ditentukan di Vue) untuk antarmuka pengguna.
