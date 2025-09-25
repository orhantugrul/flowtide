package validator

import (
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("valid struct passes validation", func(t *testing.T) {
		type User struct {
			Name  string `validate:"required,min=2,max=50"`
			Email string `validate:"required,email"`
			Age   int    `validate:"min=18,max=120"`
		}

		user := &User{
			Name:  "John Doe",
			Email: "john@example.com",
			Age:   25,
		}

		if err := Validate(user); err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("missing required field fails validation", func(t *testing.T) {
		type User struct {
			Name  string `validate:"required"`
			Email string `validate:"required"`
		}

		user := &User{Name: "John"}

		if err := Validate(user); err == nil {
			t.Error("expected validation error for missing required field")
		}
	})

	t.Run("invalid email format fails validation", func(t *testing.T) {
		type User struct {
			Email string `validate:"email"`
		}

		user := &User{
			Email: "invalid-email",
		}

		if err := Validate(user); err == nil {
			t.Error("expected validation error for invalid email format")
		}
	})

	t.Run("string length validation", func(t *testing.T) {
		type Product struct {
			Title       string `validate:"min=5,max=100"`
			Description string `validate:"max=500"`
		}

		t.Run("string too short fails", func(t *testing.T) {
			product := &Product{
				Title: "Hi", // Too short
			}

			if err := Validate(product); err == nil {
				t.Error("expected validation error for string too short")
			}
		})

		t.Run("string too long fails", func(t *testing.T) {
			product := &Product{
				Title: "This is a very long title that exceeds the maximum " +
					"allowed length of 100 characters and should fail validation",
			}

			if err := Validate(product); err == nil {
				t.Error("expected validation error for string too long")
			}
		})
	})

	t.Run("numeric range validation", func(t *testing.T) {
		type Config struct {
			Port     int     `validate:"min=1,max=65535"`
			Timeout  float64 `validate:"min=0.1,max=300.0"`
			Priority int     `validate:"oneof=1 2 3"`
		}

		t.Run("port out of range fails", func(t *testing.T) {
			config := &Config{
				Port: 99999, // Invalid port
			}

			if err := Validate(config); err == nil {
				t.Error("expected validation error for invalid port")
			}
		})

		t.Run("timeout out of range fails", func(t *testing.T) {
			config := &Config{
				Timeout: 500.0, // Too high
			}

			if err := Validate(config); err == nil {
				t.Error("expected validation error for timeout out of range")
			}
		})

		t.Run("invalid priority fails", func(t *testing.T) {
			config := &Config{
				Priority: 5, // Not in allowed values
			}

			if err := Validate(config); err == nil {
				t.Error("expected validation error for invalid priority")
			}
		})
	})

	t.Run("nested struct validation", func(t *testing.T) {
		type Address struct {
			Street string `validate:"required"`
			City   string `validate:"required"`
		}

		type Person struct {
			Name    string  `validate:"required"`
			Address Address `validate:"required"`
		}

		t.Run("valid nested struct passes", func(t *testing.T) {
			person := &Person{
				Name: "Jane",
				Address: Address{
					Street: "123 Main St",
					City:   "New York",
				},
			}

			if err := Validate(person); err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})

		t.Run("invalid nested struct fails", func(t *testing.T) {
			person := &Person{
				Name: "Jane",
				Address: Address{
					Street: "", // Missing required field
					City:   "New York",
				},
			}

			if err := Validate(person); err == nil {
				t.Error("expected validation error for invalid nested struct")
			}
		})
	})

	t.Run("slice validation", func(t *testing.T) {
		type Team struct {
			Members []string `validate:"required,min=1,dive,required"`
		}

		t.Run("empty slice fails", func(t *testing.T) {
			team := &Team{
				Members: []string{}, // Empty slice
			}

			if err := Validate(team); err == nil {
				t.Error("expected validation error for empty slice")
			}
		})

		t.Run("slice with empty strings fails", func(t *testing.T) {
			team := &Team{
				Members: []string{"John", "", "Jane"}, // Empty string in slice
			}

			if err := Validate(team); err == nil {
				t.Error("expected validation error for empty string in slice")
			}
		})

		t.Run("valid slice passes", func(t *testing.T) {
			team := &Team{
				Members: []string{"John", "Jane", "Bob"},
			}

			if err := Validate(team); err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	})

	t.Run("custom validation tags", func(t *testing.T) {
		type Order struct {
			Status   string `validate:"oneof=pending processing shipped delivered"`
			Payment  string `validate:"oneof=cash card bank_transfer"`
			Currency string `validate:"len=3"`
		}

		t.Run("invalid status fails", func(t *testing.T) {
			order := &Order{
				Status: "invalid_status",
			}

			err := Validate(order)
			if err == nil {
				t.Error("expected validation error for invalid status")
			}
		})

		t.Run("invalid payment method fails", func(t *testing.T) {
			order := &Order{
				Payment: "crypto",
			}

			if err := Validate(order); err == nil {
				t.Error("expected validation error for invalid payment method")
			}
		})

		t.Run("invalid currency length fails", func(t *testing.T) {
			order := &Order{
				Currency: "USDX", // Too long
			}

			if err := Validate(order); err == nil {
				t.Error("expected validation error for invalid currency length")
			}
		})
	})

	t.Run("nil pointer handling", func(t *testing.T) {
		type Data struct {
			Value string `validate:"required"`
		}

		var data *Data = nil

		if err := Validate(data); err == nil {
			t.Error("expected validation error for nil pointer")
		}
	})

	t.Run("empty struct passes validation", func(t *testing.T) {
		type Empty struct{}

		empty := &Empty{}

		if err := Validate(empty); err != nil {
			t.Errorf("expected no error for empty struct, got %v", err)
		}
	})

	t.Run("struct with no validation tags passes", func(t *testing.T) {
		type Simple struct {
			Name  string
			Value int
		}

		simple := &Simple{
			Name:  "test",
			Value: 42,
		}

		if err := Validate(simple); err != nil {
			t.Errorf("expected no error for struct without "+
				"validation tags, got %v", err)
		}
	})
}
