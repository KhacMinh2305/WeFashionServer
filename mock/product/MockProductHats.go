package product

// import (
// 	"WeFashionServer/infrastructure/model"
// )

// // Mock data for Hats category (mũ lưỡi trai, bucket, snapback, fedora, ...)
// var MockHatsProducts = []model.Product{
// 	// 10 sản phẩm mũ thực tế, đa dạng thương hiệu, hình ảnh, mô tả, rating, sold, liked, ShopId khác nhau
// 	// CategoryId = 6 (Hats)
// 	{Id: 1, Name: "New Era 59FIFTY", ImageUrl: "https://neweracap.com/59fifty.jpg", Description: "Mũ snapback New Era 59FIFTY, phong cách thể thao Mỹ.", Rating: 4.8, SoldAmount: 900, LikedAmount: 700, CategoryId: 6, ShopId: 1},
// 	{Id: 2, Name: "Nike Heritage86 Cap", ImageUrl: "https://nike.com/heritage86.jpg", Description: "Nike Heritage86, mũ lưỡi trai cổ điển, chất liệu cotton.", Rating: 4.7, SoldAmount: 850, LikedAmount: 650, CategoryId: 6, ShopId: 2},
// 	{Id: 3, Name: "Adidas Originals Bucket Hat", ImageUrl: "https://adidas.com/bucket-hat.jpg", Description: "Adidas Originals Bucket Hat, phong cách streetwear, logo thêu.", Rating: 4.6, SoldAmount: 800, LikedAmount: 600, CategoryId: 6, ShopId: 3},
// 	{Id: 4, Name: "Puma Classic Cap", ImageUrl: "https://puma.com/classic-cap.jpg", Description: "Puma Classic Cap, mũ lưỡi trai đơn giản, logo Puma.", Rating: 4.5, SoldAmount: 750, LikedAmount: 550, CategoryId: 6, ShopId: 4},
// 	{Id: 5, Name: "MLB Korea LA Dodgers Cap", ImageUrl: "https://mlbkorea.com/la-dodgers.jpg", Description: "MLB Korea LA Dodgers, mũ lưỡi trai phong cách Hàn Quốc.", Rating: 4.8, SoldAmount: 700, LikedAmount: 500, CategoryId: 6, ShopId: 5},
// 	{Id: 6, Name: "Kangol Bermuda Casual", ImageUrl: "https://kangol.com/bermuda-casual.jpg", Description: "Kangol Bermuda Casual, mũ bucket, chất liệu mềm mại.", Rating: 4.7, SoldAmount: 650, LikedAmount: 480, CategoryId: 6, ShopId: 6},
// 	{Id: 7, Name: "Lacoste Cotton Cap", ImageUrl: "https://lacoste.com/cotton-cap.jpg", Description: "Lacoste Cotton Cap, mũ lưỡi trai, logo cá sấu.", Rating: 4.6, SoldAmount: 600, LikedAmount: 450, CategoryId: 6, ShopId: 7},
// 	{Id: 8, Name: "Supreme Camp Cap", ImageUrl: "https://supremenewyork.com/camp-cap.jpg", Description: "Supreme Camp Cap, mũ streetwear, logo Supreme nổi bật.", Rating: 4.9, SoldAmount: 550, LikedAmount: 420, CategoryId: 6, ShopId: 8},
// 	{Id: 9, Name: "Gucci GG Canvas Baseball Cap", ImageUrl: "https://gucci.com/gg-canvas-cap.jpg", Description: "Gucci GG Canvas Cap, mũ lưỡi trai cao cấp, họa tiết GG.", Rating: 5.0, SoldAmount: 500, LikedAmount: 400, CategoryId: 6, ShopId: 9},
// 	{Id: 10, Name: "Levi's Vintage Trucker Hat", ImageUrl: "https://levis.com/vintage-trucker.jpg", Description: "Levi's Vintage Trucker Hat, mũ lưỡi trai phong cách retro.", Rating: 4.7, SoldAmount: 480, LikedAmount: 380, CategoryId: 6, ShopId: 10},
// }

