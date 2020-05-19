# Sengkalan Generator

***Sengkalan Generator*** ini merupakan hasil _porting_ dari [Sengkalan Generator](https://github.com/lantip/sengkalan) dari Python ke Go yang dibuat oleh [Mas Lantip](https://github.com/lantip).

Untuk cara kerjanya, [silakan baca tulisan Mas Lantip di blognya](https://lantip.xyz/2020/05/membuat-sengkalan/).

## Kebutuhan

- Go versi 1.14
- Go Module

## Pembangunan

```shell script
go build
```

## Menjalankan

```shell script
./sengkalan [tahun]
```

### Contoh

```shell script
./sengkalan 2020
Sengkalan versi  0.2

📅 Tahun Masehi: 2020
☀️ Surya Sengkala: Mesat Sikara Rusak Mata
📜 Makna Surya Sengkala:
   > Mesat: pergi, menghindar, melesat
   > Sikara: pengacauan, tangan, campur tangan.
   > Rusak: rusak
   > Mata: mata

📅 Tahun Jawa: 1953
🌙 Candra Sengkala: Brama Raseksa Muka Luwih
📜 Makna Candra Sengkala:
   > Brama: api
   > Raseksa: raksasa
   > Muka: wajah, depan
   > Luwih: lebih, luar biasa

```

## Akan Dilakukan

- Mengubah ke aksara Jawa

## Lisensi

- [MIT](LICENSE.md)