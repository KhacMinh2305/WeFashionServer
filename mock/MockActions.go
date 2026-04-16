package mock

// // InsertMockHandbagsProducts inserts all products in product.MockHandbagsProducts into the database
// func InsertMockHandbagsProducts() error {
// 	db := database.DB
// 	if db == nil {
// 		return fmt.Errorf("database connection is not initialized")
// 	}
// 	for _, p := range product.MockHandbagsProducts {
// 		var existing model.Product
// 		err := db.Where("id = ?", p.Id).First(&existing).Error
// 		if err == nil {
// 			// Exists, update
// 			if err := db.Model(&existing).Updates(p).Error; err != nil {
// 				return fmt.Errorf("failed to update product id %d: %w", p.Id, err)
// 			}
// 		} else {
// 			// Not found, create
// 			if err := db.Create(&p).Error; err != nil {
// 				return fmt.Errorf("failed to insert product id %d: %w", p.Id, err)
// 			}
// 		}
// 	}
// 	return nil
// }

// // InsertMockHandbagsVariants inserts all product variants in product.MockHandbagsVariants into the database
// func InsertMockHandbagsVariants() error {
// 	db := database.DB
// 	if db == nil {
// 		return fmt.Errorf("database connection is not initialized")
// 	}
// 	for _, v := range product.MockHandbagsVariants {
// 		var existing model.ProductVariant
// 		err := db.Where("sku = ?", v.Sku).First(&existing).Error
// 		if err == nil {
// 			// Exists, update
// 			if err := db.Model(&existing).Updates(v).Error; err != nil {
// 				return fmt.Errorf("failed to update variant sku %d: %w", v.Sku, err)
// 			}
// 		} else {
// 			// Not found, create
// 			if err := db.Create(&v).Error; err != nil {
// 				return fmt.Errorf("failed to insert variant sku %d: %w", v.Sku, err)
// 			}
// 		}
// 	}
// 	return nil
// }

// func InsertMockCoupons() error {
// 	db := database.DB
// 	if db == nil {
// 		return fmt.Errorf("database connection is not initialized")
// 	}
// 	for _, coupon := range coupon.MockCouponList {
// 		// Insert or update by Id (if needed, can use db.Clauses for upsert)
// 		var existing model.Coupon
// 		err := db.Where("id = ?", coupon.Id).First(&existing).Error
// 		if err == nil {
// 			// Exists, update
// 			if err := db.Model(&existing).Updates(coupon).Error; err != nil {
// 				return fmt.Errorf("failed to update coupon id %d: %w", coupon.Id, err)
// 			}
// 		} else {
// 			// Not found, create
// 			if err := db.Create(&coupon).Error; err != nil {
// 				return fmt.Errorf("failed to insert coupon id %d: %w", coupon.Id, err)
// 			}
// 		}
// 	}
// 	return nil
// }

// func InsertMockShippers() error {
// 	db := database.DB
// 	if db == nil {
// 		return fmt.Errorf("database connection is not initialized")
// 	}
// 	for _, s := range shipper.MockShipperList {
// 		var existing model.Shipper
// 		err := db.Where("id = ?", s.Id).First(&existing).Error
// 		if err == nil {
// 			// Exists, update
// 			if err := db.Model(&existing).Updates(s).Error; err != nil {
// 				return fmt.Errorf("failed to update shipper id %d: %w", s.Id, err)
// 			}
// 		} else {
// 			// Not found, create
// 			if err := db.Create(&s).Error; err != nil {
// 				return fmt.Errorf("failed to insert shipper id %d: %w", s.Id, err)
// 			}
// 		}
// 	}
// 	return nil
// }

// func InsertMockShops() error {
// 	db := database.DB
// 	if db == nil {
// 		return fmt.Errorf("database connection is not initialized")
// 	}
// 	for _, s := range shop.MockShopList {
// 		var existing model.Shop
// 		err := db.Where("id = ?", s.Id).First(&existing).Error
// 		if err == nil {
// 			// Exists, update
// 			if err := db.Model(&existing).Updates(s).Error; err != nil {
// 				return fmt.Errorf("failed to update shop id %d: %w", s.Id, err)
// 			}
// 		} else {
// 			// Not found, create
// 			if err := db.Create(&s).Error; err != nil {
// 				return fmt.Errorf("failed to insert shop id %d: %w", s.Id, err)
// 			}
// 		}
// 	}
// 	return nil
// }

// // InsertMockBackpacksProducts inserts all products in product.MockBackpacksProducts into the database
// func InsertMockBackpacksProducts() error {
// 	db := database.DB
// 	if db == nil {
// 		return fmt.Errorf("database connection is not initialized")
// 	}
// 	for _, p := range product.MockBackpacksProducts {
// 		var existing model.Product
// 		err := db.Where("id = ?", p.Id).First(&existing).Error
// 		if err == nil {
// 			// Exists, update
// 			if err := db.Model(&existing).Updates(p).Error; err != nil {
// 				return fmt.Errorf("failed to update product id %d: %w", p.Id, err)
// 			}
// 		} else {
// 			// Not found, create
// 			if err := db.Create(&p).Error; err != nil {
// 				return fmt.Errorf("failed to insert product id %d: %w", p.Id, err)
// 			}
// 		}
// 	}
// 	return nil
// }

