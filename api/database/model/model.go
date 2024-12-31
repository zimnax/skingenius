package model

type Concentration struct {
	Min float64
	Max float64
}

var concentrationMap = map[string]Concentration{
	"preservative": {0.01, 1},
	//"active_vitamins":       {0.01, 2}, // Individual actives may vary, not explicitly added
	//"active_minerals":       {0.0, 0.0}, // No range specified, default to 0
	//"active_peptides":       {0.0, 0.0}, // No range specified, default to 0
	"pH_adjuster":          {0.1, 3},
	"thickener":            {0.1, 5},
	"emollient":            {3, 20},
	"humectant":            {2, 10},
	"fragrance":            {0.01, 5},
	"emulsifier":           {1, 10},
	"colorant":             {0.01, 3},
	"solvent":              {50, 80},
	"chelating_agent":      {0.05, 0.5},
	"active_antioxidant":   {0.1, 5}, // Except Vitamin C: 5-20% (not explicitly added)
	"stabilizer":           {0.1, 3},
	"texture_enhancer":     {1, 10},
	"occlusive_agent":      {1, 15},
	"penetration_enhancer": {0.5, 5},
	"exfoliant":            {1, 10},
	"sunscreen":            {5, 25},
}
