
### 📦 Temel İşlev Genişletmeleri

+ GET
+ SET
+ DELETE
+ EXISTS
+ UPDATE
+ UPSERT
+ FLUSH
+ DUMP
+ LOAD
+ RENAME
+ COPY
+ INCR / DECR (sayı değeri arttır/azalt)

### 🧠 Veri Özellikleri

+ TTL (EXPIRE)
+ PERSIST (TTL sıfırla)
+ KEYS (listele)
+ RANDOMKEY
+ TYPE (veri tipi göster)

### 🧩 Veri Türleri (Genişletilmiş KV)

+ string (default)
+ list (push, pop)
+ set (add, remove)
+ hashmap (hset, hget)
+ json (nested destekli key-value)

🛡 Güvenlik / Yetki

+ AUTH
+ ACL (kullanıcı izinleri)
+ IP Whitelist

### 🌐 Ağ / Protokol

+ HTTP API
+ WebSocket
+ TLS
+ Proxy Desteği
+ Rate Limiting

### 🔁 Sistem İşlevleri

+ SAVE
+ RESTORE
+ SNAPSHOT
+ BACKUP
+ REPLICATE
+ LOGROTATE
+ UPTIME
+ INFO / STATS
+ VERSION

### 🧪 Debug & Test

+ DEBUG command
+ MONITOR (gelen komutları canlı izle)
+ BENCHMARK (hız testi)
+ MEMORY DUMP
+ TRACE LOG

### 🚀 Performans

+ LRU Eviction
+ LFU Eviction
+ Multi-threaded handling
+ Batching
+ WAL (Write-Ahead Log)

### 📡 Bildirim / Event

+ WATCH key
+ PUB/SUB
+ OnChange Hook
+ Webhook

### 🧰 CLI Geliştirme

+ Komut geçmişi
+ Autocomplete
+ Çoklu bağlantı yönetimi
+ Renkli çıktı
+ Shell script desteği