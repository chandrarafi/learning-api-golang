# 📚 Penjelasan Detail Syntax dan Konsep Golang

## 1. Package dan Import

```go
package main  // Deklarasi package, main = entry point aplikasi

import (
    "fmt"                    // Package untuk formatting & printing
    "api-golang/config"      // Import package lokal dari project
    db "api-golang/database" // Import dengan alias "db"
)
```

**Penjelasan:**
- `package main` - Package khusus untuk executable program
- Import path lokal menggunakan nama module dari `go.mod`
- Alias (`db`) untuk memperpendek nama package

## 2. Struct dan Type

```go
type Config struct {
    DBHost     string  // Field dengan tipe string
    DBPort     string
    DBPassword string  // Field tidak exported jika huruf kecil
}

type User struct {
    ID    int    `json:"id"`     // Struct tag untuk JSON marshaling
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

**Penjelasan:**
- `type` - Keyword untuk mendefinisikan tipe baru
- `struct` - Tipe data composite yang menggabungkan beberapa field
- Struct tag (backtick) - Metadata untuk encoding/decoding
- Field dengan huruf kapital = exported (bisa diakses dari package lain)

## 3. Method Receiver

```go
// Method dengan pointer receiver
func (u *User) Validate() error {
    // u adalah pointer, bisa modify field
    u.Name = strings.TrimSpace(u.Name)
    return nil
}

// Method dengan value receiver
func (u User) GetName() string {
    // u adalah copy, tidak bisa modify original
    return u.Name
}
```

**Penjelasan:**
- `(u *User)` - Pointer receiver, bisa modify struct
- `(u User)` - Value receiver, hanya bisa read
- Gunakan pointer receiver untuk:
  - Modify struct
  - Struct berukuran besar (efisiensi memory)

## 4. Interface

```go
// Definisi interface
type UserRepository interface {
    Create(ctx context.Context, user *models.User) error
    GetAll(ctx context.Context) ([]models.User, error)
}

// Implementasi interface (implicit)
type UserRepo struct {
    DB *pgxpool.Pool
}

func (r *UserRepo) Create(ctx context.Context, user *models.User) error {
    // Implementation
}

func (r *UserRepo) GetAll(ctx context.Context) ([]models.User, error) {
    // Implementation
}
```

**Penjelasan:**
- Interface = kontrak method yang harus diimplementasikan
- Go menggunakan implicit implementation (tidak perlu keyword `implements`)
- Berguna untuk dependency injection dan testing (mock)

## 5. Error Handling

```go
// Return error
func Connect() (*pgxpool.Pool, error) {
    pool, err := pgxpool.New(ctx, dsn)
    if err != nil {
        return nil, fmt.Errorf("buat pool gagal: %w", err)
    }
    return pool, nil
}

// Check error
pool, err := db.Connect(cfg)
if err != nil {
    log.Fatalf("Gagal: %v", err)  // %v = print error message
}
```

**Penjelasan:**
- Go tidak punya try-catch, menggunakan explicit error return
- `error` adalah interface built-in
- `%w` - Wrap error (Go 1.13+) untuk error chain
- `%v` - Print value (termasuk error message)
- `errors.New()` - Buat error baru
- `fmt.Errorf()` - Format error dengan context

## 6. Context

```go
// Buat context dengan timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()  // Harus dipanggil untuk cleanup

// Gunakan context di database operation
err := pool.Ping(ctx)

// Ambil context dari HTTP request
ctx := r.Context()
users, err := h.Service.GetAllUsers(ctx)
```

**Penjelasan:**
- `context.Context` - Membawa deadline, cancellation signal, dan values
- `context.Background()` - Root context
- `context.WithTimeout()` - Context dengan timeout otomatis
- `defer cancel()` - Cleanup resources saat function selesai
- Context digunakan untuk:
  - Timeout operation
  - Cancellation propagation
  - Request-scoped values

## 7. Defer

```go
func main() {
    pool, _ := db.Connect(cfg)
    defer pool.Close()  // Akan dipanggil saat main() selesai
    
    // ... kode lainnya ...
}  // pool.Close() dipanggil di sini

func GetAll() ([]User, error) {
    rows, err := db.Query(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()  // Cleanup rows setelah function selesai
    
    // ... iterasi rows ...
    return users, nil
}  // rows.Close() dipanggil di sini
```

**Penjelasan:**
- `defer` - Menunda eksekusi statement sampai function selesai
- Dieksekusi dalam urutan LIFO (Last In First Out)
- Berguna untuk cleanup resources (close file, connection, dll)

## 8. Pointer

```go
// Tanpa pointer - pass by value (copy)
func UpdateUser(user User) {
    user.Name = "New Name"  // Hanya modify copy
}

// Dengan pointer - pass by reference
func UpdateUser(user *User) {
    user.Name = "New Name"  // Modify original
}

// Penggunaan
user := User{Name: "John"}
UpdateUser(&user)  // & = ambil address/pointer
fmt.Println(user.Name)  // Output: "New Name"
```

**Penjelasan:**
- `*Type` - Pointer type
- `&variable` - Ambil address (pointer) dari variable
- `*pointer` - Dereference (akses value dari pointer)
- Pointer berguna untuk:
  - Modify original value
  - Efisiensi (tidak copy data besar)
  - Nil checking

## 9. Slice

```go
// Deklarasi slice
var users []User              // Nil slice
users := []User{}             // Empty slice
users := make([]User, 0, 10)  // Capacity 10

// Append ke slice
users = append(users, user)

// Iterasi slice
for i, user := range users {
    fmt.Printf("%d: %s\n", i, user.Name)
}

// Iterasi tanpa index
for _, user := range users {
    fmt.Println(user.Name)
}
```

**Penjelasan:**
- Slice = dynamic array
- `make([]Type, length, capacity)` - Buat slice dengan capacity
- `append()` - Tambah element (auto resize jika perlu)
- `range` - Iterate over slice, array, map, channel
- `_` - Blank identifier (ignore value)

## 10. JSON Encoding/Decoding

```go
// Decode JSON dari request body
var user User
err := json.NewDecoder(r.Body).Decode(&user)
if err != nil {
    // Handle error
}

// Encode struct ke JSON response
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(user)

// Marshal/Unmarshal (untuk []byte)
jsonBytes, _ := json.Marshal(user)
json.Unmarshal(jsonBytes, &user)
```

**Penjelasan:**
- `json.Decoder` - Decode dari io.Reader (stream)
- `json.Encoder` - Encode ke io.Writer (stream)
- `json.Marshal` - Convert struct → []byte
- `json.Unmarshal` - Convert []byte → struct
- Struct tag menentukan nama field di JSON

## 11. HTTP Handler

```go
// Handler function signature
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    // w = untuk write response
    // r = untuk read request
    
    // Set header
    w.Header().Set("Content-Type", "application/json")
    
    // Set status code
    w.WriteHeader(http.StatusCreated)  // 201
    
    // Write body
    json.NewEncoder(w).Encode(user)
}

