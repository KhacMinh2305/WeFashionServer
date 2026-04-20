# WeFashionServer API & Config Documentation

## 1. Cấu hình .env

| Biến                | Ý nghĩa                        | Ví dụ giá trị                 |
|---------------------|--------------------------------|-------------------------------|
| DB_HOST_DEV         | Địa chỉ host DB (dev)          | localhost                     |
| DB_PORT_DEV         | Cổng DB (dev)                  | 5432                          |
| DB_USER_DEV         | Tên user DB (dev)              | postgres                      |
| DB_PASSWORD_DEV     | Mật khẩu DB (dev)              | 123456789                     |
| DB_NAME_DEV         | Tên database (dev)             | WeFashion                     |
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
- **Response thành công:**
```json
{
  "status_code": 200,
  "time": "...",
  "data": {
    "coupons": [
      {
        "id": 1,
        "shop_id": 1,
        "name": "Giảm 10%",
        "description": "Giảm giá cho đơn hàng đầu tiên",
        "discount": 10,
        "expired_at": "2026-12-31T23:59:59Z"
      }
    ]
  }
}
```


### 3.2. Lấy coupon theo id
- **Endpoint:** `GET /api/coupons/:id`
- **Header:** `Authorization: Bearer <token>`
- **Response thành công:**
```json
{
  "status_code": 200,
  "time": "...",
  "data": {
    "id": 1,
    "shop_id": 1,
    "name": "Giảm 10%",
    "description": "Giảm giá cho đơn hàng đầu tiên",
    "discount": 10,
    "expired_at": "2026-12-31T23:59:59Z"
  }
}
```


### 3.3. Lấy coupon của shop
- **Endpoint:** `GET /api/coupons/shop?shop_id=<id>`
- **Header:** `Authorization: Bearer <token>`
- **Response thành công:**
```json
{
  "status_code": 200,
  "time": "...",
  "data": {
    "shop_id": 1,
    "coupons": [
      {
        "id": 1,
        "shop_id": 1,
        "name": "Giảm 10%",
        "description": "Giảm giá cho đơn hàng đầu tiên",
        "discount": 10,
        "expired_at": "2026-12-31T23:59:59Z"
      }
    ]
  }
}
```


### 3.4. Lấy coupon khả dụng cho user
- **Endpoint:** `GET /api/coupons/user?user_id=<id>`
- **Header:** `Authorization: Bearer <token>`
- **Response thành công:**
```json
{
  "status_code": 200,
  "time": "...",
  "data": {
    "user_id": 2,
    "coupons": [
      {
        "id": 1,
        "shop_id": -1,
        "name": "Giảm 10% toàn hệ thống",
        "description": "Áp dụng cho mọi đơn hàng",
        "discount": 10,
        "expired_at": "2026-12-31T23:59:59Z"
      }
    ]
  }
}
```

### 3.5. Lấy coupon cho đơn hàng của shop
- **Endpoint:** `GET /api/coupons/order?shop_id=<id>`
- **Header:** `Authorization: Bearer <token>`
- **Response:** Danh sách coupon áp dụng cho đơn hàng của shop.

- **Response thành công:**
```json
{
  "status_code": 200,
  "time": "...",
  "data": {
    "coupons": [
      {
        "id": 1,
        "shop_id": 1,
        "name": "Giảm 10%",
        "description": "Giảm giá cho đơn hàng đầu tiên",
        "discount": 10,
        "expired_at": "2026-12-31T23:59:59Z"
      }
    ]
  }
}
```

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

## 6. Color APIs

### 6.1. Lấy tất cả màu sắc
- **Endpoint:** `GET /api/color`
- **Header:** `Authorization: Bearer <token>`
- **Response thành công:**
```json
{
  "status_code": 200,
  "time": "...",
  "data": {
    "colors": [
      { "id": 1, "rgb": "23,1,23" },
      { "id": 2, "rgb": "23,1,23" }
    ]
  }
}
```