// // Mock data for ProductVariants (biến thể sản phẩm) cho Hats
// var MockHatsVariants = []model.ProductVariant{
// 	// SKU cho từng sản phẩm Hats, mỗi sản phẩm 3-5 biến thể (màu khác nhau, size tự do)
// 	{Sku: 6001, Amount: 20, Price: 990000, ProductId: 1, SizeId: 2, ColorId: 1},  // New Era 59FIFTY, S, Black
// 	{Sku: 6002, Amount: 15, Price: 990000, ProductId: 1, SizeId: 2, ColorId: 2},  // New Era 59FIFTY, S, White
// 	{Sku: 6003, Amount: 10, Price: 990000, ProductId: 1, SizeId: 2, ColorId: 3},  // New Era 59FIFTY, S, Red
// 	{Sku: 6004, Amount: 18, Price: 590000, ProductId: 2, SizeId: 2, ColorId: 4},  // Nike Heritage86, S, Lime
// 	{Sku: 6005, Amount: 12, Price: 590000, ProductId: 2, SizeId: 2, ColorId: 5},  // Nike Heritage86, S, Blue
// 	{Sku: 6006, Amount: 8, Price: 590000, ProductId: 2, SizeId: 2, ColorId: 6},   // Nike Heritage86, S, Yellow
// 	{Sku: 6007, Amount: 16, Price: 690000, ProductId: 3, SizeId: 2, ColorId: 7},  // Adidas Bucket, S, Cyan
// 	{Sku: 6008, Amount: 10, Price: 690000, ProductId: 3, SizeId: 2, ColorId: 8},  // Adidas Bucket, S, Magenta
// 	{Sku: 6009, Amount: 7, Price: 690000, ProductId: 3, SizeId: 2, ColorId: 9},   // Adidas Bucket, S, Silver
// 	{Sku: 6010, Amount: 14, Price: 490000, ProductId: 4, SizeId: 2, ColorId: 10}, // Puma Classic, S, Gray
// 	{Sku: 6011, Amount: 9, Price: 490000, ProductId: 4, SizeId: 2, ColorId: 1},   // Puma Classic, S, Black
// 	{Sku: 6012, Amount: 6, Price: 490000, ProductId: 4, SizeId: 2, ColorId: 2},   // Puma Classic, S, White
// 	{Sku: 6013, Amount: 13, Price: 890000, ProductId: 5, SizeId: 2, ColorId: 3},  // MLB Korea LA Dodgers, S, Red
// 	{Sku: 6014, Amount: 8, Price: 890000, ProductId: 5, SizeId: 2, ColorId: 4},   // MLB Korea LA Dodgers, S, Lime
// 	{Sku: 6015, Amount: 5, Price: 890000, ProductId: 5, SizeId: 2, ColorId: 5},   // MLB Korea LA Dodgers, S, Blue
// 	{Sku: 6016, Amount: 12, Price: 790000, ProductId: 6, SizeId: 2, ColorId: 6},  // Kangol Bermuda, S, Yellow
// 	{Sku: 6017, Amount: 7, Price: 790000, ProductId: 6, SizeId: 2, ColorId: 7},   // Kangol Bermuda, S, Cyan
// 	{Sku: 6018, Amount: 5, Price: 790000, ProductId: 6, SizeId: 2, ColorId: 8},   // Kangol Bermuda, S, Magenta
// 	{Sku: 6019, Amount: 11, Price: 990000, ProductId: 7, SizeId: 2, ColorId: 9},  // Lacoste Cotton Cap, S, Silver
// 	{Sku: 6020, Amount: 6, Price: 990000, ProductId: 7, SizeId: 2, ColorId: 10},  // Lacoste Cotton Cap, S, Gray
// 	{Sku: 6021, Amount: 4, Price: 990000, ProductId: 7, SizeId: 2, ColorId: 1},   // Lacoste Cotton Cap, S, Black
// 	{Sku: 6022, Amount: 10, Price: 1590000, ProductId: 8, SizeId: 2, ColorId: 2}, // Supreme Camp Cap, S, White
// 	{Sku: 6023, Amount: 7, Price: 1590000, ProductId: 8, SizeId: 2, ColorId: 3},  // Supreme Camp Cap, S, Red
// 	{Sku: 6024, Amount: 5, Price: 1590000, ProductId: 8, SizeId: 2, ColorId: 4},  // Supreme Camp Cap, S, Lime
// 	{Sku: 6025, Amount: 9, Price: 4990000, ProductId: 9, SizeId: 2, ColorId: 5},  // Gucci GG Canvas, S, Blue
// 	{Sku: 6026, Amount: 6, Price: 4990000, ProductId: 9, SizeId: 2, ColorId: 6},  // Gucci GG Canvas, S, Yellow
// 	{Sku: 6027, Amount: 4, Price: 4990000, ProductId: 9, SizeId: 2, ColorId: 7},  // Gucci GG Canvas, S, Cyan
// 	{Sku: 6028, Amount: 8, Price: 690000, ProductId: 10, SizeId: 2, ColorId: 8},  // Levi's Trucker, S, Magenta
// 	{Sku: 6029, Amount: 5, Price: 690000, ProductId: 10, SizeId: 2, ColorId: 9},  // Levi's Trucker, S, Silver
// 	{Sku: 6030, Amount: 3, Price: 690000, ProductId: 10, SizeId: 2, ColorId: 10}, // Levi's Trucker, S, Gray
// }