// // InsertMockBackpacksVariants inserts all product variants in product.MockBackpacksVariants into the database
// func InsertMockBackpacksVariants() error {
// 	db := database.DB
// 	if db == nil {
// 		return fmt.Errorf("database connection is not initialized")
// 	}
// 	for _, v := range product.MockBackpacksVariants {
// 		var existing model.ProductVariant
// 		err := db.Where("sku = ?", v.Sku).First(&existing).Error
// 		if err == nil {
// 			// Exists, update
// 			if err := db.Model(&existing).Updates(v).Error; err != nil {
// 				return fmt.Errorf("failed to update variant sku %d: %w", v.Sku, err)
// 			}
// 		} else {
// 			// Not found, create
// 			if err := db.Create(&v).Error; err != nil {
// 				return fmt.Errorf("failed to insert variant sku %d: %w", v.Sku, err)
// 			}
// 		}
// 	}
// 	return nil
// }

// // InsertMockHatsProducts inserts all products in product.MockHatsProducts into the database
// func InsertMockHatsProducts() error {
// 	db := database.DB
// 	if db == nil {
// 		return fmt.Errorf("database connection is not initialized")
// 	}
// 	for _, p := range product.MockHatsProducts {
// 		var existing model.Product
// 		err := db.Where("id = ?", p.Id).First(&existing).Error
// 		if err == nil {
// 			// Exists, update
// 			if err := db.Model(&existing).Updates(p).Error; err != nil {
// 				return fmt.Errorf("failed to update product id %d: %w", p.Id, err)
// 			}
// 		} else {
// 			// Not found, create
// 			if err := db.Create(&p).Error; err != nil {
// 				return fmt.Errorf("failed to insert product id %d: %w", p.Id, err)
// 			}
// 		}
// 	}
// 	return nil
// }

// // InsertMockHatsVariants inserts all product variants in product.MockHatsVariants into the database
// func InsertMockHatsVariants() error {
// 	db := database.DB
// 	if db == nil {
// 		return fmt.Errorf("database connection is not initialized")
// 	}
// 	for _, v := range product.MockHatsVariants {
// 		var existing model.ProductVariant
// 		err := db.Where("sku = ?", v.Sku).First(&existing).Error
// 		if err == nil {
// 			// Exists, update
// 			if err := db.Model(&existing).Updates(v).Error; err != nil {
// 				return fmt.Errorf("failed to update variant sku %d: %w", v.Sku, err)
// 			}
// 		} else {
// 			// Not found, create
// 			if err := db.Create(&v).Error; err != nil {
// 				return fmt.Errorf("failed to insert variant sku %d: %w", v.Sku, err)
// 			}
// 		}
// 	}
// 	return nil
// }

// func InsertMockPantsProducts() error {
// 	db := database.DB
// 	if db == nil {
// 		return fmt.Errorf("database connection is not initialized")
// 	}
// 	for _, p := range product.MockPantsProducts {
// 		var existing model.Product
// 		err := db.Where("id = ?", p.Id).First(&existing).Error
// 		if err == nil {
// 			// Exists, update
// 			if err := db.Model(&existing).Updates(p).Error; err != nil {
// 				return fmt.Errorf("failed to update product id %d: %w", p.Id, err)
// 			}
// 		} else {
// 			// Not found, create
// 			if err := db.Create(&p).Error; err != nil {
// 				return fmt.Errorf("failed to insert product id %d: %w", p.Id, err)
// 			}
// 		}
// 	}
// 	return nil
// }

// // InsertMockPantsVariants inserts all product variants in product.MockPantsVariants into the database
// func InsertMockPantsVariants() error {
// 	db := database.DB
// 	if db == nil {
// 		return fmt.Errorf("database connection is not initialized")
// 	}
// 	for _, v := range product.MockPantsVariants {
// 		var existing model.ProductVariant
// 		err := db.Where("sku = ?", v.Sku).First(&existing).Error
// 		if err == nil {
// 			// Exists, update
// 			if err := db.Model(&existing).Updates(v).Error; err != nil {
// 				return fmt.Errorf("failed to update variant sku %d: %w", v.Sku, err)
// 			}
// 		} else {
// 			// Not found, create
// 			if err := db.Create(&v).Error; err != nil {
// 				return fmt.Errorf("failed to insert variant sku %d: %w", v.Sku, err)
// 			}
// 		}
// 	}
// 	return nil
// }

