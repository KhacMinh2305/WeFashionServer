package mock

import (
	"WeFashionServer/infrastructure/model"
)

// Mock data for Shoes category (giày thể thao, giày da, sneaker, boots, ...)
var MockShoesProducts = []model.Product{
	// 30 sản phẩm giày thực tế, đa dạng thương hiệu, hình ảnh, mô tả, rating, sold, liked, ShopId khác nhau
	// CategoryId = 3 (Shoes)
	{Id: 1, Name: "Nike Air Force 1", ImageUrl: "https://static.nike.com/air-force-1.jpg", Description: "Nike Air Force 1 cổ điển, chất liệu da cao cấp, đế Air êm ái.", Rating: 4.9, SoldAmount: 2500, LikedAmount: 2100, CategoryId: 3, ShopId: 1},
	{Id: 2, Name: "Adidas Ultraboost 21", ImageUrl: "https://assets.adidas.com/ultraboost-21.jpg", Description: "Adidas Ultraboost 21, đệm Boost siêu êm, phù hợp chạy bộ.", Rating: 4.8, SoldAmount: 1800, LikedAmount: 1600, CategoryId: 3, ShopId: 2},
	{Id: 3, Name: "Converse Chuck Taylor All Star", ImageUrl: "https://converse.com/chuck-taylor.jpg", Description: "Converse Chuck Taylor All Star, phong cách cổ điển, dễ phối đồ.", Rating: 4.7, SoldAmount: 1700, LikedAmount: 1500, CategoryId: 3, ShopId: 3},
	{Id: 4, Name: "Vans Old Skool", ImageUrl: "https://vans.com/old-skool.jpg", Description: "Vans Old Skool, thiết kế trẻ trung, đế cao su bền bỉ.", Rating: 4.6, SoldAmount: 1600, LikedAmount: 1400, CategoryId: 3, ShopId: 4},
	{Id: 5, Name: "Puma Suede Classic", ImageUrl: "https://puma.com/suede-classic.jpg", Description: "Puma Suede Classic, chất liệu da lộn, phong cách retro.", Rating: 4.5, SoldAmount: 1500, LikedAmount: 1300, CategoryId: 3, ShopId: 5},
	{Id: 6, Name: "New Balance 574", ImageUrl: "https://newbalance.com/574.jpg", Description: "New Balance 574, đệm EVA nhẹ, thiết kế cổ điển.", Rating: 4.7, SoldAmount: 1400, LikedAmount: 1200, CategoryId: 3, ShopId: 6},
	{Id: 7, Name: "Reebok Classic Leather", ImageUrl: "https://reebok.com/classic-leather.jpg", Description: "Reebok Classic Leather, chất liệu da mềm, đế cao su bền.", Rating: 4.6, SoldAmount: 1300, LikedAmount: 1100, CategoryId: 3, ShopId: 7},
	{Id: 8, Name: "Fila Disruptor II", ImageUrl: "https://fila.com/disruptor-ii.jpg", Description: "Fila Disruptor II, thiết kế chunky, cá tính mạnh mẽ.", Rating: 4.5, SoldAmount: 1200, LikedAmount: 1000, CategoryId: 3, ShopId: 8},
	{Id: 9, Name: "Balenciaga Triple S", ImageUrl: "https://balenciaga.com/triple-s.jpg", Description: "Balenciaga Triple S, thiết kế hầm hố, thời trang cao cấp.", Rating: 4.8, SoldAmount: 1100, LikedAmount: 950, CategoryId: 3, ShopId: 9},
	{Id: 10, Name: "Gucci Ace Sneaker", ImageUrl: "https://gucci.com/ace-sneaker.jpg", Description: "Gucci Ace Sneaker, logo sọc đặc trưng, chất liệu cao cấp.", Rating: 4.9, SoldAmount: 1000, LikedAmount: 900, CategoryId: 3, ShopId: 10},
	{Id: 11, Name: "Nike Air Max 97", ImageUrl: "https://static.nike.com/air-max-97.jpg", Description: "Nike Air Max 97, thiết kế sóng nước, đệm Air toàn bộ.", Rating: 4.8, SoldAmount: 950, LikedAmount: 850, CategoryId: 3, ShopId: 1},
	{Id: 12, Name: "Adidas Stan Smith", ImageUrl: "https://assets.adidas.com/stan-smith.jpg", Description: "Adidas Stan Smith, phong cách tối giản, dễ phối đồ.", Rating: 4.7, SoldAmount: 900, LikedAmount: 800, CategoryId: 3, ShopId: 2},
	{Id: 13, Name: "Converse One Star", ImageUrl: "https://converse.com/one-star.jpg", Description: "Converse One Star, thiết kế trẻ trung, logo ngôi sao.", Rating: 4.6, SoldAmount: 850, LikedAmount: 750, CategoryId: 3, ShopId: 3},
	{Id: 14, Name: "Vans Sk8-Hi", ImageUrl: "https://vans.com/sk8-hi.jpg", Description: "Vans Sk8-Hi, cổ cao, bảo vệ mắt cá chân.", Rating: 4.5, SoldAmount: 800, LikedAmount: 700, CategoryId: 3, ShopId: 4},
	{Id: 15, Name: "Puma RS-X", ImageUrl: "https://puma.com/rs-x.jpg", Description: "Puma RS-X, thiết kế hiện đại, phối màu nổi bật.", Rating: 4.7, SoldAmount: 750, LikedAmount: 650, CategoryId: 3, ShopId: 5},
	{Id: 16, Name: "New Balance 997H", ImageUrl: "https://newbalance.com/997h.jpg", Description: "New Balance 997H, đệm nhẹ, phong cách retro.", Rating: 4.6, SoldAmount: 700, LikedAmount: 600, CategoryId: 3, ShopId: 6},
	{Id: 17, Name: "Reebok Club C 85", ImageUrl: "https://reebok.com/club-c-85.jpg", Description: "Reebok Club C 85, thiết kế tối giản, đế cao su bền.", Rating: 4.5, SoldAmount: 650, LikedAmount: 550, CategoryId: 3, ShopId: 7},
	{Id: 18, Name: "Fila Ray Tracer", ImageUrl: "https://fila.com/ray-tracer.jpg", Description: "Fila Ray Tracer, thiết kế chunky, phối màu trẻ trung.", Rating: 4.7, SoldAmount: 600, LikedAmount: 500, CategoryId: 3, ShopId: 8},
	{Id: 19, Name: "Balenciaga Speed Trainer", ImageUrl: "https://balenciaga.com/speed-trainer.jpg", Description: "Balenciaga Speed Trainer, thiết kế sock-fit, thời trang cao cấp.", Rating: 4.8, SoldAmount: 550, LikedAmount: 480, CategoryId: 3, ShopId: 9},
	{Id: 20, Name: "Gucci Rhyton Sneaker", ImageUrl: "https://gucci.com/rhyton-sneaker.jpg", Description: "Gucci Rhyton Sneaker, thiết kế chunky, logo nổi bật.", Rating: 4.9, SoldAmount: 500, LikedAmount: 430, CategoryId: 3, ShopId: 10},
	{Id: 21, Name: "Nike Blazer Mid '77", ImageUrl: "https://static.nike.com/blazer-mid-77.jpg", Description: "Nike Blazer Mid '77, cổ cao, phong cách vintage.", Rating: 4.7, SoldAmount: 480, LikedAmount: 400, CategoryId: 3, ShopId: 1},
	{Id: 22, Name: "Adidas NMD_R1", ImageUrl: "https://assets.adidas.com/nmd-r1.jpg", Description: "Adidas NMD_R1, đệm Boost, thiết kế hiện đại.", Rating: 4.8, SoldAmount: 460, LikedAmount: 390, CategoryId: 3, ShopId: 2},
	{Id: 23, Name: "Converse Jack Purcell", ImageUrl: "https://converse.com/jack-purcell.jpg", Description: "Converse Jack Purcell, mũi giày cười đặc trưng, phong cách cổ điển.", Rating: 4.6, SoldAmount: 440, LikedAmount: 380, CategoryId: 3, ShopId: 3},
	{Id: 24, Name: "Vans Authentic", ImageUrl: "https://vans.com/authentic.jpg", Description: "Vans Authentic, thiết kế tối giản, đế cao su bền.", Rating: 4.5, SoldAmount: 420, LikedAmount: 370, CategoryId: 3, ShopId: 4},
	{Id: 25, Name: "Puma Cali", ImageUrl: "https://puma.com/cali.jpg", Description: "Puma Cali, thiết kế trẻ trung, đế cao su dày.", Rating: 4.7, SoldAmount: 400, LikedAmount: 360, CategoryId: 3, ShopId: 5},
	{Id: 26, Name: "New Balance 327", ImageUrl: "https://newbalance.com/327.jpg", Description: "New Balance 327, thiết kế retro, phối màu nổi bật.", Rating: 4.6, SoldAmount: 380, LikedAmount: 340, CategoryId: 3, ShopId: 6},
	{Id: 27, Name: "Reebok Zig Kinetica", ImageUrl: "https://reebok.com/zig-kinetica.jpg", Description: "Reebok Zig Kinetica, đệm Zig Energy, thiết kế hiện đại.", Rating: 4.5, SoldAmount: 360, LikedAmount: 320, CategoryId: 3, ShopId: 7},
	{Id: 28, Name: "Fila Mindblower", ImageUrl: "https://fila.com/mindblower.jpg", Description: "Fila Mindblower, thiết kế logo lớn, phối màu cá tính.", Rating: 4.7, SoldAmount: 340, LikedAmount: 300, CategoryId: 3, ShopId: 8},
	{Id: 29, Name: "Balenciaga Track", ImageUrl: "https://balenciaga.com/track.jpg", Description: "Balenciaga Track, thiết kế đa lớp, thời trang cao cấp.", Rating: 4.8, SoldAmount: 320, LikedAmount: 280, CategoryId: 3, ShopId: 9},
	{Id: 30, Name: "Gucci Screener Sneaker", ImageUrl: "https://gucci.com/screener-sneaker.jpg", Description: "Gucci Screener Sneaker, phong cách vintage, logo sọc đặc trưng.", Rating: 4.9, SoldAmount: 300, LikedAmount: 260, CategoryId: 3, ShopId: 10},
}

