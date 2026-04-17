# WeFashionServer API & .env Documentation

## 1. Cấu hình .env

| Biến                | Ý nghĩa                        | Ví dụ giá trị                  |
|---------------------|--------------------------------|-------------------------------|
| DB_HOST_DEV         | Địa chỉ host DB (dev)          | localhost                     |
| DB_PORT_DEV         | Cổng DB (dev)                  | 5432                          |
| DB_USER_DEV         | Tên user DB (dev)              | postgres                      |
| DB_PASSWORD_DEV     | Mật khẩu DB (dev)              | Minh23012003?                 |
| DB_NAME_DEV         | Tên database (dev)             | WeFashion                     |
| DB_HOST_PROD        | Địa chỉ host DB (prod)         | ...                           |
| DB_PORT_PROD        | Cổng DB (prod)                 | ...                           |
| DB_USER_PROD        | Tên user DB (prod)             | ...                           |
| DB_PASSWORD_PROD    | Mật khẩu DB (prod)             | ...                           |
| DB_NAME_PROD        | Tên database (prod)            | WeFashion                     |
| JWT_PRIVATE_KEY     | Secret key cho JWT             | WeFashionJWTPrivateKey...     |
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
- **Response lỗi:**
```json
{
  "status_code": 401,
  "error": "...",
  "detail": "..."
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
