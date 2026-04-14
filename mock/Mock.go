package mock

import (
	"WeFashionServer/infrastructure/database"
	"WeFashionServer/infrastructure/model"
)

var MockCategories = []model.Category{
	{Name: "Pants"},
	{Name: "Shirts"},
	{Name: "Shoes"},
	{Name: "Handbags"},
	{Name: "Backpacks"},
	{Name: "Hats"},
	{Name: "Underwear"},
}

var MockColors = []model.Color{
	{Rgb: "0,0,0"},       // Black
	{Rgb: "255,255,255"}, // White
	{Rgb: "255,0,0"},     // Red
	{Rgb: "0,255,0"},     // Lime
	{Rgb: "0,0,255"},     // Blue
	{Rgb: "255,255,0"},   // Yellow
	{Rgb: "0,255,255"},   // Cyan
	{Rgb: "255,0,255"},   // Magenta
	{Rgb: "192,192,192"}, // Silver
	{Rgb: "128,128,128"}, // Gray
	{Rgb: "128,0,0"},     // Maroon
	{Rgb: "128,128,0"},   // Olive
	{Rgb: "0,128,0"},     // Green
	{Rgb: "128,0,128"},   // Purple
	{Rgb: "0,128,128"},   // Teal
	{Rgb: "0,0,128"},     // Navy
	{Rgb: "255,165,0"},   // Orange
	{Rgb: "255,215,0"},   // Gold
	{Rgb: "173,216,230"}, // Light Blue
	{Rgb: "0,191,255"},   // Deep Sky Blue
	{Rgb: "135,206,250"}, // Light Sky Blue
	{Rgb: "70,130,180"},  // Steel Blue
	{Rgb: "176,224,230"}, // Powder Blue
	{Rgb: "240,248,255"}, // Alice Blue
	{Rgb: "245,245,220"}, // Beige
	{Rgb: "255,228,196"}, // Bisque
	{Rgb: "255,235,205"}, // Blanched Almond
	{Rgb: "0,0,139"},     // Dark Blue
	{Rgb: "0,139,139"},   // Dark Cyan
	{Rgb: "184,134,11"},  // Dark Golden Rod
	{Rgb: "169,169,169"}, // Dark Gray
	{Rgb: "0,100,0"},     // Dark Green
	{Rgb: "139,0,139"},   // Dark Magenta
	{Rgb: "85,107,47"},   // Dark Olive Green
	{Rgb: "255,140,0"},   // Dark Orange
	{Rgb: "153,50,204"},  // Dark Orchid
	{Rgb: "139,0,0"},     // Dark Red
	{Rgb: "233,150,122"}, // Dark Salmon
	{Rgb: "143,188,143"}, // Dark Sea Green
	{Rgb: "72,61,139"},   // Dark Slate Blue
	{Rgb: "47,79,79"},    // Dark Slate Gray
	{Rgb: "0,206,209"},   // Dark Turquoise
	{Rgb: "148,0,211"},   // Dark Violet
	{Rgb: "255,20,147"},  // Deep Pink
	{Rgb: "0,191,255"},   // Deep Sky Blue
	{Rgb: "105,105,105"}, // Dim Gray
	{Rgb: "30,144,255"},  // Dodger Blue
	{Rgb: "178,34,34"},   // Fire Brick
	{Rgb: "255,250,240"}, // Floral White
	{Rgb: "34,139,34"},   // Forest Green
	{Rgb: "255,0,255"},   // Fuchsia
	{Rgb: "220,220,220"}, // Gainsboro
	{Rgb: "248,248,255"}, // Ghost White
	{Rgb: "255,215,0"},   // Gold
	{Rgb: "218,165,32"},  // Golden Rod
	{Rgb: "128,128,128"}, // Gray
	{Rgb: "0,128,0"},     // Green
	{Rgb: "173,255,47"},  // Green Yellow
	{Rgb: "240,255,240"}, // Honey Dew
	{Rgb: "255,105,180"}, // Hot Pink
	{Rgb: "205,92,92"},   // Indian Red
	{Rgb: "75,0,130"},    // Indigo
	{Rgb: "255,255,240"}, // Ivory
	{Rgb: "240,230,140"}, // Khaki
	{Rgb: "230,230,250"}, // Lavender
	{Rgb: "255,240,245"}, // Lavender Blush
	{Rgb: "124,252,0"},   // Lawn Green
	{Rgb: "255,250,205"}, // Lemon Chiffon
	{Rgb: "173,216,230"}, // Light Blue
	{Rgb: "240,128,128"}, // Light Coral
	{Rgb: "224,255,255"}, // Light Cyan
	{Rgb: "250,250,210"}, // Light Golden Rod Yellow
	{Rgb: "211,211,211"}, // Light Gray
	{Rgb: "144,238,144"}, // Light Green
	{Rgb: "255,182,193"}, // Light Pink
	{Rgb: "255,160,122"}, // Light Salmon
	{Rgb: "32,178,170"},  // Light Sea Green
	{Rgb: "135,206,250"}, // Light Sky Blue
	{Rgb: "119,136,153"}, // Light Slate Gray
	{Rgb: "176,196,222"}, // Light Steel Blue
	{Rgb: "255,255,224"}, // Light Yellow
	{Rgb: "0,255,0"},     // Lime
	{Rgb: "50,205,50"},   // Lime Green
	{Rgb: "250,240,230"}, // Linen
}

var MockSizes = []model.Size{
	// Quần áo, tất, mũ, túi
	{Name: "XS"},
	{Name: "S"},
	{Name: "M"},
	{Name: "L"},
	{Name: "XL"},
	{Name: "XXL"},
	{Name: "XXXL"},
	// Giày
	{Name: "36"},
	{Name: "37"},
	{Name: "38"},
	{Name: "39"},
	{Name: "40"},
	{Name: "41"},
	{Name: "42"},
	{Name: "43"},
	{Name: "44"},
	// Đồ lót
	{Name: "50-60kg"},
	{Name: "60-70kg"},
	{Name: "70-80kg"},
	{Name: "80-90kg"},
}

// done
func InsertMockCategories() error {
	return database.DB.Create(&MockCategories).Error
}

func InsertMockColors() error {
	return database.DB.Create(&MockColors).Error
}

func InsertMockSizes() error {
	return database.DB.Create(&MockSizes).Error
}
