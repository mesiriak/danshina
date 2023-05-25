# Danshina

The main target of this project - clicker service.

## Routes

* GET `https://host/api/v1/baby/` - get list of babies. Query params: `limit`, `offset`.
* GET `https://host/api/v1/baby/random/` - get 2 babies by default. Can be regulated by query param `size`.
* POST `https://host/api/v1/baby/` - create baby. Body `{"nickname": str, "picture": bytearray}`.
* PATCH `https://host/api/v1/baby/{uuid}` - increment baby counter.

Used CDN for pictures - **`Cloudflare`**
