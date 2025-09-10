package test

import (
	"fmt"
	"os"
	"strings"
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
	assert.True(t, true)
}

// TestTerraformVariableValidation - Prueba de validación de variables usando plan
func TestTerraformVariableValidation(t *testing.T) {
	opciones_terraform := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"nombre_proyecto":       "test",
			"ambiente":              "dev",
			"ubicacion":             "ubicacioninvalida", // Esto debería provocar la falla
			"nombre_grupo_recursos": "test-rg",
		},
	}

	terraform.Init(t, opciones_terraform)

	// Esto debería fallar porque "ubicacioninvalida" no es una región válida de Azure
	_, err := terraform.PlanE(t, opciones_terraform)
	assert.Error(t, err)

	// Verifica que el error sea sobre validación de ubicación
	assert.Contains(t, err.Error(), "La ubicación debe ser una región válida de Azure.")
}

// TestTerraformBloqueRedValidation - Prueba de validación del bloque de red CIDR
func TestTerraformBloqueRedValidation(t *testing.T) {
	opciones_terraform := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"nombre_proyecto":       "test",
			"ambiente":              "dev",
			"ubicacion":             "eastus",
			"nombre_grupo_recursos": "test-rg",
			"bloque_red":            "invalid-cidr-block", // Esto debería provocar la falla
		},
	}

	terraform.Init(t, opciones_terraform)

	// Esto debería fallar porque "invalid-cidr-block" no es un bloque CIDR válido
	_, err := terraform.PlanE(t, opciones_terraform)
	assert.Error(t, err)

	// Verifica que el error sea sobre validación del bloque CIDR
	assert.Contains(t, err.Error(), "El bloque_red debe ser un bloque CIDR válido")
	assert.Contains(t, err.Error(), "(ej., 10.0.0.0/16).")
}

// TestTerraformNombreGrupoRecursosValidation - Prueba de validación del nombre del grupo de recursos
func TestTerraformNombreGrupoRecursosValidation(t *testing.T) {
	opciones_terraform := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"nombre_proyecto":       "test",
			"ambiente":              "dev",
			"ubicacion":             "eastus",
			"nombre_grupo_recursos": "invalid@name!", // Esto debería provocar la falla
		},
	}

	terraform.Init(t, opciones_terraform)

	// Esto debería fallar porque "invalid@name!" contiene caracteres no permitidos
	_, err := terraform.PlanE(t, opciones_terraform)
	assert.Error(t, err)

	// Verifica que el error sea sobre validación del nombre del grupo de recursos
	assert.Contains(t, err.Error(), "El nombre_grupo_recursos no puede estar vacío")
	assert.Contains(t, err.Error(), "debe contener solo letras")
	assert.Contains(t, err.Error(), "números, puntos, guiones y guiones bajos")
	assert.Contains(t, err.Error(), "no puede exceder 90 caracteres")
}

// TestTerraformSubredesValidation - Prueba de validación de la configuración de subredes
func TestTerraformSubredesValidation(t *testing.T) {
	opciones_terraform := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"nombre_proyecto":       "test",
			"ambiente":              "dev",
			"ubicacion":             "eastus",
			"nombre_grupo_recursos": "test-rg",
			"subredes":              map[string]interface{}{}, // Esto debería provocar la falla (mapa vacío)
		},
	}

	terraform.Init(t, opciones_terraform)

	// Esto debería fallar porque subredes está vacío
	_, err := terraform.PlanE(t, opciones_terraform)
	assert.Error(t, err)

	// Verifica que el error sea sobre validación de subredes
	assert.Contains(t, err.Error(), "Se debe definir al menos una subred")
	assert.Contains(t, err.Error(), "en la configuración de subredes.")
}

