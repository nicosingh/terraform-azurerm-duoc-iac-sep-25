package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestTerraformInit - Prueba básica para verificar que Terraform puede inicializarse correctamente
func TestTerraformInit(t *testing.T) {
	opciones_terraform := &terraform.Options{
		// Ruta al código de Terraform
		TerraformDir: "../",
	}

	// Ejecuta terraform init (descarga proveedores y módulos)
	terraform.Init(t, opciones_terraform)

	// Si llegamos aquí sin errores, la prueba fue exitosa
	assert.True(t, true, "Terraform init completado exitosamente")
}

// TestTerraformVariableValidation - Prueba de validación de variables usando plan
func TestTerraformVariableValidation(t *testing.T) {
	opciones_terraform := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"nombre_proyecto":       "test",
			"ambiente":              "dev",
			"ubicacion":             "ubicacioninvalida", // Esto debería fallar
			"nombre_grupo_recursos": "test-rg",
		},
	}

	terraform.Init(t, opciones_terraform)

	// Esto debería fallar porque "ubicacioninvalida" no es una región válida de Azure
	_, err := terraform.PlanE(t, opciones_terraform)
	assert.Error(t, err, "Debería fallar con ubicación inválida")

	// Verifica que el error sea sobre validación (no autenticación)
	assert.Contains(t, err.Error(), "región válida de Azure", "Debería ser un error de validación sobre regiones de Azure")
}
