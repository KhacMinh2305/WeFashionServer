package mock

import (
	"WeFashionServer/infrastructure/model"
)

// Mock data for Handbags category (túi xách, túi đeo chéo, clutch, tote, ...)
var MockHandbagsProducts = []model.Product{
	// 20 sản phẩm túi xách thực tế, đa dạng thương hiệu, hình ảnh, mô tả, rating, sold, liked, ShopId khác nhau
	// CategoryId = 4 (Handbags)
	{Id: 1, Name: "Charles & Keith Quilted Shoulder Bag", ImageUrl: "https://charleskeith.com/quilted-shoulder-bag.jpg", Description: "Túi đeo vai Charles & Keith, thiết kế chần bông sang trọng.", Rating: 4.8, SoldAmount: 1200, LikedAmount: 950, CategoryId: 4, ShopId: 1},
	{Id: 2, Name: "Gucci GG Marmont Matelassé", ImageUrl: "https://gucci.com/gg-marmont.jpg", Description: "Gucci GG Marmont, chất liệu da cao cấp, logo GG nổi bật.", Rating: 4.9, SoldAmount: 1100, LikedAmount: 900, CategoryId: 4, ShopId: 2},
	{Id: 3, Name: "Louis Vuitton Neverfull MM", ImageUrl: "https://louisvuitton.com/neverfull-mm.jpg", Description: "Louis Vuitton Neverfull MM, túi tote đa năng, họa tiết monogram.", Rating: 4.9, SoldAmount: 1050, LikedAmount: 870, CategoryId: 4, ShopId: 3},
	{Id: 4, Name: "Michael Kors Jet Set Travel", ImageUrl: "https://michaelkors.com/jet-set-travel.jpg", Description: "Michael Kors Jet Set Travel, túi tote tiện dụng, chất liệu saffiano.", Rating: 4.7, SoldAmount: 980, LikedAmount: 800, CategoryId: 4, ShopId: 4},
	{Id: 5, Name: "Coach Tabby Shoulder Bag", ImageUrl: "https://coach.com/tabby-shoulder-bag.jpg", Description: "Coach Tabby Shoulder Bag, thiết kế trẻ trung, logo C đặc trưng.", Rating: 4.8, SoldAmount: 950, LikedAmount: 780, CategoryId: 4, ShopId: 5},
	{Id: 6, Name: "Prada Re-Edition 2005", ImageUrl: "https://prada.com/re-edition-2005.jpg", Description: "Prada Re-Edition 2005, túi nylon thời thượng, dây đeo đa năng.", Rating: 4.7, SoldAmount: 900, LikedAmount: 750, CategoryId: 4, ShopId: 6},
	{Id: 7, Name: "Dior Book Tote", ImageUrl: "https://dior.com/book-tote.jpg", Description: "Dior Book Tote, túi tote họa tiết thêu nổi bật.", Rating: 4.9, SoldAmount: 870, LikedAmount: 720, CategoryId: 4, ShopId: 7},
	{Id: 8, Name: "YSL Kate Monogram", ImageUrl: "https://ysl.com/kate-monogram.jpg", Description: "YSL Kate Monogram, túi đeo chéo nhỏ gọn, logo YSL kim loại.", Rating: 4.8, SoldAmount: 850, LikedAmount: 700, CategoryId: 4, ShopId: 8},
	{Id: 9, Name: "Furla Metropolis Mini", ImageUrl: "https://furla.com/metropolis-mini.jpg", Description: "Furla Metropolis Mini, túi nhỏ xinh, khóa kim loại đặc trưng.", Rating: 4.7, SoldAmount: 800, LikedAmount: 670, CategoryId: 4, ShopId: 9},
	{Id: 10, Name: "Tory Burch Perry Triple-Compartment Tote", ImageUrl: "https://toryburch.com/perry-tote.jpg", Description: "Tory Burch Perry Tote, nhiều ngăn tiện dụng, chất liệu da mềm.", Rating: 4.6, SoldAmount: 780, LikedAmount: 650, CategoryId: 4, ShopId: 10},
	{Id: 11, Name: "Marc Jacobs Snapshot", ImageUrl: "https://marcjacobs.com/snapshot.jpg", Description: "Marc Jacobs Snapshot, túi đeo chéo nhỏ, phối màu cá tính.", Rating: 4.7, SoldAmount: 760, LikedAmount: 630, CategoryId: 4, ShopId: 11},
	{Id: 12, Name: "Longchamp Le Pliage", ImageUrl: "https://longchamp.com/le-pliage.jpg", Description: "Longchamp Le Pliage, túi gấp tiện lợi, chất liệu nylon nhẹ.", Rating: 4.8, SoldAmount: 740, LikedAmount: 610, CategoryId: 4, ShopId: 12},
	{Id: 13, Name: "Celine Belt Bag", ImageUrl: "https://celine.com/belt-bag.jpg", Description: "Celine Belt Bag, thiết kế thanh lịch, dây đeo linh hoạt.", Rating: 4.9, SoldAmount: 720, LikedAmount: 600, CategoryId: 4, ShopId: 13},
	{Id: 14, Name: "Hermès Birkin 25", ImageUrl: "https://hermes.com/birkin-25.jpg", Description: "Hermès Birkin 25, biểu tượng thời trang xa xỉ, chất liệu da cao cấp.", Rating: 5.0, SoldAmount: 700, LikedAmount: 590, CategoryId: 4, ShopId: 14},
	{Id: 15, Name: "Givenchy Antigona", ImageUrl: "https://givenchy.com/antigona.jpg", Description: "Givenchy Antigona, túi xách form đứng, logo kim loại.", Rating: 4.8, SoldAmount: 680, LikedAmount: 580, CategoryId: 4, ShopId: 15},
	{Id: 16, Name: "Balenciaga Ville Top Handle", ImageUrl: "https://balenciaga.com/ville-top-handle.jpg", Description: "Balenciaga Ville Top Handle, thiết kế hiện đại, logo in nổi bật.", Rating: 4.7, SoldAmount: 660, LikedAmount: 570, CategoryId: 4, ShopId: 16},
	{Id: 17, Name: "JW Anderson Pierce Bag", ImageUrl: "https://jwanderson.com/pierce-bag.jpg", Description: "JW Anderson Pierce Bag, khóa tròn đặc trưng, thiết kế cá tính.", Rating: 4.6, SoldAmount: 640, LikedAmount: 560, CategoryId: 4, ShopId: 17},
	{Id: 18, Name: "Chanel Classic Flap Bag", ImageUrl: "https://chanel.com/classic-flap-bag.jpg", Description: "Chanel Classic Flap Bag, thiết kế chần bông, logo CC.", Rating: 5.0, SoldAmount: 620, LikedAmount: 550, CategoryId: 4, ShopId: 18},
	{Id: 19, Name: "MCM Stark Backpack", ImageUrl: "https://mcmworldwide.com/stark-backpack.jpg", Description: "MCM Stark Backpack, túi đeo lưng thời trang, họa tiết logo nổi bật.", Rating: 4.7, SoldAmount: 600, LikedAmount: 540, CategoryId: 4, ShopId: 19},
	{Id: 20, Name: "Kate Spade Margaux Satchel", ImageUrl: "https://katespade.com/margaux-satchel.jpg", Description: "Kate Spade Margaux Satchel, thiết kế nữ tính, nhiều màu sắc.", Rating: 4.8, SoldAmount: 580, LikedAmount: 530, CategoryId: 4, ShopId: 20},
}

