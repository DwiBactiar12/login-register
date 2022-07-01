# API Manage User

untuk mengunakan API ini anda perlu untuk clone di laptop anda
dengan cara https://github.com/DwiBactiar12/login-register.git
dan sesuaikan file .env sesuai laptop anda

berikut link untuk akses:

- http://localhost:(port)/registeradmin {POST}
- http://localhost:(port)/loginAdmin {POST}
  berikut login untuk user
- http://localhost:(port)/login {POST}
  berikut hanya bisa diakses jika memiliki token JWT
- http://localhost:(port)/users {POST}
- http://localhost:(port)/users {GET}
- http://localhost:(port)/users/me {GET}
- http://localhost:(port)/users/:id {GET}
  berikut hanya admin yang bisa akses
- http://localhost:(port)/users/:id {PUT}
- http://localhost:(port)/users/:id {DELETE}
- http://localhost:(port)/admin/me {GET}
