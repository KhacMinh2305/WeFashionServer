package product

// import (
// 	"WeFashionServer/infrastructure/model"
// )

// // Mock data for Shirts category (áo sơ mi, áo thun, áo polo, áo dài tay, ...)
// var MockShirtsProducts = []model.Product{
// 	// 50 sản phẩm áo thực tế, đa dạng thương hiệu, hình ảnh, mô tả, rating, sold, liked, ShopId khác nhau
// 	// CategoryId = 2 (Shirts)
// 	{Id: 1, Name: "Uniqlo Oxford Slim Fit Shirt", ImageUrl: "https://uniqlo.com/oxford-shirt.jpg", Description: "Áo sơ mi Uniqlo Oxford, vải cotton cao cấp, form slim fit, dễ phối đồ.", Rating: 4.8, SoldAmount: 1200, LikedAmount: 900, CategoryId: 2, ShopId: 1},
// 	{Id: 2, Name: "Zara Basic Poplin Shirt", ImageUrl: "https://zara.com/basic-poplin.jpg", Description: "Áo sơ mi Zara poplin, chất liệu mềm mại, phù hợp công sở và đi chơi.", Rating: 4.7, SoldAmount: 950, LikedAmount: 800, CategoryId: 2, ShopId: 2},
// 	{Id: 3, Name: "H&M Regular Fit Shirt", ImageUrl: "https://hm.com/regular-fit-shirt.jpg", Description: "Áo sơ mi H&M regular fit, phong cách trẻ trung, dễ phối.", Rating: 4.6, SoldAmount: 700, LikedAmount: 600, CategoryId: 2, ShopId: 3},
// 	{Id: 4, Name: "Lacoste Classic Polo Shirt", ImageUrl: "https://lacoste.com/classic-polo.jpg", Description: "Áo polo Lacoste cổ điển, logo cá sấu, chất liệu cao cấp.", Rating: 4.9, SoldAmount: 1100, LikedAmount: 950, CategoryId: 2, ShopId: 4},
// 	{Id: 5, Name: "Tommy Hilfiger Slim Fit Shirt", ImageUrl: "https://tommy.com/slim-fit-shirt.jpg", Description: "Áo sơ mi Tommy Hilfiger slim fit, lịch lãm, sang trọng.", Rating: 4.7, SoldAmount: 850, LikedAmount: 700, CategoryId: 2, ShopId: 5},
// 	{Id: 6, Name: "Calvin Klein White Dress Shirt", ImageUrl: "https://calvinklein.com/white-dress-shirt.jpg", Description: "Áo sơ mi trắng Calvin Klein, phù hợp công sở, chất liệu thoáng mát.", Rating: 4.6, SoldAmount: 800, LikedAmount: 650, CategoryId: 2, ShopId: 6},
// 	{Id: 7, Name: "Levi's Western Denim Shirt", ImageUrl: "https://levis.com/western-denim.jpg", Description: "Áo sơ mi denim Levi's, phong cách Mỹ, cá tính.", Rating: 4.5, SoldAmount: 600, LikedAmount: 500, CategoryId: 2, ShopId: 7},
// 	{Id: 8, Name: "Gap Linen Shirt", ImageUrl: "https://gap.com/linen-shirt.jpg", Description: "Áo sơ mi linen Gap, thoáng mát, thích hợp mùa hè.", Rating: 4.4, SoldAmount: 500, LikedAmount: 400, CategoryId: 2, ShopId: 8},
// 	{Id: 9, Name: "Pull&Bear Checked Shirt", ImageUrl: "https://pullbear.com/checked-shirt.jpg", Description: "Áo sơ mi caro Pull&Bear, trẻ trung, năng động.", Rating: 4.6, SoldAmount: 450, LikedAmount: 350, CategoryId: 2, ShopId: 9},
// 	{Id: 10, Name: "Massimo Dutti Slim Fit Shirt", ImageUrl: "https://massimodutti.com/slim-fit-shirt.jpg", Description: "Áo sơ mi Massimo Dutti slim fit, sang trọng, lịch lãm.", Rating: 4.7, SoldAmount: 400, LikedAmount: 300, CategoryId: 2, ShopId: 10},
// 	{Id: 11, Name: "Banana Republic Linen Shirt", ImageUrl: "https://bananarepublic.com/linen-shirt.jpg", Description: "Áo sơ mi linen Banana Republic, nhẹ, thoáng khí, phù hợp mùa hè.", Rating: 4.5, SoldAmount: 380, LikedAmount: 290, CategoryId: 2, ShopId: 11},
// 	{Id: 12, Name: "Abercrombie & Fitch Classic Shirt", ImageUrl: "https://abercrombie.com/classic-shirt.jpg", Description: "Áo sơ mi Abercrombie cổ điển, form rộng, trẻ trung.", Rating: 4.6, SoldAmount: 370, LikedAmount: 280, CategoryId: 2, ShopId: 12},
// 	{Id: 13, Name: "Express Slim Fit Dress Shirt", ImageUrl: "https://express.com/slim-fit-dress-shirt.jpg", Description: "Áo sơ mi Express slim fit, hiện đại, dễ phối đồ.", Rating: 4.7, SoldAmount: 360, LikedAmount: 270, CategoryId: 2, ShopId: 13},
// 	{Id: 14, Name: "Superdry Checked Shirt", ImageUrl: "https://superdry.com/checked-shirt.jpg", Description: "Áo sơ mi caro Superdry, phong cách Anh quốc.", Rating: 4.5, SoldAmount: 350, LikedAmount: 260, CategoryId: 2, ShopId: 14},
// 	{Id: 15, Name: "Mango Man Regular Fit Shirt", ImageUrl: "https://mango.com/regular-fit-shirt.jpg", Description: "Áo sơ mi Mango Man regular fit, lịch lãm, phù hợp công sở.", Rating: 4.6, SoldAmount: 340, LikedAmount: 250, CategoryId: 2, ShopId: 15},
// 	{Id: 16, Name: "Bershka Denim Shirt", ImageUrl: "https://bershka.com/denim-shirt.jpg", Description: "Áo sơ mi denim Bershka, trẻ trung, cá tính.", Rating: 4.4, SoldAmount: 330, LikedAmount: 240, CategoryId: 2, ShopId: 16},
// 	{Id: 17, Name: "GU Easy Care Shirt", ImageUrl: "https://gu-global.com/easy-care-shirt.jpg", Description: "Áo sơ mi GU easy care, chống nhăn, dễ giặt.", Rating: 4.5, SoldAmount: 320, LikedAmount: 230, CategoryId: 2, ShopId: 17},
// 	{Id: 18, Name: "Topman Slim Fit Shirt", ImageUrl: "https://topman.com/slim-fit-shirt.jpg", Description: "Áo sơ mi Topman slim fit, form ôm, cá tính.", Rating: 4.6, SoldAmount: 310, LikedAmount: 220, CategoryId: 2, ShopId: 18},
// 	{Id: 19, Name: "Routine Oxford Shirt", ImageUrl: "https://routine.com/oxford-shirt.jpg", Description: "Áo sơ mi Routine Oxford, trẻ trung, năng động.", Rating: 4.7, SoldAmount: 300, LikedAmount: 210, CategoryId: 2, ShopId: 19},
// 	{Id: 20, Name: "Canifa Cotton Shirt", ImageUrl: "https://canifa.com/cotton-shirt.jpg", Description: "Áo sơ mi cotton Canifa, mềm mại, thoáng mát.", Rating: 4.5, SoldAmount: 290, LikedAmount: 200, CategoryId: 2, ShopId: 20},
// 	{Id: 21, Name: "Owen Classic Shirt", ImageUrl: "https://owen.com/classic-shirt.jpg", Description: "Áo sơ mi Owen classic, lịch lãm, dễ phối.", Rating: 4.6, SoldAmount: 280, LikedAmount: 190, CategoryId: 2, ShopId: 21},
// 	{Id: 22, Name: "Biluxury Slim Fit Shirt", ImageUrl: "https://biluxury.com/slim-fit-shirt.jpg", Description: "Áo sơ mi Biluxury slim fit, trẻ trung, hiện đại.", Rating: 4.7, SoldAmount: 270, LikedAmount: 180, CategoryId: 2, ShopId: 22},
// 	{Id: 23, Name: "Coolmate Active Shirt", ImageUrl: "https://coolmate.me/active-shirt.jpg", Description: "Áo sơ mi Coolmate active, phù hợp vận động.", Rating: 4.5, SoldAmount: 260, LikedAmount: 170, CategoryId: 2, ShopId: 23},
// 	{Id: 24, Name: "Yody Linen Shirt", ImageUrl: "https://yody.vn/linen-shirt.jpg", Description: "Áo sơ mi linen Yody, nhẹ nhàng, thoáng mát.", Rating: 4.6, SoldAmount: 250, LikedAmount: 160, CategoryId: 2, ShopId: 24},
// 	{Id: 25, Name: "Routine Checked Shirt", ImageUrl: "https://routine.com/checked-shirt.jpg", Description: "Áo sơ mi caro Routine, trẻ trung, cá tính.", Rating: 4.7, SoldAmount: 240, LikedAmount: 150, CategoryId: 2, ShopId: 25},
// 	{Id: 26, Name: "Canifa Easy Iron Shirt", ImageUrl: "https://canifa.com/easy-iron-shirt.jpg", Description: "Áo sơ mi Canifa easy iron, dễ là, chống nhăn.", Rating: 4.5, SoldAmount: 230, LikedAmount: 140, CategoryId: 2, ShopId: 26},
// 	{Id: 27, Name: "Owen Linen Shirt", ImageUrl: "https://owen.com/linen-shirt.jpg", Description: "Áo sơ mi linen Owen, nhẹ, thoáng khí.", Rating: 4.6, SoldAmount: 220, LikedAmount: 130, CategoryId: 2, ShopId: 27},
// 	{Id: 28, Name: "Biluxury Classic Shirt", ImageUrl: "https://biluxury.com/classic-shirt.jpg", Description: "Áo sơ mi Biluxury classic, lịch lãm, sang trọng.", Rating: 4.7, SoldAmount: 210, LikedAmount: 120, CategoryId: 2, ShopId: 28},
// 	{Id: 29, Name: "Yody Checked Shirt", ImageUrl: "https://yody.vn/checked-shirt.jpg", Description: "Áo sơ mi caro Yody, trẻ trung, năng động.", Rating: 4.5, SoldAmount: 200, LikedAmount: 110, CategoryId: 2, ShopId: 29},
// 	{Id: 30, Name: "Coolmate Linen Shirt", ImageUrl: "https://coolmate.me/linen-shirt.jpg", Description: "Áo sơ mi linen Coolmate, nhẹ, thoáng khí.", Rating: 4.6, SoldAmount: 190, LikedAmount: 100, CategoryId: 2, ShopId: 30},
// 	{Id: 31, Name: "Routine Polo Shirt", ImageUrl: "https://routine.com/polo-shirt.jpg", Description: "Áo polo Routine, trẻ trung, năng động.", Rating: 4.7, SoldAmount: 180, LikedAmount: 90, CategoryId: 2, ShopId: 31},
// 	{Id: 32, Name: "Canifa Polo Shirt", ImageUrl: "https://canifa.com/polo-shirt.jpg", Description: "Áo polo Canifa, thoáng mát, dễ phối.", Rating: 4.5, SoldAmount: 170, LikedAmount: 80, CategoryId: 2, ShopId: 32},
// 	{Id: 33, Name: "Owen Polo Shirt", ImageUrl: "https://owen.com/polo-shirt.jpg", Description: "Áo polo Owen, trẻ trung, cá tính.", Rating: 4.6, SoldAmount: 160, LikedAmount: 70, CategoryId: 2, ShopId: 33},
// 	{Id: 34, Name: "Biluxury Polo Shirt", ImageUrl: "https://biluxury.com/polo-shirt.jpg", Description: "Áo polo Biluxury, lịch lãm, hiện đại.", Rating: 4.7, SoldAmount: 150, LikedAmount: 60, CategoryId: 2, ShopId: 34},
// 	{Id: 35, Name: "Yody Polo Shirt", ImageUrl: "https://yody.vn/polo-shirt.jpg", Description: "Áo polo Yody, trẻ trung, năng động.", Rating: 4.5, SoldAmount: 140, LikedAmount: 50, CategoryId: 2, ShopId: 35},
// 	{Id: 36, Name: "Coolmate Polo Shirt", ImageUrl: "https://coolmate.me/polo-shirt.jpg", Description: "Áo polo Coolmate, nhẹ, thoáng khí.", Rating: 4.6, SoldAmount: 130, LikedAmount: 40, CategoryId: 2, ShopId: 36},
// 	{Id: 37, Name: "Routine Long Sleeve Shirt", ImageUrl: "https://routine.com/long-sleeve-shirt.jpg", Description: "Áo sơ mi dài tay Routine, lịch lãm, phù hợp công sở.", Rating: 4.7, SoldAmount: 120, LikedAmount: 30, CategoryId: 2, ShopId: 37},
// 	{Id: 38, Name: "Canifa Long Sleeve Shirt", ImageUrl: "https://canifa.com/long-sleeve-shirt.jpg", Description: "Áo sơ mi dài tay Canifa, nhẹ, thoáng khí.", Rating: 4.5, SoldAmount: 110, LikedAmount: 20, CategoryId: 2, ShopId: 38},
// 	{Id: 39, Name: "Owen Long Sleeve Shirt", ImageUrl: "https://owen.com/long-sleeve-shirt.jpg", Description: "Áo sơ mi dài tay Owen, trẻ trung, hiện đại.", Rating: 4.6, SoldAmount: 100, LikedAmount: 10, CategoryId: 2, ShopId: 39},
// 	{Id: 40, Name: "Biluxury Long Sleeve Shirt", ImageUrl: "https://biluxury.com/long-sleeve-shirt.jpg", Description: "Áo sơ mi dài tay Biluxury, lịch lãm, sang trọng.", Rating: 4.7, SoldAmount: 90, LikedAmount: 8, CategoryId: 2, ShopId: 40},
// 	{Id: 41, Name: "Yody Long Sleeve Shirt", ImageUrl: "https://yody.vn/long-sleeve-shirt.jpg", Description: "Áo sơ mi dài tay Yody, trẻ trung, năng động.", Rating: 4.5, SoldAmount: 80, LikedAmount: 6, CategoryId: 2, ShopId: 41},
// 	{Id: 42, Name: "Coolmate Long Sleeve Shirt", ImageUrl: "https://coolmate.me/long-sleeve-shirt.jpg", Description: "Áo sơ mi dài tay Coolmate, nhẹ, thoáng khí.", Rating: 4.6, SoldAmount: 70, LikedAmount: 4, CategoryId: 2, ShopId: 42},
// 	{Id: 43, Name: "Routine Short Sleeve Shirt", ImageUrl: "https://routine.com/short-sleeve-shirt.jpg", Description: "Áo sơ mi ngắn tay Routine, trẻ trung, năng động.", Rating: 4.7, SoldAmount: 60, LikedAmount: 2, CategoryId: 2, ShopId: 43},
// 	{Id: 44, Name: "Canifa Short Sleeve Shirt", ImageUrl: "https://canifa.com/short-sleeve-shirt.jpg", Description: "Áo sơ mi ngắn tay Canifa, nhẹ, thoáng khí.", Rating: 4.5, SoldAmount: 50, LikedAmount: 1, CategoryId: 2, ShopId: 44},
// 	{Id: 45, Name: "Owen Short Sleeve Shirt", ImageUrl: "https://owen.com/short-sleeve-shirt.jpg", Description: "Áo sơ mi ngắn tay Owen, trẻ trung, hiện đại.", Rating: 4.6, SoldAmount: 40, LikedAmount: 1, CategoryId: 2, ShopId: 45},
// 	{Id: 46, Name: "Biluxury Short Sleeve Shirt", ImageUrl: "https://biluxury.com/short-sleeve-shirt.jpg", Description: "Áo sơ mi ngắn tay Biluxury, lịch lãm, sang trọng.", Rating: 4.7, SoldAmount: 30, LikedAmount: 1, CategoryId: 2, ShopId: 46},
// 	{Id: 47, Name: "Yody Short Sleeve Shirt", ImageUrl: "https://yody.vn/short-sleeve-shirt.jpg", Description: "Áo sơ mi ngắn tay Yody, trẻ trung, năng động.", Rating: 4.5, SoldAmount: 20, LikedAmount: 1, CategoryId: 2, ShopId: 47},
// 	{Id: 48, Name: "Coolmate Short Sleeve Shirt", ImageUrl: "https://coolmate.me/short-sleeve-shirt.jpg", Description: "Áo sơ mi ngắn tay Coolmate, nhẹ, thoáng khí.", Rating: 4.6, SoldAmount: 10, LikedAmount: 1, CategoryId: 2, ShopId: 48},
// 	{Id: 49, Name: "Routine Mandarin Collar Shirt", ImageUrl: "https://routine.com/mandarin-collar-shirt.jpg", Description: "Áo sơ mi cổ tàu Routine, trẻ trung, hiện đại.", Rating: 4.7, SoldAmount: 8, LikedAmount: 1, CategoryId: 2, ShopId: 49},
// 	{Id: 50, Name: "Canifa Mandarin Collar Shirt", ImageUrl: "https://canifa.com/mandarin-collar-shirt.jpg", Description: "Áo sơ mi cổ tàu Canifa, nhẹ, thoáng khí.", Rating: 4.5, SoldAmount: 6, LikedAmount: 1, CategoryId: 2, ShopId: 50},
// }