// Mock data for ProductVariants (biến thể sản phẩm) cho Handbags
var MockHandbagsVariants = []model.ProductVariant{
	// SKU cho từng sản phẩm Handbags, mỗi sản phẩm 2-3 biến thể (màu khác nhau, size tự do)
	{Sku: 4001, Amount: 30, Price: 1899000, ProductId: 1, SizeId: 3, ColorId: 1},   // Charles & Keith, M, Black
	{Sku: 4002, Amount: 20, Price: 1899000, ProductId: 1, SizeId: 3, ColorId: 2},   // Charles & Keith, M, White
	{Sku: 4003, Amount: 15, Price: 23990000, ProductId: 2, SizeId: 3, ColorId: 3},  // Gucci GG Marmont, M, Red
	{Sku: 4004, Amount: 10, Price: 23990000, ProductId: 2, SizeId: 3, ColorId: 4},  // Gucci GG Marmont, M, Lime
	{Sku: 4005, Amount: 12, Price: 49990000, ProductId: 3, SizeId: 4, ColorId: 5},  // LV Neverfull, L, Blue
	{Sku: 4006, Amount: 8, Price: 49990000, ProductId: 3, SizeId: 4, ColorId: 6},   // LV Neverfull, L, Yellow
	{Sku: 4007, Amount: 18, Price: 7990000, ProductId: 4, SizeId: 3, ColorId: 7},   // MK Jet Set, M, Cyan
	{Sku: 4008, Amount: 12, Price: 7990000, ProductId: 4, SizeId: 4, ColorId: 8},   // MK Jet Set, L, Magenta
	{Sku: 4009, Amount: 16, Price: 5990000, ProductId: 5, SizeId: 3, ColorId: 9},   // Coach Tabby, M, Silver
	{Sku: 4010, Amount: 10, Price: 5990000, ProductId: 5, SizeId: 4, ColorId: 10},  // Coach Tabby, L, Gray
	{Sku: 4011, Amount: 14, Price: 2990000, ProductId: 6, SizeId: 3, ColorId: 1},   // Prada Re-Edition, M, Black
	{Sku: 4012, Amount: 8, Price: 2990000, ProductId: 6, SizeId: 3, ColorId: 2},    // Prada Re-Edition, M, White
	{Sku: 4013, Amount: 13, Price: 6990000, ProductId: 7, SizeId: 4, ColorId: 3},   // Dior Book Tote, L, Red
	{Sku: 4014, Amount: 7, Price: 6990000, ProductId: 7, SizeId: 4, ColorId: 4},    // Dior Book Tote, L, Lime
	{Sku: 4015, Amount: 15, Price: 15990000, ProductId: 8, SizeId: 3, ColorId: 5},  // YSL Kate, M, Blue
	{Sku: 4016, Amount: 10, Price: 15990000, ProductId: 8, SizeId: 3, ColorId: 6},  // YSL Kate, M, Yellow
	{Sku: 4017, Amount: 12, Price: 4990000, ProductId: 9, SizeId: 2, ColorId: 7},   // Furla Metropolis, S, Cyan
	{Sku: 4018, Amount: 8, Price: 4990000, ProductId: 9, SizeId: 2, ColorId: 8},    // Furla Metropolis, S, Magenta
	{Sku: 4019, Amount: 14, Price: 3990000, ProductId: 10, SizeId: 4, ColorId: 9},  // Tory Burch Perry, L, Silver
	{Sku: 4020, Amount: 9, Price: 3990000, ProductId: 10, SizeId: 4, ColorId: 10},  // Tory Burch Perry, L, Gray
	{Sku: 4021, Amount: 13, Price: 2990000, ProductId: 11, SizeId: 2, ColorId: 1},  // Marc Jacobs Snapshot, S, Black
	{Sku: 4022, Amount: 7, Price: 2990000, ProductId: 11, SizeId: 2, ColorId: 2},   // Marc Jacobs Snapshot, S, White
	{Sku: 4023, Amount: 12, Price: 3990000, ProductId: 12, SizeId: 3, ColorId: 3},  // Longchamp Le Pliage, M, Red
	{Sku: 4024, Amount: 8, Price: 3990000, ProductId: 12, SizeId: 3, ColorId: 4},   // Longchamp Le Pliage, M, Lime
	{Sku: 4025, Amount: 11, Price: 15990000, ProductId: 13, SizeId: 4, ColorId: 5}, // Celine Belt Bag, L, Blue
	{Sku: 4026, Amount: 7, Price: 15990000, ProductId: 13, SizeId: 4, ColorId: 6},  // Celine Belt Bag, L, Yellow
	{Sku: 4027, Amount: 6, Price: 299900000, ProductId: 14, SizeId: 4, ColorId: 1}, // Hermes Birkin, L, Black
	{Sku: 4028, Amount: 4, Price: 299900000, ProductId: 14, SizeId: 4, ColorId: 2}, // Hermes Birkin, L, White
	{Sku: 4029, Amount: 10, Price: 4990000, ProductId: 15, SizeId: 3, ColorId: 3},  // Givenchy Antigona, M, Red
	{Sku: 4030, Amount: 6, Price: 4990000, ProductId: 15, SizeId: 3, ColorId: 4},   // Givenchy Antigona, M, Lime
	{Sku: 4031, Amount: 9, Price: 8990000, ProductId: 16, SizeId: 4, ColorId: 5},   // Balenciaga Ville, L, Blue
	{Sku: 4032, Amount: 5, Price: 8990000, ProductId: 16, SizeId: 4, ColorId: 6},   // Balenciaga Ville, L, Yellow
	{Sku: 4033, Amount: 8, Price: 7990000, ProductId: 17, SizeId: 3, ColorId: 7},   // JW Anderson Pierce, M, Cyan
	{Sku: 4034, Amount: 5, Price: 7990000, ProductId: 17, SizeId: 3, ColorId: 8},   // JW Anderson Pierce, M, Magenta
	{Sku: 4035, Amount: 7, Price: 29990000, ProductId: 18, SizeId: 4, ColorId: 9},  // Chanel Classic Flap, L, Silver
	{Sku: 4036, Amount: 4, Price: 29990000, ProductId: 18, SizeId: 4, ColorId: 10}, // Chanel Classic Flap, L, Gray
	{Sku: 4037, Amount: 10, Price: 5990000, ProductId: 19, SizeId: 3, ColorId: 1},  // MCM Stark Backpack, M, Black
	{Sku: 4038, Amount: 6, Price: 5990000, ProductId: 19, SizeId: 3, ColorId: 2},   // MCM Stark Backpack, M, White
	{Sku: 4039, Amount: 12, Price: 4990000, ProductId: 20, SizeId: 3, ColorId: 3},  // Kate Spade Margaux, M, Red
	{Sku: 4040, Amount: 8, Price: 4990000, ProductId: 20, SizeId: 3, ColorId: 4},   // Kate Spade Margaux, M, Lime
}
