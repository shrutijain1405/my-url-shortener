package connection

const (
    GetLongUrl   = "SELECT longUrl FROM urls WHERE shortUrl = ?"
    StoreUrlPair = "INSERT INTO urls (longUrl, shortUrl) VALUES (?, ?)"

    UpdateLongUrl    = "UPDATE urls SET longUrl = ? WHERE longUrl = ? and shortUrl = ?"
    UpdateLongUrlAll = "UPDATE urls SET longUrl = ? WHERE longUrl = ?"

    DeleteShortUrl = "DELETE FROM urls WHERE shortUrl = ?"
)
