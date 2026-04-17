# WeFashionServer API & .env Documentation

## 1. Cấu hình .env

| Biến                | Ý nghĩa                        | Ví dụ giá trị                 |
|---------------------|--------------------------------|-------------------------------|
| DB_HOST_DEV         | Địa chỉ host DB (dev)          | localhost address             |
| DB_PORT_DEV         | Cổng DB (dev)                  | Local postgres port           |
| DB_USER_DEV         | Tên user DB (dev)              | Local postgres account        |
| DB_PASSWORD_DEV     | Mật khẩu DB (dev)              | Local postgres password       |
| DB_NAME_DEV         | Tên database (dev)             | Database name                 |
| DB_HOST_PROD        | Địa chỉ host DB (prod)         | ...                           |
| DB_PORT_PROD        | Cổng DB (prod)                 | ...                           |
| DB_USER_PROD        | Tên user DB (prod)             | ...                           |
| DB_PASSWORD_PROD    | Mật khẩu DB (prod)             | ...                           |
| DB_NAME_PROD        | Tên database (prod)            | WeFashion                     |
| JWT_PRIVATE_KEY     | Secret key cho JWT             | YourSecretKeyHere             |
| ENV                 | Môi trường chạy (dev/prod)     | dev                           |

## 2. Authentication APIs

### 2.1. Lấy/gia hạn access token
- **Endpoint:** `POST /api/auth`
- **Body:**
```json
{
  "id": "<user_id>",
  "token": "<old_token hoặc rỗng>"
}
```
- **Response thành công:**
```json
{
  "status_code": 200,
  "time": "...",
  "data": {
    "token": "...",
    "created_at": "...",
    "expired_in": 3600
  }
}
```


## 3. Coupon APIs

### 3.1. Lấy tất cả coupon
- **Endpoint:** `GET /api/coupons`
- **Header:** `Authorization: Bearer <token>`
- **Response:** Danh sách coupon toàn hệ thống.

### 3.2. Lấy coupon theo id
- **Endpoint:** `GET /api/coupons/:id`
- **Header:** `Authorization: Bearer <token>`
- **Response:** Thông tin coupon theo id.

### 3.3. Lấy coupon của shop
- **Endpoint:** `GET /api/coupons/shop?shop_id=<id>`
- **Header:** `Authorization: Bearer <token>`
- **Response:** Danh sách coupon của shop.

### 3.4. Lấy coupon khả dụng cho user
- **Endpoint:** `GET /api/coupons/user?user_id=<id>`
- **Header:** `Authorization: Bearer <token>`
- **Response:** Danh sách coupon khả dụng cho user.

### 3.5. Lấy coupon cho đơn hàng của shop
- **Endpoint:** `GET /api/coupons/order?shop_id=<id>`
- **Header:** `Authorization: Bearer <token>`
- **Response:** Danh sách coupon áp dụng cho đơn hàng của shop.

## 4. Category APIs

### 4.1. Lấy tất cả category
- **Endpoint:** `GET /api/category`
- **Header:** `Authorization: Bearer <token>`
- **Response thành công:**
```json
{
  "status_code": 200,
  "time": "...",
  "data": {
    "categories": [
      { "id": 1, "name": "..." },
      { "id": 2, "name": "..." }
    ]
  }
}
```

### 4.2. Lấy category theo id
- **Endpoint:** `GET /api/category?category_id=<id>`
- **Header:** `Authorization: Bearer <token>`
- **Response thành công (tìm thấy):**
```json
{
  "status_code": 200,
  "time": "...",
  "data": { "id": 1, "name": "..." }
}
```
- **Response thành công (không tìm thấy):**
```json
{
  "status_code": 200,
  "time": "...",
  "data": null
}
```

## 5. Shop APIs

### 5.1. Lấy tất cả shop
- **Endpoint:** `GET /api/shop`
- **Header:** `Authorization: Bearer <token>`
- **Response thành công:**
```json
{
  "status_code": 200,
  "time": "...",
  "data": {
    "shops": [
      {
        "id": 1,
        "name": "...",
        "avatar_url": "...",
        "email": "...",
        "phone_number": "...",
        "bio": "...",
        "rate_amount": 10,
        "rating": 4.5,
        "followers": 100
      }
    ]
  }
}
```

### 5.2. Follow/Unfollow shop
- **Endpoint:** `PUT /api/shop/follow?shop_id=<id>&follow=<true|false>`
- **Header:** `Authorization: Bearer <token>`
- **Response thành công:**
```json
{
  "status_code": 200,
  "time": "...",
  "data": {
    "id": 1,
    "name": "...",
    "avatar_url": "...",
    "email": "...",
    "phone_number": "...",
    "bio": "...",
    "rate_amount": 10,
    "rating": 4.5,
    "followers": 101
  }
}
```

---
- Tất cả các API coupon đều yêu cầu token hợp lệ (JWT, truyền qua header Authorization).
- Nếu id shop/user không tồn tại sẽ trả về lỗi 404.
- Response lỗi dạng:
```json
{
  "status_code": <mã lỗi>,
  "error": "...",
  "detail": "..."
}
```
