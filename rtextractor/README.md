# rtextractor
Based on the CSV file in https://github.com/datanesia/jakartaopendata/blob/master/data-asmas-apbd-2015.csv, this program extracted probable RT values from text of the address (alamat_lokasi), problem (masalah), or suggestion (usulan_solusi).

## Sample case:

> **Alamat:** Jl. Cempaka Putih Utara RT.002,003,006,007 Jl. Letjend Suprapto, RT.001,002,003,0014,0015 Jl. Tembaga Raya RT.0015,0012,0011,0010,009 RW.03
> **Masalah:** Banjir dikarenakan saluran gorong-gorong tersumbat sampah
> **Solusi:** Normalisasi saluran air/got Jl. Cempaka Putih Utara RT.002,003,006,007 Jl. Letjend Suprapto, RT.001,002,003,0014,0015 Jl. Tembaga Raya RT.0015,0012,0011,0010,009 RW.03 Kel. Harapan Mulia

> **Found RT:** [002 003 006 007 001 0014 0015 0012 0011 0010 009]

Values are inserted at a new column, "lokasi_rt", right before the already-existing "lokasi_rw" column.

RTs are extracted using regex pattern "(?i)RT(\W*\d+)*", which are then filtered again to extract numerical values and to eliminate edge cases where the RT number are cojoined with the RW number.

Options:
```
-i Path to input file (mandatory)
-o Path to output file (mandatory)
```

Example usage:

```
go run rtextractor.go -i data/data-asmas-apbd-2015.csv -o data/data-asmas-apbd-2015-rt.csv
Finished processing 40511 records, found 63893 RT
```
