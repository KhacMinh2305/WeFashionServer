package mock

import (
	"WeFashionServer/infrastructure/model"
)

// Mock data for Underwears category (quần lót nam, nữ, boxer, bralette, ...)
var MockUnderwearsProducts = []model.Product{
	// 20 sản phẩm đồ lót thực tế, 10 nam, 10 nữ, đa dạng thương hiệu, hình ảnh, mô tả, rating, sold, liked, ShopId khác nhau
	// CategoryId = 7 (Underwear)
	// Nam
	{Id: 1, Name: "Calvin Klein Men's Boxer Briefs", ImageUrl: "https://calvinklein.com/men-boxer.jpg", Description: "Boxer nam Calvin Klein, cotton co giãn, thoáng khí.", Rating: 4.8, SoldAmount: 1200, LikedAmount: 950, CategoryId: 7, ShopId: 1},
	{Id: 2, Name: "Uniqlo Men AIRism Trunks", ImageUrl: "https://uniqlo.com/airism-trunks.jpg", Description: "Quần lót nam Uniqlo AIRism, siêu nhẹ, khô thoáng.", Rating: 4.7, SoldAmount: 1100, LikedAmount: 900, CategoryId: 7, ShopId: 2},
	{Id: 3, Name: "Jockey Classic Briefs", ImageUrl: "https://jockey.com/classic-briefs.jpg", Description: "Jockey Classic Briefs, thiết kế ôm sát, chất liệu mềm mại.", Rating: 4.6, SoldAmount: 1050, LikedAmount: 870, CategoryId: 7, ShopId: 3},
	{Id: 4, Name: "Bonobos Boxer Briefs", ImageUrl: "https://bonobos.com/boxer-briefs.jpg", Description: "Bonobos Boxer Briefs, vải co giãn, thoải mái vận động.", Rating: 4.7, SoldAmount: 980, LikedAmount: 800, CategoryId: 7, ShopId: 4},
	{Id: 5, Name: "Lacoste Men's Trunks", ImageUrl: "https://lacoste.com/men-trunks.jpg", Description: "Lacoste Men's Trunks, logo cá sấu, chất liệu cotton.", Rating: 4.8, SoldAmount: 950, LikedAmount: 780, CategoryId: 7, ShopId: 5},
	{Id: 6, Name: "Puma Men's Boxer Briefs", ImageUrl: "https://puma.com/men-boxer.jpg", Description: "Puma Men's Boxer Briefs, thiết kế thể thao, co giãn tốt.", Rating: 4.7, SoldAmount: 900, LikedAmount: 750, CategoryId: 7, ShopId: 6},
	{Id: 7, Name: "Under Armour Tech Boxerjock", ImageUrl: "https://underarmour.com/tech-boxerjock.jpg", Description: "Under Armour Tech Boxerjock, khô nhanh, chống mùi.", Rating: 4.8, SoldAmount: 870, LikedAmount: 720, CategoryId: 7, ShopId: 7},
	{Id: 8, Name: "Tommy Hilfiger Men's Briefs", ImageUrl: "https://tommy.com/men-briefs.jpg", Description: "Tommy Hilfiger Men's Briefs, logo thắt lưng, chất liệu mềm.", Rating: 4.7, SoldAmount: 850, LikedAmount: 700, CategoryId: 7, ShopId: 8},
	{Id: 9, Name: "Bamboo Cool Men's Trunks", ImageUrl: "https://bamboocool.com/men-trunks.jpg", Description: "Bamboo Cool Men's Trunks, vải tre tự nhiên, kháng khuẩn.", Rating: 4.6, SoldAmount: 800, LikedAmount: 670, CategoryId: 7, ShopId: 9},
	{Id: 10, Name: "Coolmate Men Everyday Briefs", ImageUrl: "https://coolmate.me/men-briefs.jpg", Description: "Coolmate Men Everyday Briefs, giá tốt, chất liệu co giãn.", Rating: 4.7, SoldAmount: 780, LikedAmount: 650, CategoryId: 7, ShopId: 10},
	// Nữ
	{Id: 11, Name: "Victoria's Secret Cotton Bikini", ImageUrl: "https://victoriassecret.com/cotton-bikini.jpg", Description: "Victoria's Secret Cotton Bikini, thiết kế nữ tính, chất liệu mềm mại.", Rating: 4.9, SoldAmount: 1200, LikedAmount: 950, CategoryId: 7, ShopId: 11},
	{Id: 12, Name: "Uniqlo Women AIRism Bikini", ImageUrl: "https://uniqlo.com/women-airism-bikini.jpg", Description: "Quần lót nữ Uniqlo AIRism, thoáng khí, nhẹ nhàng.", Rating: 4.8, SoldAmount: 1100, LikedAmount: 900, CategoryId: 7, ShopId: 12},
	{Id: 13, Name: "Bon Bon Women's Briefs", ImageUrl: "https://bonbon.com/women-briefs.jpg", Description: "Bon Bon Women's Briefs, giá tốt, chất liệu co giãn.", Rating: 4.7, SoldAmount: 1050, LikedAmount: 870, CategoryId: 7, ShopId: 13},
	{Id: 14, Name: "Triumph Amourette Maxi", ImageUrl: "https://triumph.com/amourette-maxi.jpg", Description: "Triumph Amourette Maxi, thiết kế ôm sát, ren nữ tính.", Rating: 4.8, SoldAmount: 980, LikedAmount: 800, CategoryId: 7, ShopId: 14},
	{Id: 15, Name: "Sabina Seamless Panty", ImageUrl: "https://sabina.com/seamless-panty.jpg", Description: "Sabina Seamless Panty, không viền, co giãn tốt.", Rating: 4.7, SoldAmount: 950, LikedAmount: 780, CategoryId: 7, ShopId: 15},
	{Id: 16, Name: "Sunfly Women's Boxer", ImageUrl: "https://sunfly.com/women-boxer.jpg", Description: "Sunfly Women's Boxer, vải cotton, thoáng khí.", Rating: 4.8, SoldAmount: 900, LikedAmount: 750, CategoryId: 7, ShopId: 16},
	{Id: 17, Name: "Wacoal Lace Bikini", ImageUrl: "https://wacoal.com/lace-bikini.jpg", Description: "Wacoal Lace Bikini, thiết kế ren, nữ tính, thoải mái.", Rating: 4.9, SoldAmount: 870, LikedAmount: 720, CategoryId: 7, ShopId: 17},
	{Id: 18, Name: "Jockey Women's Hipster", ImageUrl: "https://jockey.com/women-hipster.jpg", Description: "Jockey Women's Hipster, thiết kế trẻ trung, chất liệu mềm.", Rating: 4.8, SoldAmount: 850, LikedAmount: 700, CategoryId: 7, ShopId: 18},
	{Id: 19, Name: "Bon Bon Women's Boxer", ImageUrl: "https://bonbon.com/women-boxer.jpg", Description: "Bon Bon Women's Boxer, giá tốt, vải co giãn.", Rating: 4.7, SoldAmount: 800, LikedAmount: 670, CategoryId: 7, ShopId: 19},
	{Id: 20, Name: "Coolmate Women Everyday Briefs", ImageUrl: "https://coolmate.me/women-briefs.jpg", Description: "Coolmate Women Everyday Briefs, giá tốt, chất liệu co giãn.", Rating: 4.8, SoldAmount: 780, LikedAmount: 650, CategoryId: 7, ShopId: 20},
}

