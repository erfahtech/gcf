# Google Function Erfahtech - Auth

Welcome to the Google Cloud Function Erfahtech for URSMARTECOSYSTEM.MY.ID!

## Isi

- Login

  Endpoint login ini adalah titik akses atau URL dalam Aplikasi USE yang memungkinkan pengguna untuk melakukan proses login. Pada kode akan memanggil fungsi GCFHandlerLogin pada package beurse yang telah dideploy di Go.pkg.

- SignUP

  Endpoint signup ini adalah titik akses atau URL dalam Aplikasi USE yang memungkinkan pengguna untuk melakukan signup atau pendaftaran yang akan digunakan untuk login nantinya. Pada kode akan memanggil fungsi GCFHandlerSignup pada package beurse yang telah dideploy di Go.pkg.

- Send OTP

  Endpoint send-otp ini adalah titik akses atau URL dalam Aplikasi USE yang digunakan untuk melakukan pengiriman kode otp yang akan digunakan pengguna untuk reset password. Pada kode akan memanggil fungsi GCFHandlerSendOTP pada package beurse yang telah dideploy di Go.pkg.

- Verify OTP

  Endpoint verify-otp ini adalah titik akses atau URL dalam Aplikasi USE yang digunakan untuk melakukan verifikasi kode otp yang akan dimasukan pengguna untuk reset password. Pada kode akan memanggil fungsi GCFHandlerVerifyOTP pada package beurse yang telah dideploy di Go.pkg.

- Reset Password

  Endpoint reset-password ini adalah titik akses atau URL dalam Aplikasi USE yang digunakan untuk melakukan pengiriman atau update password baru yang telah dimasukan oleh pengguna. Pada kode akan memanggil fungsi GCFHandlerSendOTP pada package beurse yang telah dideploy di Go.pkg.