// Register handler
mux := http.NewServeMux()
mux.HandleFunc("POST /users", handler.CreateUser)  // Go 1.22+ routing
mux.HandleFunc("GET /users", handler.GetUsers)
mux.HandleFunc("DELETE /users/{id}", handler.DeleteUser)

// Path parameter (Go 1.22+)
id := r.PathValue("id")
```

**Penjelasan:**
- `http.ResponseWriter` - Interface untuk write HTTP response
- `*http.Request` - Struct berisi HTTP request data
- `http.NewServeMux()` - HTTP request multiplexer (router)
- Go 1.22+ mendukung method dan path parameter di routing

## 12. Middleware Pattern

```go
// Middleware function signature
func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Before handler
        start := time.Now()
        
        // Call next handler
        next.ServeHTTP(w, r)
        
        // After handler
        log.Printf("Duration: %s", time.Since(start))
    })
}

// Chain middleware
handler := Recovery(Logger(mux))
```

**Penjelasan:**
- Middleware = function yang wrap handler
- Pattern: `func(http.Handler) http.Handler`
- Bisa jalankan code sebelum dan sesudah handler
- Chain multiple middleware dengan nesting

## 13. Database Operations (pgx)

```go
// Query single row
var user User
err := db.QueryRow(ctx, "SELECT id, name FROM users WHERE id = $1", id).
    Scan(&user.ID, &user.Name)

// Query multiple rows
rows, err := db.Query(ctx, "SELECT id, name FROM users")
defer rows.Close()

for rows.Next() {
    var user User
    rows.Scan(&user.ID, &user.Name)
    users = append(users, user)
}

// Execute (INSERT, UPDATE, DELETE)
result, err := db.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
rowsAffected := result.RowsAffected()
```

**Penjelasan:**
- `$1, $2` - Placeholder untuk prepared statement
- `QueryRow()` - Untuk query yang return 1 row
- `Query()` - Untuk query yang return multiple rows
- `Exec()` - Untuk statement yang tidak return rows
- `Scan()` - Map column ke variable
- Prepared statement mencegah SQL injection

## 14. Testing

```go
// Test function naming: Test + FunctionName
func TestUser_Validate(t *testing.T) {
    // Table-driven test
    tests := []struct {
        name    string
        user    User
        wantErr bool
    }{
        {
            name:    "valid user",
            user:    User{Name: "John", Email: "john@example.com"},
            wantErr: false,
        },
        {
            name:    "empty name",
            user:    User{Name: "", Email: "john@example.com"},
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := tt.user.Validate()
            if (err != nil) != tt.wantErr {
                t.Errorf("got error = %v, want error = %v", err, tt.wantErr)
            }
        })
    }
}
```

**Penjelasan:**
- Test file: `*_test.go`
- Test function: `func TestXxx(t *testing.T)`
- `t.Run()` - Subtest
- `t.Errorf()` - Report error (test continue)
- `t.Fatalf()` - Report error dan stop test
- Table-driven test = best practice untuk multiple test cases

## 15. Dependency Injection

```go
// Bad: Hard-coded dependency
type UserService struct {
    repo *UserRepo  // Concrete type
}

// Good: Inject interface
type UserService struct {
    repo UserRepository  // Interface
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}

// Usage
repo := &UserRepo{DB: pool}
service := NewUserService(repo)  // Inject dependency
```

**Penjelasan:**
- Inject dependency melalui constructor function
- Gunakan interface untuk abstraksi
- Memudahkan testing (bisa inject mock)
- Loose coupling antar layer

## 🎯 Best Practices

1. **Error Handling**: Selalu check error, jangan ignore
2. **Context**: Gunakan context untuk timeout dan cancellation
3. **Defer**: Cleanup resources dengan defer
4. **Interface**: Gunakan interface untuk abstraksi
5. **Pointer**: Gunakan pointer untuk struct besar atau perlu modify
6. **Validation**: Validasi input sebelum process
7. **Testing**: Tulis test untuk business logic
8. **Naming**: Gunakan nama yang descriptive
9. **Package**: Organize code by feature/layer
10. **Documentation**: Comment exported functions

## 📖 Resources

- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Documentation](https://go.dev/doc/)