// // InsertMockShirtsProducts inserts all products in product.MockShirtsProducts into the database
// func InsertMockShirtsProducts() error {
// 	db := database.DB
// 	if db == nil {
// 		return fmt.Errorf("database connection is not initialized")
// 	}
// 	for _, p := range product.MockShirtsProducts {
// 		var existing model.Product
// 		err := db.Where("id = ?", p.Id).First(&existing).Error
// 		if err == nil {
// 			// Exists, update
// 			if err := db.Model(&existing).Updates(p).Error; err != nil {
// 				return fmt.Errorf("failed to update product id %d: %w", p.Id, err)
// 			}
// 		} else {
// 			// Not found, create
// 			if err := db.Create(&p).Error; err != nil {
// 				return fmt.Errorf("failed to insert product id %d: %w", p.Id, err)
// 			}
// 		}
// 	}
// 	return nil
// }

// // InsertMockShirtsVariants inserts all product variants in product.MockShirtsVariants into the database
// func InsertMockShirtsVariants() error {
// 	db := database.DB
// 	if db == nil {
// 		return fmt.Errorf("database connection is not initialized")
// 	}
// 	for _, v := range product.MockShirtsVariants {
// 		var existing model.ProductVariant
// 		err := db.Where("sku = ?", v.Sku).First(&existing).Error
// 		if err == nil {
// 			// Exists, update
// 			if err := db.Model(&existing).Updates(v).Error; err != nil {
// 				return fmt.Errorf("failed to update variant sku %d: %w", v.Sku, err)
// 			}
// 		} else {
// 			// Not found, create
// 			if err := db.Create(&v).Error; err != nil {
// 				return fmt.Errorf("failed to insert variant sku %d: %w", v.Sku, err)
// 			}
// 		}
// 	}
// 	return nil
// }

// // InsertMockShoesProducts inserts all products in product.MockShoesProducts into the database
// func InsertMockShoesProducts() error {
// 	db := database.DB
// 	if db == nil {
// 		return fmt.Errorf("database connection is not initialized")
// 	}
// 	for _, p := range product.MockShoesProducts {
// 		var existing model.Product
// 		err := db.Where("id = ?", p.Id).First(&existing).Error
// 		if err == nil {
// 			// Exists, update
// 			if err := db.Model(&existing).Updates(p).Error; err != nil {
// 				return fmt.Errorf("failed to update product id %d: %w", p.Id, err)
// 			}
// 		} else {
// 			// Not found, create
// 			if err := db.Create(&p).Error; err != nil {
// 				return fmt.Errorf("failed to insert product id %d: %w", p.Id, err)
// 			}
// 		}
// 	}
// 	return nil
// }

// // InsertMockShoesVariants inserts all product variants in product.MockShoesVariants into the database
// func InsertMockShoesVariants() error {
// 	db := database.DB
// 	if db == nil {
// 		return fmt.Errorf("database connection is not initialized")
// 	}
// 	for _, v := range product.MockShoesVariants {
// 		var existing model.ProductVariant
// 		err := db.Where("sku = ?", v.Sku).First(&existing).Error
// 		if err == nil {
// 			// Exists, update
// 			if err := db.Model(&existing).Updates(v).Error; err != nil {
// 				return fmt.Errorf("failed to update variant sku %d: %w", v.Sku, err)
// 			}
// 		} else {
// 			// Not found, create
// 			if err := db.Create(&v).Error; err != nil {
// 				return fmt.Errorf("failed to insert variant sku %d: %w", v.Sku, err)
// 			}
// 		}
// 	}
// 	return nil
// }

// // InsertMockUnderwearsProducts inserts all products in product.MockUnderwearsProducts into the database
// func InsertMockUnderwearsProducts() error {
// 	db := database.DB
// 	if db == nil {
// 		return fmt.Errorf("database connection is not initialized")
// 	}
// 	for _, p := range product.MockUnderwearsProducts {
// 		var existing model.Product
// 		err := db.Where("id = ?", p.Id).First(&existing).Error
// 		if err == nil {
// 			// Exists, update
// 			if err := db.Model(&existing).Updates(p).Error; err != nil {
// 				return fmt.Errorf("failed to update product id %d: %w", p.Id, err)
// 			}
// 		} else {
// 			// Not found, create
// 			if err := db.Create(&p).Error; err != nil {
// 				return fmt.Errorf("failed to insert product id %d: %w", p.Id, err)
// 			}
// 		}
// 	}
// 	return nil
// }

// // InsertMockUnderwearsVariants inserts all product variants in product.MockUnderwearsVariants into the database
// func InsertMockUnderwearsVariants() error {
// 	db := database.DB
// 	if db == nil {
// 		return fmt.Errorf("database connection is not initialized")
// 	}
// 	for _, v := range product.MockUnderwearsVariants {
// 		var existing model.ProductVariant
// 		err := db.Where("sku = ?", v.Sku).First(&existing).Error
// 		if err == nil {
// 			// Exists, update
// 			if err := db.Model(&existing).Updates(v).Error; err != nil {
// 				return fmt.Errorf("failed to update variant sku %d: %w", v.Sku, err)
// 			}
// 		} else {
// 			// Not found, create
// 			if err := db.Create(&v).Error; err != nil {
// 				return fmt.Errorf("failed to insert variant sku %d: %w", v.Sku, err)
// 			}
// 		}
// 	}
// 	return nil
// }