### 6.2. Lấy màu sắc theo id
- **Endpoint:** `GET /api/color?color_id=<id>`
- **Header:** `Authorization: Bearer <token>`
- **Response thành công (tìm thấy):**
```json
{
  "status_code": 200,
  "time": "...",
  "data": { "id": 1, "rgb": "23,1,23" }
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

## 7. Account APIs

### 7.1. Đăng ký tài khoản
- **Endpoint:** `POST /api/account/register`
- **Header:** `Authorization: Bearer <token>`
- **Body:**
```json
{
  "email": "user@example.com",
  "username": "username",
  "password": "password"
}
```
- **Response thành công:**
```json
{
  "status_code": 200,
  "time": "...",
  "data": {
    "id": 1,
    "name": "username",
    "avatar_url": "...",
    "email": "user@example.com",
    "phone_number": "",
    "bio": ""
  }
}
```

### 7.2. Đăng nhập tài khoản
- **Endpoint:** `POST /api/account/login`
- **Header:** `Authorization: Bearer <token>`
- **Body:**
```json
{
  "username": "username",
  "password": "password"
}
```
- **Response thành công:**
```json
{
  "status_code": 200,
  "time": "...",
  "data": {
    "id": 1,
    "name": "username",
    "avatar_url": "...",
    "email": "user@example.com",
    "phone_number": "",
    "bio": ""
  }
}
```

### 7.3. Quên mật khẩu (gửi mã xác thực)
- **Endpoint:** `POST /api/account/forgot-password`
- **Header:** `Authorization: Bearer <token>`
- **Body:**
```json
{
  "email": "user@example.com"
}
```
- **Response thành công:**
```json
{
  "status_code": 200,
  "time": "...",
  "data": {
    "credential": "code"
  }
}
```

### 7.4. Xác thực mã quên mật khẩu
- **Endpoint:** `POST /api/account/forgot-password/validate`
- **Header:** `Authorization: Bearer <token>`
- **Body:**
```json
{
  "email": "user@example.com",
  "code": "1234",
  "credential": "code"
}
```
- **Response thành công:**
```json
{
  "status_code": 200,
  "time": "...",
  "data": {
    "is_valid": true,
    "Detail": "Correct code"
  }
}
```

### 7.5. Đổi mật khẩu
- **Endpoint:** `POST /api/account/change-password`
- **Header:** `Authorization: Bearer <token>`
- **Body:**
```json
{
  "username": "username",
  "old_password": "old_raw_password",
  "new_password": "new_raw_password"
}
```
- **Response thành công:**
```json
{
  "status_code": 200,
  "time": "...",
  "data": {
    "message": "Password changed successfully"
  }
}
```

## 8. User APIs

### 8.1. Lấy thông tin user theo id
- **Endpoint:** `GET /api/user?user_id=<id>`
- **Header:** `Authorization: Bearer <token>`
- **Response thành công (tìm thấy):**
```json
{
  "status_code": 200,
  "time": "...",
  "data": {
    "id": 1,
    "name": "username",
    "avatar_url": "...",
    "email": "user@example.com",
    "phone_number": "",
    "bio": ""
  }
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

### 8.2. Cập nhật thông tin user
- **Endpoint:** `POST /api/user/:id/update`
- **Header:** `Authorization: Bearer <token>`
- **Body:**
```json
{
  "id": 1,
  "name": "username",
  "avatar_url": "...",
  "email": "user@example.com",
  "phone_number": "",
  "bio": ""
}
```
- **Response thành công:**
```json
{
  "status_code": 200,
  "time": "...",
  "data": {
    "id": 1,
    "name": "username",
    "avatar_url": "...",
    "email": "user@example.com",
    "phone_number": "",
    "bio": ""
  }
}
```

---
- Tất cả các API đều yêu cầu token hợp lệ (JWT, truyền qua header Authorization).
- Response lỗi dạng:
```json
{
  "status_code": <mã lỗi>,
  "error": "...",
  "detail": "..."
}
```