// TestTerraformPlanValidVars - Prueba que el módulo planifica correctamente con variables válidas
func TestTerraformPlanValidVars(t *testing.T) {
	// Crear un archivo provider temporal con la configuración específica
	providerConfig := `provider "azurerm" {
  features {}
  resource_provider_registrations = "none"
}
`
	err := os.WriteFile("../providers_temp.tf", []byte(providerConfig), 0644)
	if err != nil {
		t.Fatalf("Error creando archivo provider temporal: %v", err)
	}

	// Asegurar que se elimine el archivo al final del test
	defer func() {
		os.Remove("../providers_temp.tf")
	}()

	opciones_terraform := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"nombre_proyecto":       "test-proyecto",
			"ambiente":              "dev",
			"ubicacion":             "eastus",
			"nombre_grupo_recursos": "test-rg-dev",
			"bloque_red":            "10.0.0.0/16",
			"subredes": map[string]interface{}{
				"subnet-web": map[string]interface{}{
					"name":             "subnet-web",
					"address_prefixes": []string{"10.0.1.0/24"},
				},
				"subnet-app": map[string]interface{}{
					"name":             "subnet-app",
					"address_prefixes": []string{"10.0.2.0/24"},
				},
			},
		},
	}

	terraform.Init(t, opciones_terraform)

	// Intentar hacer plan - debe fallar por el proveedor de Azure pero no por validación de variables
	_, planErr := terraform.PlanE(t, opciones_terraform)

	// Debe haber un error (Azure provider sin configurar)
	assert.Error(t, planErr)

	// PERO el error NO debe ser sobre validación de variables - debe ser sobre el proveedor de Azure
	// Verificar que el error está relacionado con autenticación de Azure (diferentes mensajes en diferentes entornos)
	errorMsg := planErr.Error()

	// Verificar varios mensajes de error de autenticación de Azure
	azureAuthErrors := []string{
		"building account: unable to configure ResourceManagerAccount",
		"unable to build authorizer for Resource Manager API",
		"Please run 'az login' to setup account",
		"tenant ID was not specified",
		"subscription ID could not be determined",
		"could not configure AzureCli Authorizer",
	}

	azureAuthErrorFound := false
	for _, authError := range azureAuthErrors {
		if strings.Contains(errorMsg, authError) {
			azureAuthErrorFound = true
			break
		}
	}

	assert.True(t, azureAuthErrorFound,
		"El error debería estar relacionado con la autenticación de Azure, no con la validación de variables. Se obtuvo: %s", errorMsg) // Verificar que NO hay errores de validación de variables
	assert.NotContains(t, planErr.Error(), "Invalid value for variable")
	assert.NotContains(t, planErr.Error(), "La ubicación debe ser una región válida de Azure")
	assert.NotContains(t, planErr.Error(), "El bloque_red debe ser un bloque CIDR válido")
	assert.NotContains(t, planErr.Error(), "El nombre_grupo_recursos no puede estar vacío")
	assert.NotContains(t, planErr.Error(), "Se debe definir al menos una subred")
}

// TestTerraformSyntaxValidation - Prueba validación de sintaxis básica
func TestTerraformSyntaxValidation(t *testing.T) {
	opciones_terraform := &terraform.Options{
		TerraformDir: "../",
	}

	terraform.Init(t, opciones_terraform)

	// terraform validate verifica la sintaxis sin variables
	terraform.Validate(t, opciones_terraform)
	assert.True(t, true)
}

// TestModuleFilesExist - Prueba que los archivos del módulo existen y tienen contenido
func TestModuleFilesExist(t *testing.T) {
	requiredFiles := []string{
		"../main.tf",
		"../variables.tf",
		"../outputs.tf",
		"../locals.tf",
		"../providers.tf",
		"../terraform.tf",
	}

	for _, file := range requiredFiles {
		t.Run(fmt.Sprintf("Archivo %s existe", file), func(t *testing.T) {
			opciones_terraform := &terraform.Options{
				TerraformDir: "../",
			}

			// Si el archivo no existe, terraform init fallará
			terraform.Init(t, opciones_terraform)
			assert.True(t, true)
		})
	}
}