// Mock data for ProductVariants (biến thể sản phẩm) cho Shoes
var MockShoesVariants = []model.ProductVariant{
	// SKU cho từng sản phẩm Shoes, mỗi sản phẩm 3-5 biến thể (size, màu)
	// SizeId: 8-16 (36-44), ColorId: random 1-10
	{Sku: 3001, Amount: 50, Price: 2499000, ProductId: 1, SizeId: 8, ColorId: 1},   // Nike Air Force 1, 36, Black
	{Sku: 3002, Amount: 40, Price: 2499000, ProductId: 1, SizeId: 10, ColorId: 2},  // Nike Air Force 1, 38, White
	{Sku: 3003, Amount: 30, Price: 2499000, ProductId: 1, SizeId: 12, ColorId: 5},  // Nike Air Force 1, 40, Blue
	{Sku: 3004, Amount: 25, Price: 2499000, ProductId: 1, SizeId: 14, ColorId: 3},  // Nike Air Force 1, 42, Red
	{Sku: 3005, Amount: 45, Price: 2999000, ProductId: 2, SizeId: 9, ColorId: 2},   // Adidas Ultraboost 21, 37, White
	{Sku: 3006, Amount: 35, Price: 2999000, ProductId: 2, SizeId: 11, ColorId: 4},  // Adidas Ultraboost 21, 39, Lime
	{Sku: 3007, Amount: 28, Price: 2999000, ProductId: 2, SizeId: 13, ColorId: 6},  // Adidas Ultraboost 21, 41, Yellow
	{Sku: 3008, Amount: 38, Price: 1599000, ProductId: 3, SizeId: 8, ColorId: 1},   // Converse Chuck Taylor, 36, Black
	{Sku: 3009, Amount: 32, Price: 1599000, ProductId: 3, SizeId: 10, ColorId: 7},  // Converse Chuck Taylor, 38, Cyan
	{Sku: 3010, Amount: 24, Price: 1599000, ProductId: 3, SizeId: 12, ColorId: 8},  // Converse Chuck Taylor, 40, Magenta
	{Sku: 3011, Amount: 30, Price: 1499000, ProductId: 4, SizeId: 9, ColorId: 2},   // Vans Old Skool, 37, White
	{Sku: 3012, Amount: 22, Price: 1499000, ProductId: 4, SizeId: 11, ColorId: 3},  // Vans Old Skool, 39, Red
	{Sku: 3013, Amount: 18, Price: 1499000, ProductId: 4, SizeId: 13, ColorId: 5},  // Vans Old Skool, 41, Blue
	{Sku: 3014, Amount: 28, Price: 1399000, ProductId: 5, SizeId: 8, ColorId: 1},   // Puma Suede Classic, 36, Black
	{Sku: 3015, Amount: 20, Price: 1399000, ProductId: 5, SizeId: 10, ColorId: 9},  // Puma Suede Classic, 38, Gray
	{Sku: 3016, Amount: 16, Price: 1399000, ProductId: 5, SizeId: 12, ColorId: 4},  // Puma Suede Classic, 40, Lime
	{Sku: 3017, Amount: 26, Price: 1899000, ProductId: 6, SizeId: 9, ColorId: 2},   // New Balance 574, 37, White
	{Sku: 3018, Amount: 18, Price: 1899000, ProductId: 6, SizeId: 11, ColorId: 6},  // New Balance 574, 39, Yellow
	{Sku: 3019, Amount: 12, Price: 1899000, ProductId: 6, SizeId: 13, ColorId: 7},  // New Balance 574, 41, Cyan
	{Sku: 3020, Amount: 24, Price: 1299000, ProductId: 7, SizeId: 8, ColorId: 1},   // Reebok Classic Leather, 36, Black
	{Sku: 3021, Amount: 16, Price: 1299000, ProductId: 7, SizeId: 10, ColorId: 2},  // Reebok Classic Leather, 38, White
	{Sku: 3022, Amount: 10, Price: 1299000, ProductId: 7, SizeId: 12, ColorId: 3},  // Reebok Classic Leather, 40, Red
	{Sku: 3023, Amount: 22, Price: 1799000, ProductId: 8, SizeId: 9, ColorId: 4},   // Fila Disruptor II, 37, Lime
	{Sku: 3024, Amount: 14, Price: 1799000, ProductId: 8, SizeId: 11, ColorId: 5},  // Fila Disruptor II, 39, Blue
	{Sku: 3025, Amount: 8, Price: 1799000, ProductId: 8, SizeId: 13, ColorId: 6},   // Fila Disruptor II, 41, Yellow
	{Sku: 3026, Amount: 20, Price: 4999000, ProductId: 9, SizeId: 8, ColorId: 1},   // Balenciaga Triple S, 36, Black
	{Sku: 3027, Amount: 12, Price: 4999000, ProductId: 9, SizeId: 10, ColorId: 2},  // Balenciaga Triple S, 38, White
	{Sku: 3028, Amount: 6, Price: 4999000, ProductId: 9, SizeId: 12, ColorId: 3},   // Balenciaga Triple S, 40, Red
	{Sku: 3029, Amount: 18, Price: 5999000, ProductId: 10, SizeId: 9, ColorId: 4},  // Gucci Ace Sneaker, 37, Lime
	{Sku: 3030, Amount: 10, Price: 5999000, ProductId: 10, SizeId: 11, ColorId: 5}, // Gucci Ace Sneaker, 39, Blue
	{Sku: 3031, Amount: 5, Price: 5999000, ProductId: 10, SizeId: 13, ColorId: 6},  // Gucci Ace Sneaker, 41, Yellow
	{Sku: 3032, Amount: 16, Price: 3999000, ProductId: 11, SizeId: 8, ColorId: 1},  // Nike Air Max 97, 36, Black
	{Sku: 3033, Amount: 8, Price: 3999000, ProductId: 11, SizeId: 10, ColorId: 2},  // Nike Air Max 97, 38, White
	{Sku: 3034, Amount: 4, Price: 3999000, ProductId: 11, SizeId: 12, ColorId: 3},  // Nike Air Max 97, 40, Red
	{Sku: 3035, Amount: 14, Price: 1999000, ProductId: 12, SizeId: 9, ColorId: 4},  // Adidas Stan Smith, 37, Lime
	{Sku: 3036, Amount: 6, Price: 1999000, ProductId: 12, SizeId: 11, ColorId: 5},  // Adidas Stan Smith, 39, Blue
	{Sku: 3037, Amount: 3, Price: 1999000, ProductId: 12, SizeId: 13, ColorId: 6},  // Adidas Stan Smith, 41, Yellow
	{Sku: 3038, Amount: 12, Price: 1599000, ProductId: 13, SizeId: 8, ColorId: 1},  // Converse One Star, 36, Black
	{Sku: 3039, Amount: 6, Price: 1599000, ProductId: 13, SizeId: 10, ColorId: 2},  // Converse One Star, 38, White
	{Sku: 3040, Amount: 3, Price: 1599000, ProductId: 13, SizeId: 12, ColorId: 3},  // Converse One Star, 40, Red
	{Sku: 3041, Amount: 10, Price: 1499000, ProductId: 14, SizeId: 9, ColorId: 4},  // Vans Sk8-Hi, 37, Lime
	{Sku: 3042, Amount: 5, Price: 1499000, ProductId: 14, SizeId: 11, ColorId: 5},  // Vans Sk8-Hi, 39, Blue
	{Sku: 3043, Amount: 2, Price: 1499000, ProductId: 14, SizeId: 13, ColorId: 6},  // Vans Sk8-Hi, 41, Yellow
	{Sku: 3044, Amount: 8, Price: 1799000, ProductId: 15, SizeId: 8, ColorId: 1},   // Puma RS-X, 36, Black
	{Sku: 3045, Amount: 4, Price: 1799000, ProductId: 15, SizeId: 10, ColorId: 2},  // Puma RS-X, 38, White
	{Sku: 3046, Amount: 2, Price: 1799000, ProductId: 15, SizeId: 12, ColorId: 3},  // Puma RS-X, 40, Red
	{Sku: 3047, Amount: 6, Price: 1899000, ProductId: 16, SizeId: 9, ColorId: 4},   // New Balance 997H, 37, Lime
	{Sku: 3048, Amount: 3, Price: 1899000, ProductId: 16, SizeId: 11, ColorId: 5},  // New Balance 997H, 39, Blue
	{Sku: 3049, Amount: 1, Price: 1899000, ProductId: 16, SizeId: 13, ColorId: 6},  // New Balance 997H, 41, Yellow
	{Sku: 3050, Amount: 5, Price: 1299000, ProductId: 17, SizeId: 8, ColorId: 1},   // Reebok Club C 85, 36, Black
	{Sku: 3051, Amount: 2, Price: 1299000, ProductId: 17, SizeId: 10, ColorId: 2},  // Reebok Club C 85, 38, White
	{Sku: 3052, Amount: 1, Price: 1299000, ProductId: 17, SizeId: 12, ColorId: 3},  // Reebok Club C 85, 40, Red
	{Sku: 3053, Amount: 4, Price: 1799000, ProductId: 18, SizeId: 9, ColorId: 4},   // Fila Ray Tracer, 37, Lime
	{Sku: 3054, Amount: 2, Price: 1799000, ProductId: 18, SizeId: 11, ColorId: 5},  // Fila Ray Tracer, 39, Blue
	{Sku: 3055, Amount: 1, Price: 1799000, ProductId: 18, SizeId: 13, ColorId: 6},  // Fila Ray Tracer, 41, Yellow
	{Sku: 3056, Amount: 3, Price: 4999000, ProductId: 19, SizeId: 8, ColorId: 1},   // Balenciaga Speed Trainer, 36, Black
	{Sku: 3057, Amount: 1, Price: 4999000, ProductId: 19, SizeId: 10, ColorId: 2},  // Balenciaga Speed Trainer, 38, White
	{Sku: 3058, Amount: 1, Price: 4999000, ProductId: 19, SizeId: 12, ColorId: 3},  // Balenciaga Speed Trainer, 40, Red
	{Sku: 3059, Amount: 2, Price: 5999000, ProductId: 20, SizeId: 9, ColorId: 4},   // Gucci Rhyton Sneaker, 37, Lime
	{Sku: 3060, Amount: 1, Price: 5999000, ProductId: 20, SizeId: 11, ColorId: 5},  // Gucci Rhyton Sneaker, 39, Blue
	{Sku: 3061, Amount: 1, Price: 5999000, ProductId: 20, SizeId: 13, ColorId: 6},  // Gucci Rhyton Sneaker, 41, Yellow
	{Sku: 3062, Amount: 2, Price: 2999000, ProductId: 21, SizeId: 8, ColorId: 1},   // Nike Blazer Mid '77, 36, Black
	{Sku: 3063, Amount: 1, Price: 2999000, ProductId: 21, SizeId: 10, ColorId: 2},  // Nike Blazer Mid '77, 38, White
	{Sku: 3064, Amount: 1, Price: 2999000, ProductId: 21, SizeId: 12, ColorId: 3},  // Nike Blazer Mid '77, 40, Red
	{Sku: 3065, Amount: 2, Price: 2999000, ProductId: 22, SizeId: 9, ColorId: 4},   // Adidas NMD_R1, 37, Lime
	{Sku: 3066, Amount: 1, Price: 2999000, ProductId: 22, SizeId: 11, ColorId: 5},  // Adidas NMD_R1, 39, Blue
	{Sku: 3067, Amount: 1, Price: 2999000, ProductId: 22, SizeId: 13, ColorId: 6},  // Adidas NMD_R1, 41, Yellow
	{Sku: 3068, Amount: 2, Price: 1599000, ProductId: 23, SizeId: 8, ColorId: 1},   // Converse Jack Purcell, 36, Black
	{Sku: 3069, Amount: 1, Price: 1599000, ProductId: 23, SizeId: 10, ColorId: 2},  // Converse Jack Purcell, 38, White
	{Sku: 3070, Amount: 1, Price: 1599000, ProductId: 23, SizeId: 12, ColorId: 3},  // Converse Jack Purcell, 40, Red
	{Sku: 3071, Amount: 2, Price: 1499000, ProductId: 24, SizeId: 9, ColorId: 4},   // Vans Authentic, 37, Lime
	{Sku: 3072, Amount: 1, Price: 1499000, ProductId: 24, SizeId: 11, ColorId: 5},  // Vans Authentic, 39, Blue
	{Sku: 3073, Amount: 1, Price: 1499000, ProductId: 24, SizeId: 13, ColorId: 6},  // Vans Authentic, 41, Yellow
	{Sku: 3074, Amount: 2, Price: 1799000, ProductId: 25, SizeId: 8, ColorId: 1},   // Puma Cali, 36, Black
	{Sku: 3075, Amount: 1, Price: 1799000, ProductId: 25, SizeId: 10, ColorId: 2},  // Puma Cali, 38, White
	{Sku: 3076, Amount: 1, Price: 1799000, ProductId: 25, SizeId: 12, ColorId: 3},  // Puma Cali, 40, Red
	{Sku: 3077, Amount: 2, Price: 1899000, ProductId: 26, SizeId: 9, ColorId: 4},   // New Balance 327, 37, Lime
	{Sku: 3078, Amount: 1, Price: 1899000, ProductId: 26, SizeId: 11, ColorId: 5},  // New Balance 327, 39, Blue
	{Sku: 3079, Amount: 1, Price: 1899000, ProductId: 26, SizeId: 13, ColorId: 6},  // New Balance 327, 41, Yellow
	{Sku: 3080, Amount: 2, Price: 1299000, ProductId: 27, SizeId: 8, ColorId: 1},   // Reebok Zig Kinetica, 36, Black
	{Sku: 3081, Amount: 1, Price: 1299000, ProductId: 27, SizeId: 10, ColorId: 2},  // Reebok Zig Kinetica, 38, White
	{Sku: 3082, Amount: 1, Price: 1299000, ProductId: 27, SizeId: 12, ColorId: 3},  // Reebok Zig Kinetica, 40, Red
	{Sku: 3083, Amount: 2, Price: 1799000, ProductId: 28, SizeId: 9, ColorId: 4},   // Fila Mindblower, 37, Lime
	{Sku: 3084, Amount: 1, Price: 1799000, ProductId: 28, SizeId: 11, ColorId: 5},  // Fila Mindblower, 39, Blue
	{Sku: 3085, Amount: 1, Price: 1799000, ProductId: 28, SizeId: 13, ColorId: 6},  // Fila Mindblower, 41, Yellow
	{Sku: 3086, Amount: 2, Price: 4999000, ProductId: 29, SizeId: 8, ColorId: 1},   // Balenciaga Track, 36, Black
	{Sku: 3087, Amount: 1, Price: 4999000, ProductId: 29, SizeId: 10, ColorId: 2},  // Balenciaga Track, 38, White
	{Sku: 3088, Amount: 1, Price: 4999000, ProductId: 29, SizeId: 12, ColorId: 3},  // Balenciaga Track, 40, Red
	{Sku: 3089, Amount: 2, Price: 5999000, ProductId: 30, SizeId: 9, ColorId: 4},   // Gucci Screener Sneaker, 37, Lime
	{Sku: 3090, Amount: 1, Price: 5999000, ProductId: 30, SizeId: 11, ColorId: 5},  // Gucci Screener Sneaker, 39, Blue
	{Sku: 3091, Amount: 1, Price: 5999000, ProductId: 30, SizeId: 13, ColorId: 6},  // Gucci Screener Sneaker, 41, Yellow
}
