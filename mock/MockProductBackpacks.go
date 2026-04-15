package mock

import (
	"WeFashionServer/infrastructure/model"
)

// Mock data for Backpacks category (ba lô thời trang, balo laptop, balo du lịch, ...)
var MockBackpacksProducts = []model.Product{
	// 20 sản phẩm balo thực tế, đa dạng thương hiệu, hình ảnh, mô tả, rating, sold, liked, ShopId khác nhau
	// CategoryId = 5 (Backpacks)
	{Id: 1, Name: "Herschel Little America", ImageUrl: "https://herschel.com/little-america.jpg", Description: "Balo Herschel Little America, phong cách cổ điển, ngăn laptop tiện dụng.", Rating: 4.8, SoldAmount: 1200, LikedAmount: 950, CategoryId: 5, ShopId: 1},
	{Id: 2, Name: "The North Face Borealis", ImageUrl: "https://thenorthface.com/borealis.jpg", Description: "The North Face Borealis, balo du lịch, đệm lưng thoáng khí.", Rating: 4.9, SoldAmount: 1100, LikedAmount: 900, CategoryId: 5, ShopId: 2},
	{Id: 3, Name: "Adidas Classic Backpack", ImageUrl: "https://adidas.com/classic-backpack.jpg", Description: "Adidas Classic Backpack, thiết kế thể thao, chất liệu bền bỉ.", Rating: 4.7, SoldAmount: 1050, LikedAmount: 870, CategoryId: 5, ShopId: 3},
	{Id: 4, Name: "Nike Heritage Backpack", ImageUrl: "https://nike.com/heritage-backpack.jpg", Description: "Nike Heritage Backpack, balo thời trang, nhiều màu sắc.", Rating: 4.7, SoldAmount: 980, LikedAmount: 800, CategoryId: 5, ShopId: 4},
	{Id: 5, Name: "Fjällräven Kånken", ImageUrl: "https://fjallraven.com/kanken.jpg", Description: "Fjällräven Kånken, balo biểu tượng Thụy Điển, thiết kế tối giản.", Rating: 4.8, SoldAmount: 950, LikedAmount: 780, CategoryId: 5, ShopId: 5},
	{Id: 6, Name: "Tigernu Anti-theft Laptop Backpack", ImageUrl: "https://tigernu.com/anti-theft.jpg", Description: "Tigernu balo chống trộm, ngăn laptop riêng biệt, chống nước.", Rating: 4.7, SoldAmount: 900, LikedAmount: 750, CategoryId: 5, ShopId: 6},
	{Id: 7, Name: "MCM Stark Backpack", ImageUrl: "https://mcmworldwide.com/stark-backpack.jpg", Description: "MCM Stark Backpack, họa tiết logo nổi bật, chất liệu cao cấp.", Rating: 4.9, SoldAmount: 870, LikedAmount: 720, CategoryId: 5, ShopId: 7},
	{Id: 8, Name: "Samsonite Red Backpack", ImageUrl: "https://samsonite.com/red-backpack.jpg", Description: "Samsonite Red Backpack, balo công sở, thiết kế hiện đại.", Rating: 4.8, SoldAmount: 850, LikedAmount: 700, CategoryId: 5, ShopId: 8},
	{Id: 9, Name: "Anello Regular Backpack", ImageUrl: "https://anello.com/regular-backpack.jpg", Description: "Anello Regular Backpack, balo Nhật Bản, miệng rộng tiện lợi.", Rating: 4.7, SoldAmount: 800, LikedAmount: 670, CategoryId: 5, ShopId: 9},
	{Id: 10, Name: "Jansport Superbreak", ImageUrl: "https://jansport.com/superbreak.jpg", Description: "Jansport Superbreak, balo học sinh, nhiều màu sắc trẻ trung.", Rating: 4.6, SoldAmount: 780, LikedAmount: 650, CategoryId: 5, ShopId: 10},
	{Id: 11, Name: "Balo SimpleCarry K3", ImageUrl: "https://simplecarry.com/k3.jpg", Description: "SimpleCarry K3, balo laptop, thiết kế tối giản, nhẹ.", Rating: 4.7, SoldAmount: 760, LikedAmount: 630, CategoryId: 5, ShopId: 11},
	{Id: 12, Name: "Balo Coolbell CB8008", ImageUrl: "https://coolbell.com/cb8008.jpg", Description: "Coolbell CB8008, balo đa năng, nhiều ngăn tiện dụng.", Rating: 4.8, SoldAmount: 740, LikedAmount: 610, CategoryId: 5, ShopId: 12},
	{Id: 13, Name: "Balo Xiaomi Mi City", ImageUrl: "https://xiaomi.com/mi-city.jpg", Description: "Xiaomi Mi City, balo công nghệ, thiết kế hiện đại.", Rating: 4.9, SoldAmount: 720, LikedAmount: 600, CategoryId: 5, ShopId: 13},
	{Id: 14, Name: "Balo Targus Groove X", ImageUrl: "https://targus.com/groove-x.jpg", Description: "Targus Groove X, balo laptop, đệm lưng thoáng khí.", Rating: 4.8, SoldAmount: 700, LikedAmount: 590, CategoryId: 5, ShopId: 14},
	{Id: 15, Name: "Balo Tomtoc Travel", ImageUrl: "https://tomtoc.com/travel.jpg", Description: "Tomtoc Travel, balo du lịch, ngăn chống sốc laptop.", Rating: 4.8, SoldAmount: 680, LikedAmount: 580, CategoryId: 5, ShopId: 15},
	{Id: 16, Name: "Balo Mark Ryden MR9208", ImageUrl: "https://markryden.vn/mr9208.jpg", Description: "Mark Ryden MR9208, balo chống nước, thiết kế hiện đại.", Rating: 4.7, SoldAmount: 660, LikedAmount: 570, CategoryId: 5, ShopId: 16},
	{Id: 17, Name: "Balo Haras HR262", ImageUrl: "https://haras.com.vn/hr262.jpg", Description: "Haras HR262, balo học sinh, nhiều màu sắc trẻ trung.", Rating: 4.6, SoldAmount: 640, LikedAmount: 560, CategoryId: 5, ShopId: 17},
	{Id: 18, Name: "Balo BLH 168", ImageUrl: "https://blh.vn/168.jpg", Description: "BLH 168, balo thời trang, giá rẻ, nhiều màu sắc.", Rating: 4.7, SoldAmount: 620, LikedAmount: 550, CategoryId: 5, ShopId: 18},
	{Id: 19, Name: "Balo Mikkor The Ralph", ImageUrl: "https://mikkor.com/the-ralph.jpg", Description: "Mikkor The Ralph, balo công sở, chất liệu cao cấp.", Rating: 4.7, SoldAmount: 600, LikedAmount: 540, CategoryId: 5, ShopId: 19},
	{Id: 20, Name: "Balo Vascara City", ImageUrl: "https://vascara.com/city.jpg", Description: "Vascara City, balo nữ tính, thiết kế hiện đại.", Rating: 4.8, SoldAmount: 580, LikedAmount: 530, CategoryId: 5, ShopId: 20},
}