// // Mock data for ProductVariants (biến thể sản phẩm) cho Shirts
// var MockShirtsVariants = []model.ProductVariant{
// 	// SKU cho từng sản phẩm Shirts, mỗi sản phẩm ít nhất 2 biến thể (size, màu)
// 	// Ví dụ cho 5 sản phẩm đầu tiên, các sản phẩm còn lại sinh tương tự
// 	{Sku: 2001, Amount: 50, Price: 599000, ProductId: 1, SizeId: 3, ColorId: 1},   // Uniqlo Oxford, size M, Black
// 	{Sku: 2002, Amount: 30, Price: 599000, ProductId: 1, SizeId: 4, ColorId: 2},   // Uniqlo Oxford, size L, White
// 	{Sku: 2003, Amount: 40, Price: 599000, ProductId: 1, SizeId: 5, ColorId: 5},   // Uniqlo Oxford, size XL, Blue
// 	{Sku: 2004, Amount: 35, Price: 499000, ProductId: 2, SizeId: 3, ColorId: 1},   // Zara Poplin, size M, Black
// 	{Sku: 2005, Amount: 25, Price: 499000, ProductId: 2, SizeId: 4, ColorId: 4},   // Zara Poplin, size L, Lime
// 	{Sku: 2006, Amount: 20, Price: 499000, ProductId: 2, SizeId: 5, ColorId: 5},   // Zara Poplin, size XL, Blue
// 	{Sku: 2007, Amount: 60, Price: 399000, ProductId: 3, SizeId: 2, ColorId: 1},   // H&M Shirt, size S, Black
// 	{Sku: 2008, Amount: 45, Price: 399000, ProductId: 3, SizeId: 3, ColorId: 6},   // H&M Shirt, size M, Yellow
// 	{Sku: 2009, Amount: 30, Price: 399000, ProductId: 3, SizeId: 4, ColorId: 7},   // H&M Shirt, size L, Cyan
// 	{Sku: 2010, Amount: 55, Price: 899000, ProductId: 4, SizeId: 3, ColorId: 1},   // Lacoste Polo, size M, Black
// 	{Sku: 2011, Amount: 35, Price: 899000, ProductId: 4, SizeId: 4, ColorId: 8},   // Lacoste Polo, size L, Magenta
// 	{Sku: 2012, Amount: 25, Price: 899000, ProductId: 4, SizeId: 5, ColorId: 9},   // Lacoste Polo, size XL, Silver
// 	{Sku: 2013, Amount: 40, Price: 799000, ProductId: 5, SizeId: 3, ColorId: 10},  // Tommy Shirt, size M, Gray
// 	{Sku: 2014, Amount: 30, Price: 799000, ProductId: 5, SizeId: 4, ColorId: 11},  // Tommy Shirt, size L, Maroon
// 	{Sku: 2015, Amount: 20, Price: 799000, ProductId: 5, SizeId: 5, ColorId: 12},  // Tommy Shirt, size XL, Olive
// 	{Sku: 2016, Amount: 38, Price: 699000, ProductId: 6, SizeId: 3, ColorId: 1},   // Calvin Klein, size M, Black
// 	{Sku: 2017, Amount: 28, Price: 699000, ProductId: 6, SizeId: 4, ColorId: 2},   // Calvin Klein, size L, White
// 	{Sku: 2018, Amount: 22, Price: 699000, ProductId: 6, SizeId: 5, ColorId: 5},   // Calvin Klein, size XL, Blue
// 	{Sku: 2019, Amount: 32, Price: 749000, ProductId: 7, SizeId: 2, ColorId: 3},   // Levi's Denim, size S, Red
// 	{Sku: 2020, Amount: 24, Price: 749000, ProductId: 7, SizeId: 3, ColorId: 4},   // Levi's Denim, size M, Lime
// 	{Sku: 2021, Amount: 18, Price: 749000, ProductId: 7, SizeId: 4, ColorId: 6},   // Levi's Denim, size L, Yellow
// 	{Sku: 2022, Amount: 30, Price: 599000, ProductId: 8, SizeId: 3, ColorId: 7},   // Gap Linen, size M, Cyan
// 	{Sku: 2023, Amount: 22, Price: 599000, ProductId: 8, SizeId: 4, ColorId: 8},   // Gap Linen, size L, Magenta
// 	{Sku: 2024, Amount: 16, Price: 599000, ProductId: 8, SizeId: 5, ColorId: 9},   // Gap Linen, size XL, Silver
// 	{Sku: 2025, Amount: 28, Price: 499000, ProductId: 9, SizeId: 2, ColorId: 10},  // Pull&Bear Checked, size S, Gray
// 	{Sku: 2026, Amount: 20, Price: 499000, ProductId: 9, SizeId: 3, ColorId: 11},  // Pull&Bear Checked, size M, Maroon
// 	{Sku: 2027, Amount: 14, Price: 499000, ProductId: 9, SizeId: 4, ColorId: 12},  // Pull&Bear Checked, size L, Olive
// 	{Sku: 2028, Amount: 26, Price: 899000, ProductId: 10, SizeId: 3, ColorId: 1},  // Massimo Dutti, size M, Black
// 	{Sku: 2029, Amount: 18, Price: 899000, ProductId: 10, SizeId: 4, ColorId: 2},  // Massimo Dutti, size L, White
// 	{Sku: 2030, Amount: 12, Price: 899000, ProductId: 10, SizeId: 5, ColorId: 5},  // Massimo Dutti, size XL, Blue
// 	{Sku: 2031, Amount: 24, Price: 599000, ProductId: 11, SizeId: 3, ColorId: 3},  // Banana Republic, size M, Red
// 	{Sku: 2032, Amount: 16, Price: 599000, ProductId: 11, SizeId: 4, ColorId: 4},  // Banana Republic, size L, Lime
// 	{Sku: 2033, Amount: 10, Price: 599000, ProductId: 11, SizeId: 5, ColorId: 6},  // Banana Republic, size XL, Yellow
// 	{Sku: 2034, Amount: 22, Price: 699000, ProductId: 12, SizeId: 2, ColorId: 7},  // Abercrombie, size S, Cyan
// 	{Sku: 2035, Amount: 14, Price: 699000, ProductId: 12, SizeId: 3, ColorId: 8},  // Abercrombie, size M, Magenta
// 	{Sku: 2036, Amount: 8, Price: 699000, ProductId: 12, SizeId: 4, ColorId: 9},   // Abercrombie, size L, Silver
// 	{Sku: 2037, Amount: 20, Price: 799000, ProductId: 13, SizeId: 3, ColorId: 10}, // Express, size M, Gray
// 	{Sku: 2038, Amount: 12, Price: 799000, ProductId: 13, SizeId: 4, ColorId: 11}, // Express, size L, Maroon
// 	{Sku: 2039, Amount: 6, Price: 799000, ProductId: 13, SizeId: 5, ColorId: 12},  // Express, size XL, Olive
// 	{Sku: 2040, Amount: 18, Price: 649000, ProductId: 14, SizeId: 2, ColorId: 1},  // Superdry, size S, Black
// 	{Sku: 2041, Amount: 10, Price: 649000, ProductId: 14, SizeId: 3, ColorId: 2},  // Superdry, size M, White
// 	{Sku: 2042, Amount: 5, Price: 649000, ProductId: 14, SizeId: 4, ColorId: 3},   // Superdry, size L, Red
// 	{Sku: 2043, Amount: 16, Price: 599000, ProductId: 15, SizeId: 3, ColorId: 4},  // Mango Man, size M, Lime
// 	{Sku: 2044, Amount: 8, Price: 599000, ProductId: 15, SizeId: 4, ColorId: 5},   // Mango Man, size L, Blue
// 	{Sku: 2045, Amount: 4, Price: 599000, ProductId: 15, SizeId: 5, ColorId: 6},   // Mango Man, size XL, Yellow
// 	{Sku: 2046, Amount: 14, Price: 499000, ProductId: 16, SizeId: 2, ColorId: 7},  // Bershka, size S, Cyan
// 	{Sku: 2047, Amount: 6, Price: 499000, ProductId: 16, SizeId: 3, ColorId: 8},   // Bershka, size M, Magenta
// 	{Sku: 2048, Amount: 3, Price: 499000, ProductId: 16, SizeId: 4, ColorId: 9},   // Bershka, size L, Silver
// 	{Sku: 2049, Amount: 12, Price: 599000, ProductId: 17, SizeId: 3, ColorId: 10}, // GU, size M, Gray
// 	{Sku: 2050, Amount: 8, Price: 599000, ProductId: 17, SizeId: 4, ColorId: 11},  // GU, size L, Maroon
// 	{Sku: 2051, Amount: 4, Price: 599000, ProductId: 17, SizeId: 5, ColorId: 12},  // GU, size XL, Olive
// 	{Sku: 2052, Amount: 10, Price: 699000, ProductId: 18, SizeId: 2, ColorId: 1},  // Topman, size S, Black
// 	{Sku: 2053, Amount: 6, Price: 699000, ProductId: 18, SizeId: 3, ColorId: 2},   // Topman, size M, White
// 	{Sku: 2054, Amount: 3, Price: 699000, ProductId: 18, SizeId: 4, ColorId: 3},   // Topman, size L, Red
// 	{Sku: 2055, Amount: 8, Price: 599000, ProductId: 19, SizeId: 3, ColorId: 4},   // Routine, size M, Lime
// 	{Sku: 2056, Amount: 4, Price: 599000, ProductId: 19, SizeId: 4, ColorId: 5},   // Routine, size L, Blue
// 	{Sku: 2057, Amount: 2, Price: 599000, ProductId: 19, SizeId: 5, ColorId: 6},   // Routine, size XL, Yellow
// 	{Sku: 2058, Amount: 6, Price: 499000, ProductId: 20, SizeId: 2, ColorId: 7},   // Canifa, size S, Cyan
// 	{Sku: 2059, Amount: 3, Price: 499000, ProductId: 20, SizeId: 3, ColorId: 8},   // Canifa, size M, Magenta
// 	{Sku: 2060, Amount: 2, Price: 499000, ProductId: 20, SizeId: 4, ColorId: 9},   // Canifa, size L, Silver
// 	{Sku: 2061, Amount: 5, Price: 599000, ProductId: 21, SizeId: 3, ColorId: 10},  // Owen, size M, Gray
// 	{Sku: 2062, Amount: 3, Price: 599000, ProductId: 21, SizeId: 4, ColorId: 11},  // Owen, size L, Maroon
// 	{Sku: 2063, Amount: 2, Price: 599000, ProductId: 21, SizeId: 5, ColorId: 12},  // Owen, size XL, Olive
// 	{Sku: 2064, Amount: 4, Price: 699000, ProductId: 22, SizeId: 2, ColorId: 1},   // Biluxury, size S, Black
// 	{Sku: 2065, Amount: 2, Price: 699000, ProductId: 22, SizeId: 3, ColorId: 2},   // Biluxury, size M, White
// 	{Sku: 2066, Amount: 1, Price: 699000, ProductId: 22, SizeId: 4, ColorId: 3},   // Biluxury, size L, Red
// 	{Sku: 2067, Amount: 3, Price: 599000, ProductId: 23, SizeId: 3, ColorId: 4},   // Coolmate, size M, Lime
// 	{Sku: 2068, Amount: 2, Price: 599000, ProductId: 23, SizeId: 4, ColorId: 5},   // Coolmate, size L, Blue
// 	{Sku: 2069, Amount: 1, Price: 599000, ProductId: 23, SizeId: 5, ColorId: 6},   // Coolmate, size XL, Yellow
// 	{Sku: 2070, Amount: 2, Price: 499000, ProductId: 24, SizeId: 2, ColorId: 7},   // Yody, size S, Cyan
// 	{Sku: 2071, Amount: 1, Price: 499000, ProductId: 24, SizeId: 3, ColorId: 8},   // Yody, size M, Magenta
// 	{Sku: 2072, Amount: 1, Price: 499000, ProductId: 24, SizeId: 4, ColorId: 9},   // Yody, size L, Silver
// 	{Sku: 2073, Amount: 2, Price: 599000, ProductId: 25, SizeId: 3, ColorId: 10},  // Routine Checked, size M, Gray
// 	{Sku: 2074, Amount: 1, Price: 599000, ProductId: 25, SizeId: 4, ColorId: 11},  // Routine Checked, size L, Maroon
// 	{Sku: 2075, Amount: 1, Price: 599000, ProductId: 25, SizeId: 5, ColorId: 12},  // Routine Checked, size XL, Olive
// 	{Sku: 2076, Amount: 2, Price: 699000, ProductId: 26, SizeId: 2, ColorId: 1},   // Canifa Easy Iron, size S, Black
// 	{Sku: 2077, Amount: 1, Price: 699000, ProductId: 26, SizeId: 3, ColorId: 2},   // Canifa Easy Iron, size M, White
// 	{Sku: 2078, Amount: 1, Price: 699000, ProductId: 26, SizeId: 4, ColorId: 3},   // Canifa Easy Iron, size L, Red
// 	{Sku: 2079, Amount: 2, Price: 599000, ProductId: 27, SizeId: 3, ColorId: 4},   // Owen Linen, size M, Lime
// 	{Sku: 2080, Amount: 1, Price: 599000, ProductId: 27, SizeId: 4, ColorId: 5},   // Owen Linen, size L, Blue
// 	{Sku: 2081, Amount: 1, Price: 599000, ProductId: 27, SizeId: 5, ColorId: 6},   // Owen Linen, size XL, Yellow
// 	{Sku: 2082, Amount: 2, Price: 499000, ProductId: 28, SizeId: 2, ColorId: 7},   // Biluxury Classic, size S, Cyan
// 	{Sku: 2083, Amount: 1, Price: 499000, ProductId: 28, SizeId: 3, ColorId: 8},   // Biluxury Classic, size M, Magenta
// 	{Sku: 2084, Amount: 1, Price: 499000, ProductId: 28, SizeId: 4, ColorId: 9},   // Biluxury Classic, size L, Silver
// 	{Sku: 2085, Amount: 2, Price: 599000, ProductId: 29, SizeId: 3, ColorId: 10},  // Yody Checked, size M, Gray
// 	{Sku: 2086, Amount: 1, Price: 599000, ProductId: 29, SizeId: 4, ColorId: 11},  // Yody Checked, size L, Maroon
// 	{Sku: 2087, Amount: 1, Price: 599000, ProductId: 29, SizeId: 5, ColorId: 12},  // Yody Checked, size XL, Olive
// 	{Sku: 2088, Amount: 2, Price: 699000, ProductId: 30, SizeId: 2, ColorId: 1},   // Coolmate Linen, size S, Black
// 	{Sku: 2089, Amount: 1, Price: 699000, ProductId: 30, SizeId: 3, ColorId: 2},   // Coolmate Linen, size M, White
// 	{Sku: 2090, Amount: 1, Price: 699000, ProductId: 30, SizeId: 4, ColorId: 3},   // Coolmate Linen, size L, Red
// 	{Sku: 2091, Amount: 2, Price: 599000, ProductId: 31, SizeId: 3, ColorId: 4},   // Routine Polo, size M, Lime
// 	{Sku: 2092, Amount: 1, Price: 599000, ProductId: 31, SizeId: 4, ColorId: 5},   // Routine Polo, size L, Blue
// 	{Sku: 2093, Amount: 1, Price: 599000, ProductId: 31, SizeId: 5, ColorId: 6},   // Routine Polo, size XL, Yellow
// 	{Sku: 2094, Amount: 2, Price: 699000, ProductId: 32, SizeId: 2, ColorId: 7},   // Canifa Polo, size S, Cyan
// 	{Sku: 2095, Amount: 1, Price: 699000, ProductId: 32, SizeId: 3, ColorId: 8},   // Canifa Polo, size M, Magenta
// 	{Sku: 2096, Amount: 1, Price: 699000, ProductId: 32, SizeId: 4, ColorId: 9},   // Canifa Polo, size L, Silver
// 	{Sku: 2097, Amount: 2, Price: 599000, ProductId: 33, SizeId: 3, ColorId: 10},  // Owen Polo, size M, Gray
// 	{Sku: 2098, Amount: 1, Price: 599000, ProductId: 33, SizeId: 4, ColorId: 11},  // Owen Polo, size L, Maroon
// 	{Sku: 2099, Amount: 1, Price: 599000, ProductId: 33, SizeId: 5, ColorId: 12},  // Owen Polo, size XL, Olive
// 	{Sku: 2100, Amount: 2, Price: 699000, ProductId: 34, SizeId: 2, ColorId: 1},   // Biluxury Polo, size S, Black
// 	{Sku: 2101, Amount: 1, Price: 699000, ProductId: 34, SizeId: 3, ColorId: 2},   // Biluxury Polo, size M, White
// 	{Sku: 2102, Amount: 1, Price: 699000, ProductId: 34, SizeId: 4, ColorId: 3},   // Biluxury Polo, size L, Red
// 	{Sku: 2103, Amount: 2, Price: 599000, ProductId: 35, SizeId: 3, ColorId: 4},   // Yody Polo, size M, Lime
// 	{Sku: 2104, Amount: 1, Price: 599000, ProductId: 35, SizeId: 4, ColorId: 5},   // Yody Polo, size L, Blue
// 	{Sku: 2105, Amount: 1, Price: 599000, ProductId: 35, SizeId: 5, ColorId: 6},   // Yody Polo, size XL, Yellow
// 	{Sku: 2106, Amount: 2, Price: 699000, ProductId: 36, SizeId: 2, ColorId: 7},   // Coolmate Polo, size S, Cyan
// 	{Sku: 2107, Amount: 1, Price: 699000, ProductId: 36, SizeId: 3, ColorId: 8},   // Coolmate Polo, size M, Magenta
// 	{Sku: 2108, Amount: 1, Price: 699000, ProductId: 36, SizeId: 4, ColorId: 9},   // Coolmate Polo, size L, Silver
// 	{Sku: 2109, Amount: 2, Price: 599000, ProductId: 37, SizeId: 3, ColorId: 10},  // Routine Long Sleeve, size M, Gray
// 	{Sku: 2110, Amount: 1, Price: 599000, ProductId: 37, SizeId: 4, ColorId: 11},  // Routine Long Sleeve, size L, Maroon
// 	{Sku: 2111, Amount: 1, Price: 599000, ProductId: 37, SizeId: 5, ColorId: 12},  // Routine Long Sleeve, size XL, Olive
// 	{Sku: 2112, Amount: 2, Price: 699000, ProductId: 38, SizeId: 2, ColorId: 1},   // Canifa Long Sleeve, size S, Black
// 	{Sku: 2113, Amount: 1, Price: 699000, ProductId: 38, SizeId: 3, ColorId: 2},   // Canifa Long Sleeve, size M, White
// 	{Sku: 2114, Amount: 1, Price: 699000, ProductId: 38, SizeId: 4, ColorId: 3},   // Canifa Long Sleeve, size L, Red
// 	{Sku: 2115, Amount: 2, Price: 599000, ProductId: 39, SizeId: 3, ColorId: 4},   // Owen Long Sleeve, size M, Lime
// 	{Sku: 2116, Amount: 1, Price: 599000, ProductId: 39, SizeId: 4, ColorId: 5},   // Owen Long Sleeve, size L, Blue
// 	{Sku: 2117, Amount: 1, Price: 599000, ProductId: 39, SizeId: 5, ColorId: 6},   // Owen Long Sleeve, size XL, Yellow
// 	{Sku: 2118, Amount: 2, Price: 699000, ProductId: 40, SizeId: 2, ColorId: 7},   // Biluxury Long Sleeve, size S, Cyan
// 	{Sku: 2119, Amount: 1, Price: 699000, ProductId: 40, SizeId: 3, ColorId: 8},   // Biluxury Long Sleeve, size M, Magenta
// 	{Sku: 2120, Amount: 1, Price: 699000, ProductId: 40, SizeId: 4, ColorId: 9},   // Biluxury Long Sleeve, size L, Silver
// 	{Sku: 2121, Amount: 2, Price: 599000, ProductId: 41, SizeId: 3, ColorId: 10},  // Yody Long Sleeve, size M, Gray
// 	{Sku: 2122, Amount: 1, Price: 599000, ProductId: 41, SizeId: 4, ColorId: 11},  // Yody Long Sleeve, size L, Maroon
// 	{Sku: 2123, Amount: 1, Price: 599000, ProductId: 41, SizeId: 5, ColorId: 12},  // Yody Long Sleeve, size XL, Olive
// 	{Sku: 2124, Amount: 2, Price: 699000, ProductId: 42, SizeId: 2, ColorId: 1},   // Coolmate Long Sleeve, size S, Black
// 	{Sku: 2125, Amount: 1, Price: 699000, ProductId: 42, SizeId: 3, ColorId: 2},   // Coolmate Long Sleeve, size M, White
// 	{Sku: 2126, Amount: 1, Price: 699000, ProductId: 42, SizeId: 4, ColorId: 3},   // Coolmate Long Sleeve, size L, Red
// 	{Sku: 2127, Amount: 2, Price: 599000, ProductId: 43, SizeId: 3, ColorId: 4},   // Routine Short Sleeve, size M, Lime
// 	{Sku: 2128, Amount: 1, Price: 599000, ProductId: 43, SizeId: 4, ColorId: 5},   // Routine Short Sleeve, size L, Blue
// 	{Sku: 2129, Amount: 1, Price: 599000, ProductId: 43, SizeId: 5, ColorId: 6},   // Routine Short Sleeve, size XL, Yellow
// 	{Sku: 2130, Amount: 2, Price: 699000, ProductId: 44, SizeId: 2, ColorId: 7},   // Canifa Short Sleeve, size S, Cyan
// 	{Sku: 2131, Amount: 1, Price: 699000, ProductId: 44, SizeId: 3, ColorId: 8},   // Canifa Short Sleeve, size M, Magenta
// 	{Sku: 2132, Amount: 1, Price: 699000, ProductId: 44, SizeId: 4, ColorId: 9},   // Canifa Short Sleeve, size L, Silver
// 	{Sku: 2133, Amount: 2, Price: 599000, ProductId: 45, SizeId: 3, ColorId: 10},  // Owen Short Sleeve, size M, Gray
// 	{Sku: 2134, Amount: 1, Price: 599000, ProductId: 45, SizeId: 4, ColorId: 11},  // Owen Short Sleeve, size L, Maroon
// 	{Sku: 2135, Amount: 1, Price: 599000, ProductId: 45, SizeId: 5, ColorId: 12},  // Owen Short Sleeve, size XL, Olive
// 	{Sku: 2136, Amount: 2, Price: 699000, ProductId: 46, SizeId: 2, ColorId: 1},   // Biluxury Short Sleeve, size S, Black
// 	{Sku: 2137, Amount: 1, Price: 699000, ProductId: 46, SizeId: 3, ColorId: 2},   // Biluxury Short Sleeve, size M, White
// 	{Sku: 2138, Amount: 1, Price: 699000, ProductId: 46, SizeId: 4, ColorId: 3},   // Biluxury Short Sleeve, size L, Red
// 	{Sku: 2139, Amount: 2, Price: 599000, ProductId: 47, SizeId: 3, ColorId: 4},   // Yody Short Sleeve, size M, Lime
// 	{Sku: 2140, Amount: 1, Price: 599000, ProductId: 47, SizeId: 4, ColorId: 5},   // Yody Short Sleeve, size L, Blue
// 	{Sku: 2141, Amount: 1, Price: 599000, ProductId: 47, SizeId: 5, ColorId: 6},   // Yody Short Sleeve, size XL, Yellow
// 	{Sku: 2142, Amount: 2, Price: 699000, ProductId: 48, SizeId: 2, ColorId: 7},   // Coolmate Short Sleeve, size S, Cyan
// 	{Sku: 2143, Amount: 1, Price: 699000, ProductId: 48, SizeId: 3, ColorId: 8},   // Coolmate Short Sleeve, size M, Magenta
// 	{Sku: 2144, Amount: 1, Price: 699000, ProductId: 48, SizeId: 4, ColorId: 9},   // Coolmate Short Sleeve, size L, Silver
// 	{Sku: 2145, Amount: 2, Price: 599000, ProductId: 49, SizeId: 3, ColorId: 10},  // Routine Mandarin, size M, Gray
// 	{Sku: 2146, Amount: 1, Price: 599000, ProductId: 49, SizeId: 4, ColorId: 11},  // Routine Mandarin, size L, Maroon
// 	{Sku: 2147, Amount: 1, Price: 599000, ProductId: 49, SizeId: 5, ColorId: 12},  // Routine Mandarin, size XL, Olive
// 	{Sku: 2148, Amount: 2, Price: 699000, ProductId: 50, SizeId: 2, ColorId: 1},   // Canifa Mandarin, size S, Black
// 	{Sku: 2149, Amount: 1, Price: 699000, ProductId: 50, SizeId: 3, ColorId: 2},   // Canifa Mandarin, size M, White
// 	{Sku: 2150, Amount: 1, Price: 699000, ProductId: 50, SizeId: 4, ColorId: 3},   // Canifa Mandarin, size L, Red
// }
