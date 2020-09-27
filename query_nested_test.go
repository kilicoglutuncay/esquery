package esquery

import "testing"

func TestNested(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"nested query with bool must",
			Nested("nested_path", Bool().Must(Term("tag", "tech"))).ScoreMode("avg").IgnoreUnmapped(true),
			map[string]interface{}{
				"nested": map[string]interface{}{
					"path": "nested_path",
					"query": map[string]interface{}{
						"bool": map[string]interface{}{
							"must": []map[string]interface{}{
								{
									"term": map[string]interface{}{
										"tag": map[string]interface{}{
											"value": "tech",
										},
									},
								},
							},
						},
					},
					"score_mode":      "avg",
					"ignore_unmapped": true,
				},
			},
		},


		{
			"multi-level nested query with bool must",
			Nested("driver",
				Nested("driver.vehicle",
					Bool().Must(
						Term("driver.vehicle.make", "Powell Motors"),
						Term("driver.vehicle.model", "Canyonero"),
					),
				).IgnoreUnmapped(true),
			),
			map[string]interface{}{
				"nested": map[string]interface{}{
					"path": "driver",
					"query": map[string]interface{}{
						"nested": map[string]interface{}{
							"path": "driver.vehicle",
							"query": map[string]interface{}{
								"bool": map[string]interface{}{
									"must": []map[string]interface{}{
										{
											"term": map[string]interface{}{

												"driver.vehicle.make": map[string]interface{}{
													"value": "Powell Motors",
												},
											},
										},
										{
											"term": map[string]interface{}{

												"driver.vehicle.model": map[string]interface{}{
													"value": "Canyonero",
												},
											},
										},
									},
								},
							},
							"ignore_unmapped": true,
						},
					},
				},
			},
		},
	})
}