// Mock data for ProductVariants (biến thể sản phẩm) cho Backpacks
var MockBackpacksVariants = []model.ProductVariant{
	// SKU cho từng sản phẩm Backpacks, mỗi sản phẩm 2-3 biến thể (màu khác nhau, size tự do)
	{Sku: 5001, Amount: 35, Price: 1899000, ProductId: 1, SizeId: 4, ColorId: 1},  // Herschel Little America, L, Black
	{Sku: 5002, Amount: 20, Price: 1899000, ProductId: 1, SizeId: 4, ColorId: 2},  // Herschel Little America, L, White
	{Sku: 5003, Amount: 28, Price: 2499000, ProductId: 2, SizeId: 4, ColorId: 3},  // The North Face Borealis, L, Red
	{Sku: 5004, Amount: 15, Price: 2499000, ProductId: 2, SizeId: 4, ColorId: 4},  // The North Face Borealis, L, Lime
	{Sku: 5005, Amount: 22, Price: 999000, ProductId: 3, SizeId: 3, ColorId: 5},   // Adidas Classic, M, Blue
	{Sku: 5006, Amount: 14, Price: 999000, ProductId: 3, SizeId: 3, ColorId: 6},   // Adidas Classic, M, Yellow
	{Sku: 5007, Amount: 18, Price: 1099000, ProductId: 4, SizeId: 3, ColorId: 7},  // Nike Heritage, M, Cyan
	{Sku: 5008, Amount: 12, Price: 1099000, ProductId: 4, SizeId: 3, ColorId: 8},  // Nike Heritage, M, Magenta
	{Sku: 5009, Amount: 16, Price: 1999000, ProductId: 5, SizeId: 4, ColorId: 9},  // Fjallraven Kanken, L, Silver
	{Sku: 5010, Amount: 10, Price: 1999000, ProductId: 5, SizeId: 4, ColorId: 10}, // Fjallraven Kanken, L, Gray
	{Sku: 5011, Amount: 14, Price: 899000, ProductId: 6, SizeId: 3, ColorId: 1},   // Tigernu Anti-theft, M, Black
	{Sku: 5012, Amount: 8, Price: 899000, ProductId: 6, SizeId: 3, ColorId: 2},    // Tigernu Anti-theft, M, White
	{Sku: 5013, Amount: 13, Price: 7990000, ProductId: 7, SizeId: 4, ColorId: 3},  // MCM Stark, L, Red
	{Sku: 5014, Amount: 7, Price: 7990000, ProductId: 7, SizeId: 4, ColorId: 4},   // MCM Stark, L, Lime
	{Sku: 5015, Amount: 15, Price: 2990000, ProductId: 8, SizeId: 3, ColorId: 5},  // Samsonite Red, M, Blue
	{Sku: 5016, Amount: 10, Price: 2990000, ProductId: 8, SizeId: 3, ColorId: 6},  // Samsonite Red, M, Yellow
	{Sku: 5017, Amount: 12, Price: 1599000, ProductId: 9, SizeId: 2, ColorId: 7},  // Anello Regular, S, Cyan
	{Sku: 5018, Amount: 8, Price: 1599000, ProductId: 9, SizeId: 2, ColorId: 8},   // Anello Regular, S, Magenta
	{Sku: 5019, Amount: 14, Price: 899000, ProductId: 10, SizeId: 3, ColorId: 9},  // Jansport Superbreak, M, Silver
	{Sku: 5020, Amount: 9, Price: 899000, ProductId: 10, SizeId: 3, ColorId: 10},  // Jansport Superbreak, M, Gray
	{Sku: 5021, Amount: 13, Price: 799000, ProductId: 11, SizeId: 2, ColorId: 1},  // SimpleCarry K3, S, Black
	{Sku: 5022, Amount: 7, Price: 799000, ProductId: 11, SizeId: 2, ColorId: 2},   // SimpleCarry K3, S, White
	{Sku: 5023, Amount: 12, Price: 1099000, ProductId: 12, SizeId: 3, ColorId: 3}, // Coolbell CB8008, M, Red
	{Sku: 5024, Amount: 8, Price: 1099000, ProductId: 12, SizeId: 3, ColorId: 4},  // Coolbell CB8008, M, Lime
	{Sku: 5025, Amount: 11, Price: 1299000, ProductId: 13, SizeId: 4, ColorId: 5}, // Xiaomi Mi City, L, Blue
	{Sku: 5026, Amount: 7, Price: 1299000, ProductId: 13, SizeId: 4, ColorId: 6},  // Xiaomi Mi City, L, Yellow
	{Sku: 5027, Amount: 6, Price: 1599000, ProductId: 14, SizeId: 4, ColorId: 1},  // Targus Groove X, L, Black
	{Sku: 5028, Amount: 4, Price: 1599000, ProductId: 14, SizeId: 4, ColorId: 2},  // Targus Groove X, L, White
	{Sku: 5029, Amount: 10, Price: 1999000, ProductId: 15, SizeId: 3, ColorId: 3}, // Tomtoc Travel, M, Red
	{Sku: 5030, Amount: 6, Price: 1999000, ProductId: 15, SizeId: 3, ColorId: 4},  // Tomtoc Travel, M, Lime
	{Sku: 5031, Amount: 9, Price: 899000, ProductId: 16, SizeId: 4, ColorId: 5},   // Mark Ryden MR9208, L, Blue
	{Sku: 5032, Amount: 5, Price: 899000, ProductId: 16, SizeId: 4, ColorId: 6},   // Mark Ryden MR9208, L, Yellow
	{Sku: 5033, Amount: 8, Price: 599000, ProductId: 17, SizeId: 3, ColorId: 7},   // Haras HR262, M, Cyan
	{Sku: 5034, Amount: 5, Price: 599000, ProductId: 17, SizeId: 3, ColorId: 8},   // Haras HR262, M, Magenta
	{Sku: 5035, Amount: 7, Price: 499000, ProductId: 18, SizeId: 4, ColorId: 9},   // BLH 168, L, Silver
	{Sku: 5036, Amount: 4, Price: 499000, ProductId: 18, SizeId: 4, ColorId: 10},  // BLH 168, L, Gray
	{Sku: 5037, Amount: 10, Price: 1599000, ProductId: 19, SizeId: 3, ColorId: 1}, // Mikkor The Ralph, M, Black
	{Sku: 5038, Amount: 6, Price: 1599000, ProductId: 19, SizeId: 3, ColorId: 2},  // Mikkor The Ralph, M, White
	{Sku: 5039, Amount: 12, Price: 1299000, ProductId: 20, SizeId: 3, ColorId: 3}, // Vascara City, M, Red
	{Sku: 5040, Amount: 8, Price: 1299000, ProductId: 20, SizeId: 3, ColorId: 4},  // Vascara City, M, Lime
}
