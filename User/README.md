# Google Function Erfahtech - User

Welcome to the Google Cloud Function Erfahtech for URSMARTECOSYSTEM.MY.ID!

## Isi

- Insert Device
  Endpoint insert-device ini adalah titik akses atau URL dalam Aplikasi USE yang digunakan user untuk melakukan penambahan device. Pada kode akan memanggil fungsi GCFInsertDevice pada package beurse yang telah dideploy di Go.pkg.

- Get Device
  Endpoint get-device ini adalah titik akses atau URL dalam Aplikasi USE yang digunakan untuk mendapatkan device yang dimiliki oleh pengguna. Pada kode akan memanggil fungsi GCFGetDevice pada package beurse yang telah dideploy di Go.pkg.

- Get Profile
  Endpoint get-profile ini adalah titik akses atau URL dalam Aplikasi USE untuk mendapatkan data profile pengguna. Pada kode akan memanggil fungsi GCFGetUserByEmail pada package beurse yang telah dideploy di Go.pkg.

- Update Device
  Endpoint update-device ini adalah titik akses atau URL dalam Aplikasi USE yang digunakan memperbaharui data device pada menu setting, memperbaharui jika pengguna ingin merubah nama atau topic device yang dimiliki. Pada kode akan memanggil fungsi GCFHandlerUpdateDevice pada package beurse yang telah dideploy di Go.pkg.

- Update Status Device
  Endpoint update-statusdevice ini adalah titik akses atau URL dalam Aplikasi USE yang digunakan memperbaharui status device pada menu kontrol, memperbaharui apakah status device ini on/off. Pada kode akan memanggil fungsi GCFHandlerUpdateStatusDevice pada package beurse yang telah dideploy di Go.pkg.

- Delete Device
  Endpoint delete-device ini adalah titik akses atau URL dalam Aplikasi USE yang digunakan menghapus device yang dimiliki oleh pengguna. Pada kode akan memanggil fungsi GCFHandlerUpdateDevice pada package beurse yang telah dideploy di Go.pkg.