// Mock data for ProductVariants (biến thể sản phẩm) cho Underwears
var MockUnderwearsVariants = []model.ProductVariant{
	// SKU cho từng sản phẩm Underwears, mỗi sản phẩm 2-3 biến thể (size đồ lót, màu đa dạng)
	// SizeId: 17-20 (50-60kg, 60-70kg, 70-80kg, 80-90kg), ColorId: random 1-10
	{Sku: 7001, Amount: 30, Price: 299000, ProductId: 1, SizeId: 17, ColorId: 1},  // CK Boxer, 50-60kg, Black
	{Sku: 7002, Amount: 20, Price: 299000, ProductId: 1, SizeId: 18, ColorId: 2},  // CK Boxer, 60-70kg, White
	{Sku: 7003, Amount: 25, Price: 199000, ProductId: 2, SizeId: 17, ColorId: 3},  // Uniqlo AIRism, 50-60kg, Red
	{Sku: 7004, Amount: 15, Price: 199000, ProductId: 2, SizeId: 18, ColorId: 4},  // Uniqlo AIRism, 60-70kg, Lime
	{Sku: 7005, Amount: 22, Price: 159000, ProductId: 3, SizeId: 19, ColorId: 5},  // Jockey Briefs, 70-80kg, Blue
	{Sku: 7006, Amount: 12, Price: 159000, ProductId: 3, SizeId: 20, ColorId: 6},  // Jockey Briefs, 80-90kg, Yellow
	{Sku: 7007, Amount: 18, Price: 249000, ProductId: 4, SizeId: 17, ColorId: 7},  // Bonobos Boxer, 50-60kg, Cyan
	{Sku: 7008, Amount: 10, Price: 249000, ProductId: 4, SizeId: 18, ColorId: 8},  // Bonobos Boxer, 60-70kg, Magenta
	{Sku: 7009, Amount: 16, Price: 299000, ProductId: 5, SizeId: 19, ColorId: 9},  // Lacoste Trunks, 70-80kg, Silver
	{Sku: 7010, Amount: 8, Price: 299000, ProductId: 5, SizeId: 20, ColorId: 10},  // Lacoste Trunks, 80-90kg, Gray
	{Sku: 7011, Amount: 14, Price: 199000, ProductId: 6, SizeId: 17, ColorId: 1},  // Puma Boxer, 50-60kg, Black
	{Sku: 7012, Amount: 7, Price: 199000, ProductId: 6, SizeId: 18, ColorId: 2},   // Puma Boxer, 60-70kg, White
	{Sku: 7013, Amount: 13, Price: 299000, ProductId: 7, SizeId: 19, ColorId: 3},  // UA Boxerjock, 70-80kg, Red
	{Sku: 7014, Amount: 6, Price: 299000, ProductId: 7, SizeId: 20, ColorId: 4},   // UA Boxerjock, 80-90kg, Lime
	{Sku: 7015, Amount: 15, Price: 249000, ProductId: 8, SizeId: 17, ColorId: 5},  // Tommy Briefs, 50-60kg, Blue
	{Sku: 7016, Amount: 8, Price: 249000, ProductId: 8, SizeId: 18, ColorId: 6},   // Tommy Briefs, 60-70kg, Yellow
	{Sku: 7017, Amount: 12, Price: 199000, ProductId: 9, SizeId: 19, ColorId: 7},  // Bamboo Cool, 70-80kg, Cyan
	{Sku: 7018, Amount: 6, Price: 199000, ProductId: 9, SizeId: 20, ColorId: 8},   // Bamboo Cool, 80-90kg, Magenta
	{Sku: 7019, Amount: 10, Price: 149000, ProductId: 10, SizeId: 17, ColorId: 9}, // Coolmate Briefs, 50-60kg, Silver
	{Sku: 7020, Amount: 5, Price: 149000, ProductId: 10, SizeId: 18, ColorId: 10}, // Coolmate Briefs, 60-70kg, Gray
	// Nữ
	{Sku: 7021, Amount: 28, Price: 399000, ProductId: 11, SizeId: 17, ColorId: 1}, // VS Cotton Bikini, 50-60kg, Black
	{Sku: 7022, Amount: 18, Price: 399000, ProductId: 11, SizeId: 18, ColorId: 2}, // VS Cotton Bikini, 60-70kg, White
	{Sku: 7023, Amount: 24, Price: 199000, ProductId: 12, SizeId: 19, ColorId: 3}, // Uniqlo AIRism, 70-80kg, Red
	{Sku: 7024, Amount: 14, Price: 199000, ProductId: 12, SizeId: 20, ColorId: 4}, // Uniqlo AIRism, 80-90kg, Lime
	{Sku: 7025, Amount: 20, Price: 159000, ProductId: 13, SizeId: 17, ColorId: 5}, // Bon Bon Briefs, 50-60kg, Blue
	{Sku: 7026, Amount: 10, Price: 159000, ProductId: 13, SizeId: 18, ColorId: 6}, // Bon Bon Briefs, 60-70kg, Yellow
	{Sku: 7027, Amount: 16, Price: 299000, ProductId: 14, SizeId: 19, ColorId: 7}, // Triumph Maxi, 70-80kg, Cyan
	{Sku: 7028, Amount: 8, Price: 299000, ProductId: 14, SizeId: 20, ColorId: 8},  // Triumph Maxi, 80-90kg, Magenta
	{Sku: 7029, Amount: 14, Price: 249000, ProductId: 15, SizeId: 17, ColorId: 9}, // Sabina Panty, 50-60kg, Silver
	{Sku: 7030, Amount: 7, Price: 249000, ProductId: 15, SizeId: 18, ColorId: 10}, // Sabina Panty, 60-70kg, Gray
	{Sku: 7031, Amount: 12, Price: 199000, ProductId: 16, SizeId: 19, ColorId: 1}, // Sunfly Boxer, 70-80kg, Black
	{Sku: 7032, Amount: 6, Price: 199000, ProductId: 16, SizeId: 20, ColorId: 2},  // Sunfly Boxer, 80-90kg, White
	{Sku: 7033, Amount: 10, Price: 299000, ProductId: 17, SizeId: 17, ColorId: 3}, // Wacoal Lace Bikini, 50-60kg, Red
	{Sku: 7034, Amount: 5, Price: 299000, ProductId: 17, SizeId: 18, ColorId: 4},  // Wacoal Lace Bikini, 60-70kg, Lime
	{Sku: 7035, Amount: 8, Price: 249000, ProductId: 18, SizeId: 19, ColorId: 5},  // Jockey Hipster, 70-80kg, Blue
	{Sku: 7036, Amount: 4, Price: 249000, ProductId: 18, SizeId: 20, ColorId: 6},  // Jockey Hipster, 80-90kg, Yellow
	{Sku: 7037, Amount: 7, Price: 199000, ProductId: 19, SizeId: 17, ColorId: 7},  // Bon Bon Boxer, 50-60kg, Cyan
	{Sku: 7038, Amount: 3, Price: 199000, ProductId: 19, SizeId: 18, ColorId: 8},  // Bon Bon Boxer, 60-70kg, Magenta
	{Sku: 7039, Amount: 9, Price: 149000, ProductId: 20, SizeId: 19, ColorId: 9},  // Coolmate Briefs, 70-80kg, Silver
	{Sku: 7040, Amount: 5, Price: 149000, ProductId: 20, SizeId: 20, ColorId: 10}, // Coolmate Briefs, 80-90kg, Gray
}
